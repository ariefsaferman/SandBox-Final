CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    description VARCHAR NULL,
    quantity INTEGER NULL
);

CREATE TABLE product_categories (
    id SERIAL PRIMARY KEY,
    name VARCHAR NOT NULL ,
    description VARCHAR NULL
);

INSERT INTO product_categories(name, description) VALUES('panadol', 'obat sakit kepala')