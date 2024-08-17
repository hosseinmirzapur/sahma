package config

import "os"

var app_root = os.Getenv("APP_ROOT") + "/storage"

func GetStoragePath(disk string) string {
	switch disk {
	case "local":
		return app_root
	case "image":
		return app_root + "/image"
	case "word":
		return app_root + "/word"
	case "csv":
		return app_root + "/csv"
	case "voice":
		return app_root + "/voice"
	case "video":
		return app_root + "/video"
	case "pdf":
		return app_root + "/pdf"
	case "zip":
		return app_root + "/zip"
	case "excel":
		return app_root + "/excel"
	case "public":
		return app_root
	default:
		return app_root
	}
}
