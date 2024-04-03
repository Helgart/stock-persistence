package database

import (
	"database/sql"
	"fmt"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
)

func TestDatabase_Ping(t *testing.T) {
	tests := []struct {
		name    string
		setup   func(mock sqlmock.Sqlmock)
		wantErr bool
	}{
		{
			name: "Success",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectPing().WillReturnError(nil)
			},
			wantErr: false,
		},
		{
			name: "Failure",
			setup: func(mock sqlmock.Sqlmock) {
				mock.ExpectPing().WillReturnError(sql.ErrConnDone)
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.MonitorPingsOption(true))

			if err != nil {
				t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
			}

			tt.setup(mock)

			database := NewDatabase(db)

			if err := database.Ping(); (err != nil) != tt.wantErr {
				t.Errorf("Database.Ping() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("Database.Ping() there were unfulfilled expectations: %s", err)
			}
		})
	}
}

func TestConnexionError_Error(t *testing.T) {
	tests := []struct {
		name string
		err  ConnexionError
		want string
	}{
		{
			name: "Error Message",
			err: ConnexionError{
				PreviousError:      fmt.Errorf("previousError"),
				PsqlInfo:           "psqlInfo",
				DatabaseParameters: DatabaseParameters{User: "user", Host: "host", Port: 9999, Name: "name"},
			},
			want: fmt.Sprintf(connexionErrorMessage, "user", "host", 9999, "name", "previousError"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.err.Error(); got != tt.want {
				t.Errorf("ConnexionError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}
