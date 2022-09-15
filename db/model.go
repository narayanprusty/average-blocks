package db

type User struct {
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

type APIKey struct {
	Key  string `mapstructure:"apiKey"`
	User *User  `mapstructure:"username" pg:"rel:has-one"`
}
