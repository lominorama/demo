CREATE USER 'demo'@'%' IDENTIFIED BY 'demo';
CREATE DATABASE demo;
GRANT USAGE ON *.* TO `demo`@`%`;
GRANT ALL PRIVILEGES ON `demo`.* TO `demo`@`%`;
FLUSH PRIVILEGES;
USE demo;
CREATE TABLE IF NOT EXISTS images (id INT NOT NULL, name VARCHAR(255) NOT NULL, CONSTRAINT PK_ID PRIMARY KEY (id));
INSERT INTO images VALUES(1, "goat.jpg");
INSERT INTO images VALUES(2, "monkey.jpg");
INSERT INTO images VALUES(3, "thanks.jpg");