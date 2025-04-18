-- +goose Up
-- SQL in this section is executed when the migration is applied.

-- Check if goose_db_version table exists, if not create it
DO $$ 
BEGIN
    IF NOT EXISTS (SELECT FROM information_schema.tables WHERE table_name = 'goose_db_version') THEN
        CREATE TABLE goose_db_version (
            id SERIAL PRIMARY KEY,
            version_id INT NOT NULL,
            is_applied BOOLEAN NOT NULL,
            applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        );
    END IF;
END $$;

-- Now create the products table
CREATE TABLE products (
    id        BIGSERIAL PRIMARY KEY,
    name      TEXT NOT NULL,
    price     FLOAT NOT NULL,
    category  TEXT NOT NULL,
    UNIQUE(name, category)
);

-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS goose_db_version;
