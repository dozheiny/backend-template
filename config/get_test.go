package config_test

import (
	"github.com/dozheiny/backend-template/config"
	"os"
	"strings"
	"testing"
)

func TestGetSuccessFul(t *testing.T) {
	t.Parallel()

	if err := os.Setenv("env_test", "successful"); err != nil {
		t.Errorf("Got error on set config, %v", err.Error())
	}

	value, err := config.Get("env_test")
	if err != nil {
		t.Errorf("Got error on config.Get package: %v", err.Error())
	}

	if strings.Compare(value, "successful") != 0 {
		t.Errorf("Value is not compare, excepted: %s, actual: %s",
			"successful", value)
	}

}

func TestGetUnsuccessful(t *testing.T) {
	t.Parallel()

	_, err := config.Get("not_env_test")
	if err == nil {
		t.Errorf("Didn't Got error on config.Get package")
	}
}
