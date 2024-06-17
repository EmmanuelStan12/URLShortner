CREATE DATABASE url_shortner_db;

CREATE USER '<prod-user>'@'localhost' IDENTIFIED BY '<prod-password>';

GRANT ALL PRIVILEGES ON url_shortner_db.* TO '<prod-user>'@'localhost';

FLUSH PRIVILEGES;