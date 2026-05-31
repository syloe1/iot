-- 设备表
CREATE TABLE `device` (
  `id` bigint NOT NULL AUTO_INCREMENT,
  `product_id` bigint NOT NULL,
  `name` varchar(128) NOT NULL,
  `device_key` varchar(64) NOT NULL COMMENT '设备唯一标识',
  `status` tinyint NOT NULL DEFAULT 0 COMMENT '0:离线 1:在线',
  `last_online_time` datetime DEFAULT NULL,
  `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `uk_device_key` (`device_key`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;