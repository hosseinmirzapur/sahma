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
	"strings"
	"time"

	"gorm.io/gorm"
)

type EntityGroup struct {
	ID                  uint
	UserID              uint
	ParentFolderID      *uint
	Name                string
	Type                string
	TranscriptionResult *string
	TranscriptionAt     *string
	Status              EntityGroupStatus
	Meta                *string
	FileLocation        string
	Description         *string
	ArchivedAt          *string
	ResultLocation      *string
	NumberOfTry         uint
	DeletedAt           time.Time
	Slug                *string
	CreatedAt           time.Time
	UpdatedAt           time.Time
	Activities          []*Activity `gorm:"polymorphic:Activity;"`
	Entities            []*Entity
	Folders             []*Folder
	User                User
}

type ResultLocation struct {
	WavLocation        string `json:"wav_location"`
	ConvertedWordToPDF string `json:"converted_word_to_pdf"`
	PdfLocation        string `json:"pdf_location"`
}

type Meta struct {
	TifConvertedPngLocation string `json:"tif_converted_png_location"`
}

type FileData struct {
	Type    string
	Content string
	Name    string
}

// Different entity group status types
type EntityGroupStatus string

const (
	STATUS_WAITING_FOR_TRANSCRIPTION    = "WAITING_FOR_TRANSCRIPTION"
	STATUS_TRANSCRIBED                  = "TRANSCRIBED"
	STATUS_WAITING_FOR_AUDIO_SEPARATION = "WAITING_FOR_AUDIO_SEPARATION"
	STATUS_WAITING_FOR_SPLIT            = "WAITING_FOR_SPLIT"
	STATUS_WAITING_FOR_WORD_EXTRACTION  = "WAITING_FOR_WORD_EXTRACTION"
	STATUS_WAITING_FOR_RETRY            = "WAITING_FOR_RETRY"
	STATUS_REJECTED                     = "REJECTED"
	STATUS_ZIPPED                       = "ZIPPED"
	STATUS_REPORT                       = "REPORT"
)

func (eg *EntityGroup) GetEntityGroupDepartments() ([]Department, error) {
	var departments []Department
	err := globals.
		GetDB().
		Select("departments.id", "departments.name").
		Joins("department_files ON department_files.department_id = departments.id").
		Where("department_files.entity_group_id", eg.ID).
		Distinct().
		Find(&departments).
		Error
	if err != nil {
		return nil, err
	}

	return departments, nil
}

func (eg *EntityGroup) TextSearch(query *gorm.DB, searchTxt string) *gorm.DB {
	return query.Where("transcription_result LIKE %?%", searchTxt)
}

// Always create EntityGroup with this function
func CreateWithSlug(eg *EntityGroup) error {
	err := globals.GetDB().Transaction(func(tx *gorm.DB) error {
		slug, err := eg.GetEntityGroupID()
		if err != nil {
			return err
		}
		eg.Slug = slug
		return globals.GetDB().Create(&eg).Error
	})

	return err
}

func (eg *EntityGroup) GetEntityGroupID() (*string, error) {
	// converting entity group id to string
	idStr := strconv.Itoa(int(eg.ID))

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

func ConvertObfuscatedIdToEntityGroupId(obfuscated string) (*uint, error) {
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

// Fetch all the records between now and 15 minutes ago
func (eg *EntityGroup) AvailableNow(query *gorm.DB) *gorm.DB {
	fifteenMinutesAgo := time.Now().Add(-15 * time.Minute)
	return query.Where("created_at BETWEEN ? AND ?", fifteenMinutesAgo, time.Now())
}

func (eg *EntityGroup) GetFileData(getWavFile bool) (*string, error) {
	data := ""

	var (
		resultLocation ResultLocation
		meta           Meta
	)

	if getWavFile {
		if eg.ResultLocation == nil {
			data = config.GetStoragePath("voice") + "/"
		}
		err := helper.FromJSON([]byte(*eg.ResultLocation), &resultLocation)
		if err != nil {
			return nil, err
		}

		data = config.GetStoragePath("voice") + "/" + resultLocation.WavLocation
	} else {
		if strings.Contains(eg.Name, "tif") {
			err := helper.FromJSON([]byte(*eg.Meta), &meta)
			if err != nil {
				return nil, err
			}
			data = config.GetStoragePath(eg.Type) + "/" + meta.TifConvertedPngLocation
		} else {
			disk := ""
			fileLocation := ""
			if eg.Type == "word" {
				disk = "pdf"
				fileLocation = resultLocation.ConvertedWordToPDF
			} else {
				disk = eg.Type
				fileLocation = eg.FileLocation
			}
			data = config.GetStoragePath(disk) + "/" + fileLocation
		}
	}

	if !helper.FileExists(data) {
		return nil, fmt.Errorf("failed to get entity data. EntityGroup id: #%d", eg.ID)
	}

	return &data, nil
}

func (eg *EntityGroup) GetTranscribedFileData() (*string, error) {
	var (
		resultLocation ResultLocation
		err            error
	)
	fileLocation := ""
	data := ""
	if eg.ResultLocation != nil {
		err = helper.FromJSON([]byte(*eg.ResultLocation), &resultLocation)
		if err != nil {
			return nil, err
		}
		if eg.Type == "word" && eg.Status != STATUS_TRANSCRIBED {
			fileLocation = resultLocation.ConvertedWordToPDF
		} else {
			fileLocation = resultLocation.PdfLocation
		}
	}
	data = config.GetStoragePath("pdf") + "/" + fileLocation
	if !helper.FileExists(data) {
		return nil, fmt.Errorf("failed to get entity data. EntityGroup id: #%d", eg.ID)
	}

	return &data, nil

}

func (eg *EntityGroup) GetFileExtension() (*string, error) {
	ext := helper.FileExtension(eg.FileLocation)
	if ext == "" {
		return nil, fmt.Errorf("could not resolve file extension")
	}
	return &ext, nil
}

func (eg *EntityGroup) GetHtmlEmbeddableFileData(isBase64 bool) (*string, error) {
	if !helper.FileExists(eg.FileLocation) {
		return nil, fmt.Errorf("there is no file for EntityGroup id: #%d", eg.ID)
	}

	supportedTypes := []string{"voice", "image", "pdf", "video", "word"}
	if !isBase64 && slices.Contains(supportedTypes, eg.Type) {
		return eg.GetFileData(false)
	}

	result := ""
	fileFormat := helper.FileExtension(eg.FileLocation)
	fileData, err := eg.GetFileData(false)
	base64FileData := helper.Base64Encode(*fileData)
	if err != nil {
		return nil, err
	}
	switch eg.Type {
	case "voice":
		result = "data:audio/" + fileFormat + ";base64," + base64FileData
	case "image":
		result = "data:image/" + fileFormat + ";base64," + base64FileData
	case "pdf":
		result = "data:application/pdf;base64," + base64FileData
	case "video":
		result = "data:video/" + fileFormat + ";base64," + base64FileData
	case "word":
		result = "data:application/pdf;base64," + base64FileData
	default:
		result = ""
	}

	if result == "" {
		return nil, fmt.Errorf("file type not supported")
	}

	return &result, nil
}

func (eg *EntityGroup) GetHtmlEmbeddableTranscribedFileData(isBase64 bool) (*string, error) {
	if eg.ResultLocation == nil {
		return nil, fmt.Errorf("there is no file specified with this EntityGroup id: #%d", eg.ID)
	}

	var resultLocation ResultLocation
	helper.FromJSON([]byte(*eg.ResultLocation), &resultLocation)
	if !helper.FileExists(resultLocation.PdfLocation) {
		return nil, fmt.Errorf("there is no file specified with this EntityGroup id: #%d", eg.ID)
	}

	supportedTypes := []string{"image", "pdf", "word"}
	result := ""
	if slices.Contains(supportedTypes, eg.Type) {
		data, err := eg.GetTranscribedFileData()
		if err != nil {
			return nil, err
		}
		if isBase64 {
			result = "data:application/pdf;base64," + helper.Base64Encode(*data)
		} else {
			result = *data
		}
	}

	return &result, nil
}

func (eg *EntityGroup) GenerateFileDataForEmbedding(isBase64 bool) (*FileData, error) {
	var fileData FileData

	supportedTypes := []string{"image", "pdf", "word"}
	if slices.Contains(supportedTypes, eg.Type) && eg.TranscriptionResult != nil {
		if eg.Type == "image" && eg.Status == STATUS_TRANSCRIBED {
			fileData.Type = "image"
			fileData.Name = eg.Name
		} else {
			fileData.Type = "pdf"
			fileData.Name = helper.FileName(eg.Name) + ".pdf"
		}

		content, err := eg.GetHtmlEmbeddableTranscribedFileData(isBase64)
		if err != nil {
			return nil, err
		}
		fileData.Content = *content
	} else {
		fileData.Type = eg.Type
		content, err := eg.GetHtmlEmbeddableFileData(isBase64)
		if err != nil {
			return nil, err
		}
		fileData.Content = *content
	}

	return &fileData, nil
}

func (eg *EntityGroup) GetFileSizeHumanReadable(sizeInBytes float64) string {
	units := []string{"B", "KB", "MB", "GB", "TB"}

	unitIndex := 0
	for sizeInBytes >= 1024 && unitIndex < len(units)-1 {
		sizeInBytes /= 1024
		unitIndex++
	}

	return fmt.Sprintf("%.2f %s", sizeInBytes, units[unitIndex])
}
