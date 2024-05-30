create table student(id uuid primary key not null default gen_random_uuid (), name varchar,age int)

create table curse(id uuid primary key not null default gen_random_uuid (), name varchar,student_id uuid referances student(id))