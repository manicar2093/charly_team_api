package config

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	godotenv.Load("../../.env")
	StartConfig()

	os.Exit(m.Run())
}

func TestGetDBConnectionURL(t *testing.T) {

	t.Run("if DB_URL is set should return the env variable value", func(t *testing.T) {
		expected := "my_setted_url"
		os.Setenv("DB_URL", expected)
		StartConfig()

		got := DBConnectionURL()

		assert.Equal(t, expected, got, "unexpected url")
	})
}
