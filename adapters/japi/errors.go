package japi

import "errors"

var ApiError = errors.New("jupiter-api")
var StatusNotCreated = errors.New("resource not created")
var StatusNotOK = errors.New("response not ok")
var JsonError = errors.New("json/decode")
