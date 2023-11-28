package util

const (
	FEMALE = 0
	MALE   = 1
	OTHER  = 2
)

func GenderTransform(gender string) int {
	switch gender {
	case "男":
		return MALE
	case "女":
		return FEMALE
	default:
		return OTHER
	}
}
