CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE users (
   id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
   email TEXT NOT NULL UNIQUE,
   name TEXT NOT NULL DEFAULT '',
   created_at TIMESTAMPTZ NOT NULL DEFAULT (now()),
   updated_at TIMESTAMPTZ NOT NULL DEFAULT (now())
   
);