package database

import (
	"fmt"
	"os"
	"testing"
)

func TestNewDatabaseParameters(t *testing.T) {
	cases := []struct {
		name                   string
		host                   string
		port                   uint
		user, password, dbname string
		expected               DatabaseParameters
	}{
		{
			name:     "default",
			host:     "localhost",
			port:     5432,
			user:     "root",
			password: "secret",
			dbname:   "testdb",
			expected: DatabaseParameters{
				Host:     "localhost",
				Port:     5432,
				User:     "root",
				Password: "secret",
				Name:     "testdb",
			},
		},
		// Additional test cases...
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			actual := NewDatabaseParameters(c.host, c.port, c.user, c.password, c.dbname)
			if actual != c.expected {
				t.Errorf("NewDatabaseParameters(%v, %v, %v, %v, %v) == %v, expected %v", c.host, c.port, c.user, c.password, c.dbname, actual, c.expected)
			}
		})
	}
}

func TestNewDatabaseParametersFromEnv(t *testing.T) {
	cases := []struct {
		name        string
		envValues   map[string]string
		expected    DatabaseParameters
		expectError bool
	}{
		{
			name: "default",
			envValues: map[string]string{
				"DATABASE_HOST":     "localhost",
				"DATABASE_PORT":     "5432",
				"DATABASE_USER":     "root",
				"DATABASE_PASSWORD": "secret",
				"DATABASE_NAME":     "testdb",
			},
			expected: DatabaseParameters{
				Host:     "localhost",
				Port:     5432,
				User:     "root",
				Password: "secret",
				Name:     "testdb",
			},
			expectError: false,
		},
		// Additional test cases...
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// Set environment variables
			for k, v := range c.envValues {
				err := os.Setenv(k, v)

				if err != nil {
					t.Errorf(fmt.Sprintf("Expected no error on seeting env, getting %s", err.Error()))
				}

				defer func(k string) {
					_ = os.Unsetenv(k)

					if err != nil {
						t.Errorf(fmt.Sprintf("Expected no error on cleaning up env, getting %s", err.Error()))
					}
				}(k)
			}

			actual, err := NewDatabaseParametersFromEnv()

			if c.expectError {
				if err == nil {
					t.Errorf("Expected an error but got nil")
					return
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect error, but got %v", err)
					return
				}
				if actual != c.expected {
					t.Errorf("NewDatabaseParametersFromEnv() == %v, expected %v", actual, c.expected)
				}
			}
		})
	}
}
