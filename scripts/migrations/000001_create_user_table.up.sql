CREATE TABLE IF NOT EXISTS users (
 id BIGINT auto_increment primary key,
 email varchar(250) not null unique ,
 password varchar(500) not null ,
 created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
 updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
 created_by longtext not null ,
 updated_by longtext not null
) engine = innoDB;