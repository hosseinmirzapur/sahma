package models

import (
	"fmt"
	"sahma/internal/config"
	"sahma/internal/helper"
	"time"
)

type Entity struct {
	ID                 uint
	EntityGroupID      *uint
	Type               string
	TransciptionResult *string
	FileLocation       string
	ResultLocation     *string
	Meta               *string
	CreatedAt          time.Time
	UpdatedAt          *time.Time
	EntityGroup        *EntityGroup
}

func (ent *Entity) GetFileData() (*string, error) {
	data := config.GetStoragePath(ent.Type) + "/" + ent.FileLocation
	if !helper.FileExists(data) {
		return nil, fmt.Errorf("no file found for this entity with id: #%d", ent.ID)
	}

	return &data, nil
}

func (ent *Entity) GetFileExtension() (*string, error) {
	ext := helper.FileExtension(ent.FileLocation)
	if ext == "" {
		return nil, fmt.Errorf("could not resolve file extension")
	}
	return &ext, nil
}
