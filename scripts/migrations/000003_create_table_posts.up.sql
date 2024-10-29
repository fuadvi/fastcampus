CREATE TABLE IF NOT EXISTS posts(
    id BIGINT auto_increment primary key,
    user_id BIGINT NOT NULL,
    post_title VARCHAR(250) NOT NULL,
    post_content LONGTEXT NOT NULL,
    post_hash_tag    LONGTEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by longtext not null ,
    updated_by longtext not null
);
