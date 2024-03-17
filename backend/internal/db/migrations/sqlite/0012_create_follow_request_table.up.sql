CREATE TABLE IF NOT EXISTS FollowRequest (
    request_id INTEGER PRIMARY KEY AUTOINCREMENT,
    follower_id INTEGER NOT NULL,
    followee_id INTEGER NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (follower_id) REFERENCES User(user_id),
    FOREIGN KEY (followee_id) REFERENCES User(user_id)
);