package dotenv

// Опишите свои поля в формате:
//
//	SecretAPIKey `env:"SECRET_API_KEY"`
//
// Соответственно в .env файле должен быть ключ SECRET_API_KEY=supersecret_api_key
// //
type Env struct {
	BackendPort string `env:"BACKEND_PORT"`

	SqliteConnection string `env:"SQLITE_CONNECTION_FILE"`

	CookieName string `env:"COOKIE_NAME"`
}
