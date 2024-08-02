package domain

import "errors"

var (
	ErrNotFound     = errors.New("not found")
	NoBakcNeedMatch = errors.New("no bakc need match")
)
