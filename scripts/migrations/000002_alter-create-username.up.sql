ALTER TABLE users
ADD username varchar(100);

ALTER TABLE users
ADD CONSTRAINT unique_username unique (username);
