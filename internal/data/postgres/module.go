package postgres

import (
	"database/sql"

	"github.com/binsabit/go-bank-api/internal/data"
)

type PostgresStore struct {
	db *sql.DB
}

//const conString = "postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName";

func NewPostgreStore() (*PostgresStore, error) {
	//hardcoded db connection
	connStr := "postgres://postgres:admin@localhost/pqgotest?sslmode=disabled"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (s PostgresStore) Init() error {
	return nil
}

func (s PostgresStore) CreateAccount(*data.Account) error {
	return nil

}
func (s PostgresStore) DeleteAccount(int) error {
	return nil

}
func (s PostgresStore) UpdateAccount(*data.Account) error {
	return nil

}
func (s PostgresStore) GetAccountByID(ID int) (*data.Account, error) {
	return nil, nil
}
