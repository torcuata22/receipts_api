package main

import "errors"

var (
	ErrInvalidJSON     = errors.New("invalid JSON format")
	ErrReceiptNotFound = errors.New("receipt not found")
)
