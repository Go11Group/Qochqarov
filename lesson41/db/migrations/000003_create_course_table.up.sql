create table if not exists course(
    id serial,
    person_id int references person(id),
    name varchar, 
    price int
)
