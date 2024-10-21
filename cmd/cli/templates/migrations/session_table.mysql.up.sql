ALTER TABLE `sessions` DROP FOREIGN KEY `sessions_expiry_idx`;
DROP TABLE IF EXISTS sessions;

CREATE TABLE sessions (
	token CHAR(43) PRIMARY KEY,
	data BLOB NOT NULL,
	expiry TIMESTAMP(6) NOT NULL
);

CREATE INDEX sessions_expiry_idx ON sessions (expiry);
