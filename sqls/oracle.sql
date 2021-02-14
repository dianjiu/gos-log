DROP TABLE t_client;

CREATE TABLE t_client(
    id NUMBER DEFAULT 1000 NOT NULL,
    ip NVARCHAR2(32),
    port NVARCHAR2(32),
    vkey NVARCHAR2(32),
    desc NVARCHAR2(32),
    zip NVARCHAR2(32) DEFAULT 1,
    status VARCHAR2(1) DEFAULT 0,
    created_by VARCHAR2(32),
    created_time DATE,
    updated_by VARCHAR2(32),
    updated_time DATE,
    PRIMARY KEY (id)
);

COMMENT ON TABLE t_client IS '客户端表';
COMMENT ON COLUMN t_client.id IS '主键';
COMMENT ON COLUMN t_client.ip IS 'IP';
COMMENT ON COLUMN t_client.port IS 'Port';
COMMENT ON COLUMN t_client.vkey IS '密钥';
COMMENT ON COLUMN t_client.desc IS '备注';
COMMENT ON COLUMN t_client.zip IS '0-不压缩1-压缩';
COMMENT ON COLUMN t_client.status IS '0-无效1-有效';
COMMENT ON COLUMN t_client.created_by IS '创建人';
COMMENT ON COLUMN t_client.created_time IS '创建时间';
COMMENT ON COLUMN t_client.updated_by IS '更新人';
COMMENT ON COLUMN t_client.updated_time IS '更新时间';

COMMENT ON TABLE t_client IS '客户端表';

DROP TABLE t_item;

CREATE TABLE t_item(
    id NUMBER DEFAULT 1000 NOT NULL,
    client_id NUMBER NOT NULL,
    item_name NVARCHAR2(512),
    item_desc NVARCHAR2(512),
    log_path NVARCHAR2(512),
    log_prefix NVARCHAR2(128),
    log_suffix NVARCHAR2(128),
    status VARCHAR2(1) DEFAULT 1,
    created_by VARCHAR2(32),
    created_time DATE,
    updated_by VARCHAR2(32),
    updated_time DATE,
    PRIMARY KEY (id)
);

COMMENT ON TABLE t_item IS '项目日志表';
COMMENT ON COLUMN t_item.id IS '主键';
COMMENT ON COLUMN t_item.client_id IS '客户端ID';
COMMENT ON COLUMN t_item.item_name IS '项目名称';
COMMENT ON COLUMN t_item.item_desc IS '项目描述';
COMMENT ON COLUMN t_item.log_path IS '日志路径';
COMMENT ON COLUMN t_item.log_prefix IS '日志前缀';
COMMENT ON COLUMN t_item.log_suffix IS '日志后缀';
COMMENT ON COLUMN t_item.status IS '0-无效1-有效';
COMMENT ON COLUMN t_item.created_by IS '创建人';
COMMENT ON COLUMN t_item.created_time IS '创建时间';
COMMENT ON COLUMN t_item.updated_by IS '更新人';
COMMENT ON COLUMN t_item.updated_time IS '更新时间';

CREATE INDEX index_client_id ON t_item(client_id);

COMMENT ON TABLE t_item IS '项目日志表';

