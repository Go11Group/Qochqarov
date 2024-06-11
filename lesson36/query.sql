create table problems(id serial,name varchar,text varchar)

create table users(id serial primary key,first_name varchar,last_name varchar,age int,email varchar)

create table solved_problems(id serial primary key,name varchar, degre varchar,user_id int references users(id))