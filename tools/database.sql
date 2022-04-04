CREATE TABLE member (
    id int NOT NULL AUTO_INCREMENT,
    name varchar(255) NOT NULL,
    email varchar(255),
    PRIMARY KEY (id)
);

INSERT INTO member (`name`, `email`)
VALUES ('Fajrin', 'fajrin.arif@raksharing.com');

INSERT INTO member (name, email)
VALUES ('Jhon', 'jhon@raksharing.com');