CREATE TABLE
    IF NOT EXISTS favorites (
        id BIGSERIAL PRIMARY KEY,
        client_id INTEGER NOT NULL,
        product_id INTEGER NOT NULL,
        CONSTRAINT fk_client FOREIGN KEY (client_id) REFERENCES clients (id) ON DELETE CASCADE
    );