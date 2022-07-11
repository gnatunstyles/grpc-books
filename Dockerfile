FROM mysql:latest

ENV MYSQL_ROOT_PASSWORD=root

COPY ./test_migration.sql /docker-entrypoint-initdb.d/