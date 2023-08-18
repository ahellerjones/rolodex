from sqlalchemy import create_engine, ForeignKey, Column, Integer, String, ForeignKeyConstraint

engine = create_engine('sqlite:///sql_app.db', echo = True)
from sqlalchemy.ext.declarative import declarative_base
Base = declarative_base()
from sqlalchemy.orm import relationship


# These are our data models which dictate how 
# data will be organized in our databases.
class User(Base):
    # Tells SQLAlchemy the table to use
    __tablename__ = "users"
    id = Column(Integer, primary_key=True, index=True)
    username = Column(String, unique=True)
    hashed_password = Column(String)
    #contacts = relationship("Contact", back_populates="user")
    # This is what relates a user to contacts,
    # Each user will contain a list of contacts


class Contact (Base): 
    __tablename__ = "contacts"
    id = Column(Integer, primary_key=True, index=True)
    name = Column(String,index=True) # What does index do?
    address = Column(String,index=True)
    phoneNumber = Column(String,index=True)
    email = Column(String,index=True)
    birthday = Column(String,index=True)
    user_id = Column(Integer,ForeignKey('users.id') )
    #user = relationship('User', back_populates='contacts')



