CREATE TABLE `product` (
                           `id` bigint NOT NULL AUTO_INCREMENT,
                           `name` varchar(128) NOT NULL,
                           `description` varchar(512) DEFAULT '',
                           `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                           `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                           PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;