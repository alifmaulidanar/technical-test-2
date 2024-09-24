CREATE DATABASE IF NOT EXISTS go_technical_test;

USE go_technical_test;

CREATE TABLE users (
    user_id CHAR(36) NOT NULL,
    first_name VARCHAR(50) NOT NULL,
    last_name VARCHAR(50) NOT NULL,
    phone_number VARCHAR(15) NOT NULL UNIQUE,
    address TEXT NOT NULL,
    pin VARCHAR(255) NOT NULL,
    balance DECIMAL(15, 2) DEFAULT 0,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (user_id)
);

CREATE TABLE transactions (
    transaction_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    transaction_type ENUM('CREDIT', 'DEBIT') NOT NULL,
    amount DECIMAL(15, 2) NOT NULL,
    remarks TEXT,
    balance_before DECIMAL(15, 2) NOT NULL,
    balance_after DECIMAL(15, 2) NOT NULL,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    transaction_reference CHAR(36) NOT NULL,
    PRIMARY KEY (transaction_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

CREATE TABLE tokens (
    token_id CHAR(36) NOT NULL,
    user_id CHAR(36) NOT NULL,
    access_token TEXT NOT NULL,
    refresh_token TEXT NOT NULL,
    created_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (token_id),
    FOREIGN KEY (user_id) REFERENCES users (user_id)
);

SELECT * FROM users;

SELECT * FROM transactions;

SELECT * FROM tokens;

DROP TABLE tokens;

DROP TABLE transactions;

DROP TABLE users;

TRUNCATE TABLE tokens;

TRUNCATE TABLE transactions;

SET FOREIGN_KEY_CHECKS = 0;

TRUNCATE TABLE users;

SET FOREIGN_KEY_CHECKS = 1;