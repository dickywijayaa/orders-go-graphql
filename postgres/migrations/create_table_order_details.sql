CREATE TABLE order_details
(
    id      	BIGSERIAL       PRIMARY KEY,
	order_id	INTEGER    NOT NULL,
	seller_id	INTEGER    NOT NULL,
	item_id		INTEGER 	NOT NULL,
	item_name 	VARCHAR(225) NOT NULL,
	item_price	INTEGER    NOT NULL,
    item_quantity INTEGER    NOT NULL,
    item_weight INTEGER NOT NULL,
    shipping_cost FLOAT NOT NULL,
	shipping_method_id INTEGER NOT NULL,
	created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);