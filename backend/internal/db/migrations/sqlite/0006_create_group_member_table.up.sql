CREATE TABLE IF NOT EXISTS GroupMember (
	group_id INTEGER NOT NULL,
	user_id INTEGER NOT NULL,
	FOREIGN KEY (group_id) REFERENCES UserGroup(group_id),
	FOREIGN KEY (user_id) REFERENCES User(user_id)
);