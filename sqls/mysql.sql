-- 删除客户端表
drop table if exists t_client;
-- 创建客户端表
CREATE TABLE t_client(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键' ,
    ip VARCHAR(32)    COMMENT 'IP' ,
    port VARCHAR(32)    COMMENT 'Port' ,
    vkey VARCHAR(32)    COMMENT '密钥' ,
    info VARCHAR(32)    COMMENT '备注' ,
    zip VARCHAR(32)   DEFAULT 1 COMMENT '压缩 0-不压缩1-压缩' ,
    status VARCHAR(1)   DEFAULT 0 COMMENT '状态 0-无效1-有效' ,
    online VARCHAR(1)   DEFAULT 0 COMMENT '状态 0-离线1-在线' ,
    created_by VARCHAR(32)    COMMENT '创建人' ,
    created_time DATETIME    COMMENT '创建时间' ,
    updated_by VARCHAR(32)    COMMENT '更新人' ,
    updated_time DATETIME    COMMENT '更新时间' ,
    PRIMARY KEY (id)
) COMMENT = '客户端表 ';

ALTER TABLE t_client COMMENT '客户端表';


-- 删除项目表
drop table if exists t_item;
-- 新建项目表
CREATE TABLE t_item(
    id BIGINT NOT NULL AUTO_INCREMENT COMMENT '主键' ,
    client_id BIGINT NOT NULL   COMMENT '客户端ID' ,
    item_name VARCHAR(512)    COMMENT '项目名称' ,
    item_desc VARCHAR(512)    COMMENT '项目描述' ,
    log_path VARCHAR(512)    COMMENT '日志路径' ,
    log_prefix VARCHAR(128)    COMMENT '日志前缀' ,
    log_suffix VARCHAR(128)    COMMENT '日志后缀' ,
    status VARCHAR(1)   DEFAULT 1 COMMENT '状态 0-无效1-有效' ,
    created_by VARCHAR(32)    COMMENT '创建人' ,
    created_time DATETIME    COMMENT '创建时间' ,
    updated_by VARCHAR(32)    COMMENT '更新人' ,
    updated_time DATETIME    COMMENT '更新时间' ,
    PRIMARY KEY (id)
) COMMENT = '项目日志表 ';

ALTER TABLE t_item COMMENT '项目日志表';