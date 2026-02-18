CREATE TABLE daily_active_addresses (
    id SERIAL PRIMARY KEY,
    date DATE NOT NULL,
    chain VARCHAR(50) NOT NULL,
    count BIGINT NOT NULL,
    UNIQUE(date, chain)
);