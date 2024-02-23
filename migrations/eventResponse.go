package migrations

const EventResponse = `
	CREATE TABLE IF NOT EXISTS EventResponse (
		event_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		response TEXT NOT NULL,
		FOREIGN KEY (event_id) REFERENCES Event(event_id),
		FOREIGN KEY (user_id) REFERENCES User(user_id)
	);
`