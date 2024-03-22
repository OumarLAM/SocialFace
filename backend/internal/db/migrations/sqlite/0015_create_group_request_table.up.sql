CREATE TABLE IF NOT EXISTS GroupRequest (
    request_id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id INTEGER NOT NULL,
    group_id INTEGER NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending',
    FOREIGN KEY (user_id) REFERENCES User(user_id),
    FOREIGN KEY (group_id) REFERENCES UserGroup(group_id)
);