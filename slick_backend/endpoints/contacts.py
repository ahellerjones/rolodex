from fastapi import Depends, HTTPException
import os, sys
from slick_backend.utils import get_db, oauth2_scheme
import crud 
from sqlalchemy.orm import Session
import schemas
from fastapi import APIRouter
from typing_extensions import Annotated

router = APIRouter()
@router.post("/users/{user_id}/contacts/", response_model=schemas.Contact)
def post_contact(user_id: int, contact: schemas.ContactCreate, db: Session = Depends(get_db)):
    return crud.create_contact(db, contact=contact, user_id=user_id)



# Get a list of contacts for a specific user
@router.get("/users/{user_id}/contacts/", response_model=list[schemas.Contact])
def read_contacts_for_user(token: Annotated[str, Depends(oauth2_scheme)], user_id: int, db: Session = Depends(get_db)): 
    contacts = crud.read_contacts_for_user(db, user_id)
    if contacts.count() == 0: 
        return []
    else: 
        return contacts.all()