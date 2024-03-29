CREATE TABLE IF NOT EXISTS Comment (
	comment_id INTEGER PRIMARY KEY AUTOINCREMENT,
	user_id INTEGER NOT NULL,
	post_id INTEGER NOT NULL,
	content TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	image_gif TEXT,
	FOREIGN KEY (user_id) REFERENCES User(user_id),
	FOREIGN KEY (post_id) REFERENCES Post(post_id)
);