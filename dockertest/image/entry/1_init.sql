CREATE DATABASE IF NOT EXISTS employees;
USE employees;

CREATE TABLE IF NOT EXISTS employees(
       emp_no INT PRIMARY KEY NOT NULL,
       birth_date DATE NOT NULL,
       first_name VARCHAR(32) NOT NULL,
       last_name BINARY(32) NOT NULL,
       gender ENUM('M', 'F') NOT NULL,
       hire_date DATE NOT NULL
);

CREATE TABLE IF NOT EXISTS salaries(
       emp_no INT NOT NULL,
       salary int(11) NOT NULL,
       from_date DATE NOT NULL,
       to_date DATE NOT NULL,
       CONSTRAINT fk_emp_no
                  FOREIGN KEY (emp_no)
                  REFERENCES employees (emp_no)
                  ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS titles(
       emp_no INT NOT NULL,
       title VARCHAR(50) NOT NULL,
       from_date DATE NOT NULL,
       to_date DATE NOT NULL,
       CONSTRAINT fk_emp_no
                  FOREIGN KEY (emp_no)
                  REFERENCES employees (emp_no)
                  ON DELETE CASCADE       
);

LOAD DATA LOCAL INFILE
     "/docker/data/employees.csv"
     INTO TABLE employees
     CHARACTER SET utf8
     FIELDS TERMINATED BY ','
     ENCLOSED BY '"'
     (emp_no, birth_date, first_name, last_name, gender, hire_date);

LOAD DATA LOCAL INFILE
     "/docker/data/salaries.csv"
     INTO TABLE salaries
     CHARACTER SET utf8
     FIELDS TERMINATED BY ','
     ENCLOSED BY '"'
     (emp_no, salary, from_date, to_date);

LOAD DATA LOCAL INFILE
     "/docker/data/titles.csv"
     INTO TABLE titles
     CHARACTER SET utf8
     FIELDS TERMINATED BY ','
     ENCLOSED BY '"'
     (emp_no, title, from_date, to_date);

