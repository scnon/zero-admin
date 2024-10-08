CREATE TABLE
    attr_name (
        id INT PRIMARY KEY AUTO_INCREMENT COMMENT '属性值ID',
        name VARCHAR(255) NOT NULL COMMENT '属性值',
        create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (id)
    ) COMMENT '属性值表';