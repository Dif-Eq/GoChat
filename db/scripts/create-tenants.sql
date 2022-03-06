CREATE TABLE IF NOT EXISTS tenants (
   id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
   name VARCHAR NOT NULL
);
