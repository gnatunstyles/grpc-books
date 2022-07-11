CREATE DATABASE db;
use db;

CREATE TABLE IF NOT EXISTS books(
    id MEDIUMINT AUTO_INCREMENT,
    name CHAR(255),
    author CHAR(255),
    PRIMARY KEY (id)
);

INSERT INTO books (name, author)
    VALUES
        ('Harry Potter and the goblet of fire', 'Rowling'),
        ('Harry Potter and the order of the phoenix', 'Rowling'),
        ('Moidodyr', 'Chukovsky'),
        ('Animal Farm', 'orwell'),
        ('1984', 'orwell');