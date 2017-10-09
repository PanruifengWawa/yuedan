create database if not exists yuedan;

use yuedan;
drop table if exists t_user;
create table t_user(
	id serial primary key,
	user_name varchar(16) not null unique,
	password varchar(64) not null
)engine=innodb default charset=utf8;

drop table if exists t_token;
create table t_token(
	id serial primary key,
	token_str char(36) not null unique,
	user_id bigint(20) unsigned not null unique,
	login_date timestamp not null,
	user_type tinyint(1) not null default 1,
	foreign key(user_id) references t_user(id) on delete cascade
)engine=innodb default charset=utf8;