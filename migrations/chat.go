package migrations

const Chat = `
	CREATE TABLE IF NOT EXISTS Chat (
		chat_id INTEGER PRIMARY KEY AUTOINCREMENT,
		sender_id INTEGER NOT NULL,
		recipient_id INTEGER NOT NULL,
		message TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (sender_id) REFERENCES User(user_id),
		FOREIGN KEY (recipient_id) REFERENCES User(user_id)
	);
`