package request

import (
	"errors"
	"strings"
)

type CBool bool

var ErrNotBoolean = errors.New("not a boolean")

func (b *CBool) UnmarshalJSON(data []byte) error {
	str := strings.Trim(string(data), `"`)
	if str == "true" || str == "1" {
		*b = true
		return nil
	}
	if str == "false" || str == "0" {
		*b = false
		return nil
	}
	return ErrNotBoolean
}
