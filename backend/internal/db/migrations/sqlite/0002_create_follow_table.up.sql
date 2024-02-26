CREATE TABLE IF NOT EXISTS Follow (
	follower_id INTEGER NOT NULL,
	followee_id INTEGER NOT NULL,
	FOREIGN KEY (follower_id) REFERENCES User(user_id),
	FOREIGN KEY (followee_id) REFERENCES User(user_id)
);