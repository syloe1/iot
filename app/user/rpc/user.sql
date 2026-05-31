CREATE TABLE user (
                         id bigint NOT NULL AUTO_INCREMENT,
                         username varchar(64) NOT NULL,
                         password varchar(128) NOT NULL,
                         created_at datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         PRIMARY KEY (id),
                         UNIQUE KEY uk_username (username)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;