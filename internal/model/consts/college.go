package consts

import (
	"github.com/pkg/errors"
)

//go:generate stringer -type College -linecomment

// College 学院
type College int

const (
	MARXIST                            College = iota // 马克思主义学院
	SOCIOLOGY                                         // 生命科学学院
	NEWS_AND_PROPAGATION                              // 新闻与传播学院
	FOREIGN_LANGUAGE                                  // 外国语学院
	DRAMATIC_ART                                      // 戏剧艺术学院
	LIFE_SCIENCE                                      // 生命科学学院
	PALEONTOLOGY                                      // 古生物学院
	TOURISM_MANAGEMENT                                // 旅游管理学院
	INTERNATIONAL_BUSINESS                            // 国际商学院
	EDUCATION_SCIENCE                                 // 教育科学学院
	SPORTS_SCIENCE                                    // 体育科学学院
	MUSIC                                             // 音乐学院
	MATHEMATICS_AND_SYSTEMS_SCIENCE                   // 数学与系统科学学院
	CHEMISTRY_AND_CHEMICAL_ENGINEERING                // 化学化工学院
	SOFTWARE                                          // 软件学院
	LAW                                               // 法学院
	PRESCHOOL_AND_ELEMENTARY_EDUCATION                // 学前与初等教育学院
	LITERATURE                                        // 文学院
	FINE_ARTS_AND_DESIGN                              // 美术与设计学院
	PHYSICAL_SCIENCE_AND_TECHNOLOGY                   // 物理科学与技术学院
	CEREALS                                           // 粮食学院
	MANAGEMENT                                        // 管理学院
)

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
