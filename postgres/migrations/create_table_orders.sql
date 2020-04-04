CREATE TABLE orders
(
    id      	BIGSERIAL       PRIMARY KEY,
	buyer_id	INTEGER    NOT NULL,
	total_price	FLOAT    NOT NULL,
	shipping_method_id INTEGER NOT NULL,
	total_shipping_cost	FLOAT NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);