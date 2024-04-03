package postgres

import (
	"database/sql"
	"fmt"
	"github.com/Helgart/stock-persistence/infrastructure/database"

	_ "github.com/lib/pq"
)

const infoFormat = "host=%s port=%d user=%s password=%s dbname=%s sslmode=disable"
const driver = "postgres"

func NewDatabase(databaseParameters database.DatabaseParameters) (database.Database, error) {
	psqlInfo := fmt.Sprintf(
		infoFormat,
		databaseParameters.Host,
		databaseParameters.Port,
		databaseParameters.User,
		databaseParameters.Password,
		databaseParameters.Name,
	)

	connexion, err := sql.Open(driver, psqlInfo)

	if err != nil {
		return database.Database{}, database.ConnexionError{
			PreviousError:      err,
			PsqlInfo:           psqlInfo,
			DatabaseParameters: databaseParameters,
		}
	}

	if err := connexion.Ping(); err != nil {
		return database.Database{}, database.ConnexionError{
			PreviousError:      err,
			PsqlInfo:           psqlInfo,
			DatabaseParameters: databaseParameters,
		}
	}

	return database.NewDatabase(connexion), nil
}
