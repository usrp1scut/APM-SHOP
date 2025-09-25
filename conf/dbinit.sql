create database ordersvc;
use ordersvc;
CREATE TABLE T_ORDER (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    order_id VARCHAR(255) NOT NULL,
    ctime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    utime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    sku_id BIGINT NOT NULL,
    num INT NOT NULL,
    price INT NOT NULL,
    uid BIGINT NOT NULL,
    constraint order_pk2
    	unique (order_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
create database skusvc;
use skusvc;
create table t_sku (
	id BIGINT AUTO_INCREMENT PRIMARY KEY,
	name varchar(10) not null,
	price INT NOT NULL,
	ctime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    utime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    num int null
);
create database usrsvc;
use usrsvc;
create table t_user (
	id BIGINT AUTO_INCREMENT PRIMARY KEY,
	name varchar(10) not null,
	ctime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    utime TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);