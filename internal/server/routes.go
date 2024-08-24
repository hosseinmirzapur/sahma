package server

import (
	"sahma/internal/handlers"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// Register user web routes
	userWebRoutes(r)
}

func userWebRoutes(r *gin.Engine) {
	// Register guest routes
	guest := r.Group("/guest")
	guestRoutes(guest)

	// Register auth routes
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
	r.GET("/", handlers.NotificationHandler().Index)
	r.POST("/create", handlers.NotificationHandler().CreateAction)
}

func departmentRoutes(r *gin.RouterGroup) {
	r.GET("/list", handlers.DepartmentHandler().Index)
	r.POST("/create", handlers.DepartmentHandler().Create)
	r.POST("/edit/:department", handlers.DepartmentHandler().Edit)
	r.POST("/delete/:department", handlers.DepartmentHandler().Delete)
}

func reportRoutes(r *gin.RouterGroup) {
	r.GET("/users", handlers.ReportHandler().UsersReport)
	r.GET("/create-excel-users", handlers.ReportHandler().CreateExcelUserReport)
	r.GET("/total-uploaded-files", handlers.ReportHandler().TotalUploadedFiles)
	r.GET("/total-uploaded-files-type", handlers.ReportHandler().TotalUploadedFilesByType)
	r.GET("/total-uploaded-transcribed-files", handlers.ReportHandler().TotalTranscribedFiles)
}

func dashboardRoutes(r *gin.RouterGroup) {
	r.GET("/", handlers.DashboardHandler().Index)
	r.POST("/copy", handlers.DashboardHandler().Copy)
	r.POST("/move", handlers.DashboardHandler().Move)
	r.POST("/delete", handlers.DashboardHandler().PermanentDelete)
	r.GET("/trash", handlers.DashboardHandler().TrashList)
	r.POST("/trash", handlers.DashboardHandler().TrashAction)
	r.POST("/trash-retrieve", handlers.DashboardHandler().TrashRetrieve)
	r.GET("/archive", handlers.DashboardHandler().ArchiveList)
	r.POST("/archive", handlers.DashboardHandler().ArchiveAction)
	r.POST("/archive-retrieve", handlers.DashboardHandler().ArchiveRetrieve)
	r.POST("/create-zip", handlers.DashboardHandler().CreateZip)
	r.GET("/search", handlers.DashboardHandler().SearchForm)
	r.POST("/search", handlers.DashboardHandler().SearchAction)

	// Register folder routes
	folder := r.Group("/folder")
	folderRoutes(folder)

	// Register file routes
	file := r.Group("/file")
	fileRoutes(file)
}

func folderRoutes(r *gin.RouterGroup) {
	r.GET("/show/:folderID", handlers.FolderHandler().Show)
	r.POST("/create-root", handlers.FolderHandler().CreateRoot)
	r.POST("/create/:folderID", handlers.FolderHandler().Create)
	r.POST("/rename/:folderID", handlers.FolderHandler().Rename)
}

func fileRoutes(r *gin.RouterGroup) {
	r.GET("/show/:fileID", handlers.FileHandler().Show)
	r.POST("/add-description/:fileID", handlers.FileHandler().AddDescription)
	r.POST("/transcribe-file/:fileID", handlers.FileHandler().Transcribe)
	r.GET("/download-original-file/:fileID", handlers.FileHandler().DownloadOriginalFile)
	r.GET("/download-searchable-file/:fileID", handlers.FileHandler().DownloadSearchableFile)
	r.GET("/download-word-file/:fileID", handlers.FileHandler().DownloadWordFile)
	r.POST("/rename/:fileID", handlers.FileHandler().Rename)
	r.GET("/print/:fileID", handlers.FileHandler().PrintOriginalFile)
	r.POST("/upload/:folderID", handlers.FileHandler().Upload)
	r.GET("/upload-root/:fileID", handlers.FileHandler().UploadRoot)
}

func userManagementRoutes(r *gin.RouterGroup) {
	r.GET("/", handlers.UserHandler().Index)
	r.GET("/:user", handlers.UserHandler().UserInfo)
	r.POST("/create-user", handlers.UserHandler().Create)
	r.POST("/delete-user/:user", handlers.UserHandler().Block)
	r.POST("/edit-user/:user", handlers.UserHandler().Edit)
	r.POST("/search", handlers.UserHandler().Search)
}

func apiRoutes(r *gin.RouterGroup) {

}
