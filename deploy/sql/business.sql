CREATE TABLE
    business (
        id INT AUTO_INCREMENT COMMENT '商家ID',
        phone VARCHAR(20) NOT NULL DEFAULT '' COMMENT '手机号',
        tg_id VARCHAR(20) NOT NULL DEFAULT '' COMMENT 'TG ID',
        admin_id bigint NOT NULL COMMENT '管理后台ID',
        create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (id),
        UNIQUE KEY (admin_id),
        UNIQUE KEY (phone)
    ) COMMENT '商家表';