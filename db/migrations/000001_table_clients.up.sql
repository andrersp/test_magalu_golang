CREATE TABLE
    IF NOT EXISTS clients (
        id BIGSERIAL PRIMARY KEY NOT NULL,
        name varchar(255) NOT NULL,
        email varchar(255) NOT NULL UNIQUE
    );

INSERT INTO
    clients (name, email)
VALUES
    ('John Doe', 'email@mail.com');