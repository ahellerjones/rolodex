from sqlalchemy.orm import Session
from . import models, schemas

# MODELS ARE WHAT GOES INTO THE DBs
# SCHEMAS ARE THE DATA THAT COMES IN OR OUT 
def get_user(db: Session, user_id: int):
    return db.query(models.User).filter(models.User.id == user_id).first()


def get_user_by_username(db: Session, username: str):
    return db.query(models.User).filter(models.User.username == username).first()

def get_users(db: Session, skip: int = 0, limit: int = 100):
    return db.query(models.User).offset(skip).limit(limit).all()

def create_user(db: Session, user: schemas.UserCreate): 
    db_user_obj = models.User(username=user.username, 
    hashed_password=user.password)
    db.add(db_user_obj)
    db.commit()
    db.refresh(db_user_obj) # I think this updates the user's ID etc. 

def get_contacts_for_user(db: Session, user_id: int): 
    # I think this is how we do this. 
    return db.query(models.Contact).filter(models.Contact.owner.id == user_id)
