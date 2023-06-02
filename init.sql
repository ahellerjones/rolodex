-- Table schema creation
-- Dim Users table

-- userID (PK)
-- firstname
-- lastname
-- login_name
-- password
-- email

CREATE TABLE users(
    userID int primary key,
    firstname varchar(256),
    lastname varchar(256), 
    login_name varchar(64) unique not null,
    password varchar(256) not null,
    email varchar(256) unique not null
)

-- Mod Contacts table

-- contactID (PK)
-- userID (FK, on delete cascade)
-- firstname
-- lastname
-- address
-- phone number
-- email
-- birthday

CREATE TABLE contacts(
    contactID int primary key,
    userID int foreign key ON DELETE cascade,
    firstname varchar(256),
    lastname varchar(256),
    address varchar(256),
    phonenumber int, 
    email varchar(256), 
    birthday varchar(256)
)