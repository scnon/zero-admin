CREATE TABLE
    sys_menu (
        id INT AUTO_INCREMENT,
        parent_id INT,
        tenant_id INT NOT NULL DEFAULT 0,
        name VARCHAR(255) NOT NULL DEFAULT '',
        path VARCHAR(255) NOT NULL DEFAULT '',
        title VARCHAR(255) NOT NULL DEFAULT '',
        sort INT NOT NULL DEFAULT 0,
        create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (id)
    ) COMMENT '菜单表';