CREATE TABLE carts
(
    id      	BIGSERIAL       PRIMARY KEY,
	buyer_id	INTEGER    NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    deleted_at TIMESTAMP WITH TIME ZONE DEFAULT NULL
);