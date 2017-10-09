drop database yuedan;
create database if not exists yuedan;

use yuedan;
drop table if exists user;
create table user(
	id serial primary key,
	student_id varchar(16) not null unique,
	student_password varchar(64) not null,
	password varchar(64) not null,
	phone varchar(11) not null,
	photo varchar(256),
	nickname varchar(64),
	real_name varchar(64),
	sex char(1),
	birth date,
	email varchar(64),
	school varchar(64),
	school_part varchar(64),
	college varchar(64),
	dormitory varchar(64),
	specialty varchar(64),
	education varchar(64),
	school_year char(4),
	identity varchar(20),
	hobby varchar(512),
	is_single char(1),
	ancestral_home varchar(64),
	lvl int not null default 0,
	exp int not null default 0,
	user_type tinyint(1) not null default 1
)engine=innodb default charset=utf8;

drop table if exists token;
create table token(
	id serial primary key,
	token_str char(36) not null unique,
	user_id bigint(20) unsigned not null unique,
	login_date timestamp not null,
	user_type tinyint(1) not null default 1,
	foreign key(user_id) references user(id) on delete cascade
)engine=innodb default charset=utf8;

