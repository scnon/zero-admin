CREATE TABLE
    sys_role (
        id BIGINT AUTO_INCREMENT COMMENT '角色ID',
        name VARCHAR(64) NOT NULL default '' COMMENT '角色名称',
        sort TINYINT NOT NULL DEFAULT 0 COMMENT '排序',
        status TINYINT NOT NULL DEFAULT 0 COMMENT '状态(0:启用,1:禁用)',
        creator BIGINT NOT NULL DEFAULT 0 COMMENT '创建人 user_id',
        updater BIGINT NOT NULL DEFAULT 0 COMMENT '修改人 user_id',
        remark VARCHAR(512) NOT NULL DEFAULT '' COMMENT '备注',
        tenant_id BIGINT NOT NULL DEFAULT 0 COMMENT '租户ID',
        create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL COMMENT '创建时间',
        update_time TIMESTAMP DEFAULT NULL ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (id)
    ) COMMENT '角色表';