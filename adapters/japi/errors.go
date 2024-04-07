package japi

import "errors"

var ApiError = errors.New("jupiter-api")
var StatusNotCreated = errors.New("resource not created")
var StatusNotOK = errors.New("response not ok")
var StatusNotDeleted = errors.New("resource not deleted")
var JsonError = errors.New("json/decode")
