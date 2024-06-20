create table if not exists person(
    id serial Primary key,
    name varchar,
    age int, 
    year int,
    password varchar
)