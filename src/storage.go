package main

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	TEST_HOST     = "test"
	TEST_PASSWORD = "test"
	TEST_DBNAME   = "test"
	TEST_USER     = "test"
	TEST_SSLMODE  = "disable"
)

type Storage interface {
	CreateCommandSubset(subsetName string) error
	CreateCommandSet(subsetName string, commandset CommandSet) error
	GetCommandSubsets(subsetName string) ([]CommandSubset, error)
	GetCommandSet(subsetName, shortDescription string)
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
		user := os.Getenv("POSTGRES_USER")
		password := os.Getenv("POSTGRES_PASSWORD")
		host := os.Getenv("POSTGRES_HOST")
		db := os.Getenv("POSTGRES_DB")
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

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatalln(err)
		return nil, err
	}

	if db.AutoMigrate(&CommandSet{}, &CommandSubset{}) != nil {
		return nil, err
	}

	return &PostgresStorage{
		ps.config,
		db,
	}, nil
}
