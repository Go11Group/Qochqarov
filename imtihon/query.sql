create table if not exists
users(id uuid primary key not null default gen_random_uuid(),
name varchar(100),email varchar(100) unique,
birthday timestamp,password varchar unique,created_at timestamp default CURRENT_TIMESTAMP,
update_at timestamp default CURRENT_TIMESTAMP, delete_at timestamp )


create table 
Courses(id uuid primary key not null default gen_random_uuid(),
title varchar(100),description text,created_at timestamp default CURRENT_TIMESTAMP,
update_at timestamp default CURRENT_TIMESTAMP, delete_at timestamp )


create table 
Lessons(id uuid primary key not null default gen_random_uuid(),
course_id uuid references Courses(id),title varchar(100),
content text,created_at timestamp default CURRENT_TIMESTAMP,
update_at timestamp default CURRENT_TIMESTAMP, delete_at timestamp)


create table Enrollments(id uuid primary key not null default gen_random_uuid(),
user_id uuid references users(id),course_id uuid references Courses(id),
enrollment_date timestamp,created_at timestamp default CURRENT_TIMESTAMP,
update_at timestamp default CURRENT_TIMESTAMP, delete_at timestamp)





INSERT INTO users (name, email, birthday, password)
VALUES
('John Doe', 'john@example.com', TIMESTAMP '1990-01-01', 'password1'),
('Jane Smith', 'jane@example.com', TIMESTAMP '1985-05-20', 'password2'),
('Alice Johnson', 'alice@example.com', TIMESTAMP '1982-11-15', 'password3'),
('Bob Williams', 'bob@example.com', TIMESTAMP '1978-07-10', 'password4'),
('Emma Brown', 'emma@example.com', TIMESTAMP '1992-03-25', 'password5'),
('Michael Davis', 'michael@example.com', TIMESTAMP '1989-09-05', 'password6'),
('Emily Wilson', 'emily@example.com', TIMESTAMP '1995-12-30', 'password7'),
('David Moore', 'david@example.com', TIMESTAMP '1987-02-17', 'password8'),
('Sarah Taylor', 'sarah@example.com', TIMESTAMP '1983-06-22', 'password9'),
('Andrew Anderson', 'andrew@example.com', TIMESTAMP '1991-08-18', 'password10'),
('Olivia Martinez', 'olivia@example.com', TIMESTAMP '1984-04-12', 'password11'),
('James Thompson', 'james@example.com', TIMESTAMP '1980-10-07', 'password12'),
('Isabella Harris', 'isabella@example.com', TIMESTAMP '1993-11-29', 'password13'),
('William Nelson', 'william@example.com', TIMESTAMP '1981-01-14', 'password14'),
('Sophia White', 'sophia@example.com', TIMESTAMP '1986-08-01', 'password15'),
('Matthew Martin', 'matthew@example.com', TIMESTAMP '1994-02-22', 'password16'),
('Charlotte Jackson', 'charlotte@example.com', TIMESTAMP '1979-05-18', 'password17'),
('Daniel Adams', 'daniel@example.com', TIMESTAMP '1988-07-04', 'password18'),
('Oliver Walker', 'oliver@example.com', TIMESTAMP '1980-09-28', 'password20');



INSERT INTO Courses (title, description)
VALUES
('Mathematics', 'This course covers various topics in mathematics including algebra, calculus, and geometry.'),
('Physics', 'This course explores the fundamental principles of physics such as mechanics, thermodynamics, and electromagnetism.'),
('Biology', 'This course focuses on the study of living organisms and their interactions with each other and their environments.'),
('Chemistry', 'This course examines the composition, structure, properties, and reactions of matter.'),
('Computer Science', 'This course introduces students to the fundamental concepts and practices of computer science.'),
('Literature', 'This course explores various literary genres and works of literature from different cultures and time periods.'),
('History', 'This course examines significant events, developments, and personalities in human history.'),
('Art', 'This course explores different forms of artistic expression including painting, sculpture, and architecture.'),
('Music', 'This course introduces students to the elements of music and explores various musical styles and genres.'),
('Psychology', 'This course examines the mind and behavior, covering topics such as perception, cognition, and personality.'),
('Economics', 'This course explores the principles of economics including microeconomics and macroeconomics.'),
('Sociology', 'This course examines human society and social behavior, covering topics such as culture, institutions, and social change.'),
('Political Science', 'This course explores the theory and practice of politics, government, and political systems.'),
('Anthropology', 'This course examines human culture, society, and biology across time and space.'),
('Geography', 'This course explores the physical and human geography of the Earth, including maps, landscapes, and regions.'),
('Environmental Science', 'This course examines the interactions between humans and the environment, covering topics such as ecology and sustainability.'),
('Philosophy', 'This course explores fundamental questions about existence, knowledge, values, reason, mind, and language.'),
('Religious Studies', 'This course examines religious beliefs, practices, and traditions from a comparative and critical perspective.'),
('Health Education', 'This course focuses on promoting health and preventing disease through education and behavior change.'),
('Foreign Languages', 'This course introduces students to the study of foreign languages and cultures.');



INSERT INTO Lessons (course_id, title, content)
VALUES
('cc6101f3-237b-4093-8c8c-5f090a3e881b', 'Lesson 1: Introduction to Algebra', 'This lesson covers basic concepts of algebra including variables, equations, and expressions.'),
('568eb178-5537-4192-864f-e3f6c33058ba', 'Lesson 1: Introduction to Mechanics', 'This lesson introduces the fundamental principles of mechanics such as force, motion, and energy.'),
('3ffd2046-65da-4584-b3ca-3262cf217529', 'Lesson 1: Introduction to Cells', 'This lesson provides an overview of cells, their structure, and their functions in living organisms.'),
('77da741d-3b78-4909-9911-5eb94f928c13', 'Lesson 1: Introduction to Chemical Reactions', 'This lesson explores different types of chemical reactions including synthesis, decomposition, and combustion.'),
('74d43e00-8e10-4e86-884a-1bef4a84bab6', 'Lesson 1: Introduction to Programming', 'This lesson covers basic programming concepts such as variables, data types, and control structures.'),
('080cfab0-0d4e-4c37-8433-fee0af3d3316', 'Lesson 1: Introduction to Poetry', 'This lesson explores various forms of poetry including sonnets, haikus, and ballads.'),
('5aa22b28-e860-4484-977e-fcc95b67fbbb', 'Lesson 1: Ancient Civilizations', 'This lesson examines the rise and fall of ancient civilizations such as Mesopotamia, Egypt, and Greece.'),
('142b08e4-0f65-4585-85ad-46d8ae08a76f', 'Lesson 1: Introduction to Music Theory', 'This lesson introduces basic music theory concepts including notes, scales, and chords.'),
('a4835bf5-bd22-46d2-9559-ddb0b28c32c3', 'Lesson 1: Introduction to Behaviorism', 'This lesson explores the principles of behaviorism and its applications in psychology.'),
('c36c57c2-5634-4d80-ae67-0f3489ebeb0d', 'Lesson 1: Supply and Demand', 'This lesson examines the basic principles of supply and demand in economics.'),
('05ec3e6c-e016-485c-8ba1-4b3b835f5cba', 'Lesson 1: Introduction to Social Institutions', 'This lesson explores the role of social institutions such as family, education, and religion in society.'),
('beee46ef-05fd-4724-8f8f-b014d50a41a4', 'Lesson 1: Introduction to Political Theory', 'This lesson introduces fundamental political theories such as liberalism, conservatism, and socialism.'),
('34529a95-d129-4785-b720-71d982cd2ec9', 'Lesson 1: Introduction to Cultural Anthropology', 'This lesson provides an overview of cultural anthropology and its methods of study.'),
('34bdf8d0-c820-4b6b-8cc2-98d366ca9894', 'Lesson 1: Introduction to Physical Geography', 'This lesson covers basic concepts of physical geography including landforms, climate, and ecosystems.'),
('0cb15dce-8caa-4553-af98-c8ea1bdeabf6', 'Lesson 1: Environmental Issues', 'This lesson examines various environmental issues such as pollution, climate change, and deforestation.'),
('ff3c4974-54bc-46ca-8227-911a6353d165', 'Lesson 1: Introduction to Epistemology', 'This lesson explores different theories of knowledge and the nature of belief and justification.'),
('bb30b390-3e91-4d54-a7cb-1f51e668d009', 'Lesson 1: Comparative Religion', 'This lesson compares different religious beliefs, practices, and traditions from around the world.'),
('6b55172e-1969-4107-8a28-f0e6f5826a87', 'Lesson 1: Personal Health', 'This lesson focuses on personal health habits and lifestyle choices for improving overall well-being.'),
('c355998d-9dc7-486e-a7eb-d851c3cdce55', 'Lesson 1: Introduction to Language Learning', 'This lesson provides an overview of.')


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