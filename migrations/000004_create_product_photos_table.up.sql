CREATE TABLE product_photos (
                                id BIGSERIAL PRIMARY KEY,
                                product_id BIGINT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
                                file_id TEXT NOT NULL
);