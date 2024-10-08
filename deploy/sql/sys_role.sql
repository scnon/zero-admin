CREATE TABLE
    sys_role (
        id INT PRIMARY KEY AUTO_INCREMENT COMMENT '角色ID',
        name VARCHAR(255) NOT NULL default '' COMMENT '角色名称',
        description VARCHAR(255) '角色描述',
        create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (id)
    ) COMMENT '角色表';