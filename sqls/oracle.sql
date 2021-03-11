-- 删除表
drop table if exists t_client;

-- 创建表
CREATE TABLE t_client(
    id NUMBER(11) PRIMARY KEY,
    ip VARCHAR2(32),
    port VARCHAR2(32),
    vkey VARCHAR2(32),
    info VARCHAR2(32),
    zip VARCHAR2(32),
    online VARCHAR2(1),
    status VARCHAR2(1),
    created_by VARCHAR2(32),
    created_time DATE,
    updated_by VARCHAR2(32),
    updated_time DATE
);

-- 创建序列
CREATE SEQUENCE CLIENT_ID_SEQ
INCREMENT BY 1
START WITH 100
MAXVALUE 999999999
NOCYCLE
NOCACHE;
-- 创建触发器
create or replace trigger CLIENT_ID_SEQ_TRG
before
insert on t_client
for each row
 begin
    select CLIENT_ID_SEQ.nextval into :new.id from dual;
 end;

-- 备注客户端表
comment on table t_client is '客户端表';
comment on column  t_client.id is '主键';
comment on column t_client.ip is '服务';
comment on column t_client.port is '端口';
comment on column t_client.vkey is '密钥';
comment on column t_client.info is '备注';
comment on column t_client.zip is '压缩 0-默认 1-压缩';
comment on column t_client.online is '在线 0-离线 1-在线';
comment on column t_client.status is '状态 0-无效 1-有效';
comment on column t_client.created_by is '创建人';
comment on column t_client.created_time is '创建时间';
comment on column t_client.updated_by is '更新人';
comment on column t_client.updated_time is '更新时间';

-- 删除项目表
drop table if exists t_item;

-- 创建表
CREATE TABLE t_item(
    id NUMBER(11) PRIMARY KEY,
    client_id NUMBER NOT NULL,
    item_name VARCHAR2(512),
    item_desc VARCHAR2(512),
    log_path VARCHAR2(512),
    log_prefix VARCHAR2(128),
    log_suffix VARCHAR2(128),
    status VARCHAR2(1),
    created_by VARCHAR2(32),
    created_time DATE,
    updated_by VARCHAR2(32),
    updated_time DATE
);

-- 创建序列
CREATE SEQUENCE ITEM_ID_SEQ
INCREMENT BY 1
START WITH 100
MAXVALUE 999999999
NOCYCLE
NOCACHE;

-- 创建触发器
create or replace trigger ITEM_ID_SEQ_TRG
before
insert on t_item
for each row
 begin
    select ITEM_ID_SEQ.nextval into :new.id from dual;
 end;

-- 备注项目表
comment on table t_item is '项目表';
comment on column t_item.id  is '主键';
comment on column t_item.client_id  is '客户端编码';
comment on column t_item.item_name  is '项目名称';
comment on column t_item.item_desc  is '项目描述';
comment on column t_item.log_path  is '日志路径';
comment on column t_item.log_prefix  is '日志前缀';
comment on column t_item.log_suffix  is '日志后缀';
comment on column t_item.status is '状态 0-无效 1-有效';
comment on column t_item.created_by is '创建人';
comment on column t_item.created_time is '创建时间';
comment on column t_item.updated_by is '更新人';
comment on column t_item.updated_time is '更新时间';



