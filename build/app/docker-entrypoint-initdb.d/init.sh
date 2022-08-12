#!/bin/bash
source .env
set -a

# todo: have to change test stuff to variables
psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE USER tester WITH PASSWORD 'test_password';
    CREATE DATABASE vuegodb_test;
    GRANT ALL PRIVILEGES ON DATABASE vuegodb_test TO tester;
EOSQL
set +a
