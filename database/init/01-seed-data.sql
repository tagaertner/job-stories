-- Wait for tables to exist (created by GORM)
DO $$ 
BEGIN
    WHILE NOT EXISTS (SELECT 1 FROM information_schema.tables WHERE table_name = 'users') LOOP
        PERFORM pg_sleep(1);
    END LOOP;
END $$;

-- Only insert data, don't create tables
INSERT INTO users (id, name, email, role, active) VALUES 
('1', 'John Doe', 'john@example.com', 'customer', true),
('2', 'Jane Smith', 'jane@example.com', 'admin', true),
('3', 'Bob Wilson', 'bob@example.com', 'customer', true),
('4', 'Alice Johnson', 'alice@example.com', 'customer', true),
('5', 'Mike Chen', 'mike@example.com', 'admin', true),
('6', 'Sarah Wilson', 'sarah@example.com', 'customer', true),
('7', 'David Brown', 'david@example.com', 'customer', false),
('8', 'Emma Davis', 'emma@example.com', 'customer', true),
('9', 'James Miller', 'james@example.com', 'admin', true),
('10', 'Lisa Garcia', 'lisa@example.com', 'customer', true)
ON CONFLICT (id) DO NOTHING;

INSERT INTO products (id, name, description, price, inventory) VALUES
('1', 'MacBook Pro', '14-inch MacBook Pro with M3 chip', 1999.99, 10),
('2', 'iPhone 15', 'Latest iPhone with A17 chip', 999.99, 25),
('3', 'AirPods Pro', 'Wireless earbuds with noise cancellation', 249.99, 50),
('4', 'iPad Pro', '12.9-inch iPad Pro with M2 chip', 1099.99, 15),
('5', 'Apple Watch', 'Series 9 GPS + Cellular', 499.99, 30),
('6', 'Magic Keyboard', 'Wireless keyboard for Mac', 199.99, 20),
('7', 'Studio Display', '27-inch 5K Retina display', 1599.99, 8),
('8', 'AirTag', 'Bluetooth tracking device', 29.99, 100),
('9', 'HomePod mini', 'Smart speaker with Siri', 99.99, 25),
('10', 'Mac Studio', 'Compact pro desktop with M2 Max', 3999.99, 5),
('11', 'iPhone 14', 'Previous generation iPhone', 699.99, 40),
('12', 'MacBook Air M2', '13-inch lightweight laptop', 1199.99, 12),
('13', 'Magic Mouse', 'Wireless multi-touch mouse', 79.99, 35),
('14', 'Apple Pencil', '2nd generation stylus for iPad', 129.99, 45),
('15', 'Mac mini', 'Compact desktop computer', 599.99, 18)
ON CONFLICT (id) DO NOTHING;

INSERT INTO orders (id, user_id, product_id, quantity, total_price, status) VALUES
('1', '1', '1', 1, 1999.99, 'completed'),
('2', '2', '2', 2, 1999.98, 'pending'),
('3', '3', '3', 1, 249.99, 'shipped'),
('4', '4', '4', 1, 1099.99, 'shipped'),
('5', '5', '5', 2, 999.98, 'completed'),
('6', '6', '6', 1, 199.99, 'pending'),
('7', '4', '7', 1, 1599.99, 'completed'),
('8', '8', '8', 4, 119.96, 'shipped'),
('9', '9', '9', 1, 99.99, 'cancelled'),
('10', '10', '10', 1, 3999.99, 'pending'),
('11', '6', '11', 1, 699.99, 'completed'),
('12', '4', '12', 1, 1199.99, 'shipped'),
('13', '8', '13', 2, 159.98, 'completed'),
('14', '5', '14', 3, 389.97, 'pending'),
('15', '10', '15', 1, 599.99, 'shipped'),
('16', '7', '2', 1, 999.99, 'cancelled'),
('17', '6', '3', 2, 499.98, 'completed'),
('18', '4', '1', 1, 1999.99, 'pending'),
('19', '8', '8', 10, 299.90, 'completed'),
('20', '10', '5', 1, 499.99, 'shipped')
ON CONFLICT (id) DO NOTHING;