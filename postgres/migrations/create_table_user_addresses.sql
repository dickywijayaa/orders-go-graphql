CREATE TABLE user_addresses
(
    id          BIGSERIAL   PRIMARY KEY,
	user_id		integer     NOT NULL,
    address     TEXT        NOT NULL,
    province_id integer     NOT NULL,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    deleted_at  TIMESTAMP WITH TIME ZONE DEFAULT NULL
);