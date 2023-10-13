package consts

import "github.com/pkg/errors"

// Grade 年级
type Grade int

// TODO 研究生
const (
	FRESHMAN Grade = iota
	SOPHOMORE
	JUNIOR
	SENIOR
	GRADUATES
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
	case grade == "毕业生":
		return GRADUATES, nil
	}

	var g Grade
	return g, errors.Errorf("not a valid grade: %q", g)
}

func (g Grade) String() string {
	switch g {
	case FRESHMAN:
		return "大一"
	case SOPHOMORE:
		return "大二"
	case JUNIOR:
		return "大三"
	case SENIOR:
		return "大四"
	case GRADUATES:
		return "毕业生"
	default:
		return "不支持的级别"
	}
}

func CheckGradeFuncName() string {
	return "grade"
}
