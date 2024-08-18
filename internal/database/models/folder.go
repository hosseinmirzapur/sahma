package models

import (
	"encoding/base64"
	"fmt"
	"os"
	"sahma/internal/globals"
	"sahma/internal/helper"
	"slices"
	"strconv"
	"time"

	"gorm.io/gorm"
)

type Folder struct {
	ID             uint
	Name           string
	UserID         *uint
	ParentFolderID *uint
	DeletedAt      *string
	Meta           *string
	ArchivedAt     *string
	Slug           *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Activities     []*Activity `gorm:"polymorphic:Activity"`
	EntityGroups   []*EntityGroup
	User           *User
}

func CreateFolderWithSlug(folder *Folder) error {
	err := globals.GetDB().Transaction(func(tx *gorm.DB) error {
		globals.GetDB().Create(&folder)

		slug, err := folder.GetFolderID()
		if err != nil {
			return err
		}
		folder.Slug = slug
		return globals.GetDB().Save(&folder).Error
	})

	return err
}

func (folder *Folder) GetFolderID() (*string, error) {
	// converting folder id to string
	idStr := strconv.Itoa(int(folder.ID))

	/**
	This string operation does the same as
	str_pad((string)$this->id, 12, '0', STR_PAD_LEFT)
	which php does
	*/
	paddedID := fmt.Sprintf("%012s", idStr)

	// This will encrypt the paddedID with the 32 bytes app key in .env
	appKey := os.Getenv("APP_KEY")
	encryptedID, err := helper.Encrypt(paddedID, appKey)
	if err != nil {
		return nil, err
	}

	// returning the base64 encoded version of the encryptedID
	base64EncodedStr := base64.StdEncoding.EncodeToString([]byte(encryptedID))
	return &base64EncodedStr, nil
}

func (folder *Folder) ConvertObfuscatedIdToFolderId(obfuscated string) (*uint, error) {
	// base64_decode the string provided
	base64Decoded, err := base64.StdEncoding.DecodeString(obfuscated)
	if err != nil {
		return nil, err
	}

	// decrypt the base64 decoded string using app key
	appKey := os.Getenv("APP_KEY")
	decryptedID, err := helper.Decrypt(string(base64Decoded), appKey)
	if err != nil {
		return nil, err
	}

	// convert the decrypted string to int
	intID, err := strconv.Atoi(decryptedID)
	if err != nil {
		return nil, err
	}

	uintID := uint(intID)
	return &uintID, nil
}

func (f *Folder) ParentFolder() (*Folder, error) {
	var folder *Folder
	err := globals.GetDB().Where("id = ?", f.ParentFolderID).First(&folder).Error
	if err != nil {
		return nil, err
	}

	return folder, nil
}

func (f *Folder) SubFolders(breadcrumbs []uint, currentFolderID *int) ([]map[string]interface{}, error) {
	var folders []Folder
	err := globals.
		GetDB().
		Preload("SubFolders").
		Where("parent_folder_id = ?", f.ID).
		Where("deleted_at = NULL").
		Find(&folders).
		Error
	if err != nil {
		return nil, err
	}

	result := make([]map[string]interface{}, 0)

	for _, folder := range folders {
		subFolders, err := folder.SubFolders(breadcrumbs, currentFolderID)
		if err != nil {
			return nil, err
		}

		slug, err := folder.GetFolderID()
		if err != nil {
			return nil, err
		}

		isOpen := slices.Contains(
			breadcrumbs, folder.ID) ||
			(currentFolderID != nil && folder.ID == uint(*currentFolderID))
		result = append(result, map[string]interface{}{
			"id":               folder.ID,
			"name":             folder.Name,
			"parent_folder_id": folder.ParentFolderID,
			"slug":             *slug,
			"subFolders":       subFolders,
			"isOpen":           isOpen,
		})
	}

	return result, nil
}

func (f *Folder) TempDeleteSubFoldersAndFiles(folder Folder, user User) error {
	var folders []Folder
	err := globals.
		GetDB().
		Preload("User").
		Preload("SubFolders").
		Where("parent_folder_id = ?", folder.ID).
		Find(&folders).
		Error
	if err != nil {
		return err
	}

	return nil // todo: complete this method
}
