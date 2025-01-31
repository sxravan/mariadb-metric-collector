CREATE TABLE IF NOT EXISTS user (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL
);

CREATE TABLE IF NOT EXISTS booking (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    booking_date DATETIME NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user(id)
);

INSERT INTO user (name, email) VALUES
('John Doe', 'john@example.com'),
('Jane Smith', 'jane@example.com'),
('Alice Johnson', 'alice@example.com'),
('Bob Williams', 'bob@example.com');

INSERT INTO booking (user_id, booking_date) VALUES
(1, '2025-02-15 14:30:00'),
(2, '2025-03-10 10:00:00'),
(3, '2025-04-05 08:45:00'),
(4, '2025-02-25 16:00:00');

GRANT ALL PRIVILEGES ON *.* TO `go_user`@`%` IDENTIFIED BY PASSWORD '*80BC6B1CFF248AD20446E9DB1AF2EFD0A30B9E36' WITH GRANT OPTION
FLUSH PRIVILEGES;