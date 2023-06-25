package app

import "os"

func getAppEnv() string {
	return os.Getenv("APP_ENV")
}

func IsDevelopment() bool {
	return getAppEnv() == "development"
}

func IsLocal() bool {
	return os.Getenv("LOCAL") == "true"
}
