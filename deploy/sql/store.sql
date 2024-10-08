CREATE TABLE
    store (
        id INT AUTO_INCREMENT COMMENT '店铺ID',
        business_id INT NOT NULL COMMENT '商家ID',
        name VARCHAR(255) NOT NULL COMMENT '店铺名称',
        phone VARCHAR(20) NOT NULL COMMENT '店铺电话',
        status TINYINT NOT NULL DEFAULT 1 COMMENT '店铺状态 0:禁用 1:启用',
        address VARCHAR(255) NOT NULL COMMENT '店铺地址',
        start_time TIME NOT NULL DEFAULT '9:00' COMMENT '营业开始时间',
        end_time TIME NOT NULL DEFAULT '22:00' COMMENT '营业结束时间',
        create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (id)
    ) COMMENT '店铺表';