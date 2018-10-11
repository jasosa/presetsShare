package datastore

import "errors"

//ErrPresetNotFoundInDataStore the preset was not found in the datastore
var ErrPresetNotFoundInDataStore = errors.New("preset not found in datastore")
