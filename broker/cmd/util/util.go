package util

import (
	"errors"
	"fmt"
	"log"
	"os"
)

var ErrAssertEnv error = errors.New("[ErrAssertEnv]: ")

func AssertEnv(name string) string {
	envVar := os.Getenv(name)

	if envVar == "" {
		log.Fatal(fmt.Errorf("%wNo env variable %s was provided", ErrAssertEnv, name))
	}

	return envVar
}
