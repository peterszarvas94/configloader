package envloader

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strings"
)

var envPath string

func File(path string) {
	envPath = path
}

func Loader(variables interface{}, keys ...string) error {
	if envPath != "" {
		err := loadEnvFromFile(envPath)
		if err != nil {
			return err
		}
	}

	v := reflect.ValueOf(variables)
	if v.Kind() != reflect.Ptr || v.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("initEnv expects a pointer to a struct")
	}

	structValue := v.Elem()

	for i := 0; i < structValue.NumField(); i++ {
		field := structValue.Field(i)
		fieldType := structValue.Type().Field(i)

		if field.Kind() != reflect.String {
			return fmt.Errorf("Field %s must be of type string", fieldType.Name)
		}

		envVarName := strings.ToUpper(fieldType.Name)
		envVarValue, found := os.LookupEnv(envVarName)

		if !found {
			return fmt.Errorf("Environment variable %s not found", envVarName)
		}

		field.SetString(envVarValue)
	}

	return nil
}

func loadEnvFromFile(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err
	}

	for _, line := range lines {
		// Skip empty lines and comments
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return fmt.Errorf("Invalid line: %s", line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		// strip quotes
		if len(value) >= 2 && value[0] == '"' && value[len(value)-1] == '"' {
			value = value[1 : len(value)-1]
		}

		os.Setenv(key, value)
	}

	return nil
}
