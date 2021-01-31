CREATE DATABASE IF NOT EXISTS mock;
USE mock;

CREATE TABLE IF NOT EXISTS mock(
       id INT PRIMARY KEY NOT NULL,
       height FLOAT NOT NULL,
       gender CHAR(1) NOT NULL,
       day_of_the_week BINARY(3) NOT NULL,
       first_name VARCHAR(128) NOT NULL,
       last_name VARBINARY(128) NOT NULL,
       date1 DATE NOT NULL,
       date2 DATETIME NOT NULL
);

CREATE TABLE IF NOT EXISTS mock2(
       id INT PRIMARY KEY NOT NULL,
       company VARCHAR(128) NOT NULL,
       country VARCHAR(128) NOT NULL,
       from_date DATE NOT NULL
);

LOAD DATA LOCAL INFILE
     "/docker/data/mock.csv"
     INTO TABLE mock
     CHARACTER SET utf8
     FIELDS TERMINATED BY ','
     ENCLOSED BY '"'
     (id, height, gender, day_of_the_week, date1, date2, first_name, last_name);

LOAD DATA LOCAL INFILE
     "/docker/data/mock2.csv"
     INTO TABLE mock2
     CHARACTER SET utf8
     FIELDS TERMINATED BY ','
     ENCLOSED BY '"'
     (id, from_date, company, country);

