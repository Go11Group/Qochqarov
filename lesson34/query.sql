create table product(id uuid primary key not null default gen_random_uuid(),name varchar, price int, year int)

create table users(id uuid primary key default gen_random_uuid() not null,name varchar,age int,year int,product_id uuid references product(id))

INSERT INTO product (name, price, year) VALUES
    ('Product 1', 100, 2020),
    ('Product 2', 200, 2021),
    ('Product 3', 150, 2019),
    ('Product 4', 180, 2022),
    ('Product 5', 220, 2018),
    ('Product 6', 130, 2023),
    ('Product 7', 170, 2017),
    ('Product 8', 250, 2024),
    ('Product 9', 90, 2016),
    ('Product 10', 300, 2025),
    ('Product 11', 120, 2015),
    ('Product 12', 190, 2026),
    ('Product 13', 230, 2014),
    ('Product 14', 140, 2027),
    ('Product 15', 160, 2013),
    ('Product 16', 270, 2028),
    ('Product 17', 110, 2012),
    ('Product 18', 280, 2029),
    ('Product 19', 240, 2011),
    ('Product 20', 200, 2030);

INSERT INTO users (name, age, year, product_id) VALUES
    ('User 1', 25, 1997, '1a9d2598-4f09-4c7e-9c7f-35474add5a12'),
    ('User 2', 30, 1992, 'db2635ca-cf0f-4642-b310-3780727b52c0'),
    ('User 3', 35, 1987, '162eedd3-4895-4269-9c3a-dddca19366ad'),
    ('User 4', 40, 1982, '4d82dd44-1a4f-4b31-ab4f-79e8f6c6f0aa'),
    ('User 5', 22, 2000, '59d8d5a3-d2ae-49ef-be5f-5920970d4442'),
    ('User 6', 28, 1994, '9db94ac4-a184-4dca-a1d3-00bcb36d7f58'),
    ('User 7', 32, 1990, '0c43fbee-1937-425d-86aa-5296ad1f65f2'),
    ('User 8', 37, 1985, 'b868a51c-f52c-44f7-b148-46c87e1fa293'),
    ('User 9', 27, 1995, '7d81a963-bf68-4ac0-a1ed-0168d602909b'),
    ('User 10', 33, 1989, '387f8a78-f9f2-4a9c-8b9a-86c5ad1e4984'),
    ('User 11', 29, 1993, '507bb28a-1b26-4373-9bba-f119a957b01f'),
    ('User 12', 34, 1988, '48308b15-15c7-4ced-9d4d-7ddb60cf2437'),
    ('User 13', 26, 1996, '40bddaaa-4c3b-423b-ad20-cc4d487e36b6'),
    ('User 14', 31, 1991, '9028c30e-f865-4bc4-99c7-d81e8e721fc5'),
    ('User 15', 36, 1986, 'f19c377a-0b7e-4fde-9271-6b1499817634'),
    ('User 16', 23, 1999, '828b4ce6-b88f-495d-a0e7-40c16e29c36c'),
    ('User 17', 38, 1984, '70319e5e-e911-465c-83b2-0744ab8f3e47'),
    ('User 18', 24, 1998, 'ddf665dc-12a8-4eeb-9255-0f4c1526e408'),
    ('User 19', 39, 1983, 'daca3f36-a13b-4613-9d73-07c061d53493'),
    ('User 20', 21, 2001, '39b89165-e4a9-4b3b-bb03-dcad54f00ec2');
