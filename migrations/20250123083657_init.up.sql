CREATE TABLE Roles
(
    ID       SERIAL PRIMARY KEY,
    NameRole VARCHAR(40) NOT NULL
);

CREATE TABLE Employees
(
    ID         SERIAL PRIMARY KEY,
    Surname    VARCHAR(40) NOT NULL,
    FirstName  VARCHAR(40) NOT NULL,
    MiddleName VARCHAR(40) NOT NULL,
    Login      VARCHAR(40) NOT NULL,
    Password   VARCHAR(64) NOT NULL,
    Role_ID    INT REFERENCES Roles (ID) NOT NULL
);

CREATE TABLE Clients
(
    ID          SERIAL PRIMARY KEY,
    CompanyName TEXT NOT NULL,
    ContactPerson TEXT NOT NULL,
    Email VARCHAR(50) NOT NULL,
    NumberPhone VARCHAR(30) NOT NULL
);

CREATE TABLE Products
(
    ID          SERIAL PRIMARY KEY,
    NameProduct VARCHAR(50)    NOT NULL,
    Article     VARCHAR(40)    NOT NULL,
    Quantity    INT            NOT NULL,
    Price       DECIMAL(12, 2) NOT NULL
);

CREATE TABLE Storages
(
    ID              SERIAL PRIMARY KEY,
    Product_ID      INT REFERENCES Products (ID) NOT NULL,
    ProductLocation TEXT NOT NULL,
    QuantityStorage INT NOT NULL,
    QuantityReserved INT NOT NULL
);

CREATE TABLE Orders
(
    ID         SERIAL PRIMARY KEY,
    Product_ID INT REFERENCES Products (ID),
    Client_ID  INT REFERENCES Clients (ID),
    OrderDate  TIMESTAMP      NOT NULL,
    Status     VARCHAR(30),
    Quantity   INT            NOT NULL,
    TotalPrice DECIMAL(12, 2) NOT NULL
);

CREATE TABLE Payments
(
    ID       SERIAL PRIMARY KEY,
    Status   VARCHAR(30),
    Date     TIMESTAMP NOT NULL,
    Order_ID INT REFERENCES Orders (ID)
);

CREATE TABLE Drivers
(
    ID          SERIAL PRIMARY KEY,
    Surname     VARCHAR(40) NOT NULL,
    MiddleName  VARCHAR(40) NOT NULL,
    FirstName   VARCHAR(40) NOT NULL,
    NumberPhone VARCHAR(30) NOT NULL
);

CREATE TABLE Deliveries
(
    ID        SERIAL PRIMARY KEY,
    Order_ID  INT REFERENCES Orders (ID) NOT NULL,
    Transport TEXT NOT NULL,
    Route TEXT NOT NULL,
    Status VARCHAR(30) NOT NULL,
    Driver_ID INT REFERENCES Drivers(ID) NOT NULL
);
