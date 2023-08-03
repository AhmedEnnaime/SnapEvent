FROM postgres:15

WORKDIR /docker-entrypoint-initdb.d

EXPOSE 5432