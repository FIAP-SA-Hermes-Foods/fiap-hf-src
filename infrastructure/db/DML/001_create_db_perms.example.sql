CREATE DATABASE db_name;

SELECT session_user, current_database();

DROP USER IF EXISTS db_user;

CREATE USER $DB_USER LOGIN PASSWORD 'db_password';

GRANT ALL PRIVILEGES ON DATABASE db_name TO db_user;

GRANT USAGE, CREATE ON SCHEMA public TO db_user;

ALTER DATABASE db_name OWNER TO db_user;

