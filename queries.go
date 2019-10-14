package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const createTable = `
	DROP TABLE IF EXISTS ACCOUNT;
	CREATE TABLE ACCOUNT
	(
		ID serial,
		NAME varchar(50),
		MAIL_ADDRESS varchar(50),
		LANG varchar(5)
	)
`

type Account struct {
	ID          int    "ID"
	Name        string "NAME"
	MailAddress string "MAIL_ADDRESS"
	Lang        string "LANG"
}

func connectPostgres() (*sql.DB, error) {
	connStr := "postgres://restuu:postgres@localhost:5432/postgres?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	fmt.Println("Ping OK")

	return db, nil
}

func insertAccount(db *sql.DB) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	const insertRow = `
		INSERT INTO ACCOUNT (NAME, MAIL_ADDRESS, LANG)
		VALUES ('john', 'john@email.com', 'en')
	`

	_, err = tx.Exec(insertRow)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func getAccount(db *sql.DB) ([]Account, error) {
	const selectRow = `
		SELECT * FROM ACCOUNT
		ORDER BY ID
	`

	rows, err := db.Query(selectRow)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var accounts []Account
	for rows.Next() {
		var account Account
		if err = rows.Scan(&account.ID, &account.Name, &account.MailAddress, &account.Lang); err != nil {
			return nil, err
		}

		accounts = append(accounts, account)

	}

	return accounts, nil
}

func main() {
	db, err := connectPostgres()
	if err != nil {
		return
	}
	defer db.Close()

	if _, err = db.Exec(createTable); err != nil {
		fmt.Println("1", err)
		return
	}

	err = insertAccount(db)
	if err != nil {
		fmt.Println("2", err)
		return
	}

	accounts, err := getAccount(db)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("accounts", accounts)

	if _, err = db.Exec("DROP TABLE ACCOUNT"); err != nil {
		fmt.Println("", err)
		return
	}
}
