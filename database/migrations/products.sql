CREATE TABLE Products (
    ID varchar(100) not null primary key,
    Name varchar(100) not null,
    Description varchar(100),
    Price int not null,
    Quantity int not null,
    CreatedAt timestamp default current_timestamp,
    UpdateAt timestamp default current_timestamp
)