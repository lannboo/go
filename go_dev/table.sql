create table tbl_app_info (
    app_id int auto_increment primary key,
    app_name varchar(1024) not null,
    app_type varchar(64) not null,
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    develop_path VARCHAR(256) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 auto_increment=1;

CREATE TABLE tbl_app_ip (
    app_id INT,
    ip VARCHAR(64),
    Key app_id_ip_index (app_id, ip)
)  ENGINE=InnoDB DEFAULT CHARSET=utf8 auto_increment=1;



CREATE TABLE tbl_log_info (
    log_id int auto_increment primary key,
    app_id VARCHAR(1024) not NULL,
    log_path VARCHAR(64) not null,
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    status TINYINT DEFAULT 1
) ENGINE=InnoDB DEFAULT CHARSET=utf8 auto_increment=1;



mysql > alter table tbl_log_info add column topic VARCHAR(1024) not NULL;