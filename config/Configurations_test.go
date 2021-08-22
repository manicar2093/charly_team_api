package config

import (
	"testing"
)

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
