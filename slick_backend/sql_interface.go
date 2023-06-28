package main

//TODO: Need to fix all of the silly f strings, go doesn't support that but does have variations of it
import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type SQLiteHandler struct {
	db *sql.DB
}

func NewSQLiteHandler(dbPath string) (*SQLiteHandler, error) {
	// pass in "./rolodex.db"
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	// Do we need to initializeTables() here?
	return &SQLiteHandler{db: db}, nil
}

func (handler *SQLiteHandler) initializeTables() error {
	users, err := handler.db.Prepare(`
		CREATE TABLE IF NOT EXISTS 
		users(
			user_id AUTO_INCREMENT primary key,
			login_name varchar(64) unique not null,
			password varchar(256) not null
		)`)
	if err != nil {
		return err
	}
	_, err = users.Exec()

	contacts, err := handler.db.Prepare(`
		CREATE TABLE IF NOT EXISTS 
		contacts(
			contact_id AUTO_INCREMENT primary key,
			user_id int, 
			name varchar(256),
			address varchar(256),
			phone_number varchar(64), 
			email varchar(256), 
			birthday varchar(256),
			FOREIGN KEY (userID) REFERENCES users(userID) ON DELETE CASCADE
		)`)
	if err != nil {
		return err
	}
	_, err = contacts.Exec()

	return err
}

func (handler *SQLiteHandler) CheckUsernameExists(logininfo LoginInfo) (bool, error) {
	// checks for only username match
	stmt, err := handler.db.Prepare(`
	SELECT login_name FROM users
	WHERE login_name = ?
	`)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	var username string
	_, err = stmt.Query(logininfo.Username).Scan(&username)
	if err != nil {
		return false, err
	}

	return &username, nil // username exists
}

func 

func (handler *SQLiteHandler) InsertUser(logininfo LoginInfo) (int, error) {
	stmt, err := handler.db.Prepare(`
	INSERT INTO users VALUES(
		DEFAULT, {logininfo.user}, {logininfo.password}
	)`)
	if err != nil {
		return 0, err
	}
	_, err = stmt.Exec() // I have no idea what this is supposed to do
	if err != nil {
		return 0, err
	}
	// TODO: Need to then perform a query to figure out what the userID was for the last operation

	return 0, nil
}

func (handler *SQLiteHandler) InsertContact(contact Contact) (int, error) {
	stmt, err := handler.db.Prepare(`
	INSERT INTO contacts VALUES(
		DEFAULT, ?, ?, ?, ?, ?, ?
	)
	`)
	if err != nil { // TODO: This probably gives SQLnoRows, just watch out
		return 0, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(contact.ContactId.UserId, contact.Name, contact.Address,
	contat.PhoneNumber, contact.Email, contact.Birthday)
	if err != nil {
		return 0, err
	}

	var lastID int
	qry, err = handler.db.QueryRow(`
	SELECT MAX(contact_id) FROM contacts
	`
	).Scan(&lastId)
	if err != nil {
		return 0, err
	}

	return &lastID, nil
}

func (handler *SQLiteHandler) GetContacts(userID int) ([]Contact, error) {
	stmt, err := handler.db.Prepare(`
	SELECT * FROM contacts
	WHERE userID = ?
	ORDER BY name desc
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	qry, err := handler.db.Query(userID)

	var key int
	var name string
	var address string
	var phoneNumber string
	var email string
	var birthday string
	contact_slice := []Contact{}
	// iterate through each and create a contact
	for qry.Next() {
		err := qry.Scan(&key, &name, &address, &phoneNumber, &email, &birthday)
		if err != nil {
			return nil, err
		}
		// return userID
		// Fill return struct with the values.
		contact_slice = append(contact_slice, Contact{
			ContactID: Identification{
				UserID: userID,
				Key:    key,
			},
			Name:        name,
			Address:     address,
			PhoneNumber: phoneNumber,
			Email:       email,
			Birthday:    birthday,
		})
	}
	return contact_slice, nil
}

func (handler *SQLiteHandler) DeleteContact(contact Contact) (int, error) {
	stmt, err := handler.db.Prepare(`
	DELETE FROM contacts 
	WHERE contact_id = ?
	`)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	_, err := stmt.Exec(contact.ContactID.Key)
	if err != nil {
		return -1, err
	}
	defer _.Close()

	// Confirm ID was deleted
	prep_empty, err := handler.db.Prepare(`
	SELECT user_id FROM contacts WHERE user_id = ?
	`)
	if err != nil {
		return -1, err
	}
	defer prep_empty.Close()

	empty, err := handler.db.QueryRow(contact.ContactID.Key)
	if err != nil {
		if err == sql.ErrNoRows{
			return contact.ContactID.Key, nil
		} else {return -1, err}
	}

	// can just return given contact ID since error would be thrown if there was a problem
}
