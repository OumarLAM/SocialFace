package main

import (
	"log"
	"net/http"

	"github.com/OumarLAM/SocialFace/internal/controllers"
	"github.com/OumarLAM/SocialFace/internal/db/sqlite"
	"github.com/OumarLAM/SocialFace/internal/middlewares"
)

func main() {
	// Connect to the database
	db, err := sqlite.ConnectDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer db.Close()

	// Apply migrations
	if err := sqlite.MigrateDB(db); err != nil {
		log.Fatalf("failed to apply migrations: %v", err)
	}

	// Start background goroutine to clear expired sessions
	go sqlite.ClearExpiredSessions()

	// Initialize router
	router := http.NewServeMux()

	// Register authentication endpoints
	router.HandleFunc("/register", controllers.RegisterHandler)
	router.HandleFunc("/login", controllers.LoginHandler)
	router.HandleFunc("/logout", controllers.LogoutHandler)

	// Register profile endpoints
	router.HandleFunc("/profile/info", middlewares.AuthMiddleware(controllers.ProfileInfoHandler))
	router.HandleFunc("/profile/privacy", middlewares.AuthMiddleware(controllers.UpdateProfilePrivacyHandler))

	// Register user activity endpoints
	router.HandleFunc("/activity/posts", middlewares.AuthMiddleware(controllers.FetchPostsHandler))
	router.HandleFunc("/activity/comments", middlewares.AuthMiddleware(controllers.FetchCommentsHandler))
	router.HandleFunc("/activity/likes", middlewares.AuthMiddleware(controllers.FetchLikesHandler))
	router.HandleFunc("/activity/followers", middlewares.AuthMiddleware(controllers.FetchFollowersHandler))
	router.HandleFunc("/activity/following", middlewares.AuthMiddleware(controllers.FetchFollowingHandler))

	// Register endpoints for users to follow and unfollow other users
	router.HandleFunc("/user/follow", middlewares.AuthMiddleware(controllers.FollowUserHandler))
	router.HandleFunc("/user/unfollow", middlewares.AuthMiddleware(controllers.UnfollowUserHandler))
	router.HandleFunc("/user/acceptFollow", middlewares.AuthMiddleware(controllers.AcceptFollowRequestHandler))
	router.HandleFunc("/user/declineFollow", middlewares.AuthMiddleware(controllers.DeclineFollowRequestHandler))

	// Endpoints to create posts, comments and likes
	router.HandleFunc("/post/create", middlewares.AuthMiddleware(controllers.CreatePostHandler))
	router.HandleFunc("/post/comment", middlewares.AuthMiddleware(controllers.CreateCommentHandler))
	router.HandleFunc("/user/like", middlewares.AuthMiddleware(controllers.LikePostHandler))
	router.HandleFunc("/user/dislike", middlewares.AuthMiddleware(controllers.DislikePostHandler))

	// Endpoints to create groups
	router.HandleFunc("/group/create", middlewares.AuthMiddleware(controllers.CreateGroupHandler))
    // router.HandleFunc("/group/join", middlewares.AuthMiddleware(controllers.JoinGroupHandler))
    // router.HandleFunc("/group/leave", middlewares.AuthMiddleware(controllers.LeaveGroupHandler))
    // router.HandleFunc("/group/invite", middlewares.AuthMiddleware(controllers.InviteUserToGroupHandler))
    // router.HandleFunc("/group/acceptInvite", middlewares.AuthMiddleware(controllers.AcceptGroupInviteHandler))
    // router.HandleFunc("/group/declineInvite", middlewares.AuthMiddleware(controllers.DeclineGroupInviteHandler))

    // // Endpoints to manage groups
    // router.HandleFunc("/group/info", middlewares.AuthMiddleware(controllers.FetchGroupInfoHandler))
    // router.HandleFunc("/group/members", middlewares.AuthMiddleware(controllers.FetchGroupMembersHandler))
	// Start server
	log.Println("Starting server on :8080")
	http.ListenAndServe(":8080", router)
}
