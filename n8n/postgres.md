# Postgresql
Create database, username, password, connect database and set permission

> sudo -u postgres psql

> CREATE ROLE n8n_user LOGIN PASSWORD 'someStrongPasswordHere';

> CREATE DATABASE n8n_db OWNER n8n_user;

> \c n8n_db

> ALTER SCHEMA public OWNER TO n8n_user;

> GRANT ALL ON SCHEMA public TO n8n_user;

> \q