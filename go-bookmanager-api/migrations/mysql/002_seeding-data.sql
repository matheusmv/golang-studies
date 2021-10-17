USE bookmanager;

INSERT INTO authors (first_name, last_name) VALUES
("John", "Ronald Reuel Tolkien"),
("Douglas", "Adams"),
("Arthur", "Conan Doyle"),
("Clive", "Staples Lewis"),
("Bram", "Stoker");

INSERT INTO books (isbn, title) VALUES
("978-8595086357", "Box Trilogia O Senhor dos Anéis"),
("978-8530601492", "O guia definitivo do mochileiro das galáxias"),
("978-8595080836", "Box Sherlock Holmes"),
("978-8578270698", "As Crônicas de Nárnia"),
("978-8566636239", "Drácula");

INSERT INTO book_authors (book_id, author_id) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 5);
