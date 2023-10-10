package consts

// College 学院
type College int

const (
	MARXIST = iota
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

// CheckFuncName 校验学院函数名
func (c College) CheckFuncName() string {
	return "college"
}

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
		return ""
	}
}

var CollegeList = map[string]College{
	"法学院":       LAW,
	"古生物学院":     PALEONTOLOGY,
	"管理学院":      MANAGEMENT,
	"国际商学院":     INTERNATIONAL_BUSINESS,
	"化学化工学院":    CHEMISTRY_AND_CHEMICAL_ENGINEERING,
	"教育科学学院":    EDUCATION_SCIENCE,
	"粮食学院":      CEREALS,
	"旅游管理学院":    TOURISM_MANAGEMENT,
	"马克思主义学院":   MARXIST,
	"美术与设计学院":   FINE_ARTS_AND_DESIGN,
	"软件学院":      SOFTWARE,
	"社会学学院":     SOCIOLOGY,
	"生命科学学院":    LIFE_SCIENCE,
	"数学与系统科学学院": MATHEMATICS_AND_SYSTEMS_SCIENCE,
	"体育科学学院":    SPORTS_SCIENCE,
	"外国语学院":     FOREIGN_LANGUAGE,
	"文学院":       LITERATURE,
	"物理科学与技术学院": PHYSICAL_SCIENCE_AND_TECHNOLOGY,
	"戏剧艺术学院":    DRAMATIC_ART,
	"新闻与传播学院":   NEWS_AND_PROPAGATION,
	"学前与初等教育学院": PRESCHOOL_AND_ELEMENTARY_EDUCATION,
	"音乐学院":      MUSIC,
}
