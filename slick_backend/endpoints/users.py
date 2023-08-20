import crud
import schemas
import utils
from fastapi import APIRouter, Depends, HTTPException
from sqlalchemy.orm import Session
from typing_extensions import Annotated

router = APIRouter()
# Create a new user by issuing a POST on /users
# @router.post("/users")
# async def post_users(user: schemas.UserCreate, db: Session = Depends(utils.get_db)):
#     print("Hello")
#     db_user = crud.read_user_by_username(db, username=user.username)
#     if db_user:
#         raise HTTPException(status_code=400, detail="username already registered")
#     return crud.create_user(db=db,user=user)


# Read a specific user using id with GET /users/{id}
# @router.get("/users/{user_id}", response_model=schemas.User)
# def read_user(user_id: int, db: Session = Depends(utils.get_db)):
#     db_user = crud.read_user(db, user_id=user_id)
#     if db_user is None:
#         print("poop")
#         raise HTTPException(status_code=404, detail="User not found")
#     return db_user

# @router.post("/users/{user_id}/contacts/", response_model=schemas.Contact)
# def post_contact(user_id: int, contact: schemas.ContactCreate, db: Session = Depends(utils.get_db)):
#     return crud.create_contact(db, contact=contact, user_id=user_id)

@router.get("/users/me")
async def read_users_me(current_user: Annotated[schemas.User, Depends(utils.get_current_user)]):
    return current_user
