CREATE TABLE IF NOT EXISTS refresh_tokens(
      id BIGINT auto_increment primary key,
      user_id BIGINT NOT NULL,
      refresh_token TEXT NOT NULL,
      expired_at TIMESTAMP NOT NULL,
      created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
      created_by longtext not null ,
      updated_by longtext not null,
      CONSTRAINT fk_user_id_user_refresh_token FOREIGN KEY (user_id) REFERENCES users(id)

);
