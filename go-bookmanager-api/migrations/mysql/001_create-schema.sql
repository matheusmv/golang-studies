USE bookmanager;

-- DOWN
DROP TABLE IF EXISTS book_authors;
DROP TABLE IF EXISTS authors;
DROP TABLE IF EXISTS books;

-- UP
CREATE TABLE `books` (
    id INT NOT NULL AUTO_INCREMENT,
    isbn VARCHAR(20) NOT NULL,
    title VARCHAR(255) NOT NULL,

    PRIMARY KEY (uuid)
);

CREATE TABLE `authors` (
    id INT NOT NULL AUTO_INCREMENT,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,

    PRIMARY KEY (id)
);

CREATE TABLE `book_authors` (
    book_id INT NOT NULL,
    author_id INT NOT NULL,
    
    PRIMARY KEY (book_id, author_id),
    INDEX author_id (author_id)
);
