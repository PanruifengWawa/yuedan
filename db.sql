create database if not exists yuedan;

use yuedan;
drop table if exists t_user;
create table t_user(
	id serial primary key,
	user_name varchar(16) not null unique,
	password varchar(64) not null
)engine=innodb default charset=utf8;