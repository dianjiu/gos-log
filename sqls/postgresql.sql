-- 删除客户端表
drop table if exists t_client;
-- 创建客户端表
create table t_client(
    id bigserial not null, --主键
    ip varchar(32) ,  --服务
    port varchar(32)    , --端口
    vkey varchar(64)    , --密钥
    info text    ,  --备注
    zip varchar(32)   default '1' , --压缩 0-默认 1-压缩
    online varchar(2)   default '0' , -- 在线 0-离线 1-在线
    status varchar(2)   default '1' , -- 状态 0-无效 1-有效
    created_by varchar(32)    ,-- 创建人
    created_time timestamptz default now() not null ,--创建时间
    updated_by varchar(32)   ,-- 更新人
    updated_time timestamptz default now() not null ,--更新时间
    constraint pk_t_client_id primary key (id) --设置主键
) ;
-- 备注客户端表
comment on table t_client is '客户端表';
comment on column t_client.id is '主键';
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
comment on constraint pk_t_client_id on t_client is '主键,唯一约束';

-- 删除项目表
drop table if exists t_item;
-- 新建项目表
create table t_item(
    id bigserial not null , -- '主键'
    client_id bigint not null  , --'客户端id'
    item_name varchar(512)   ,-- '项目名称'
    item_desc varchar(512)   ,--  '项目描述' 
    log_path varchar(512) not null   ,--  '日志路径' 
    log_prefix varchar(128)     ,--  '日志前缀'
    log_suffix varchar(128)    ,--  '日志后缀' 
    status varchar(1)   default '1' ,--  '状态 0-无效1-有效' 
    created_by varchar(32)    ,--  '创建人' 
    created_time timestamptz default now() not null    , --  '创建时间'
    updated_by varchar(32)    ,--  '更新人' 
    updated_time timestamptz default now() not null    ,--  '更新时间' 
    constraint pk_t_item_id primary key (id) --设置主键
) ;
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
comment on constraint pk_t_item_id on t_item is '主键,唯一约束';

