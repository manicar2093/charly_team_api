package config

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	godotenv.Load("../.env.example")
	StartConfig()

	os.Exit(m.Run())
}

func TestGetEnvOrPanica(t *testing.T) {

	t.Run("should panic if env variable does not exists", func(t *testing.T) {
		defer func() {
			if r := recover(); r == nil {
				t.Fatal("should panic. env variable does not exists")
			}
		}()
		GetEnvOrPanic("NOT_EXISTS")
	})
}

func TestGetDBConnectionURL(t *testing.T) {

	t.Run("if DB_URL is not set should return generated ULR", func(t *testing.T) {
		expected_content := "localhost"
		got := DBConnectionURL()

		assert.Contains(t, got, expected_content, "incorrect generated URL")
	})

	t.Run("if DB_URL is set should return the env variable value", func(t *testing.T) {
		expected := "my_setted_url"
		os.Setenv("DB_URL", expected)
		StartConfig()

		got := DBConnectionURL()

		assert.Equal(t, expected, got, "unexpected url")
	})
}
