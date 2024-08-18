package models

import (
	"encoding/base64"
	"fmt"
	"os"
	"sahma/internal/config"
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
	DeletedAt      *time.Time
	DeletedBy      *uint
	Meta           *string
	ArchivedAt     *string
	Slug           *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Activities     []*Activity `gorm:"polymorphic:Activity"`
	EntityGroups   []*EntityGroup
	User           *User
}

type BreadCrumb struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
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
	// Fetch all sub folders
	var subFolders []Folder
	err := globals.
		GetDB().
		Where("parent_folder_id = ?", folder.ID).
		Find(&subFolders).
		Error
	if err != nil {
		return err
	}

	// This should be a transaction to abide ACID rules
	now := time.Now()
	err = globals.GetDB().Transaction(func(tx *gorm.DB) error {
		for _, sf := range subFolders {
			sf.DeletedAt = &now
			sf.DeletedBy = &user.ID
			err = globals.GetDB().Save(&sf).Error
			if err != nil {
				return err
			}
			err = f.TempDeleteSubFoldersAndFiles(sf, user)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return err
	}

	// Fetch all entity groups
	var egs []EntityGroup
	err = globals.
		GetDB().
		Where("parent_folder_id = ?", folder.ID).
		Find(&egs).
		Error
	if err != nil {
		return err
	}

	// This DB operation must also be a transaction to abide ACID rules
	err = globals.GetDB().Transaction(func(tx *gorm.DB) error {
		for _, eg := range egs {
			eg.DeletedAt = &now
			eg.DeletedBy = &user.ID
			err = globals.GetDB().Save(&eg).Error
			if err != nil {
				return err
			}
		}
		return nil
	})

	return nil
}

func (f *Folder) RetrieveSubFoldersAndFiles(folder Folder, user User) error {
	// Fetch all subfolders
	var subFolders []Folder
	err := globals.GetDB().Where("parent_folder_id = ?", folder.ID).Find(&subFolders).Error
	if err != nil {
		return err
	}

	// Just like the function above these batch DB operations should be transaction as well
	err = globals.GetDB().Transaction(func(tx *gorm.DB) error {
		for _, sf := range subFolders {
			sf.DeletedAt = nil
			sf.DeletedBy = nil
			err = globals.GetDB().Save(&sf).Error
			if err != nil {
				return err
			}
		}
		return nil
	})

	var egs []EntityGroup
	err = globals.GetDB().Where("parent_folder_id = ?", folder.ID).Find(&egs).Error
	if err != nil {
		return err
	}

	err = globals.GetDB().Transaction(func(tx *gorm.DB) error {
		for _, eg := range egs {
			eg.DeletedAt = nil
			eg.DeletedBy = nil
			err = globals.GetDB().Save(&eg).Error
			if err != nil {
				return err
			}
		}
		return nil
	})

	return nil
}

func (f *Folder) GetParentFolders(folder Folder, breadcrumbs []BreadCrumb) (interface{}, error) {
	// debug statement to see function call
	config.Logger().Infof("Debug: Entering GetParentFolders for folder '%s'", folder.Name)

	if len(breadcrumbs) == 0 {
		slug, err := folder.GetFolderID()
		if err != nil {
			return nil, err
		}
		breadcrumbs = append(breadcrumbs, BreadCrumb{
			Name: folder.Name,
			Slug: *slug,
			ID:   folder.ID,
		})
	}

	if folder.ParentFolderID == nil {
		parentFolder, err := folder.ParentFolder()
		if err != nil {
			return nil, err
		}

		pfSlug, err := parentFolder.GetFolderID()
		if err != nil {
			return nil, err
		}

		breadcrumbs = append(breadcrumbs, BreadCrumb{
			Name: parentFolder.Name,
			Slug: *pfSlug,
			ID:   parentFolder.ID,
		})

		return f.GetParentFolders(*parentFolder, breadcrumbs)
	}
	config.Logger().Infof("Debug: Breadcrumbs for '%s': %v", folder.Name, breadcrumbs)
	return breadcrumbs, nil
}

func (f *Folder) GetAllSubFoldersID(folder Folder, arrayIDs []uint) ([]uint, error) {
	var subFolders []Folder
	err := globals.GetDB().Where("parent_folder_id = ?", folder.ID).Find(&subFolders).Error
	if err != nil {
		return nil, err
	}

	for _, sf := range subFolders {
		arrayIDs = append(arrayIDs, sf.ID)
		ids, err := f.GetAllSubFoldersID(sf, arrayIDs)
		if err != nil {
			return nil, err
		}
		arrayIDs = append(arrayIDs, ids...)
	}

	return arrayIDs, nil
}

func (f *Folder) ReplicateSubFoldersAndFiles(newFolder Folder) error {
	// Fetch folders
	var folders []Folder
	err := globals.
		GetDB().
		Where("parent_folder_id = ?", f.ID).
		Where("deleted_at = NULL").
		Find(&folders).
		Error
	if err != nil {
		return err
	}

	// Fetch files
	var files []EntityGroup
	err = globals.
		GetDB().
		Where("parent_folder_id", f.ID).
		Find(&files).
		Error
	if err != nil {
		return err
	}

	// Create entity groups and department files in one transaction
	// If one fails, or anything goes wrong the transaction is reverted
	err = globals.GetDB().Transaction(func(tx *gorm.DB) error {
		data := make([]map[string]interface{}, 0)
		for _, file := range files {
			attr, err := file.GetAttributes()
			if err != nil {
				return err
			}
			data = append(data, attr, map[string]interface{}{
				"user_id":          newFolder.UserID,
				"parent_folder_id": newFolder.ID,
			})

			dataBytes, err := helper.ToJSON(data)
			if err != nil {
				return err
			}

			var newEntityGroup EntityGroup
			err = helper.FromJSON(dataBytes, &newEntityGroup)
			if err != nil {
				return err
			}

			err = CreateEntityGroupWithSlug(&newEntityGroup)
			if err != nil {
				return err
			}

			departments, err := file.GetEntityGroupDepartments()
			if err != nil {
				return err
			}

			for _, dep := range departments {
				var departmentFile DepartmentFile
				departmentFile.EntityGroupID = newEntityGroup.ID
				departmentFile.DepartmentID = dep.ID

				err = globals.GetDB().Create(&departmentFile).Error
				if err != nil {
					return err
				}
			}
		}

		for _, folder := range folders {
			var newF Folder
			newF.Name = folder.Name
			newF.UserID = newFolder.UserID
			newF.ParentFolderID = &newFolder.ID
			err = globals.GetDB().Create(&newF).Error
			if err != nil {
				return err
			}

			err = folder.ReplicateSubFoldersAndFiles(newF)
			if err != nil {
				return err
			}
		}

		return nil
	})
	return err
}

// todo: implement RetrieveSubFoldersAndFilesForDownload function
