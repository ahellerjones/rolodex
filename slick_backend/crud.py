from sqlalchemy.orm import Session
from . import models, schemas

# MODELS ARE HOW THE DATA IS ORGANIZED IN THE DBs
# SCHEMAS ARE THE DATA THAT COMES IN OR OUT 
# These functions transform data schemas into db models and place them into the dbs.
# Each crud operation function gets
# an instance of the db,
# Special parameters which are parsed from the href
# And the schema which the data comes in on. Ez. 
# And then the methods return the created db object.
# Which actually seems kind of fucked up but oh well. I like 'database schemas' but yolo

def create_user(db: Session, user: schemas.UserCreate): 
    db_user_obj = models.User(username=user.username, 
    hashed_password=user.password)
    db.add(db_user_obj)
    db.commit()
    db.refresh(db_user_obj) # I think this updates the user's ID etc. 
    return db_user_obj
    
def read_user(db: Session, user_id: int):
    return db.query(models.User).filter(models.User.id == user_id).first()

def read_user_by_username(db: Session, username: str):
    return db.query(models.User).filter(models.User.username == username).first()

def read_users(db: Session, skip: int = 0, limit: int = 100):
    return db.query(models.User).offset(skip).limit(limit).all()

def create_contact(db: Session, contact: schemas.ContactCreate, user_id: int): 
    db_contact_obj = models.Contact(**contact.dict(), owner_id=user_id) # Unfurl the contact into a dict and put it into a Contact model
    db.add(db_contact_obj)
    db.commit()
    db.refresh(db_contact_obj)
    return db_contact_obj

def read_contacts_for_user(db: Session, user_id: int): 
    # I think this is how we do this. 
    return db.query(models.Contact).filter(models.Contact.owner_id == user_id).all()

def read_contact_for_user_id(db: Session, user_id: int, contact_id: int): 
    return db.query(models.Contact).filter(models.Contact.id == contact_id, models.Contact.user_id == user_id) # I have no idea what im doing 

# def update_contact_for_user_id(db: Session, user_id: int, contact_id: int): 
