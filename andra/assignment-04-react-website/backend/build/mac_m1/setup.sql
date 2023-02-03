CREATE TABLE IF NOT EXISTS "wallets_tab"(
    "id" BIGSERIAL PRIMARY KEY,
    "balance" NUMERIC NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL,
    "deleted_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "users_tab"(
    "id" BIGSERIAL PRIMARY KEY,
    "email" VARCHAR UNIQUE NOT NULL,
    "password" VARCHAR NOT NULL,
    "first_name" VARCHAR NOT NULL,
    "last_name" VARCHAR NOT NULL,
    "wallet_id" BIGINT UNIQUE NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL,
    "deleted_at" TIMESTAMP,
    CONSTRAINT fk_wallet FOREIGN KEY("wallet_id") REFERENCES "wallets_tab"("id")
);

CREATE TABLE IF NOT EXISTS "source_of_funds_tab"(
    "id" BIGSERIAL PRIMARY KEY,
    "name" VARCHAR UNIQUE NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL,
    "deleted_at" TIMESTAMP
);

CREATE TABLE IF NOT EXISTS "transactions_tab"(
    "id" BIGSERIAL PRIMARY KEY,
    "to" BIGINT,
    "amount" NUMERIC NOT NULL,
    "description" VARCHAR,
    "source_of_fund_id" BIGINT,
    "wallet_id" BIGINT NOT NULL,
    "created_at" TIMESTAMP NOT NULL,
    "updated_at" TIMESTAMP NOT NULL,
    "deleted_at" TIMESTAMP,
    CONSTRAINT fk_wallet FOREIGN KEY("wallet_id") REFERENCES "wallets_tab"("id"),
    CONSTRAINT fk_wallet_to FOREIGN KEY("to") REFERENCES "wallets_tab"("id"),
    CONSTRAINT fk_source_of_fund FOREIGN KEY("source_of_fund_id") REFERENCES "source_of_funds_tab"("id")
);

INSERT INTO "source_of_funds_tab" ("name", "created_at", "updated_at")
VALUES
    ('Bank Transfer', NOW(), NOW()),
    ('Credit Card', NOW(), NOW()),
    ('Cash', NOW(), NOW());

INSERT INTO "wallets_tab" ("balance", "created_at", "updated_at")
VALUES
    (21000000, NOW(), NOW()),
    (21000000, NOW(), NOW()),
    (20580000, NOW(), NOW());

INSERT INTO "users_tab" ("email", "password", "first_name", "last_name", "wallet_id", "created_at", "updated_at")
VALUES
    ('acong@mail.com', '$2a$10$/3Z1rguEv6DFFGab8jzjUON0TSlndPMQDxrP/H521k1qXinr9vjWe', 'Acong', 'Suherman', 1, NOW(), NOW()),
    ('djoko@mail.com', '$2a$10$/3Z1rguEv6DFFGab8jzjUON0TSlndPMQDxrP/H521k1qXinr9vjWe', 'Djoko', 'Susanto', 2, NOW(), NOW()),
    ('sitorus@mail.com', '$2a$10$/3Z1rguEv6DFFGab8jzjUON0TSlndPMQDxrP/H521k1qXinr9vjWe', 'Sitorus', 'Hermanto', 3, NOW(), NOW());

INSERT INTO "transactions_tab" ("to", "amount", "description", "source_of_fund_id", "wallet_id", "created_at", "updated_at")
VALUES
    (NULL, 1000000, 'Top up from Bank Transfer', 1, 1, NOW(), NOW()),
    (NULL, 2000000, 'Top up from Bank Transfer', 1, 1, NOW(), NOW()),
    (NULL, 3000000, 'Top up from Credit Card', 2, 1, NOW(), NOW()),
    (NULL, 4000000, 'Top up from Credit Card', 2, 1, NOW(), NOW()),
    (NULL, 5000000, 'Top up from Cash', 3, 1, NOW(), NOW()),
    (NULL, 6000000, 'Top up from Cash', 3, 1, NOW(), NOW()),
    (NULL, 1000000, 'Top up from Bank Transfer', 1, 2, NOW(), NOW()),
    (NULL, 2000000, 'Top up from Bank Transfer', 1, 2, NOW(), NOW()),
    (NULL, 3000000, 'Top up from Credit Card', 2, 2, NOW(), NOW()),
    (NULL, 4000000, 'Top up from Credit Card', 2, 2, NOW(), NOW()),
    (NULL, 5000000, 'Top up from Cash', 3, 2, NOW(), NOW()),
    (NULL, 6000000, 'Top up from Cash', 3, 2, NOW(), NOW()),
    (NULL, 1000000, 'Top up from Bank Transfer', 1, 3, NOW(), NOW()),
    (NULL, 2000000, 'Top up from Bank Transfer', 1, 3, NOW(), NOW()),
    (NULL, 3000000, 'Top up from Credit Card', 2, 3, NOW(), NOW()),
    (NULL, 4000000, 'Top up from Credit Card', 2, 3, NOW(), NOW()),
    (NULL, 5000000, 'Top up from Cash', 3, 3, NOW(), NOW()),
    (NULL, 6000000, 'Top up from Cash', 3, 3, NOW(), NOW()),
    (2, 10000, 'Breakfast', NULL, 1, NOW(), NOW()),
    (3, 10000, 'Breakfast', NULL, 1, NOW(), NOW()),
    (2, 10000, 'Breakfast', NULL, 1, NOW(), NOW()),
    (3, 10000, 'Breakfast', NULL, 1, NOW(), NOW()),
    (2, 10000, 'Breakfast', NULL, 1, NOW(), NOW()),
    (3, 10000, 'Breakfast', NULL, 1, NOW(), NOW()),
    (2, 10000, 'Breakfast', NULL, 1, NOW(), NOW()),
    (3, 20000, 'Lunch', NULL, 1, NOW(), NOW()),
    (2, 20000, 'Lunch', NULL, 1, NOW(), NOW()),
    (3, 20000, 'Lunch', NULL, 1, NOW(), NOW()),
    (2, 20000, 'Lunch', NULL, 1, NOW(), NOW()),
    (3, 20000, 'Lunch', NULL, 1, NOW(), NOW()),
    (2, 20000, 'Lunch', NULL, 1, NOW(), NOW()),
    (3, 20000, 'Lunch', NULL, 1, NOW(), NOW()),
    (2, 30000, 'Dinner', NULL, 1, NOW(), NOW()),
    (3, 30000, 'Dinner', NULL, 1, NOW(), NOW()),
    (2, 30000, 'Dinner', NULL, 1, NOW(), NOW()),
    (3, 30000, 'Dinner', NULL, 1, NOW(), NOW()),
    (2, 30000, 'Dinner', NULL, 1, NOW(), NOW()),
    (3, 30000, 'Dinner', NULL, 1, NOW(), NOW()),
    (2, 30000, 'Dinner', NULL, 1, NOW(), NOW());
