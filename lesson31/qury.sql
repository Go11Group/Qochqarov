create table persons(id uuid primary key,last_name varchar,first_name varchar,age int,courdse int, phone varchar )

create index id_index on persons(id)

create index last_nmame_index on persons(last_name)

create index age_index on persons(age)
