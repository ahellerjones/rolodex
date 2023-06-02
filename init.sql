-- Table schema creation
-- Dim Users table

CREATE TABLE users(
    userID int AUTO_INCREMENT primary key,
    login_name varchar(64) unique not null,
    password varchar(256) not null
    )
-- Mod Contacts table

CREATE TABLE contacts(
    contactID int AUTO_INCREMENT primary key,
    userID int, 
    name varchar(256),
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

-- create new contact - default autoincrements, userID is from frontend, pass null for values not used
INSERT INTO contacts VALUES(
    DEFAULT, {userID}, {name}, {address}, {phonenumber}, {email}, {birthday}
)

-- update contact - pass in all old data, but replace new data 
UPDATE contacts
SET name = {new}, address = {new}, phonenumber = {new}, email = {new}, birthday = {new}
WHERE contactID = {contactID}

-- get all contacts
SELECT * FROM contacts
WHERE userID = {userID}