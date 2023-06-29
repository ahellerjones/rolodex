from sqlalchemy import Boolean, Column, ForeignKey, Integer, String
from sqlalchemy.orm import relationship

from .database import Base

# These are our data models which dictate how 
# data will be organized in our databases.
class User(Base):
    # Tells SQLAlchemy the table to use
    __tablename__ = "users"
    id = Column(Integer, primary_key=True, index=True)
    username = Column(String, unique=True, index=True)
    hashed_password = Column(String)
    # This is what relates a user to contacts,
    # Each user will contain a list of contacts
    contacts = relationship("Contact", back_populates="owner")


class Contact (Base): 
    __tablename__ = "contacts"
    id = Column(Integer, primary_key=True, index=True)
    name = Column(String,index=True) # What does index do?
    address = Column(String,index=True)
    phoneNumber = Column(String,index=True)
    email = Column(String,index=True)
    birthday = Column(String,index=True)
    owner = relationship("User", back_populates="contacts")




