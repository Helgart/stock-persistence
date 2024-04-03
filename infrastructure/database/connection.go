package database

import (
	"database/sql"
	"fmt"
)

// connexionErrorMessage represents the error message format for a database connection error.
const connexionErrorMessage = "Database connexion error (%s@%s:%d/%s) : %s"

// Database is a type that represents a database connection.
type Database struct {
	connexion *sql.DB
}

// DatabaseParameters represents the parameters used to establish a database connection.
type DatabaseParameters struct {
	Host     string
	Port     uint
	User     string
	Password string
	Name     string
}

// ConnexionError represents an error that occurred during database connection.
type ConnexionError struct {
	PreviousError      error
	PsqlInfo           string
	DatabaseParameters DatabaseParameters
}

// Error returns the error message representation of a ConnexionError.
func (connexionError ConnexionError) Error() string {
	return fmt.Sprintf(
		connexionErrorMessage,
		connexionError.DatabaseParameters.User,
		connexionError.DatabaseParameters.Host,
		connexionError.DatabaseParameters.Port,
		connexionError.DatabaseParameters.Name,
		connexionError.PreviousError.Error(),
	)
}

// NewDatabase creates a new Database struct with the provided connection.
// It takes a pointer to sql.DB and returns a Database instance.
func NewDatabase(connexion *sql.DB) Database {
	return Database{connexion: connexion}
}

// Ping pings the database to check its availability.
// It returns an error if the ping fails.
func (database *Database) Ping() error {
	err := database.connexion.Ping()

	if err != nil {
		return err
	}

	return nil
}
