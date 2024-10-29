CREATE TABLE IF NOT EXISTS comments(
    id BIGINT auto_increment primary key,
    user_id BIGINT NOT NULL,
    post_id BIGINT NOT NULL,
    comment_content LONGTEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by longtext not null ,
    updated_by longtext not null,
    CONSTRAINT fk_post_id_comments FOREIGN KEY (post_id) REFERENCES posts(id),
    CONSTRAINT fk_user_id_comments FOREIGN KEY (user_id) REFERENCES users(id)

);