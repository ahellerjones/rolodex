-- Table schema creation
-- Dim Users table

CREATE TABLE users(
    userID int AUTO_INCREMENT primary key,
    login_name varchar(64) unique not null,
    password varchar(256) not null
    )
-- Mod Contacts table

CREATE TABLE contacts(
    contactID int primary key,
    userID int, 
    firstname varchar(256),
    lastname varchar(256),
    address varchar(256),
    phonenumber varchar(64), 
    email varchar(256), 
    birthday varchar(256),
    FOREIGN KEY contacts(userID) REFERENCES users(userID) ON DELETE CASCADE
)

-- update users table with SIGN-UP
-- should return success
INSERT INTO users VALUES(
    DEFAULT, login_name, password
)

-- check users for login_name and password
-- if none returned, not a user
SELECT userID FROM users
WHERE {login_data} = login_name AND {password} = password

-- create new contact
INSERT INTO contacts VALUES(
    
)