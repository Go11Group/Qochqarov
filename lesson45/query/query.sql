create table Book(Book_id uuid primary key not null default gen_random_uuid(), Title varchar, Author varchar,Year int,borrow bool default false)


INSERT INTO Book (Book_id, Title, Author, Year) VALUES
('B001', 'To Kill a Mockingbird', 'Harper Lee', 1960),
('B002', '1984', 'George Orwell', 1949),
('B003', 'Pride and Prejudice', 'Jane Austen', 1813),
('B004', 'The Great Gatsby', 'F. Scott Fitzgerald', 1925),
('B005', 'Moby-Dick', 'Herman Melville', 1851),
('B006', 'War and Peace', 'Leo Tolstoy', 1869),
('B007', 'The Catcher in the Rye', 'J.D. Salinger', 1951),
('B008', 'The Lord of the Rings', 'J.R.R. Tolkien', 1954),
('B009', 'Harry Potter and the Philosophers Stone', 'J.K. Rowling', 1997),
('B010', 'The Hobbit', 'J.R.R. Tolkien', 1937),
('B011', 'Crime and Punishment', 'Fyodor Dostoevsky', 1866),
('B012', 'The Odyssey', 'Homer', -800),
('B013', 'Brave New World', 'Aldous Huxley', 1932),
('B014', 'Wuthering Heights', 'Emily Brontë', 1847),
('B015', 'Jane Eyre', 'Charlotte Brontë', 1847),
('B016', 'The Divine Comedy', 'Dante Alighieri', 1320),
('B017', 'The Brothers Karamazov', 'Fyodor Dostoevsky', 1880),
('B018', 'Madame Bovary', 'Gustave Flaubert', 1857),
('B019', 'The Iliad', 'Homer', -750),
('B020', 'One Hundred Years of Solitude', 'Gabriel Garcia Marquez', 1967);