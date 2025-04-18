-- +goose Up

-- Now create the orders table
CREATE TABLE orders (
    id                  TEXT PRIMARY KEY,
    orders              JSONB NOT NULL,
    products            JSONB NOT NULL,
    coupon_code         TEXT DEFAULT '',
     order_status       TEXT DEFAULT 'ordered',
    ordered_by          TEXT DEFAULT ''
   
);
-- +goose Down
-- SQL in this section is executed when the migration is rolled back.
DROP TABLE IF EXISTS orders;
