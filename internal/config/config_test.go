package config_test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
	"github.com/manicar2093/health_records/internal/config"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {

	godotenv.Load("../../.env")
	config.StartConfig()

	os.Exit(m.Run())
}

func TestGetDBConnectionURL(t *testing.T) {

	t.Run("if DB_URL is set should return the env variable value", func(t *testing.T) {
		expected := "my_setted_url"
		os.Setenv("DB_URL", expected)
		config.StartConfig()

		got := config.DBConnectionURL()

		assert.Equal(t, expected, got, "unexpected url")
	})
}
