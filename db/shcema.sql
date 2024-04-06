CREATE TYPE account_status AS ENUM (
    'active',
    'inactive',
    'dormant'
    );

CREATE TYPE customer_status AS ENUM (
    'active',
    'inactive'
    );

CREATE TYPE gender AS ENUM (
    'm',
    'f'
    );
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE accounts (
                          account_id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
                          account_type varchar NOT NULL,
                          account_number varchar(10) NOT NULL,
                          balance bigint NOT NULL DEFAULT 0,
                          currency varchar NOT NULL,
                          created_at timestamptz DEFAULT now(),
                          status account_status NOT NULL DEFAULT 'active',
                          comment varchar
);
