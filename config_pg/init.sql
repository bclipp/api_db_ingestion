CREATE TABLE customers (
    id integer PRIMARY KEY,
    first_name character varying(255),
    last_name  character varying(255),
    email  character varying(255),
    latitude float,
    longitude float
);

COPY customers(id,first_name,last_name,email,latitude,longitude)
FROM 'customers.csv' DELIMITER ',' CSV HEADER;

CREATE TABLE stores (
    id integer PRIMARY KEY,
    name character varying(255),
    latitude float,
    longitude float
);

COPY customers(id,name,latitude,longitude)
FROM 'stores.csv' DELIMITER ',' CSV HEADER;