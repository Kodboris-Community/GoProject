CREATE TYPE member_status AS ENUM (
    'active',
    'inactive',
    'dormant'
    );

CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE members (
                          member_id uuid PRIMARY KEY DEFAULT uuid_generate_v4 (),
                          first_name varchar(10) NOT NULL,
                          last_name varchar(10) NOT NULL,
                          created_at timestamptz DEFAULT now(),
                          status member_status NOT NULL DEFAULT 'active',
                          comment varchar
);


