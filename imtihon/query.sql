create table if not exists
users(id uuid primary key not null default gen_random_uuid(),
name varchar(100),email varchar(100) unique,
birthday timestamp,password varchar unique,created_at timestamp default CURRENT_TIMESTAMP,
update_at timestamp default CURRENT_TIMESTAMP, delete_at bigint default 0 )


create table 
Courses(id uuid primary key not null default gen_random_uuid(),
title varchar(100),description text,created_at timestamp default CURRENT_TIMESTAMP,
update_at timestamp default CURRENT_TIMESTAMP, delete_at bigint default 0 )


create table 
Lessons(id uuid primary key not null default gen_random_uuid(),
course_id uuid references Courses(id),title varchar(100),
content text,created_at timestamp default CURRENT_TIMESTAMP,
update_at timestamp default CURRENT_TIMESTAMP, delete_at bigint default 0 )


create table Enrollments(id uuid primary key not null default gen_random_uuid(),
user_id uuid references users(id),course_id uuid references Courses(id),
enrollment_date timestamp,created_at timestamp default CURRENT_TIMESTAMP,
update_at timestamp default CURRENT_TIMESTAMP, delete_at bigint default 0 )


INSERT INTO users (name, email, birthday, password)
VALUES
    ('John Doe', 'john@example.com', '1990-01-01', 'password1'),
    ('Jane Smith', 'jane@example.com', '1985-03-15', 'password2'),
    ('Alice Johnson', 'alice@example.com', '1982-07-20', 'password3'),
    ('Michael Brown', 'michael@example.com', '1978-09-25', 'password4'),
    ('Emily Davis', 'emily@example.com', '1995-12-10', 'password5'),
    ('Daniel Wilson', 'daniel@example.com', '1987-05-18', 'password6'),
    ('Sarah Martinez', 'sarah@example.com', '1983-02-28', 'password7'),
    ('Christopher Taylor', 'chris@example.com', '1992-08-05', 'password8'),
    ('Amanda Anderson', 'amanda@example.com', '1979-11-14', 'password9'),
    ('Matthew Thomas', 'matthew@example.com', '1993-04-22', 'password10'),
    ('Jennifer Hernandez', 'jennifer@example.com', '1981-06-09', 'password11'),
    ('Andrew Young', 'andrew@example.com', '1988-10-17', 'password12'),
    ('Jessica Lee', 'jessica@example.com', '1991-03-30', 'password13'),
    ('James Clark', 'james@example.com', '1977-07-07', 'password14'),
    ('Megan Walker', 'megan@example.com', '1984-09-02', 'password15'),
    ('Justin Lewis', 'justin@example.com', '1996-01-11', 'password16'),
    ('Lauren King', 'lauren@example.com', '1980-04-26', 'password17'),
    ('Brandon Hill', 'brandon@example.com', '1989-08-13', 'password18'),
    ('Ashley Scott', 'ashley@example.com', '1994-12-05', 'password19'),
    ('Kate Brown', 'kate@example.com', '1998-11-30', 'password20');



INSERT INTO Courses (title, description)
VALUES
    ('Course 1', 'Description of course 1'),
    ('Course 2', 'Description of course 2'),
    ('Course 3', 'Description of course 3'),
    ('Course 4', 'Description of course 4'),
    ('Course 5', 'Description of course 5'),
    ('Course 6', 'Description of course 6'),
    ('Course 7', 'Description of course 7'),
    ('Course 8', 'Description of course 8'),
    ('Course 9', 'Description of course 9'),
    ('Course 10', 'Description of course 10'),
    ('Course 11', 'Description of course 11'),
    ('Course 12', 'Description of course 12'),
    ('Course 13', 'Description of course 13'),
    ('Course 14', 'Description of course 14'),
    ('Course 15', 'Description of course 15'),
    ('Course 16', 'Description of course 16'),
    ('Course 17', 'Description of course 17'),
    ('Course 18', 'Description of course 18'),
    ('Course 19', 'Description of course 19'),
    ('Course 20', 'Description of course 20');




INSERT INTO Lessons (course_id, title, content)
VALUES
    ('1705ef15-416d-4efb-ae3c-e9f7cb655867', 'Lesson 1', 'Content of Lesson 1'),
    ('c2a733d9-2c0a-4af2-b934-501873c4d55d', 'Lesson 2', 'Content of Lesson 2'),
    ('23ef1c7a-19c0-4d64-a40a-9c48cc385405', 'Lesson 3', 'Content of Lesson 3'),
    ('cde80c00-e3cf-40ad-a307-39efb73be4dd', 'Lesson 4', 'Content of Lesson 4'),
    ('4ad8a916-c789-4a6a-ad52-03f7a1c8a37a', 'Lesson 5', 'Content of Lesson 5'),
    ('bcc26c9f-44db-447d-9eaf-b9895174255c', 'Lesson 6', 'Content of Lesson 6'),
    ('721e7534-a19f-45db-8b13-b13d5b344c06', 'Lesson 7', 'Content of Lesson 7'),
    ('81e8e28e-e5fa-4775-9c79-4905cc04b2a9', 'Lesson 8', 'Content of Lesson 8'),
    ('a69379fc-f9bb-4c25-842e-2f403ff3c7c7', 'Lesson 9', 'Content of Lesson 9'),
    ('b60778a3-9079-42f4-831e-8843ecfe3d25', 'Lesson 10', 'Content of Lesson 10'),
    ('afd4a1df-e775-41ff-b033-ef5d55ce615c', 'Lesson 11', 'Content of Lesson 11'),
    ('51693a92-b06e-4de8-9c4d-bae97263990e', 'Lesson 12', 'Content of Lesson 12'),
    ('c17f4b3a-1329-4dfa-bdaa-472a1ee77173', 'Lesson 13', 'Content of Lesson 13'),
    ('8b8d1f46-9e3e-4f01-a8fa-ad7cf3bf27e3', 'Lesson 14', 'Content of Lesson 14'),
    ('8872f7aa-cd55-48ad-8df7-5cdaa5cc1ff3', 'Lesson 15', 'Content of Lesson 15'),
    ('0993174e-ec0f-4696-9664-7119cb88ae4b', 'Lesson 16', 'Content of Lesson 16'),
    ('3ff22a3e-e6ba-41db-9212-d558be8a617a', 'Lesson 17', 'Content of Lesson 17'),
    ('947f5359-c579-4ed8-882f-a97203dcda85', 'Lesson 18', 'Content of Lesson 18'),
    ('63beaf7f-3c6a-44bf-a3bf-5134f89d8838', 'Lesson 19', 'Content of Lesson 19'),
    ('692cfa6d-aae5-45a9-b7f8-843354861ed0', 'Lesson 20', 'Content of Lesson 20');


INSERT INTO Enrollments (user_id, course_id, enrollment_date)
VALUES
('b290de5b-9aea-4c10-b810-a2b521639244', 'cc6101f3-237b-4093-8c8c-5f090a3e881b', '2024-06-12 14:11:29.585241'),
('4204c10e-76d3-412f-a462-5ae4a06e8a9b', '568eb178-5537-4192-864f-e3f6c33058ba', '2024-06-12 14:11:29.585241'),
('96b29c9c-ea05-499a-b232-aeb9d8ac0633', '3ffd2046-65da-4584-b3ca-3262cf217529', '2024-06-12 14:11:29.585241'),
('1acaa9b9-56de-4a33-8450-603e00731ec9', '77da741d-3b78-4909-9911-5eb94f928c13', '2024-06-12 14:11:29.585241'),
('2e9b9d4d-09ca-4c1f-9009-dc79f4f77e21', '74d43e00-8e10-4e86-884a-1bef4a84bab6', '2024-06-12 14:11:29.585241'),
('86708cf6-841b-4e5c-99f7-298091479179', '080cfab0-0d4e-4c37-8433-fee0af3d3316', '2024-06-12 14:11:29.585241'),
('13af3a45-529b-4811-8ff9-05545bd350d8', '5aa22b28-e860-4484-977e-fcc95b67fbbb', '2024-06-12 14:11:29.585241'),
('921f1256-7942-4def-859a-2da9dc8ca362', '142b08e4-0f65-4585-85ad-46d8ae08a76f', '2024-06-12 14:11:29.585241'),
('30de93ba-ee17-4c49-9138-587de18e27d7', 'a4835bf5-bd22-46d2-9559-ddb0b28c32c3', '2024-06-12 14:11:29.585241'),
('e1328cdb-ee97-4df0-b725-40607ef5993d', 'c36c57c2-5634-4d80-ae67-0f3489ebeb0d', '2024-06-12 14:11:29.585241'),
('65cb9ca6-7456-4593-8422-995e51f197b0', '05ec3e6c-e016-485c-8ba1-4b3b835f5cba', '2024-06-12 14:11:29.585241'),
('860fe264-0030-4457-b487-5649a2e71ccf', 'beee46ef-05fd-4724-8f8f-b014d50a41a4', '2024-06-12 14:11:29.585241'),
('09c14eae-60b7-4732-bf25-e5a97c7be2ad', '34529a95-d129-4785-b720-71d982cd2ec9', '2024-06-12 14:11:29.585241'),
('a82bf455-53f3-4a69-9276-4a241e87b2ad', '34bdf8d0-c820-4b6b-8cc2-98d366ca9894', '2024-06-12 14:11:29.585241'),
('ec901c73-065a-4008-972c-a7898d3b9383', '0cb15dce-8caa-4553-af98-c8ea1bdeabf6', '2024-06-12 14:11:29.585241'),
('404ebbd5-58a1-42a4-beb8-11b838de4ede', 'ff3c4974-54bc-46ca-8227-911a6353d165', '2024-06-12 14:11:29.585241'),
('2f7904b3-dbc0-4641-853a-d414d50ef06f', 'bb30b390-3e91-4d54-a7cb-1f51e668d009', '2024-06-12 14:11:29.585241');