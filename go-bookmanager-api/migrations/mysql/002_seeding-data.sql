USE bookmanager;

INSERT INTO authors (first_name, last_name, created_at, updated_at) VALUES
("John", "Ronald Reuel Tolkien", now(), now()),
("Douglas", "Adams", now(), now()),
("Arthur", "Conan Doyle", now(), now()),
("Clive", "Staples Lewis", now(), now()),
("Bram", "Stoker", now(), now());

INSERT INTO books (isbn, title, created_at, updated_at) VALUES
("978-8595086357", "Box Trilogia O Senhor dos Anéis", now(), now()),
("978-8530601492", "O guia definitivo do mochileiro das galáxias", now(), now()),
("978-8595080836", "Box Sherlock Holmes", now(), now()),
("978-8578270698", "As Crônicas de Nárnia", now(), now()),
("978-8566636239", "Drácula", now(), now());

INSERT INTO book_authors (book_id, author_id) VALUES
(1, 1),
(2, 2),
(3, 3),
(4, 4),
(5, 5);
