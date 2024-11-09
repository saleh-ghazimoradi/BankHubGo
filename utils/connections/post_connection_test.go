package utils

import (
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestPostgresURI(t *testing.T) {
	cfg := PostConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "testuser",
		Password: "testpass",
		Database: "testdb",
		SSLMode:  "disable",
	}

	expectedURI := "host=localhost port=5432 user=testuser password=testpass dbname=testdb sslmode=disable"
	actualURI := PostgresURI(cfg)
	assert.Equal(t, expectedURI, actualURI, "The URI does not match the expected format")
}

func TestPostgresUrl(t *testing.T) {
	cfg := PostConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "testuser",
		Password: "testpass",
		Database: "testdb",
		SSLMode:  "disable",
	}

	expectedUrl := "postgres://testuser:testpass@localhost:5432/testdb?sslmode=disable"
	actualUrl := PostgresUrl(cfg)
	assert.Equal(t, expectedUrl, actualUrl, "The URL does not match the expected format")
}

func TestPostConnection(t *testing.T) {
	cfg := PostConfig{
		Host:         "localhost",
		Port:         "5432",
		User:         "invaliduser",
		Password:     "invalidpass",
		Database:     "invaliddb",
		SSLMode:      "disable",
		MaxOpenConns: 5,
		MaxIdleConns: 5,
		MaxIdleTime:  time.Minute,
		Timeout:      1 * time.Millisecond,
	}

	db, err := PostConnection(cfg)
	assert.Nil(t, db, "Expected no database connection")
	assert.Error(t, err, "Expected an error due to connection timeout or invalid credentials")
}
