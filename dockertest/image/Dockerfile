FROM mysql:5.7

ENV MYSQL_ROOT_PASSWORD=toor
ENV MYSQL_ROOT_HOST="%"
ENV MYSQL_DATABASE="mock"
ENV TZ="Asia/Tokyo"

ADD data/ /docker/data
ADD entry/ /docker-entrypoint-initdb.d

CMD ["mysqld"]

EXPOSE 3306
