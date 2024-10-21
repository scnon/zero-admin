CREATE TABLE
    product (
        id INT AUTO_INCREMENT NOT NULL COMMENT '商品ID',
        store_id INT NOT NULL COMMENT '店铺ID',
        cate_id INT NOT NULL COMMENT '分类ID',
        name VARCHAR(255) NOT NULL COMMENT '商品名称',
        price DECIMAL(10, 2) NOT NULL COMMENT '价格',
        image VARCHAR(255) NOT NULL COMMENT '图片',
        unit VARCHAR(20) NOT NULL COMMENT '单位',
        stock INT NOT NULL COMMENT '库存',
        create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
        update_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (id)
    ) COMMENT '商品表';