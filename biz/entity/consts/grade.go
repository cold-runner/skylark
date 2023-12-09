package consts

import "github.com/pkg/errors"

//go:generate stringer -type Grade -linecomment

// Grade 年级
type Grade int

// TODO 研究生
const (
	FRESHMAN  Grade = iota // 大一
	SOPHOMORE              // 大二
	JUNIOR                 // 大三
	SENIOR                 // 大四
	GRADUATES              // 毕业生
)

func ParseGrade(grade string) (Grade, error) {
	switch {
	case grade == "大一":
		return FRESHMAN, nil
	case grade == "大二":
		return SOPHOMORE, nil
	case grade == "大三":
		return JUNIOR, nil
	case grade == "大四":
		return SENIOR, nil
	case grade == "研究生":
		return GRADUATES, nil
	}

	var g Grade
	return g, errors.Errorf("not a valid grade: %q", g)
}
func CheckGradeFuncName() string {
	return "grade"
}
