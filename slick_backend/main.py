from fastapi import Depends, FastAPI, HTTPException
from fastapi.responses import HTMLResponse
from sqlalchemy.orm import Session
from fastapi.staticfiles import StaticFiles
import uvicorn

from . import crud, models, schemas
from .database import SessionLocal, engine

app = FastAPI()


models.Base.metadata.create_all(bind=engine)

#app.mount("/", StaticFiles(directory="sexy_frontend/dist", html=True), name="dist")

# Dependency
def get_db():
    db = SessionLocal()
    try:
        yield db
    finally:
        db.close()
# Create a new user by issuing a POST on /users
@app.post("/users")
async def post_users(user: schemas.UserCreate, db: Session = Depends(get_db)):
    print("Hello")
    db_user = crud.read_user_by_username(db, username=user.username)
    if db_user:
        raise HTTPException(status_code=400, detail="username already registered")
    return crud.create_user(db=db,user=user)

# Read a specific user using id with GET /users/{id}
@app.get("/users/{user_id}", response_model=schemas.User)
def read_user(user_id: int, db: Session = Depends(get_db)):
    db_user = crud.read_user(db, user_id=user_id)
    if db_user is None:
        print("poop")
        raise HTTPException(status_code=404, detail="User not found")
    return db_user

@app.post("/users/{user_id}/contacts/", response_model=schemas.Contact)
def post_contact(user_id: int, contact: schemas.ContactCreate, db: Session = Depends(get_db)):
    return crud.create_contact(db, contact=contact, user_id=user_id)

@app.get("/users/{user_id}/contacts/", response_model=schemas.Contact)
def read_contacts_for_user(user_id: int, db: Session = Depends(get_db)
):
    return crud.read_contacts_for_user(db, user_id)
    #return crud.create_contact(db=db, contact=contact, user_id=user_id)
