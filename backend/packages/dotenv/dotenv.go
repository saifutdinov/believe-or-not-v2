package dotenv

import (
	"errors"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func LoadEnvFile(path ...string) (*Env, error) {

	var filepath string
	if len(path) > 0 {
		filepath = path[0]
	} else {
		filepath = ".env"
	}

	parsedfile, err := readAndParse(filepath)
	if err != nil {
		return nil, err
	}

	return fillStruct(parsedfile)
}

func readAndParse(filepath string) (map[string]string, error) {
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	rows := strings.Split(string(file), "\n")

	if len(rows) == 0 {
		return nil, errors.New("no values in .env file")
	}

	result := make(map[string]string)
	for _, row := range rows {
		if row == "" || row[0] == '#' {
			continue
		}

		parts := strings.Split(row, "=")
		if len(parts) != 2 {
			continue
		}
		result[parts[0]] = parts[1]
	}
	return result, nil
}

func fillStruct(parsedfile map[string]string) (*Env, error) {
	env := new(Env)
	v := reflect.ValueOf(env).Elem()
	t := v.Type()

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		structField := t.Field(i)
		tag := structField.Tag.Get("env")
		if tag == "" {
			continue
		}
		val, ok := parsedfile[tag]
		if !ok {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(val)
		case reflect.Int, reflect.Int64:
			intVal, err := strconv.Atoi(val)
			if err != nil {
				return nil, err
			}
			field.SetInt(int64(intVal))
		case reflect.Bool:
			boolVal, err := strconv.ParseBool(val)
			if err != nil {
				return nil, err
			}
			field.SetBool(boolVal)
		}
	}
	return env, nil
}
