from typing import Optional
from fastapi import FastAPI, HTTPException
from pydantic import BaseModel
from sql_interface.py import SQLiteHandler

class Login(BaseModel):
    username: str
    passw: str
    loginOrSignup: bool


db = SQLiteHandler("foo")
app = FastAPI()

@app.get("/")
async def root():
    return {"message": "Hello World"}



@app.post("/login")
async def post_login(login: Login):
    if (login.loginOrSignup): 
        # Trying to login
        userID = db.checkUsernamePassword(login.username, login.passw) 
        if userID is not None: 
            # Succesfully logged in 
            return {"UserID": f"{userID}"}
        else: 
            raise HTTPException(status_code=410, detail="Bad Username or password")
    else: 
        # Want to sign up 
        if not db.checkUsername(login.username): 
            raise HTTPException(status_code=410, detail="Username already exists")
        else: 
            userID = db.insertNewUser(login.username, login.passw)
            # Succesfully made new user
            return {"UserID": f"{userID}"}

class Identification(BaseModel): 
    UserID: int 
    Key: int

class Contact(BaseModel):
    identification: Identification
    name: str 
    address: str
    phoneNumber: str
    email: str 
    birthday: str 
@app.post("/contacts")
async def post_contacts(contact: Contact):
    identification = db.insertContact(contact)
    return {identification}
