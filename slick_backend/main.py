from fastapi import Depends, FastAPI, HTTPException, status
from fastapi.responses import HTMLResponse
from sqlalchemy.orm import Session
from fastapi.staticfiles import StaticFiles
import uvicorn
from fastapi.security import OAuth2PasswordBearer, OAuth2PasswordRequestForm
import os, sys
file_dir = os.path.dirname(__file__)
sys.path.append(file_dir)
from utils import get_db, oauth2_scheme, pwd_context
from slick_backend.endpoints import users, contacts, login
from database import SessionLocal, engine
import crud, models, schemas 


app = FastAPI()
app.include_router(users.router)
app.include_router(contacts.router)
app.include_router(login.router)

#app.mount("/", StaticFiles(directory="sexy_frontend/dist", html=True), name="dist")


