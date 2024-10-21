CREATE TABLE
    sys_menu (
        id BIGINT AUTO_INCREMENT COMMENT '菜单ID',
        parent_id BIGINT NOT NULL DEFAULT 0 COMMENT '父ID',
        tenant_id BIGINT NOT NULL DEFAULT 0 COMMENT '租户ID',
        name VARCHAR(64) NOT NULL DEFAULT '' COMMENT '名称',
        icon VARCHAR(128) NOT NULL DEFAULT '' COMMENT '图标',
        path VARCHAR(255) NOT NULL DEFAULT '' COMMENT '路径',
        title VARCHAR(255) NOT NULL DEFAULT '' COMMENT '标题',
        sort TINYINT NOT NULL DEFAULT 0 COMMENT '排序',
        type TINYINT NOT NULL DEFAULT 0 COMMENT '类型(1:菜单,2:按钮,3:外链)',
        permission VARCHAR(255) NOT NULL DEFAULT '' COMMENT '权限标识',
        status TINYINT NOT NULL DEFAULT 0 COMMENT '状态(0:启用,1:禁用)',
        visible TINYINT NOT NULL DEFAULT 0 COMMENT '是否隐藏(0:隐藏,1:显示)',
        creator BIGINT NOT NULL DEFAULT 0 COMMENT '创建人',
        updater BIGINT NOT NULL DEFAULT 0 COMMENT '修改人',
        create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
        update_time TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
        PRIMARY KEY (id)
    ) COMMENT '菜单表';