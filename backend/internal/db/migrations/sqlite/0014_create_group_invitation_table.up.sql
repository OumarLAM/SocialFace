CREATE TABLE IF NOT EXISTS GroupInvitation (
    invitation_id INTEGER PRIMARY KEY AUTOINCREMENT,
    group_id INTEGER NOT NULL,
    inviter_id INTEGER NOT NULL,
    invitee_id INTEGER NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending',
    FOREIGN KEY (group_id) REFERENCES UserGroup(group_id),
    FOREIGN KEY (inviter_id) REFERENCES User(user_id),
    FOREIGN KEY (invitee_id) REFERENCES User(user_id)
);
