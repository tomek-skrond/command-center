package main

type StorageConfig struct {
	user     string
	password string
	host     string
	dbname   string
	sslmode  string
}

func NewStorageConfig(user, password, host, dbname, sslmode string) StorageConfig {
	return StorageConfig{
		user:     user,
		password: password,
		host:     host,
		dbname:   dbname,
		sslmode:  sslmode,
	}
}
