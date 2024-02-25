CREATE TABLE shorteners
(
    id   VARCHAR(255) PRIMARY KEY,
    target_url  VARCHAR(255),
    expiry_date INT,
    clicks      TEXT NOT NULL,
    created_at  INT,
    updated_at  INT
);

CREATE TABLE IF NOT EXISTS chefs (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    busy BOOLEAN
);

CREATE TABLE IF NOT EXISTS menus (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255),
    duration INT
);

CREATE TABLE IF NOT EXISTS orders (
    id INT AUTO_INCREMENT PRIMARY KEY,
    pizza VARCHAR(255),
    duration INT,
    status VARCHAR(255),
    created_at INT
);
