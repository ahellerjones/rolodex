from utils import get_db, oauth2_scheme, pwd_context
# Here we use a fastAPI shortcut in the Annotated, we really are just creating a OAuth2Password.. object and putting it 
# into the form_data object. 
@app.post("/token") 
async def login(form_data: Annotated[OAuth2PasswordRequestForm, Depends()], db: Session = Depends(get_db)):
    db_user = crud.read_user_by_username(db, username=form_data.username)
    if not db_user:
        raise HTTPException(status_code=401, detail="Incorrect Username or Password")
    if not form_data.password == db_user.hashed_password: 
        raise HTTPException(status_code=401, detail="Incorrect Username or Password")
    return {"access_token": db_user.username, "token_type": "bearer"}
