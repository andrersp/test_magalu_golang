package config

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

const (
	localEnvironment = "LOCAL"
	environmentKEY   = "ENVIRONMENT"
)

func init() {
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
}

func ReadConfigFromEnv(config interface{}) error {
	bindEnvs(config)

	if errU := viper.Unmarshal(config); errU != nil {
		return errU
	}

	return validator.New().Struct(config)
}

func ReadConfigFromEnvFile() error {
	err := godotenv.Load()
	if err != nil && strings.EqualFold(os.Getenv(environmentKEY), localEnvironment) {
		return fmt.Errorf("error loading .env file: %w", err)
	}

	return nil
}

func bindEnvs(config interface{}, parts ...string) {
	vc := reflect.ValueOf(config)

	if vc.Kind() == reflect.Ptr {
		vc = vc.Elem()
	}

	for i := 0; i < vc.NumField(); i++ {
		field := vc.Field(i)
		structField := vc.Type().Field(i)
		tv, ok := structField.Tag.Lookup("mapstructure")

		if !ok {
			continue
		}

		if field.Kind() == reflect.Struct {
			bindEnvs(field.Interface(), append(parts, tv)...)
		} else {
			_ = viper.BindEnv(strings.Join(append(parts, tv), "."))
		}
	}
}
