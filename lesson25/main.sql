create table cars(id serial primary key, name varchar,probegi int ,year int);

create table users(id serial primary key, name varchar, cars_id int references cars(id));

create table user_cars(id serial primary key,user_id int references users(id),cars_id int references cars(id));


INSERT INTO cars (name, probegi, year, user_id) VALUES
('Toyota Camry', 50000, 2019),
('Honda Civic', 40000, 2018),
('Ford Mustang', 60000, 2020),
('Chevrolet Malibu', 45000, 2017),
('BMW 3 Series', 55000, 2019),
('Audi A4', 48000, 2018),
('Mercedes-Benz C-Class', 52000, 2020),
('Volkswagen Jetta', 42000, 2016),
('Nissan Altima', 49000, 2018),
('Hyundai Sonata', 47000, 2021,),
('Toyota Corolla', 51000, 2017),
('Honda Accord', 43000, 2020),
('Ford Fusion', 59000, 2016),
('Chevrolet Cruze', 46000, 2019),
('BMW 5 Series', 54000, 2018),
('Audi A6', 50000, 2017),
('Mercedes-Benz E-Class', 48000, 2021),
('Volkswagen Passat', 47000, 2019),
('Nissan Maxima', 52000, 2018),
('Hyundai Elantra', 44000, 2020,);

INSERT INTO users (name, cars_id) VALUES
('John'),
('Alice'),
('Bob'),
('Emily'),
('Michael'),
('Sarah'),
('David'),
('Emma'),
('Daniel'),
('Olivia');

INSERT INTO user_cars (user_id, cars_id) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 5),
(6, 6),
(7, 7),
(8, 8),
(9, 9),
(10, 10),
(1, 11),
(2, 12),
(3, 13),
(4, 14),
(5, 15),
(6, 16),
(7, 17),
(8, 18),
(9, 19),
(10, 20),
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 5),
(6, 6),
(7, 7),
(8, 8),
(9, 9),
(10, 10);




