CREATE DATABASE wallet_db_arief; 

CREATE TABLE users(
    id PRIMARY KEY NOT NULL SERIAL, 
    name VARCHAR NOT NULL, 
    email VARCHAR NOT NULL,
    phone VARCHAR(15) NOT NULL,
    password VARCHAR,
    created_at timestamp,
    deleted_at timestamp,
    updated_at timestamp
); 

CREATE TABLE transactions(
    id PRIMARY KEY NOT NULL SERIAL, 
    sender INT, 
    receiver INT, 
    amount BIGINT CHECK (amount >= 0), 
    description VARCHAR, 
    source_of_fund_id INT,
    created_at timestamp,
    deleted_at timestamp,
    updated_at timestamp
);

CREATE TABLE wallets (
    id PRIMARY KEY NOT NULL SERIAL, 
    wallet_number INT, 
    balance INT, 
    user_id INT, 
    created_at timestamp,
    deleted_at timestamp,
    updated_at timestamp
)

INSERT INTO user(name, email, phone, password) VALUES ('arief', 'arief@gmail.com', '12345', 'password');
INSERT INTO user(name, email, phone, password) VALUES ('andra', 'andra@gmail.com', '6789', 'password');
INSERT INTO user(name, email, phone, password) VALUES ('rai', 'rai@gmail.com', '11111', 'password');
INSERT INTO user(name, email, phone, password) VALUES ('javin', 'javin@gmail.com', '22222', 'password');
INSERT INTO user(name, email, phone, password) VALUES ('merisa', 'merisa@gmail.com', '33333', 'password');

INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (1, 0, 50000, 'Top up from Bank', 1)
INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (1, 700002, 10000, 'Paylatter', 4)
INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (1, 700003, 10000, 'Paylatter', 4)
INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (1, 700004, 10000, 'Paylatter', 4)
INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (1, 700005 10000, 'Paylatter', 4)
INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (2, 0, 10000, 'Top up from Bank', 1)
INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (2, 0, 10000, 'Top up from Bank', 1)
INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (2, 0, 10000, 'Top up from Bank', 1)
INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (2, 0, 10000, 'Top up from Bank', 1)
INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (2, 0, 10000, 'Top up from Bank', 1)
INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (3, 0, 10000, 'Top up from Bank', 1)
INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (3, 0, 10000, 'Top up from Bank', 1)
INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (3, 0, 10000, 'Top up from Bank', 1)
INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (3, 0, 10000, 'Top up from Bank', 1)
INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (3, 0, 10000, 'Top up from Bank', 1)
INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (4, 0, 10000, 'Top up from Bank', 1)
INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (4, 0, 10000, 'Top up from Bank', 1)
INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (4, 0, 10000, 'Top up from Bank', 1)
INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (4, 0, 10000, 'Top up from Bank', 1)
INSERT INTO transactions(sender, receiver, amount, description, source_of_fund_id) VALUES (4, 0, 10000, 'Top up from Bank', 1)