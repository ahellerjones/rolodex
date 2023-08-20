from fastapi import Depends, FastAPI, HTTPException, status
from fastapi.responses import HTMLResponse
from sqlalchemy.orm import Session
from fastapi.staticfiles import StaticFiles
import uvicorn
from fastapi.security import OAuth2PasswordBearer, OAuth2PasswordRequestForm
from typing_extensions import Annotated
import os, sys
file_dir = os.path.dirname(__file__)
sys.path.append(file_dir)
from utils import get_db, oauth2_scheme, pwd_context
from slick_backend.endpoints import users, contacts
from database import SessionLocal, engine
import crud, models, schemas 


app = FastAPI()
app.include_router(users.router)
app.include_router(contacts.router)

#app.mount("/", StaticFiles(directory="sexy_frontend/dist", html=True), name="dist")


# # Create a new user by issuing a POST on /users
# @app.post("/users")
# async def post_users(user: schemas.UserCreate, db: Session = Depends(get_db)):
#     print("Hello")
#     db_user = crud.read_user_by_username(db, username=user.username)
#     if db_user:
#         raise HTTPException(status_code=400, detail="username already registered")
#     return crud.create_user(db=db,user=user)

# Login
@app.post("/login")
async def post_users(user: schemas.UserCreate, db: Session = Depends(get_db)):
    db_user = crud.read_user_by_username(db, username=user.username)
    if db_user:
        if user.password == db_user.hashed_password: 
            #TODO This seems like the place where we would demarcate whether a user is active or not
            # Which TBH is kind of stupid and they don't do anything with it in the docs, so I'm not going to worry about it.  
            return {"status": "OK"}
        raise HTTPException(status_code=401, detail="Incorrect Password")
    raise HTTPException(status_code=401, detail="Username not found")

# # Read a specific user using id with GET /users/{id}
# @app.get("/users/{user_id}", response_model=schemas.User)
# def read_user(user_id: int, db: Session = Depends(get_db)):
#     db_user = crud.read_user(db, user_id=user_id)
#     if db_user is None:
#         print("poop")
#         raise HTTPException(status_code=404, detail="User not found")
#     return db_user

# @app.post("/users/{user_id}/contacts/", response_model=schemas.Contact)
# def post_contact(user_id: int, contact: schemas.ContactCreate, db: Session = Depends(get_db)):
#     return crud.create_contact(db, contact=contact, user_id=user_id)



# # Get a list of contacts for a specific user
# @app.get("/users/{user_id}/contacts/", response_model=list[schemas.Contact])
# def read_contacts_for_user(token: Annotated[str, Depends(oauth2_scheme)], user_id: int, db: Session = Depends(get_db)): 
#     contacts = crud.read_contacts_for_user(db, user_id)
#     if contacts.count() == 0: 
#         return []
#     else: 
#         return contacts.all()

# A method for decoding a token into a user.
# I think in the future we're going to want to query the db. 
def fake_decode_token(token): 
    # We need a way to get the actual user using this given token. 
    # What we can do is simply create a map on every login which associates a token with 
    # Each logged in user, the next section I believe tells how to actually handle the tokens. 
    return schemas.User(username=token+"fakedecoded", id=5)

async def get_current_user(token: Annotated[str, Depends(oauth2_scheme)]):
    user = fake_decode_token(token)
    if not user:
            raise HTTPException(
                status_code=status.HTTP_401_UNAUTHORIZED,
                detail="Invalid authentication credentials",
                headers={"WWW-Authenticate": "Bearer"},
            )
    return user

# Note this doesn't currently actually work because we have no system of matching tokens to users... yet 
@app.get("/users/me")
async def read_users_me(current_user: Annotated[schemas.User, Depends(get_current_user)]):
    return current_user

def verify_password(plain_password, hashed_password):
    return pwd_context.verify(plain_password, hashed_password)


def get_password_hash(password):
    return pwd_context.hash(password)
