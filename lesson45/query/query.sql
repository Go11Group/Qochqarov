create table Book(Book_id uuid primary key not null default gen_random_uuid(), Title varchar, Author varchar,Year int,user_Id varchar default null,borrow bool default false)


INSERT INTO Book (Title, Author, Year, user_Id, borrow) VALUES
('To Kill a Mockingbird', 'Harper Lee', 1960, 'U001', false),
('1984', 'George Orwell', 1949, 'U002', true),
('Pride and Prejudice', 'Jane Austen', 1813, NULL, false),
('The Great Gatsby', 'F. Scott Fitzgerald', 1925, 'U003', true),
('Moby-Dick', 'Herman Melville', 1851, NULL, false),
('War and Peace', 'Leo Tolstoy', 1869, 'U004', true),
('The Catcher in the Rye', 'J.D. Salinger', 1951, 'U005', false),
('The Lord of the Rings', 'J.R.R. Tolkien', 1954, NULL, false),
('Harry Potter and the Philosophers Stone', 'J.K. Rowling', 1997, 'U006', true),
('The Hobbit', 'J.R.R. Tolkien', 1937, NULL, false),
('Crime and Punishment', 'Fyodor Dostoevsky', 1866, 'U007', true),
('The Odyssey', 'Homer', -800, NULL, false),
('Brave New World', 'Aldous Huxley', 1932, 'U008', true),
('Wuthering Heights', 'Emily Brontë', 1847, NULL, false),
('Jane Eyre', 'Charlotte Brontë', 1847, 'U009', true);