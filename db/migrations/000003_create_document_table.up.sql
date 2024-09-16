CREATE TABLE document (
    id SERIAL PRIMARY KEY,
    uuid UUID NOT NULL,
    familyId INTEGER REFERENCES family(id) ON DELETE
    SET
        NULL,
        memberId INTEGER REFERENCES member(id) ON DELETE
    SET
        NULL,
        category VARCHAR(100) NOT NULL,
        isThumbnailReady BOOLEAN NOT NULL DEFAULT FALSE,
        isProcessed BOOLEAN NOT NULL DEFAULT FALSE,
        isActive BOOLEAN NOT NULL DEFAULT TRUE
);