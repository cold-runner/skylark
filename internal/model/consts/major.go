package consts

import "github.com/pkg/errors"

// Major 专业
type Major int

// TODO 专升本专业
const (
	MAJOR_LAW Major = iota
	MAJOR_PALEONTOLOGY
	MAJOR_ADMINISTRATION
	MAJOR_LABOR_AND_SOCIAL_SECURITY
	MAJOR_INTERNATIONAL_ECONOMY_AND_TRADING
	MAJOR_FINANCE
	MAJOR_MARKETING
	MAJOR_LOGISTICS_MANAGEMENT
	MAJOR_CHEMISTRY_NORMAL
	MAJOR_ENERGY_CHEMICAL_ENGINEERING
	MAJOR_PEDAGOGY_NORMAL
	MAJOR_APPLIED_PSYCHOLOGY
	MAJOR_GRAIN_ENGINEERING
	MAJOR_FOOD_SCIENCE_AND_ENGINEERING
	MAJOR_TOURISM_MANAGEMENT
	MAJOR_HOTEL_MANAGEMENT
	MAJOR_HISTORY_NORMAL
	MAJOR_IDEOLOGICAL_AND_POLITICAL_EDUCATION_NORMAL
	MAJOR_FINE_ARTS_NORMAL
	MAJOR_VISUAL_COMMUNICATION_DESIGN
	MAJOR_ENVIRONMENTAL_DESIGN
	MAJOR_PUBLIC_ART
	MAJOR_PAINTING
	MAJOR_SOFTWARE_ENGINEERING
	MAJOR_NETWORK_ENGINEERING
	MAJOR_COMPUTER_SCIENCE_AND_TECHNOLOGY
	MAJOR_COMPUTER_SCIENCE_AND_TECHNOLOGY_NORMAL
	MAJOR_SOCIOLOGY
	MAJOR_SOCIAL_WORK
	MAJOR_BIOLOGICAL_SCIENCES_NORMAL
	MAJOR_ENVIRONMENTAL_ECOLOGICAL_ENGINEERING
	MAJOR_MATHEMATICS_AND_APPLIED_MATHEMATICS_NORMAL
	MAJOR_DATA_SCIENCE_AND_BIG_DATA_TECHNOLOGY
	MAJOR_PHYSICAL_EDUCATION
	MAJOR_SPORTS_TRAINING
	MAJOR_WUSHU_AND_NATIONAL_TRADITIONAL_SPORTS
	MAJOR_ENGLISH_NORMAL
	MAJOR_JAPANESE
	MAJOR_RUSSIAN
	MAJOR_FRENCH
	MAJOR_TRANSLATION
	MAJOR_GERMAN
	MAJOR_CHINESE_LANGUAGE_AND_LITERATURE_NORMAL
	MAJOR_TEACHING_CHINESE_TO_SPEAKERS_OF_OTHER_LANGUAGES_NORMAL
	MAJOR_PHYSICS_NORMAL
	MAJOR_ELECTRONIC_INFORMATION_ENGINEERING
	MAJOR_PERFORMANCE
	MAJOR_DANCE_PERFORMANCE
	MAJOR_MUSIC_PERFORMANCE_ART
	MAJOR_DRAMA_FILM_AND_TELEVISION_ART_DESIGN
	MAJOR_BROADCASTING_AND_HOSTING_ART
	MAJOR_JOURNALISM
	MAJOR_EDUCATIONAL_TECHNOLOGY_NORMAL
	MAJOR_INTERNET_AND_NEW_MEDIA
	MAJOR_RADIO_AND_TELEVISION_DIRECTOR
	MAJOR_PRESCHOOL_EDUCATION_NORMAL
	MAJOR_PRIMARY_EDUCATION_NORMAL
	MAJOR_MUSICOLOGY_NORMAL
	MAJOR_MUSIC_PERFORMANCE_MUSIC
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

func (m Major) String() string {
	switch m {
	case MAJOR_LAW:
		return "法学"
	case MAJOR_PALEONTOLOGY:
		return "古生物学"
	case MAJOR_ADMINISTRATION:
		return "行政管理"
	case MAJOR_LABOR_AND_SOCIAL_SECURITY:
		return "劳动与社会保障"
	case MAJOR_INTERNATIONAL_ECONOMY_AND_TRADING:
		return "国际经济与贸易"
	case MAJOR_FINANCE:
		return "金融学"
	case MAJOR_MARKETING:
		return "市场营销"
	case MAJOR_LOGISTICS_MANAGEMENT:
		return "物流管理"
	case MAJOR_CHEMISTRY_NORMAL:
		return "化学（师范）"
	case MAJOR_ENERGY_CHEMICAL_ENGINEERING:
		return "能源化学工程"
	case MAJOR_PEDAGOGY_NORMAL:
		return "教育学（师范）"
	case MAJOR_APPLIED_PSYCHOLOGY:
		return "应用心理学（师范）"
	case MAJOR_GRAIN_ENGINEERING:
		return "粮食工程"
	case MAJOR_FOOD_SCIENCE_AND_ENGINEERING:
		return "食品科学与工程"
	case MAJOR_TOURISM_MANAGEMENT:
		return "旅游管理"
	case MAJOR_HOTEL_MANAGEMENT:
		return "酒店管理"
	case MAJOR_HISTORY_NORMAL:
		return "历史学（师范类）"
	case MAJOR_IDEOLOGICAL_AND_POLITICAL_EDUCATION_NORMAL:
		return "思想政治教育（师范类）"
	case MAJOR_FINE_ARTS_NORMAL:
		return "美术学（师范）"
	case MAJOR_VISUAL_COMMUNICATION_DESIGN:
		return "视觉传达设计"
	case MAJOR_ENVIRONMENTAL_DESIGN:
		return "环境设计"
	case MAJOR_PUBLIC_ART:
		return "公共艺术"
	case MAJOR_PAINTING:
		return "绘画"
	case MAJOR_SOFTWARE_ENGINEERING:
		return "软件工程"
	case MAJOR_NETWORK_ENGINEERING:
		return "网络工程"
	case MAJOR_COMPUTER_SCIENCE_AND_TECHNOLOGY:
		return "计算机科学与技术"
	case MAJOR_COMPUTER_SCIENCE_AND_TECHNOLOGY_NORMAL:
		return "计算机科学与技术（师范）"
	case MAJOR_SOCIOLOGY:
		return "社会学"
	case MAJOR_SOCIAL_WORK:
		return "社会工作"
	case MAJOR_BIOLOGICAL_SCIENCES_NORMAL:
		return "生物科学专业（师范）"
	case MAJOR_ENVIRONMENTAL_ECOLOGICAL_ENGINEERING:
		return "环境生态工程"
	case MAJOR_MATHEMATICS_AND_APPLIED_MATHEMATICS_NORMAL:
		return "数学与应用数学（师范）"
	case MAJOR_DATA_SCIENCE_AND_BIG_DATA_TECHNOLOGY:
		return "数据科学与大数据技术"
	case MAJOR_PHYSICAL_EDUCATION:
		return "体育教育"
	case MAJOR_SPORTS_TRAINING:
		return "运动训练"
	case MAJOR_WUSHU_AND_NATIONAL_TRADITIONAL_SPORTS:
		return "武术与民族传统体育"
	case MAJOR_ENGLISH_NORMAL:
		return "英语（师范）"
	case MAJOR_JAPANESE:
		return "日语"
	case MAJOR_RUSSIAN:
		return "俄语"
	case MAJOR_FRENCH:
		return "法语"
	case MAJOR_TRANSLATION:
		return "翻译"
	case MAJOR_GERMAN:
		return "德语"
	case MAJOR_CHINESE_LANGUAGE_AND_LITERATURE_NORMAL:
		return "汉语言文学（师范）"
	case MAJOR_TEACHING_CHINESE_TO_SPEAKERS_OF_OTHER_LANGUAGES_NORMAL:
		return "汉语国际教育（师范）"
	case MAJOR_PHYSICS_NORMAL:
		return "物理学（师范）"
	case MAJOR_ELECTRONIC_INFORMATION_ENGINEERING:
		return "电子信息工程"
	case MAJOR_PERFORMANCE:
		return "表演"
	case MAJOR_DANCE_PERFORMANCE:
		return "舞蹈表演"
	case MAJOR_MUSIC_PERFORMANCE_ART:
		return "音乐表演（戏剧艺术学院）"
	case MAJOR_DRAMA_FILM_AND_TELEVISION_ART_DESIGN:
		return "戏剧影视美术设计"
	case MAJOR_BROADCASTING_AND_HOSTING_ART:
		return "播音与主持艺术"
	case MAJOR_JOURNALISM:
		return "新闻学"
	case MAJOR_EDUCATIONAL_TECHNOLOGY_NORMAL:
		return "教育技术学（师范）"
	case MAJOR_INTERNET_AND_NEW_MEDIA:
		return "网络与新媒体"
	case MAJOR_RADIO_AND_TELEVISION_DIRECTOR:
		return "广播电视编导"
	case MAJOR_PRESCHOOL_EDUCATION_NORMAL:
		return "学前教育（师范）"
	case MAJOR_PRIMARY_EDUCATION_NORMAL:
		return "小学教育（师范)"
	case MAJOR_MUSICOLOGY_NORMAL:
		return "音乐学（师范）"
	case MAJOR_MUSIC_PERFORMANCE_MUSIC:
		return "音乐表演（音乐学院）"
	default:
		return "未知专业类型"
	}
}

func CheckMajorFuncName() string {
	return "major"
}
