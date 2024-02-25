CREATE TABLE shorteners
(
    id   VARCHAR(255) PRIMARY KEY,
    target_url  VARCHAR(255),
    expiry_date INT,
    clicks      TEXT NOT NULL,
    created_at  INT,
    updated_at  INT
);