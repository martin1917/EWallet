CREATE TABLE wallets
(
    id uuid PRIMARY KEY,
    balance DECIMAL NOT NULL
);

CREATE TABLE transactions
(
    id uuid PRIMARY KEY,
    from_wallet uuid NOT NULL,
    to_wallet uuid NOT NULL,
    amount DECIMAL NOT NULL,
    date_time TIMESTAMP NOT NULL
);

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

insert into wallets(id, balance) values 
    (uuid_generate_v4(), 100.00),
    (uuid_generate_v4(), 100.00),
    (uuid_generate_v4(), 100.00),
    (uuid_generate_v4(), 100.00),
    (uuid_generate_v4(), 100.00),
    (uuid_generate_v4(), 100.00),
    (uuid_generate_v4(), 100.00),
    (uuid_generate_v4(), 100.00),
    (uuid_generate_v4(), 100.00),
    (uuid_generate_v4(), 100.00);