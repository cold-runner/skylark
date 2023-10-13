package consts

import (
	"github.com/pkg/errors"
)

// Gender 性别
type Gender int

const (
	MALE Gender = iota
	FEMALE
	UNKNOWN
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

func (g Gender) String() string {
	switch g {
	case MALE:
		return "男"
	case FEMALE:
		return "女"
	case UNKNOWN:
		return "其他"
	default:
		return "不支持的性别类型"
	}
}

func CheckGenderFuncName() string {
	return "gender"
}
