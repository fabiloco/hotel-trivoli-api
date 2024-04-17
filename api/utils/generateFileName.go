package utils

import (
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"github.com/google/uuid"
)

func GenerateFileName(file *multipart.FileHeader) (string, error) {
  fileExt := filepath.Ext(file.Filename)

  if fileExt == "" || 
    fileExt != ".png"   &&
    fileExt != ".jpg"   &&
    fileExt != ".jpeg"  &&
    fileExt != ".webp" {
    return "", errors.New(fmt.Sprintf("File extension '%s' is not valid.", fileExt)) 
  }

  fileId := uuid.NewString()
  return fileId + filepath.Ext(file.Filename), nil
}
