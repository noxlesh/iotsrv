-- mysql 8.0 scheme
create database if not exists iot;
use iot;

drop tables if exists humidity, temperature, device;

create table if not exists device (
    id bigint auto_increment,
    name varchar(30) not null unique,
    created_at datetime not null,
    primary key(id)
);

create table if not exists temperature (
    id bigint auto_increment,
    device_id bigint not null,
    value double not null,
    created_at datetime not null,
    primary key(id),
    foreign key(device_id) references device (id) on delete cascade
);

create table if not exists humidity (
    id bigint auto_increment,
    device_id bigint not null,
    value double not null,
    created_at datetime not null,
    primary key(id),
	foreign key(device_id) references device (id) on delete cascade
);