package migrations

const Group = `
	CREATE TABLE IF NOT EXISTS Group (
		group_id INTEGER PRIMARY KEY AUTOINCREMENT,
		creator_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		FOREIGN KEY (creator_id) REFERENCES User(user_id)
	);
`