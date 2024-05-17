CREATE TABLE orders(
    ID bigint not null primary key,
    Customer varchar(100) not null,
    Item varchar(100),
    Price int,
    Quantity int not null,
    CreatedAt timestamp default current_timestamp,
    UpdatedAt timestamp default current_timestamp
)