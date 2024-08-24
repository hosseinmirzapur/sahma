package server

import (
	"sahma/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// register user web routes
	userWebRoutes(r)
}

func userWebRoutes(r *gin.Engine) {
	// register guest routes
	guest := r.Group("/guest")
	guestRoutes(guest)

	// register auth routes
	auth := r.Group("/auth")
	authRoutes(auth)
}

func guestRoutes(r *gin.RouterGroup) {
	r.GET("/", handlers.AuthHandler().Index)
	r.GET("/login", handlers.AuthHandler().LoginPage)
	r.POST("/login", handlers.AuthHandler().LoginAction)
	r.POST("/logout", handlers.AuthHandler().Logout)
}

func authRoutes(r *gin.RouterGroup) {
	// Register profile routes
	profile := r.Group("/profile")
	profileRoutes(profile)

	// Register cartable routes
	cartable := r.Group("/cartable")
	cartableRoutes(cartable)

	// Register notification routes
	notification := r.Group("/notifications")
	notificationRoutes(notification)

	// Register department routes
	department := r.Group("/department")
	departmentRoutes(department)

	// Register report routes
	report := r.Group("/report")
	reportRoutes(report)

	// Register dashboard routes
	dashboard := r.Group("/dashboard")
	dashboardRoutes(dashboard)

	// Register user management routes
	userManagement := r.Group("/user-management")
	userManagementRoutes(userManagement)

	// Register api routes
	api := r.Group("/api")
	apiRoutes(api)

}

func profileRoutes(r *gin.RouterGroup) {
	r.GET("/", handlers.ProfileHandler().Show)
}

func cartableRoutes(r *gin.RouterGroup) {
	r.GET("/inbox-list", handlers.LetterHandler().Inbox)
	r.GET("/draft-list", handlers.LetterHandler().GetDraftedLetters)
	r.GET("/submit-list", handlers.LetterHandler().GetSubmittedLetters)
	r.GET("/deleted-list", handlers.LetterHandler().GetDeletedLetters)
	r.GET("/archived-list", handlers.LetterHandler().GetArchivedLetters)
	r.GET("/submit-form", handlers.LetterHandler().SubmitForm)
	r.POST("/submit-action", handlers.LetterHandler().SubmitAction)
	r.GET("/show/:letter", handlers.LetterHandler().Show)
	r.POST("/sign/:letter", handlers.LetterHandler().SignAction)
	r.POST("/refer/:letter", handlers.LetterHandler().ReferAction)
	r.POST("/reply/:letter", handlers.LetterHandler().ReplyAction)
	r.GET("/download-attachment/:letterAttachment", handlers.LetterHandler().DownloadAttachment)
	r.POST("/draft", handlers.LetterHandler().DraftAction)
	r.GET("/show-draft/:letter", handlers.LetterHandler().ShowDrafted)
	r.POST("/submit-draft/:letter", handlers.LetterHandler().SubmitDrafted)
	r.POST("/archive", handlers.LetterHandler().Archive)
	r.POST("/temp-delete", handlers.LetterHandler().TempDelete)
	r.POST("/submit-reminder/:letter", handlers.LetterHandler().SubmitNotification)
}

func notificationRoutes(r *gin.RouterGroup) {

}

func departmentRoutes(r *gin.RouterGroup) {

}

func reportRoutes(r *gin.RouterGroup) {

}

func dashboardRoutes(r *gin.RouterGroup) {

}

func userManagementRoutes(r *gin.RouterGroup) {

}

func apiRoutes(r *gin.RouterGroup) {

}
