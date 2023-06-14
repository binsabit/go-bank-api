package postgres

import (
	"database/sql"

	"github.com/binsabit/go-bank-api/internal/data"
	_ "github.com/lib/pq"
)

type PostgresStore struct {
	db *sql.DB
}

//const conString = "postgres://YourUserName:YourPassword@YourHostname:5432/YourDatabaseName";

func NewPostgreStore() (*PostgresStore, error) {
	//hardcoded db connection
	connStr := "postgres://postgres:admin@localhost:5432/gobank?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (s PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `
		create table if not exists account(
			id serial primary key,
			first_name varchar(50),
			last_name varchar(50),
			number serial,
			balance serial,
			created_at timestamp
		);`
	_, err := s.db.Query(query)
	return err
}

func (s PostgresStore) CreateAccount(*data.Account) error {
	query := `
		insert into account (first_name,last_name, number, balance, created_at) 
		values($1,$2,$3,$4,$5)
	`

	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil

}
func (s PostgresStore) DeleteAccount(id int) error {

	query := `delete from account where id = $1;`

	_, err := s.db.Exec(query, id)
	if err != nil {
		return err
	}

	return nil

}
func (s PostgresStore) UpdateAccount(id int, acc *data.Account) error {
	query := `update account
		set first_name = $1,
		set last_name = $2,
		number = $3,
		balance = $4
		where id = $5;	
	`
	_, err := s.db.Exec(query, acc.FirstName, acc.LastName, acc.Number, acc.Balance, id)

	if err != nil {
		return err
	}
	return nil

}
func (s PostgresStore) GetAccountByID(ID int) (*data.Account, error) {
	query := `select (id,first_name, last_name, number, balance, created_at) from account where id=$1`

	var acc data.Account

	err := s.db.QueryRow(query, ID).Scan(acc.FirstName, acc.LastName, acc.Number, acc.Balance, acc.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &acc, nil
}
