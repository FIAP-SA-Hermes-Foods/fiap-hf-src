FROM postgres:latest

COPY ${PWD}/infrastructure/scripts/init-database.sh /docker-entrypoint-initdb.d/

RUN chmod +x /docker-entrypoint-initdb.d/init-database.sh

EXPOSE 5432
