package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteHandler struct {
	db *sql.DB
}

func main() {
	//db = dbConnection("./rolodex.db")
	}	

	func dbConnection(dbPath string) (*SQLiteHandler, error) {
		// pass in "./rolodex.db"
		db, err := sql.Open("sqlite3", dbPath)
		if err != nil {
			return nil, err
		}
	
		return &SQLiteHandler{db: db}, nil
	}
	
	func (handler *SQLiteHandler) initializeTables() error {
		users, err := handler.db.Prepare("
			CREATE TABLE IF NOT EXISTS 
			users(
				userID AUTO_INCREMENT primary key,
				login_name varchar(64) unique not null,
				password varchar(256) not null
			)")
		if err != nil {
			return err
		}
		_, err = users.Exec()

		contacts, err := handler.db.Prepare("
			CREATE TABLE IF NOT EXISTS 
			contacts(
				contactID AUTO_INCREMENT primary key,
				userID int, 
				name varchar(256),
				address varchar(256),
				phone_number varchar(64), 
				email varchar(256), 
				birthday varchar(256),
				FOREIGN KEY (userID) REFERENCES users(userID) ON DELETE CASCADE
			)")
		if err != nil {
			return err
		}
		_, err = contacts.Exec()

		return err
	}

	func (handler, *SQLiteHandler) checkUserExists(logininfo LoginInfo) error { 
		// checks username and password for existence
		check, err := handler.db.Query(f"
		SELECT userID FROM users
 		WHERE login_name = {logininfo.user} AND password = {logininfo.password}
		")
		defer check.Close()

		var userID int
		for check.Next() {
			err := rows.Scan(&userID)
			if err != nil {
				return err
			}
			return userID 
		}
		// if no user and password match, user does not exist
		return nil
	}

	func (handler, *SQLiteHandler) checkUsernameExists(logininfo LoginInfo) error{
		// checks for only username match
		check, err := handler.db.Query(f"
		SELECT {logininfo.user} FROM users
 		WHERE login_name = {logininfo.user}
		")
		defer check.Close()

		var username string
		for check.Next() {
			err := rows.Scan(&username)
			if err != nil {
				return err
			}
			return username 
		}
		// if no username, then username does not exist
		return nil
	}

	func (handler, *SQLiteHandler) insertUser(logininfo LoginInfo) error{
		insert, err := handler.db.Prepare("
		INSERT INTO users VALUES(
			DEFAULT, {logininfo.user}, {logininfo.password}
		)")
		if err != nil {
			return err
		}
		_, err = contacts.Exec()
		if err != nil {
			return err
		}
		id, err = _.LastInsertID()
		if err != nil {
			return err
		}

		return id
	}

	func (handler, *SQLiteHandler) deleteUser(logininfo LoginInfo) error{
		delete, err := handler.db.Prepare(F"
		DELETE FROM users
		WHERE userID = {logininfo.userID}
		)")
		if err != nil {
			return err
		}
		_, err = update.Exec()
		if err != nil {
			return err
		}
		id, err = _.LastInsertID()
		if err != nil {
			return err
		}
		return id
	}



	





	}

