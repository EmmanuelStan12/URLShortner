CREATE DATABASE IF NOT EXISTS url_shortner_db_test;

CREATE USER 'test'@'localhost' IDENTIFIED BY 'test123';

GRANT ALL PRIVILEGES ON url_shortner_db_test.* TO 'test'@'localhost';

FLUSH PRIVILEGES;