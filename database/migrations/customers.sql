CREATE TABLE Customers(
    ID varchar(100) not null primary key,
    Name varchar(50) not null,
    Email varchar(100) not null unique,
    Password varchcar(225) not null,
    Token varchar(100), 
    CreatedAt timestamp default current_timestamp,
    UpdatedAt timestamp default current_timestamp
)