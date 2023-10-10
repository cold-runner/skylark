package consts

// Grade 年级
type Grade int

const (
	FRESHMAN Grade = iota
	SOPHOMORE
	JUNIOR
	SENIOR
	GRADUATES
)

var GradeList = map[string]Grade{
	"大一":  FRESHMAN,
	"大二":  SOPHOMORE,
	"大三":  JUNIOR,
	"大四":  SENIOR,
	"毕业生": GRADUATES,
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
		return ""
	}
}

func (g Grade) CheckFuncName() string {
	return "grade"
}
