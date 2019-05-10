package blog

import "gopkg.in/go-playground/validator.v8"

// 验证器

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}
