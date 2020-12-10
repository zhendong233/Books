DROP TABLE IF EXISTS `book`;
CREATE TABLE `book` (
    `book_id` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
    `book_name` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL,
    `author` varchar(36) CHARACTER SET utf8mb4 COLLATE utf8mb4_bin NOT NULL, 
    `created_at` datetime(3) NOT NULL
);