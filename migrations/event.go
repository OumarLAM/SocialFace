package migrations

const Event = `
	CREATE TABLE IF NOT EXISTS Event (
		event_id INTEGER PRIMARY KEY AUTOINCREMENT,
		group_id INTEGER NOT NULL,
		title TEXT NOT NULL,
		description TEXT NOT NULL,
		date_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (group_id) REFERENCES Group(group_id)
	);
`