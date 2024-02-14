package integration

import "github.com/joho/godotenv"

// InitEnv initializes the environment.
//
// It does not take any parameters and returns an error.
func InitEnv() error {
	return godotenv.Load()
}
