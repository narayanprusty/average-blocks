package db

type User struct {
	Id       int64
	Username string
	Password string
}

type APIKey struct {
	Id     int64
	Key    string
	UserId int64
	User   *User `pg:"rel:has-one"`
}
