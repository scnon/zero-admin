CREATE TABLE
    IF NOT EXISTS sys_user (
        id BIGINT AUTO_INCREMENT NOT NULL COMMENT '主键',
        username VARCHAR(32) NOT NULL DEFAULT '' COMMENT '用户名',
        password VARCHAR(128) NOT NULL DEFAULT '' COMMENT '密码',
        nickname VARCHAR(32) DEFAULT '' NOT NULL COMMENT '昵称',
        avatar VARCHAR(128) DEFAULT '' NOT NULL COMMENT '头像',
        status TINYINT DEFAULT 1 NOT NULL DEFAULT 0 COMMENT '状态(1:正常,0:禁用)',
        sort TINYINT DEFAULT 1 NOT NULL DEFAULT 0 COMMENT '排序',
        remark VARCHAR(512) NOT NULL DEFAULT '' COMMENT '备注',
        department_id BIGINT NOT NULL DEFAULT 0 COMMENT '部门ID',
        tenant_id BIGINT NOT NULL DEFAULT 0 COMMENT '租户ID',
        creator BIGINT NOT NULL DEFAULT 0 COMMENT '创建人 user_id',
        updater BIGINT NOT NULL DEFAULT 0 COMMENT '修改人 user_id',
        create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间',
        update_time TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
        is_deleted TINYINT DEFAULT 0 NOT NULL COMMENT '是否删除(1:已删除,0:未删除)',
        CONSTRAINT AK_username UNIQUE (username),
        PRIMARY KEY (id)
    ) COMMENT '用户信息';