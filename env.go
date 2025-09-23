package flint

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

type Env struct{}

func EnvLoad(file string) *Env {
	if file == "" {
		file = ".env"
	}

	f, err := os.Open(file)
	if err != nil {
		LogError("EnvLoad()", ".env file not found.")
		return &Env{}
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			continue
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		os.Setenv(key, value)
	}

	if err := scanner.Err(); err != nil {
		log.Println("Error reading .env file:", err)
	}

	return &Env{}
}

func (e *Env) GetEnv(key string, defaults ...string) string {
	if len(defaults) > 1 {
		log.Fatal("Env.Getenv: sadece bir opsiyonel default değer girebilirsiniz")
	}

	value := os.Getenv(key)
	if value != "" {
		return value
	}

	if len(defaults) == 1 {
		return defaults[0]
	}

	return ""
}

func (e *Env) GetInt(key string, defaults ...int) int {
	val := e.GetEnv(key)
	if val == "" && len(defaults) == 1 {
		return defaults[0]
	}
	i, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("Env.GetInt: %v is not an int", val)
	}
	return i
}

func (e *Env) GetBool(key string, defaults ...bool) bool {
	val := e.GetEnv(key)
	if val == "" && len(defaults) == 1 {
		return defaults[0]
	}
	b, err := strconv.ParseBool(val)
	if err != nil {
		log.Fatalf("Env.GetBool: %v is not a bool", val)
	}
	return b
}

func (e *Env) GetFloat(key string, defaults ...float64) float64 {
	val := e.GetEnv(key)
	if val == "" && len(defaults) == 1 {
		return defaults[0]
	}
	f, err := strconv.ParseFloat(val, 64)
	if err != nil {
		log.Fatalf("Env.GetFloat: %v is not a float", val)
	}
	return f
}

func (e *Env) GetChar(key string, defaults ...string) {

}
