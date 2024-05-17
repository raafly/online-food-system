-- Active: 1715908777371@@127.0.0.1@5432@ordering
CREATE TABLE Customers(
    ID varchar(100) not null primary key,
    Username varchar(50) not null,
    Email varchar(100) not null unique,
    Password varchar(225) not null,
    Token varchar(100), 
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp
)