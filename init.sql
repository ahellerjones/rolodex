-- Table schema creation
-- Dim Users table

CREATE TABLE users(
    userID AUTO_INCREMENT primary key,
    login_name varchar(64) unique not null,
    password varchar(256) not null
    )
-- Mod Contacts table

CREATE TABLE contacts(
    contactID AUTO_INCREMENT primary key,
    userID int, 
    name varchar(256),
    address varchar(256),
    phone_number varchar(64), 
    email varchar(256), 
    birthday varchar(256),
    FOREIGN KEY (userID) REFERENCES users(userID) ON DELETE CASCADE
)

-- update users table with SIGN-UP
-- check if user exists 
SELECT userID FROM users
WHERE userID = {userID}

-- should return success
INSERT INTO users VALUES(
    DEFAULT, login_name, password
)

-- check users for login_name and password
-- if none returned, not a user
SELECT userID FROM users
WHERE login_name = {arg_login_name} AND password = {arg_password}

-- delete account (users)
DELETE FROM users
WHERE login_name = {arg_login_name} AND password = {arg_password}


-- create new contact - default autoincrements, userID is from frontend, pass null for values not used
INSERT INTO contacts VALUES(
    DEFAULT, {arg_userID}, {arg_name}, {arg_address}, {arg_phone_number}, {arg_email}, {arg_birthday}
)

-- get all contacts
SELECT * FROM contacts
WHERE userID = {arg_userID}

-- update contact - pass in all old data, but replace new data 
UPDATE contacts
SET name = {arg_new}, address = {arg_new}, phone_number = {arg_new}, email = {arg_new}, birthday = {arg_new}
WHERE contactID = {arg_contactID}

-- delete contacts - based on contact ID
DELETE FROM contacts
WHERE contactID = {arg_contactID}

