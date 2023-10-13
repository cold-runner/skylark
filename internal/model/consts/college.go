package consts

import (
	"github.com/pkg/errors"
)

// College 学院
type College int

const (
	MARXIST College = iota
	SOCIOLOGY
	NEWS_AND_PROPAGATION
	FOREIGN_LANGUAGE
	DRAMATIC_ART
	LIFE_SCIENCE
	PALEONTOLOGY
	TOURISM_MANAGEMENT
	INTERNATIONAL_BUSINESS
	EDUCATION_SCIENCE
	SPORTS_SCIENCE
	MUSIC
	MATHEMATICS_AND_SYSTEMS_SCIENCE
	CHEMISTRY_AND_CHEMICAL_ENGINEERING
	SOFTWARE
	LAW
	PRESCHOOL_AND_ELEMENTARY_EDUCATION
	LITERATURE
	FINE_ARTS_AND_DESIGN
	PHYSICAL_SCIENCE_AND_TECHNOLOGY
	CEREALS
	MANAGEMENT
)

func (c College) String() string {
	switch c {
	case LAW:
		return "法学院"
	case PALEONTOLOGY:
		return "古生物学院"
	case MANAGEMENT:
		return "管理学院"
	case INTERNATIONAL_BUSINESS:
		return "国际商学院"
	case CHEMISTRY_AND_CHEMICAL_ENGINEERING:
		return "化学化工学院"
	case EDUCATION_SCIENCE:
		return "教育科学学院"
	case CEREALS:
		return "粮食学院"
	case TOURISM_MANAGEMENT:
		return "旅游管理学院"
	case MARXIST:
		return "马克思主义学院"
	case FINE_ARTS_AND_DESIGN:
		return "美术与设计学院"
	case SOFTWARE:
		return "软件学院"
	case SOCIOLOGY:
		return "社会学学院"
	case LIFE_SCIENCE:
		return "生命科学学院"
	case MATHEMATICS_AND_SYSTEMS_SCIENCE:
		return "数学与系统科学学院"
	case SPORTS_SCIENCE:
		return "体育科学学院"
	case FOREIGN_LANGUAGE:
		return "外国语学院"
	case LITERATURE:
		return "文学院"
	case PHYSICAL_SCIENCE_AND_TECHNOLOGY:
		return "物理科学与技术学院"
	case DRAMATIC_ART:
		return "戏剧艺术学院"
	case NEWS_AND_PROPAGATION:
		return "新闻与传播学院"
	case PRESCHOOL_AND_ELEMENTARY_EDUCATION:
		return "学前与初等教育学院"
	case MUSIC:
		return "音乐学院"
	default:
		return "未知的学院"
	}
}

func ParseCollege(college string) (College, error) {
	switch {
	case college == "法学院":
		return LAW, nil
	case college == "古生物学院":
		return PALEONTOLOGY, nil
	case college == "管理学院":
		return MANAGEMENT, nil
	case college == "国际商学院":
		return INTERNATIONAL_BUSINESS, nil
	case college == "化学化工学院":
		return CHEMISTRY_AND_CHEMICAL_ENGINEERING, nil
	case college == "教育科学学院":
		return EDUCATION_SCIENCE, nil
	case college == "粮食学院":
		return CEREALS, nil
	case college == "旅游管理学院":
		return TOURISM_MANAGEMENT, nil
	case college == "马克思主义学院":
		return MARXIST, nil
	case college == "美术与设计学院":
		return FINE_ARTS_AND_DESIGN, nil
	case college == "软件学院":
		return SOFTWARE, nil
	case college == "社会学学院":
		return SOCIOLOGY, nil
	case college == "生命科学学院":
		return LIFE_SCIENCE, nil
	case college == "数学与系统科学学院":
		return MATHEMATICS_AND_SYSTEMS_SCIENCE, nil
	case college == "体育科学学院":
		return SPORTS_SCIENCE, nil
	case college == "外国语学院":
		return FOREIGN_LANGUAGE, nil
	case college == "文学院":
		return LITERATURE, nil
	case college == "物理科学与技术学院":
		return PHYSICAL_SCIENCE_AND_TECHNOLOGY, nil
	case college == "戏剧艺术学院":
		return DRAMATIC_ART, nil
	case college == "新闻与传播学院":
		return NEWS_AND_PROPAGATION, nil
	case college == "学前与初等教育学院":
		return PRESCHOOL_AND_ELEMENTARY_EDUCATION, nil
	case college == "音乐学院":
		return MUSIC, nil
	}

	var c College
	return c, errors.Errorf("not a valid college: %q", c)
}

// CheckCollegeFuncName 校验学院函数名
func CheckCollegeFuncName() string {
	return "college"
}
