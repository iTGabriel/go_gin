FROM mariadb:latest

LABEL "name"="gogin-mariadb"

ENV MARIADB_USER=my-user
ENV MARIADB_PASSWORD=my-secret
ENV MARIADB_ROOT_PASSWORD=root-secret

EXPOSE 3306