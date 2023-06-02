package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccounts() ([]*Account, error)
	GetAccountByID(int) (*Account, error)
	GetAccountByNumber(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {

	connStr := "user=postgres dbname=postgres password=msr sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {

	return s.CreateAccountTabe()

}

func (s *PostgresStore) CreateAccountTabe() error {
	query := `create table  if not exists account(
			id serial primary key,
			first_name varchar(50),
			last_name varchar(50),
			number serial,
			encrypted_password varchar(100),
			balance serial,
			create_at timestamp
		)`

	_, err := s.db.Exec(query)

	return err

}

func (s *PostgresStore) CreateAccount(acc *Account) error {

	query := `insert into account(first_name,last_name,number,encrypted_password,balance,create_at)
	values($1,$2,$3,$4,$5,$6)
	`
	resp, err := s.db.Query(
		query,
		acc.FirstName,
		acc.LastName,
		acc.Number,
		acc.Encryptedpassword,
		acc.Balance,
		acc.CreatedAt)

	if err != nil {
		return err
	}

	fmt.Printf("%+v\n", resp)

	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {

	_, err := s.db.Query("delete from account where id = $1", id)
	return err
}

func (s *PostgresStore) GetAccountByNumber(number int) (*Account, error) {

	rows, err := s.db.Query("select * from account  where  number = $1", number)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("Account with Number [%d] Not found", number)
}

func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {

	rows, err := s.db.Query("select * from account  where  id = $1", id)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		return scanIntoAccount(rows)
	}

	return nil, fmt.Errorf("Account %d not found", id)
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("select * from account")

	if err != nil {
		return nil, err
	}

	accounts := []*Account{}

	for rows.Next() {
		account, err := scanIntoAccount(rows)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, account)
	}

	return accounts, nil
}

func scanIntoAccount(row *sql.Rows) (*Account, error) {
	account := new(Account)
	err := row.Scan(
		&account.ID,
		&account.FirstName,
		&account.LastName,
		&account.Number,
		&account.Encryptedpassword,
		&account.Balance,
		&account.CreatedAt,
	)

	return account, err
}