package models

import (
	"database/sql"
	"fmt"

	"github.com/OumarLAM/SocialFace/internal/db/sqlite"
)

type GroupInvitation struct {
	InvitationID int    `json:"invitation_id"`
	GroupID      int    `json:"group_id"`
	InviterID    int    `json:"inviter_id"`
	InviteeID    int    `json:"invitee_id"`
	Status       string `json:"status"`
}

func SendGroupInvitation(groupID, inviterID, inviteeID int) error {
	isGroupMember, err := IsGroupMember(inviterID, groupID)
	if err != nil {
        return fmt.Errorf("failed to check group member")
    }
	if !isGroupMember {
        return fmt.Errorf("you must be a member of the group to send invitation")
    }

	db, err := sqlite.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("INSERT INTO GroupInvitation (group_id, inviter_id, invitee_id, status) VALUES (?, ?, ?, ?)", groupID, inviterID, inviteeID, "pending")
	if err != nil {
		return fmt.Errorf("failed to send group invitation: %v", err)
	}

	return nil
}

func AcceptGroupInvitation(invitationID, inviteeID int) error {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("UPDATE GroupInvitation SET status = 'accepted' WHERE invitation_id = ? AND invitee_id = ?", invitationID, inviteeID)
	if err != nil {
		return fmt.Errorf("failed to accept group invitation: %v", err)
	}

	invitation := GroupInvitation{InvitationID: invitationID}
	groupID, err := GetGroupIDByInvitation(invitation)
	if err != nil {
		return fmt.Errorf("failed to get group ID for group invitation: %v", err)
	}

	err = AddGroupMember(groupID, inviteeID)
	if err != nil {
		return fmt.Errorf("failed to add user to group: %v", err)
	}

	return nil
}

func DeclineGroupInvitation(invitationID, inviteeID int) error {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	_, err = db.Exec("DELETE FROM GroupInvitation WHERE invitation_id = ? AND invitee_id = ?", invitationID, inviteeID)
	if err != nil {
		return fmt.Errorf("failed to decline group invitation: %v", err)
	}

	return nil
}

func GetGroupIDByInvitation(invitation GroupInvitation) (int, error) {
	db, err := sqlite.ConnectDB()
	if err != nil {
		return 0, fmt.Errorf("failed to connect to database: %v", err)
	}
	defer db.Close()

	var groupID int
	err = db.QueryRow("SELECT group_id FROM GroupInvitation WHERE invitation_id = ?", invitation.InvitationID).Scan(&groupID)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, fmt.Errorf("invitation not found: %v", err)
		}
		return 0, fmt.Errorf("failed to get group ID: %v", err)
	}

	return groupID, nil
}
