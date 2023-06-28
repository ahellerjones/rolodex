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
	stmt, err := handler.db.Query(`
	SELECT ? FROM users
	WHERE login_name = ?
	`)
	defer stmt.Close()
	if err != nil {
		return false, err
	}

	user := logininfo.user
	_, err = stmt.Exec(user)

	/*var username string
	for check.Next() {
		err := check.Scan(&username)
		if err != nil && err != sql.ErrNoRows {
			return false, err
		}
		return true, nil // Username exists
	}*/
	// if no username, then username does not exist
	return false, nil // username does not exist
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

// TODO: Do we even need this?
// I dont even think I have logic to perform this
func (handler *SQLiteHandler) DeleteUser(logininfo LoginInfo) (int, error) {
	stmt, err := handler.db.Prepare(`
	DELETE FROM users
	WHERE userID = {logininfo.userID}
	)`)
	if err != nil {
		return 0, err
	}
	_, err = stmt.Exec()
	if err != nil {
		return 0, err
	}
	// Do some shit like finding the userID that you just used or whatever
	return 0, err
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

	var maxID int
	qry, err = handler.db.QueryRow(`
	SELECT MAX(user_id) FROM contacts
	`
	)
	if err != nil {
		return 0, err
	}

	
	// fetch ID of last record inserted
	// make sure this is the contact ID and not the user ID
	//TODO: Find the last ID used
	// why
	return 0, nil
}

func (handler *SQLiteHandler) GetContacts(userID int) ([]Contact, error) {
	stmt, err := handler.db.Query(`
	SELECT * FROM contacts
	WHERE userID = {userID}
	ORDER BY name desc
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var key int
	var name string
	var address string
	var phoneNumber string
	var email string
	var birthday string
	contact_slice := []Contact{}
	// iterate through each and create a contact
	for stmt.Next() {
		err := stmt.Scan(&key, &name, &address, &phoneNumber, &email, &birthday)
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
	query, err := handler.db.Prepare(`
	SELECT contact_id FROM contacts
	WHERE contact_id = ?
	`)
	if err != nil {
		return -1, err
	}
	defer query.close()

	id_ret, err := query.Query(contact.ContactID)
	if err != nil {
		return -1, err
	}
	stmt, err := handler.db.Prepare(`
	DELETE FROM contacts 
	WHERE contact_id = ?
	`)
	if err != nil {
		return -1, err
	}
	_, err = stmt.Exec(contact.key)
	if err != nil {
		return -1, err
	}
	// TODO: Figure out the id that u deleted and return it
	return -1, nil

}
