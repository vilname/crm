FROM postgres:16-alpine

COPY development/db/init.sql /docker-entrypoint-initdb.d/
RUN chmod a+r /docker-entrypoint-initdb.d/init.sql