package database

import (
	"fmt"
	"os"
	"strconv"
)

// databaseParametersErrorMessage represents the error message when database parameters couldn't be created.
const databaseParametersErrorMessage = "Couldn't create database parameters : %s"

// DatabaseParameters represents the parameters used to establish a database connection.
type DatabaseParameters struct {
	Host     string
	Port     uint
	User     string
	Password string
	Name     string
}

// DatabaseParametersError represents an error that occurs when there is a problem with the database parameters.
type DatabaseParametersError struct {
	PreviousError error
}

// Error returns the error message of a DatabaseParametersError.
func (databaseParametersError DatabaseParametersError) Error() string {
	return fmt.Sprintf(databaseParametersErrorMessage, databaseParametersError.PreviousError)
}

// NewDatabaseParameters is a function that creates a new DatabaseParameters object with the provided parameters.
func NewDatabaseParameters(host string, port uint, user, password, name string) DatabaseParameters {
	return DatabaseParameters{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Name:     name,
	}
}

// NewDatabaseParametersFromEnv creates a new DatabaseParameters struct based on environment variables.
// It retrieves the values for host, port, user, password, and name from the environment variables:
// - DATABASE_HOST
// - DATABASE_PORT
// - DATABASE_USER
// - DATABASE_PASSWORD
// - DATABASE_NAME
// It returns the created DatabaseParameters struct and any error encountered during the process.
func NewDatabaseParametersFromEnv() (DatabaseParameters, error) {
	port, err := strconv.ParseUint(os.Getenv("DATABASE_PORT"), 10, 32)

	if err != nil {
		return DatabaseParameters{}, DatabaseParametersError{PreviousError: err}
	}

	return NewDatabaseParameters(
		os.Getenv("DATABASE_HOST"),
		uint(port),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
	), nil
}
