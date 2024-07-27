package main

import (
	"fmt"
	"os"

	"gorm.io/gorm"
)

const (
	TEST_HOST     = "test"
	TEST_PASSWORD = "test"
	TEST_DBNAME   = "test"
	TEST_USER     = "test"
	TEST_SSLMODE  = "disable"
)

type Storage interface {
	NewCommandSubset(subsetName string)
	NewCommandSet(subsetName string, commands []CommandSet)
}

type PostgresStorage struct {
	config StorageConfig
	db     *gorm.DB
}

func WithTestEnv() func(*PostgresStorage) {
	return func(s *PostgresStorage) {
		s.config = NewStorageConfig(
			TEST_USER,
			TEST_PASSWORD,
			TEST_HOST,
			TEST_DBNAME,
			TEST_SSLMODE,
		)
	}
}

func WithProdEnv() func(*PostgresStorage) {
	return func(s *PostgresStorage) {
		user := os.Getenv("USER")
		password := os.Getenv("PASSWORD")
		host := os.Getenv("HOST")
		db := os.Getenv("DB")
		sslmode := os.Getenv("SSLMODE")

		s.config = NewStorageConfig(
			user,
			password,
			host,
			db,
			sslmode,
		)
	}
}

func NewPostgresStorage(opts ...func(*PostgresStorage)) (*PostgresStorage, error) {

	ps := &PostgresStorage{}

	for _, o := range opts {
		o(ps)
	}

	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=5432 sslmode=%s TimeZone=Europe/Warsaw",
		ps.config.host,
		ps.config.user,
		ps.config.password,
		ps.config.dbname,
		ps.config.sslmode,
	)

	fmt.Println(connStr)

	return &PostgresStorage{
		ps.config,
		nil,
	}, nil
}
