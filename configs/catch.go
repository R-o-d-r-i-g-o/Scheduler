package configs

import (
	"encoding/json"
	"os"
	"reflect"
	"scheduler/constants"
	"scheduler/util"

	"github.com/joho/godotenv"
)

func loadStructWithEnvVars(configStructure interface{}) {
	reflectElement := reflect.ValueOf(configStructure).Elem()

	environments := make(map[string]string)
	loadEnvFile()

	elements := reflectElement.Type()
	for i := 0; i < elements.NumField(); i++ {
		field := elements.Field(i)
		fieldName, envVarName := field.Name, field.Tag.Get("env")

		envVarValue := os.Getenv(envVarName)
		environments[fieldName] = envVarValue
	}

	parsed := util.ParseMapToJson(environments)
	json.Unmarshal([]byte(parsed), &configStructure)
}

func loadEnvFile() {
	if err := godotenv.Load(".env"); err != nil {
		panic(constants.FAIL_LOAD_ENV)
	}
}
