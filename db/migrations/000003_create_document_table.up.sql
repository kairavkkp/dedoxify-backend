CREATE TABLE document (
    id SERIAL PRIMARY KEY,
    uuid UUID NOT NULL,
    family_id INTEGER REFERENCES family(id) ON DELETE
    SET
        NULL,
        member_id INTEGER REFERENCES member(id) ON DELETE
    SET
        NULL,
        category VARCHAR(100) NOT NULL,
        is_thumbnail_ready BOOLEAN NOT NULL DEFAULT FALSE,
        is_processed BOOLEAN NOT NULL DEFAULT FALSE,
        is_active BOOLEAN NOT NULL DEFAULT TRUE,
        created_at TIMESTAMP NOT NULL DEFAULT NOW(),
        updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);