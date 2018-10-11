package service

import (
	"fmt"
)

//ErrPresetNotFound represents an error when the preset was not found in the datastore
type ErrPresetNotFound struct {
	message string
}

//newErrPresetNotFound returns a new ErrPresetNotFound error
func newErrPresetNotFound(message string) *ErrPresetNotFound {
	return &ErrPresetNotFound{
		message: message,
	}
}

func (e *ErrPresetNotFound) Error() string {
	return e.message
}

//ErrUnknown represents an error when the preset was not found in the datastore
type ErrUnknown struct {
	message string
}

//newErrUnknown returns a new ErrUnknown error
func newErrUnknown(message string) *ErrUnknown {
	return &ErrUnknown{
		message: message,
	}
}

func (e *ErrUnknown) Error() string {
	return fmt.Sprintf("unknown error getting the preset. %s", e.message)
}

// ErrUploadingPreset represents an error while trying to upload a preset in the
type ErrUploadingPreset struct {
	message string
}

//newErrUploadingPreset returns a new ErrUploadingPreset error
func newErrUploadingPreset(message string) *ErrUploadingPreset {
	return &ErrUploadingPreset{
		message: message,
	}
}

func (e *ErrUploadingPreset) Error() string {
	return fmt.Sprintf("error uploading preset. %s", e.message)
}
