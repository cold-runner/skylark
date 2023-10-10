package consts

// Gender 性别
type Gender int

const (
	MALE Gender = iota
	FEMALE
	UNKNOWN
)

var GenderList = map[string]Gender{
	"男":  MALE,
	"女":  FEMALE,
	"其他": UNKNOWN,
}

func (g Gender) CheckFuncName() string {
	return "gender"
}

func (g Gender) String() string {
	switch g {
	case MALE:
		return "男"
	case FEMALE:
		return "女"
	default:
		return "其他"
	}
}
