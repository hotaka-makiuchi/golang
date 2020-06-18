CREATE DATABASE IF NOT EXISTS gormsample CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
CREATE USER IF NOT EXISTS 'gormsample'@'%' IDENTIFIED BY 'password';
GRANT ALL PRIVILEGES ON gormsample.* TO 'gormsample'@'%';
