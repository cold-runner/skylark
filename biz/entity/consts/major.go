package consts

import "github.com/pkg/errors"

//go:generate stringer -type Major -linecomment

// Major 专业
type Major int

// TODO 专升本专业
const (
	MAJOR_LAW                                                    Major = iota // 法学
	MAJOR_PALEONTOLOGY                                                        // 古生物学
	MAJOR_ADMINISTRATION                                                      // 行政管理
	MAJOR_LABOR_AND_SOCIAL_SECURITY                                           // 劳动与社会保障
	MAJOR_INTERNATIONAL_ECONOMY_AND_TRADING                                   // 国际经济与贸易
	MAJOR_FINANCE                                                             // 金融学
	MAJOR_MARKETING                                                           // 市场营销
	MAJOR_LOGISTICS_MANAGEMENT                                                // 物流管理
	MAJOR_CHEMISTRY_NORMAL                                                    // 化学（师范）
	MAJOR_ENERGY_CHEMICAL_ENGINEERING                                         // 能源化学工程
	MAJOR_PEDAGOGY_NORMAL                                                     // 教育学（师范）
	MAJOR_APPLIED_PSYCHOLOGY                                                  // 应用心理学（师范）
	MAJOR_GRAIN_ENGINEERING                                                   // 粮食工程
	MAJOR_FOOD_SCIENCE_AND_ENGINEERING                                        // 食品科学与工程
	MAJOR_TOURISM_MANAGEMENT                                                  // 旅游管理
	MAJOR_HOTEL_MANAGEMENT                                                    // 酒店管理
	MAJOR_HISTORY_NORMAL                                                      // 历史学（师范类）
	MAJOR_IDEOLOGICAL_AND_POLITICAL_EDUCATION_NORMAL                          // 思想政治教育（师范类）
	MAJOR_FINE_ARTS_NORMAL                                                    // 美术学（师范）
	MAJOR_VISUAL_COMMUNICATION_DESIGN                                         // 视觉传达设计
	MAJOR_ENVIRONMENTAL_DESIGN                                                // 环境设计
	MAJOR_PUBLIC_ART                                                          // 公共艺术
	MAJOR_PAINTING                                                            // 绘画
	MAJOR_SOFTWARE_ENGINEERING                                                // 软件工程
	MAJOR_NETWORK_ENGINEERING                                                 // 网络工程
	MAJOR_COMPUTER_SCIENCE_AND_TECHNOLOGY                                     // 计算机科学与技术
	MAJOR_COMPUTER_SCIENCE_AND_TECHNOLOGY_NORMAL                              // 计算机科学与技术（师范）
	MAJOR_SOCIOLOGY                                                           // 社会学
	MAJOR_SOCIAL_WORK                                                         // 社会工作
	MAJOR_BIOLOGICAL_SCIENCES_NORMAL                                          // 生物科学专业（师范）
	MAJOR_ENVIRONMENTAL_ECOLOGICAL_ENGINEERING                                // 环境生态工程
	MAJOR_MATHEMATICS_AND_APPLIED_MATHEMATICS_NORMAL                          // 数学与应用数学（师范）
	MAJOR_DATA_SCIENCE_AND_BIG_DATA_TECHNOLOGY                                // 数据科学与大数据技术
	MAJOR_PHYSICAL_EDUCATION                                                  // 体育教育
	MAJOR_SPORTS_TRAINING                                                     // 运动训练
	MAJOR_WUSHU_AND_NATIONAL_TRADITIONAL_SPORTS                               // 武术与民族传统体育
	MAJOR_ENGLISH_NORMAL                                                      // 英语（师范）
	MAJOR_JAPANESE                                                            // 日语
	MAJOR_RUSSIAN                                                             // 俄语
	MAJOR_FRENCH                                                              // 法语
	MAJOR_TRANSLATION                                                         // 翻译
	MAJOR_GERMAN                                                              // 德语
	MAJOR_CHINESE_LANGUAGE_AND_LITERATURE_NORMAL                              // 汉语言文学（师范）
	MAJOR_TEACHING_CHINESE_TO_SPEAKERS_OF_OTHER_LANGUAGES_NORMAL              // 汉语国际教育（师范）
	MAJOR_PHYSICS_NORMAL                                                      // 物理学（师范）
	MAJOR_ELECTRONIC_INFORMATION_ENGINEERING                                  // 电子信息工程
	MAJOR_PERFORMANCE                                                         // 表演
	MAJOR_DANCE_PERFORMANCE                                                   // 舞蹈表演
	MAJOR_MUSIC_PERFORMANCE_ART                                               // 音乐表演（戏剧艺术学院）
	MAJOR_DRAMA_FILM_AND_TELEVISION_ART_DESIGN                                // 戏剧影视美术设计
	MAJOR_BROADCASTING_AND_HOSTING_ART                                        // 播音与主持艺术
	MAJOR_JOURNALISM                                                          // 新闻学
	MAJOR_EDUCATIONAL_TECHNOLOGY_NORMAL                                       // 教育技术学（师范）
	MAJOR_INTERNET_AND_NEW_MEDIA                                              // 网络与新媒体
	MAJOR_RADIO_AND_TELEVISION_DIRECTOR                                       // 广播电视编导
	MAJOR_PRESCHOOL_EDUCATION_NORMAL                                          // 学前教育（师范）
	MAJOR_PRIMARY_EDUCATION_NORMAL                                            // 小学教育（师范)
	MAJOR_MUSICOLOGY_NORMAL                                                   // 音乐学（师范）
	MAJOR_MUSIC_PERFORMANCE_MUSIC                                             // 音乐表演（音乐学院）
)

func ParseMajor(major string) (Major, error) {
	switch {
	case major == "法学":
		return MAJOR_LAW, nil
	case major == "古生物学":
		return MAJOR_PALEONTOLOGY, nil
	case major == "行政管理":
		return MAJOR_ADMINISTRATION, nil
	case major == "劳动与社会保障":
		return MAJOR_LABOR_AND_SOCIAL_SECURITY, nil
	case major == "国际经济与贸易":
		return MAJOR_INTERNATIONAL_ECONOMY_AND_TRADING, nil
	case major == "金融学":
		return MAJOR_FINANCE, nil
	case major == "市场营销":
		return MAJOR_MARKETING, nil
	case major == "物流管理":
		return MAJOR_LOGISTICS_MANAGEMENT, nil
	case major == "化学（师范）":
		return MAJOR_CHEMISTRY_NORMAL, nil
	case major == "能源化学工程":
		return MAJOR_ENERGY_CHEMICAL_ENGINEERING, nil
	case major == "教育学（师范）":
		return MAJOR_PEDAGOGY_NORMAL, nil
	case major == "应用心理学（师范）":
		return MAJOR_APPLIED_PSYCHOLOGY, nil
	case major == "粮食工程":
		return MAJOR_GRAIN_ENGINEERING, nil
	case major == "食品科学与工程":
		return MAJOR_FOOD_SCIENCE_AND_ENGINEERING, nil
	case major == "旅游管理":
		return MAJOR_TOURISM_MANAGEMENT, nil
	case major == "酒店管理":
		return MAJOR_HOTEL_MANAGEMENT, nil
	case major == "历史学（师范类）":
		return MAJOR_HISTORY_NORMAL, nil
	case major == "思想政治教育（师范）":
		return MAJOR_IDEOLOGICAL_AND_POLITICAL_EDUCATION_NORMAL, nil
	case major == "美术学（师范）":
		return MAJOR_FINE_ARTS_NORMAL, nil
	case major == "视觉传达设计":
		return MAJOR_VISUAL_COMMUNICATION_DESIGN, nil
	case major == "环境设计":
		return MAJOR_ENVIRONMENTAL_DESIGN, nil
	case major == "公共艺术":
		return MAJOR_PUBLIC_ART, nil
	case major == "绘画":
		return MAJOR_PAINTING, nil
	case major == "软件工程":
		return MAJOR_SOFTWARE_ENGINEERING, nil
	case major == "网络工程":
		return MAJOR_NETWORK_ENGINEERING, nil
	case major == "计算机科学与技术":
		return MAJOR_COMPUTER_SCIENCE_AND_TECHNOLOGY, nil
	case major == "计算机科学与技术（师范）":
		return MAJOR_COMPUTER_SCIENCE_AND_TECHNOLOGY_NORMAL, nil
	case major == "社会学":
		return MAJOR_SOCIOLOGY, nil
	case major == "社会工作":
		return MAJOR_SOCIAL_WORK, nil
	case major == "生物科学专业（师范）":
		return MAJOR_BIOLOGICAL_SCIENCES_NORMAL, nil
	case major == "环境生态工程":
		return MAJOR_ENVIRONMENTAL_ECOLOGICAL_ENGINEERING, nil
	case major == "数学与应用数学（师范）":
		return MAJOR_MATHEMATICS_AND_APPLIED_MATHEMATICS_NORMAL, nil
	case major == "数据科学与大数据技术":
		return MAJOR_DATA_SCIENCE_AND_BIG_DATA_TECHNOLOGY, nil
	case major == "体育教育":
		return MAJOR_PHYSICAL_EDUCATION, nil
	case major == "运动训练":
		return MAJOR_SPORTS_TRAINING, nil
	case major == "武术与民族传统体育":
		return MAJOR_WUSHU_AND_NATIONAL_TRADITIONAL_SPORTS, nil
	case major == "英语（师范）":
		return MAJOR_ENGLISH_NORMAL, nil
	case major == "日语":
		return MAJOR_JAPANESE, nil
	case major == "俄语":
		return MAJOR_RUSSIAN, nil
	case major == "法语":
		return MAJOR_FRENCH, nil
	case major == "翻译":
		return MAJOR_TRANSLATION, nil
	case major == "德语":
		return MAJOR_GERMAN, nil
	case major == "汉语言文学（师范）":
		return MAJOR_CHINESE_LANGUAGE_AND_LITERATURE_NORMAL, nil
	case major == "汉语国际教育（师范）":
		return MAJOR_TEACHING_CHINESE_TO_SPEAKERS_OF_OTHER_LANGUAGES_NORMAL, nil
	case major == "物理学（师范）":
		return MAJOR_PHYSICS_NORMAL, nil
	case major == "电子信息工程":
		return MAJOR_ELECTRONIC_INFORMATION_ENGINEERING, nil
	case major == "表演":
		return MAJOR_PERFORMANCE, nil
	case major == "舞蹈表演":
		return MAJOR_DANCE_PERFORMANCE, nil
	case major == "音乐表演（戏剧艺术学院）":
		return MAJOR_MUSIC_PERFORMANCE_ART, nil
	case major == "戏剧影视美术设计":
		return MAJOR_DRAMA_FILM_AND_TELEVISION_ART_DESIGN, nil
	case major == "播音与主持艺术":
		return MAJOR_BROADCASTING_AND_HOSTING_ART, nil
	case major == "新闻学":
		return MAJOR_JOURNALISM, nil
	case major == "教育技术学（师范）":
		return MAJOR_EDUCATIONAL_TECHNOLOGY_NORMAL, nil
	case major == "网络与新媒体":
		return MAJOR_INTERNET_AND_NEW_MEDIA, nil
	case major == "广播电视编导":
		return MAJOR_RADIO_AND_TELEVISION_DIRECTOR, nil
	case major == "学前教育（师范）":
		return MAJOR_PRESCHOOL_EDUCATION_NORMAL, nil
	case major == "小学教育（师范)":
		return MAJOR_PRIMARY_EDUCATION_NORMAL, nil
	case major == "音乐学（师范）":
		return MAJOR_MUSICOLOGY_NORMAL, nil
	case major == "音乐表演（音乐学院）":
		return MAJOR_MUSIC_PERFORMANCE_MUSIC, nil
	}

	var m Major
	return m, errors.Errorf("not a valid major %q", m)
}

func CheckMajorFuncName() string {
	return "major"
}
