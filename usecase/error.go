package usecase

import "errors"

var (
	ErrInvalidPhoneNumber  = errors.New("invalid phone number")
	ErrInvalidEmail        = errors.New("invalid email")
	ErrInvalidLinkedInURL  = errors.New("invalid linkedin url")
	ErrInvalidPortfolioURL = errors.New("invalid portfolio url")
)
