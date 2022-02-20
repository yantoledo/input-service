package config

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

var skipChecks = map[string]bool{
	"AWS_TOKEN": true,
}

// GetString ...
func GetString(name string, defaultValue ...string) string {
	value := os.Getenv(name)
	if value == "" && len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return value
}

// MustGetString ...
func MustGetString(name string) string {
	value := os.Getenv(name)
	if value == "" && !skipChecks[name] {
		fmt.Printf("%s can't be empty\n", name)
		os.Exit(1)
	}
	return value
}

// GetInt ...
func GetInt(name string, defaultValue ...int) int {
	value, err := strconv.Atoi(os.Getenv(name))
	if err != nil && len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return value
}

// MustGetInt ...
func MustGetInt(name string) int {
	value, err := strconv.Atoi(os.Getenv(name))
	if err != nil {
		fmt.Printf("%s must contain a float value!\n", name)
		os.Exit(1)
	}
	return value
}

// GetFloat ...
func GetFloat(name string, defaultValue ...float64) float64 {
	value, err := strconv.ParseFloat(os.Getenv(name), 64)
	if err != nil && len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return value
}

// MustGetFloat ...
func MustGetFloat(name string) float64 {
	value, err := strconv.ParseFloat(os.Getenv(name), 64)
	if err != nil {
		fmt.Printf("%s must contain a float value!\n", name)
		os.Exit(1)
	}
	return value
}

// GetBool ...
func GetBool(name string, defaultValue ...bool) bool {
	value, err := strconv.ParseBool(os.Getenv(name))
	if err != nil && len(defaultValue) > 0 {
		value = defaultValue[0]
	}
	return value
}

// MustGetBool ...
func MustGetBool(name string) bool {
	value, err := strconv.ParseBool(os.Getenv(name))
	if err != nil {
		fmt.Printf("%s must contain a boolean value! (true or false)\n", name)
		os.Exit(1)
	}
	return value
}

// MustGetTime ...
func MustGetTime(name, format string) time.Time {
	value, err := time.Parse(format, os.Getenv(name))
	if err != nil {
		fmt.Printf("%s must contain a date value!\n", name)
		os.Exit(1)
	}

	value = time.Date(value.Year(), value.Month(), value.Day(), value.Hour(), value.Minute(), value.Second(), value.Nanosecond(), value.Location())
	return value
}
