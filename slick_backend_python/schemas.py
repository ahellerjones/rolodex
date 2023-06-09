from pydantic import BaseModel

# First we create a BaseModel class 
# So that we can access the attributes of the model. 
# These are our database schemas and dictates how 
# our API will work. 

# We start off with the BaseModel and create a Base class
# this is data that every model will inherit. 
class ContactBase(BaseModel):
    name: str 
    address: str
    phoneNumber: str
    email: str 
    birthday: str 

# Next we define what data a Contact needs to know in order 
# to properly create it. Thankfully, everything within 
# The base class contains all the data needed to make a Contact.
class ContactCreate(ContactBase): 
    pass 

# Next we need to define what kind of data will come back 
# When we read. We want to include the id of contact 
# and the owner's id of the contact when we return. 
class Contact(ContactBase):
    id: int
    owner_id: int

    class Config:
        orm_mode = True

# It is left as an exercise to the reader to figure out 
# Why the members are used as they are within the following 
# User classes. 
class UserBase(BaseModel):
    username: str

# Note, when we create, we need to supply the password
class UserCreate(UserBase):
    password: str


# However, when we return a user from the API we don't include 
# the pword, only the user ID and the contacts list. 
class User(UserBase):
    id: int
    contacts: list[Contact] = []

    class Config:
        # This tells Pydantic model to read the data
        # NOT as a dictionary, but as an ORM model 
        # (basically just a class) 
        orm_mode = True