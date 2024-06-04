BEGIN;

create table kitob(id serial primary key, name varchar, author varchar);
insert into kitob(name,author) values('vatan','ali');
insert into kitob(name,author) values('ona','karim');
insert into kitob(name,author) values('bhayot','toshmat');
insert into kitob(name,author) values('ddmm','dmkf');

update kitob set name='Eshamt' where id=2;