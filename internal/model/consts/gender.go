package consts

import (
	"github.com/pkg/errors"
)

//go:generate stringer -type Gender -linecomment

// Gender 性别
type Gender int

const (
	MALE    Gender = iota // 男
	FEMALE                // 女
	UNKNOWN               // 其他
)

func ParseGender(gender string) (Gender, error) {
	switch {
	case gender == "男":
		return MALE, nil
	case gender == "女":
		return FEMALE, nil
	case gender == "其他":
		return UNKNOWN, nil
	}
	var g Gender
	return g, errors.Errorf("not a valid gender: %q", g)
}

func CheckGenderFuncName() string {
	return "gender"
}
