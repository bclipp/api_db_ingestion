CREATE TABLE customers (
    id integer PRIMARY KEY,
    first_name VARCHAR(255),
    last_name  VARCHAR(255),
    email  VARCHAR(255),
    latitude FLOAT,
    longitude FLOAT,
    block_id INT,
    state_fips INT,
    state_code VARCHAR(10),
    block_pop INT
);

COPY customers(id,first_name,last_name,email,latitude,longitude)
FROM '/tmp/customers.csv' DELIMITER ',' CSV HEADER;

CREATE TABLE stores (
    id INT PRIMARY KEY,
    name VARCHAR(255),
    latitude FLOAT,
    longitude FLOAT
);

COPY stores(id,name,latitude,longitude)
FROM '/tmp/stores.csv' DELIMITER ',' CSV HEADER;
