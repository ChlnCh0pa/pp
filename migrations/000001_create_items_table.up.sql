CREATE TABLE IF NOT EXISTS items (
    id BIGSERIAL PRIMARY KEY,
    sku VARCHAR(50) NOT NULL UNIQUE,
    name VARCHAR(255) NOT NULL,
    unit VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS shelves (
    id BIGSERIAL PRIMARY KEY,
    code VARCHAR(50) NOT NULL UNIQUE,
    capacity INT NOT NULL
);

CREATE TABLE IF NOT EXISTS stocks (
    item_id BIGINT NOT NULL REFERENCES items(id),
    shelf_id BIGINT NOT NULL REFERENCES shelves(id),
    quantity INT NOT NULL CHECK (quantity >= 0),
    PRIMARY KEY (item_id, shelf_id)
);

CREATE TABLE IF NOT EXISTS book_entries (
    id BIGSERIAL PRIMARY KEY,
    item_id BIGINT NOT NULL REFERENCES items(id),
    shelf_id BIGINT NOT NULL REFERENCES shelves(id),
    type VARCHAR(20) NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);