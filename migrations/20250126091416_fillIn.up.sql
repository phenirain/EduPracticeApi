INSERT INTO roles (role_name)
VALUES
('Manager'),
('StorageManager'),
('SalesManager'),
('Logician');

INSERT INTO product_categories(category_name)
VALUES
('Семена'),
('Удобрения'),
('Техника');

INSERT INTO employees (full_name, login, password, role_id) VALUES
('Alice Johnson', 'alice', '123', 1),
('Bob Smith', 'bob', '123', 2),
('Charlie Brown', 'charlie', '123', 3),
('Diana Prince', 'diana', 'securepass', 2),
('Eve Adams', 'eve', 'adminpass', 1),
('Frank Miller', 'frank', 'managerpass', 2),
('Grace Hopper', 'grace', 'driverpass', 3),
('Hannah Lee', 'hannah', 'supportpass', 4),
('Ian Curtis', 'ian', 'manager2023', 2),
('Julia Roberts', 'julia', 'adminsecure', 1);

INSERT INTO clients (company_name, contact_person, email, telephone_number) VALUES
('TechCorp', 'John Doe', 'johndoe@techcorp.com', '+123456789'),
('HomeDesigns', 'Sarah Connor', 'sarah@homedesigns.com', '+987654321'),
('FashionWorld', 'Tom Hanks', 'tom@fashionworld.com', '+112233445'),
('ToyPlanet', 'Emma Watson', 'emma@toyplanet.com', '+223344556'),
('BookLovers', 'James Bond', 'james@booklovers.com', '+334455667'),
('AutoParts', 'Bruce Wayne', 'bruce@autoparts.com', '+445566778'),
('HealthyFoods', 'Clark Kent', 'clark@healthyfoods.com', '+556677889'),
('SmartTech', 'Natasha Romanoff', 'natasha@smarttech.com', '+667788990'),
('GreenEnergy', 'Tony Stark', 'tony@greenenergy.com', '+778899001'),
('BuildIt', 'Steve Rogers', 'steve@buildit.com', '+889900112');

INSERT INTO products (product_name, article, category_id, quantity, price, location, reserved_quantity) VALUES
('Пшеница элитная', 'SE001', 1, 500, 150.00, 'Склад A', 50),
('Ячмень посевной', 'SE002', 1, 300, 120.00, 'Склад A', 30),
('Кукуруза сахарная', 'SE003', 1, 200, 180.00, 'Склад B', 20),
('Подсолнечник масличный', 'SE004', 1, 400, 250.00, 'Склад B', 40),
('Рапс посевной', 'SE005', 1, 350, 220.00, 'Склад C', 30),
('Комплексное удобрение NPK', 'FE001', 2, 100, 4500.00, 'Склад D', 10),
('Калийная соль', 'FE002', 2, 200, 3500.00, 'Склад D', 20),
('Азотное удобрение (селитра)', 'FE003', 2, 150, 3800.00, 'Склад E', 15),
('Микроудобрения', 'FE004', 2, 50, 5500.00, 'Склад E', 5),
('Трактор МТЗ-82', 'EQ001', 3, 10, 1500000.00, 'Склад F', 2),
('Сеялка зерновая', 'EQ002', 3, 5, 750000.00, 'Склад F', 1),
('Опрыскиватель навесной', 'EQ003', 3, 8, 300000.00, 'Склад G', 1),
('Комбайн зерноуборочный', 'EQ004', 3, 3, 4500000.00, 'Склад G', 0),
('Плуги для трактора', 'EQ005', 3, 15, 250000.00, 'Склад H', 3);

INSERT INTO orders (client_id, product_id, order_date, status, quantity, total_price) VALUES
(1, 1, '2025-01-01 08:00:00', 'Delivering', 50, 7500.00),
(1, 2, '2025-01-01 08:30:00', 'Delivering', 30, 3600.00),
(2, 3, '2025-01-02 09:00:00', 'Canceled', 20, 3600.00),
(2, 4, '2025-01-02 09:30:00', 'Completed', 40, 10000.00),
(3, 5, '2025-01-03 10:00:00', 'Completed', 30, 6600.00),
(3, 6, '2025-01-03 10:30:00', 'Completed', 10, 45000.00),
(4, 7, '2025-01-04 11:00:00', 'Completed', 20, 70000.00),
(4, 8, '2025-01-04 11:30:00', 'Completed', 15, 57000.00),
(5, 9, '2025-01-05 12:00:00', 'Completed', 2, 3000000.00),
(5, 10, '2025-01-05 12:30:00', 'Completed', 1, 750000.00),
(6, 11, '2025-01-06 13:00:00', 'Completed', 1, 300000.00),
(6, 12, '2025-01-06 13:30:00', 'Completed', 1, 4500000.00),
(7, 13, '2025-01-07 14:00:00', 'Delivering', 3, 750000.00);


INSERT INTO drivers (full_name) VALUES
('John Smith'),
('Emma Stone'),
('Michael Johnson'),
('Sarah Davis'),
('James Anderson'),
('Emily Taylor'),
('Robert Wilson'),
('Sophia Martinez'),
('David Harris'),
('Olivia Clark');

INSERT INTO deliveries (order_id, driver_id, delivery_date, transport, route, status) VALUES
(1, 1, '2025-02-02 10:00:00', 'Truck', 'Route A', 'Scheduled'),
(2, 2, '2025-02-03 12:00:00', 'Van', 'Route B', 'OnTheWay'),
(3, 3, '2025-01-04 14:00:00', 'Truck', 'Route C', 'Canceled'),
(4, 4, '2025-01-05 09:00:00', 'Bike', 'Route D', 'Completed'),
(5, 5, '2025-01-06 11:00:00', 'Car', 'Route E', 'Completed'),
(6, 6, '2025-01-07 13:00:00', 'Truck', 'Route F', 'Completed'),
(7, 7, '2025-01-08 15:00:00', 'Van', 'Route G', 'Completed'),
(8, 8, '2025-01-10 10:00:00', 'Car', 'Route H', 'Completed'),
(9, 9, '2025-01-20 17:00:00', 'Truck', 'Route I', 'Completed'),
(10, 10, '2025-01-30 18:00:00', 'Bike', 'Route J', 'Completed');




