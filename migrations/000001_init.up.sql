CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    name varchar(50),
    surname varchar(50),
    patronymic varchar(50),
    age int,
    gender varchar(50),
    nationality varchar(50)
);
