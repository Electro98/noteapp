package utils

import "errors"

const ContextKeyDB = "DB"

var ErrNoFoundInDB = errors.New("Item was not found in DB")

const FilterLimitMAX = 50
