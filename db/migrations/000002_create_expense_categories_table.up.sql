CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE Categories
(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    category_name VARCHAR(255) NOT NULL,
    user_id       UUID          NOT NULL,
    FOREIGN KEY (user_id) REFERENCES Users(id)
);

CREATE TABLE Expenses
(
    id UUID DEFAULT uuid_generate_v4() PRIMARY KEY,
    name        VARCHAR(255)   NOT NULL,
    amount      DECIMAL(10, 2) NOT NULL,
    category_id UUID            NOT NULL,
    FOREIGN KEY (category_id) REFERENCES Categories(id)
);
