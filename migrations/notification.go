package migrations

const Notification = `
	CREATE TABLE IF NOT EXISTS Notification (
		notification_id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (user_id) REFERENCES User(user_id)
	);
`