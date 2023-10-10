package consts

// Province 省份
type Province int

const (
	ZHE_JIANG Province = iota
	SHANG_HAI
	BEI_JING
	TIAN_JIN
	CHONG_QING
	HEI_LONG_JIANG
	JI_LIN
	LIAO_NING
	NEI_MENG_GU
	HE_BEI
	XIN_JIANG
	GAN_SU
	QING_HAI
	SH_AN_XI
	NING_XIA
	HE_NAN
	SHAN_DONG
	SHAN_XI
	AN_HUI
	HU_BEI
	HU_NAN
	JIANG_SU
	SI_CHUAN
	GUI_ZHOU
	YUN_NAN
	GUANG_XI
	XI_ZANG
	JIANG_XI
	GUANG_DONG
	FU_JIAN
	TAI_WAN
	HAI_NAN
	XIANG_GANG
	AO_MEN
)

var ProvinceList = map[string]Province{
	"浙江":  ZHE_JIANG,
	"上海":  SHANG_HAI,
	"北京":  BEI_JING,
	"天津":  TIAN_JIN,
	"重庆":  CHONG_QING,
	"黑龙江": HEI_LONG_JIANG,
	"吉林":  JI_LIN,
	"辽宁":  LIAO_NING,
	"内蒙古": NEI_MENG_GU,
	"河北":  HE_BEI,
	"新疆":  XIN_JIANG,
	"甘肃":  GAN_SU,
	"青海":  QING_HAI,
	"陕西":  SH_AN_XI,
	"宁夏":  NING_XIA,
	"河南":  HE_NAN,
	"山东":  SHAN_DONG,
	"山西":  SHAN_XI,
	"安徽":  AN_HUI,
	"湖北":  HU_BEI,
	"湖南":  HU_NAN,
	"江苏":  JIANG_SU,
	"四川":  SI_CHUAN,
	"贵州":  GUI_ZHOU,
	"云南":  YUN_NAN,
	"广西":  GUANG_XI,
	"西藏":  XI_ZANG,
	"江西":  JIANG_XI,
	"广东":  GUANG_DONG,
	"福建":  FU_JIAN,
	"台湾":  TAI_WAN,
	"海南":  HAI_NAN,
	"香港":  XIANG_GANG,
	"澳门":  AO_MEN,
}

func (p Province) String() string {
	switch p {
	case ZHE_JIANG:
		return "浙江"
	case SHANG_HAI:
		return "上海"
	case BEI_JING:
		return "北京"
	case TIAN_JIN:
		return "天津"
	case CHONG_QING:
		return "重庆"
	case HEI_LONG_JIANG:
		return "黑龙江"
	case JI_LIN:
		return "吉林"
	case LIAO_NING:
		return "辽宁"
	case NEI_MENG_GU:
		return "内蒙古"
	case HE_BEI:
		return "河北"
	case XIN_JIANG:
		return "新疆"
	case GAN_SU:
		return "甘肃"
	case QING_HAI:
		return "青海"
	case SH_AN_XI:
		return "陕西"
	case NING_XIA:
		return "宁夏"
	case HE_NAN:
		return "河南"
	case SHAN_DONG:
		return "山东"
	case SHAN_XI:
		return "山西"
	case AN_HUI:
		return "安徽"
	case HU_BEI:
		return "湖北"
	case HU_NAN:
		return "湖南"
	case JIANG_SU:
		return "江苏"
	case SI_CHUAN:
		return "四川"
	case GUI_ZHOU:
		return "贵州"
	case YUN_NAN:
		return "云南"
	case GUANG_XI:
		return "广西"
	case XI_ZANG:
		return "西藏"
	case JIANG_XI:
		return "江西"
	case GUANG_DONG:
		return "广东"
	case FU_JIAN:
		return "福建"
	case TAI_WAN:
		return "台湾"
	case HAI_NAN:
		return "海南"
	case XIANG_GANG:
		return "香港"
	case AO_MEN:
		return "澳门"
	default:
		return ""
	}
}

func (p Province) CheckFuncName() string {
	return "province"
}
