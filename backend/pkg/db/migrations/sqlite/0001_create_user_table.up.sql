CREATE TABLE IF NOT EXISTS User (
    user_id INTEGER PRIMARY KEY AUTOINCREMENT,
    email TEXT NOT NULL,
    password TEXT NOT NULL,
    firstname TEXT NOT NULL,
    lastname TEXT NOT NULL,
    date_of_birth TEXT NOT NULL,
    avatar_image TEXT,
    nickname TEXT,
    about_me TEXT,
    profile_type TEXT NOT NULL
);