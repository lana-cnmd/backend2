CREATE TABLE IF NOT EXISTS addresses
(
    id SERIAL PRIMARY KEY,
    country VARCHAR(100) NOT NULL,
    city VARCHAR(100) NOT NULL,
    street VARCHAR(150) NOT NULL
);

CREATE TABLE IF NOT EXISTS suppliers
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    address_id INT NOT NULL,
    phone_number VARCHAR(50) NOT NULL,
    FOREIGN KEY (address_id) REFERENCES addresses(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS clients
(
    id SERIAL PRIMARY KEY,
    client_name VARCHAR(50) NOT NULL,
    client_surname VARCHAR(70) NOT NULL,
    birthday DATE NOT NULL,
    gender VARCHAR(1) NOT NULL CHECK (gender IN ('M', 'F')),
    registration_date DATE NOT NULL,
    address_id INT NOT NULL,
    FOREIGN KEY (address_id) REFERENCES addresses(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS images
(
    id UUID PRIMARY KEY,
    image BYTEA NOT NULL
);


CREATE TABLE IF NOT EXISTS products
(
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    category VARCHAR(100) NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    available_stock INT NOT NULL,
    last_update_date DATE NOT NULL,
    supplier_id INT NOT NULL,
    image_id UUID,
    FOREIGN KEY (supplier_id) REFERENCES suppliers(id) ON DELETE CASCADE,
    FOREIGN KEY(image_id) REFERENCES images(id) ON DELETE CASCADE
);
