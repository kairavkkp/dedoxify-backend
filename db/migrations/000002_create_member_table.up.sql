CREATE TABLE member (
    id SERIAL PRIMARY KEY,
    uuid UUID NOT NULL,
    firstName VARCHAR(255) NULL,
    lastName VARCHAR(255) NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    familyId INTEGER REFERENCES family(id) ON DELETE
    SET
        NULL,
        isActive BOOLEAN NOT NULL DEFAULT TRUE
);