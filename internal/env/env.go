package env

import (
	"os"
	"strconv"
)

func GetString(key, fallback string) string {
	value,found := os.LookupEnv(key)
	if !found {
		return fallback
	}
	return value 
}

func GetInt(key string,fallback int) int {
	value,found := os.LookupEnv(key)
	if !found{
		return fallback
	}
	valueInt, err := strconv.Atoi(value)
	if err != nil  {
		return fallback
	}
	return valueInt
}

func GetBool(key string,fallback bool) bool {
	value,found := os.LookupEnv(key)
	if !found {
		return fallback
	}
	valueBool,err := strconv.ParseBool(value)
	if err != nil {
		return fallback
	}
	return valueBool
}