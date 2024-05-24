create table author(id serial primary key,name varchar);

insert into author(name) values('ali'),('vali'),('karim'),('eshmat'),('toshmat'),('donyor'),('elyor');

create table books(id serial primary key,page int,author_id int, foreign key (author_id) references author(id))

insert into books(name,page,author_id) values('vatan',4233,7);

insert into books(name,page,author_id) values('vatan',4233,7),('hayot',43,6),('ona',659,5);

insert into books(name,page,author_id) values('vatan',4233,7),('hayot',43,6),('ona',659,5),('hayot',543,4),('fahod va shirin',443,3),('mehr',443,2),('qasos',433,6);

insert into books(name,page,author_id) values('ona',659,5),('hayot',543,4),('fahod va shirin',443,3),('mehr',443,2),('qasos',433,6);

select * from books

select b.id, b.name,b.page,a.name from books as b join author as a on a.id=b.author_id;