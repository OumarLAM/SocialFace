package migrations

const Post = `
	CREATE TABLE IF NOT EXISTS Post (
		post_id INTEGER PRIMARY KEY AUTOINCREMENT,
		user_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		privacy TEXT NOT NULL,
		image_gif TEXT,
		FOREIGN KEY (user_id) REFERENCES User(user_id)
	);
`