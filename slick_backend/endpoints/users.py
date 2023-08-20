from fastapi import Depends, HTTPException
import os, sys
# Get the parent directory
parent_dir = os.path.dirname(os.path.realpath(__file__))

# Add the parent directory to sys.path
sys.path.append(parent_dir)
from slick_backend.utils import get_db
import crud 
from sqlalchemy.orm import Session
import schemas
from fastapi import APIRouter
from typing_extensions import Annotated

router = APIRouter()
# Create a new user by issuing a POST on /users
@router.post("/users")
async def post_users(user: schemas.UserCreate, db: Session = Depends(get_db)):
    print("Hello")
    db_user = crud.read_user_by_username(db, username=user.username)
    if db_user:
        raise HTTPException(status_code=400, detail="username already registered")
    return crud.create_user(db=db,user=user)


# Read a specific user using id with GET /users/{id}
@router.get("/users/{user_id}", response_model=schemas.User)
def read_user(user_id: int, db: Session = Depends(get_db)):
    db_user = crud.read_user(db, user_id=user_id)
    if db_user is None:
        print("poop")
        raise HTTPException(status_code=404, detail="User not found")
    return db_user

@router.post("/users/{user_id}/contacts/", response_model=schemas.Contact)
def post_contact(user_id: int, contact: schemas.ContactCreate, db: Session = Depends(get_db)):
    return crud.create_contact(db, contact=contact, user_id=user_id)


