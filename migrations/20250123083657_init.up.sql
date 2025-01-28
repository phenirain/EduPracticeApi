CREATE TABLE IF NOT EXISTS roles
(
    id       SERIAL PRIMARY KEY,
    role_name VARCHAR(40) NOT NULL
);

CREATE TABLE IF NOT EXISTS employees
(
    id         SERIAL PRIMARY KEY,
    full_name    VARCHAR(100) NOT NULL,
    login      VARCHAR(40) NOT NULL,
    password   VARCHAR(64) NOT NULL,
    role_id    INT REFERENCES roles (id) NOT NULL
);

CREATE TABLE IF NOT EXISTS clients
(
    id          SERIAL PRIMARY KEY,
    company_name TEXT NOT NULL,
    contact_person VARCHAR(100) NOT NULL,
    email VARCHAR(50) NOT NULL,
    telephone_number VARCHAR(30) NOT NULL
);

CREATE TABLE IF NOT EXISTS product_categories(
    id SERIAL PRIMARY KEY,
    category_name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS products
(
    id          SERIAL PRIMARY KEY,
    product_name VARCHAR(50)    NOT NULL,
    article     VARCHAR(40)    NOT NULL,
    category_id INT REFERENCES product_categories (id) NOT NULL,
    quantity    INT            NOT NULL,
    price       DECIMAL(12, 2) NOT NULL,
    location VARCHAR(50),
    reserved_quantity INT NOT NULL
);

CREATE TABLE IF NOT EXISTS orders
(
    id         SERIAL PRIMARY KEY,
    client_id  INT REFERENCES clients (id),
    product_id INT REFERENCES products (id),
    order_date  TIMESTAMP      NOT NULL,
    status     VARCHAR(30),
    quantity   INT            NOT NULL,
    total_price DECIMAL(12, 2) NOT NULL
);

CREATE TABLE IF NOT EXISTS drivers
(
    id          SERIAL PRIMARY KEY,
    full_name VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS deliveries
(
    id        SERIAL PRIMARY KEY,
    order_id  INT REFERENCES orders (ID) NOT NULL,
    driver_id INT REFERENCES drivers (id),
    delivery_date TIMESTAMP NOT NULL,
    transport TEXT NOT NULL,
    route TEXT NOT NULL,
    status VARCHAR(30) NOT NULL
);
