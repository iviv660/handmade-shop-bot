CREATE TABLE orders (
                        id SERIAL PRIMARY KEY,
                        user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
                        product_id INT NOT NULL REFERENCES products(id) ON DELETE CASCADE,
                        status TEXT NOT NULL DEFAULT 'pending',
                        total_price NUMERIC(10,2) NOT NULL DEFAULT 0, -- 👈 как в Go
                        quantity INT,
                        payment_id TEXT,                              -- 👈 добавлено
                        created_at TIMESTAMP DEFAULT now(),
                        updated_at TIMESTAMP DEFAULT now()
);
