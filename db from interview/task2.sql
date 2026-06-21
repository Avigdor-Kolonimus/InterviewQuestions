-- You need to design a database model for a library.
-- There are three entities:
-- * Author
-- * Book
-- * Reader
--
-- Each physical book exists as a single copy and can be borrowed by only one reader at a time.
-- Create the database tables for the library while taking this requirement into account.
-- 
-- Then write the following SQL queries:
-- Select the titles of all books that are currently checked out.
-- Select the titles of all books in the library that have more than three authors.
-- Select the names of the top 3 most-read authors at the current moment.

-- Author
CREATE TABLE authors(
    author_id CHAR(12) PRIMARY KEY,
    author_name VARCHAR(255) NOT NULL
);

-- Reader
CREATE TABLE readers (
    reader_id CHAR(12) PRIMARY KEY,
    reader_name VARCHAR(255) NOT NULL
);

-- Book
CREATE TABLE books (
    book_id CHAR(12) PRIMARY KEY,
    book_title VARCHAR(255) NOT NULL,
    reader_id CHAR(12) NULL,
    FOREIGN KEY (reader_id) REFERENCES readers(reader_id)
);

-- connection between Book and Author
CREATE TABLE book_authors (
    book_id CHAR(12) NOT NULL,
    author_id CHAR(12) NOT NULL,
    PRIMARY KEY (book_id, author_id),
    FOREIGN KEY (book_id) REFERENCES books(book_id),
    FOREIGN KEY (author_id) REFERENCES authors(author_id)
);

-- Select the titles of all books that are currently checked out.
SELECT b.book_title 
FROM books as b 
WHERE reader_id IS NOT NULL;

-- Select the titles of all books in the library that have more than three authors.
SELECT b.book_title 
FROM books b 
JOIN book_authors ba ON b.book_id = ba.book_id 
GROUP BY b.book_id, b.book_title 
HAVING COUNT(ba.author_id) > 3;

-- Select the names of the top 3 most-read authors at the current moment.
SELECT a.author_name
FROM authors as a
JOIN book_authors ba ON a.author_id = ba.author_id 
JOIN books b ON ba.book_id = b.book_id 
WHERE b.reader_id IS NOT NULL 
GROUP BY a.author_id, a.author_name
ORDER BY COUNT(*) DESC
LIMIT 3;