CREATE TABLE
    customer (
        id INT PRIMARY KEY AUTO_INCREMENT COMMENT '顾客ID',
        phone VARCHAR(20) NOT NULL DEFAULT '' COMMENT '手机号',
        tg_id VARCHAR(20) NOT NULL DEFAULT '' COMMENT 'Telegram ID',
        tg_username VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'Telegram 用户名',
        tg_first_name VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'Telegram 名',
        tg_last_name VARCHAR(255) NOT NULL DEFAULT '' COMMENT 'Telegram 姓',
        tg_language_code VARCHAR(10) NOT NULL DEFAULT '' COMMENT 'Telegram 语言',
        create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (id),
        UNIQUE KEY (tg_id),
        UNIQUE KEY (phone)
    ) COMMENT '顾客表';