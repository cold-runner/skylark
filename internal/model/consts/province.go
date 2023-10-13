package consts

import "github.com/pkg/errors"

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

func ParseProvince(province string) (Province, error) {
	switch {
	case province == "浙江":
		return ZHE_JIANG, nil
	case province == "上海":
		return SHANG_HAI, nil
	case province == "北京":
		return BEI_JING, nil
	case province == "天津":
		return TIAN_JIN, nil
	case province == "重庆":
		return CHONG_QING, nil
	case province == "黑龙江":
		return HEI_LONG_JIANG, nil
	case province == "吉林":
		return JI_LIN, nil
	case province == "辽宁":
		return LIAO_NING, nil
	case province == "内蒙古":
		return NEI_MENG_GU, nil
	case province == "河北":
		return HE_BEI, nil
	case province == "新疆":
		return XIN_JIANG, nil
	case province == "甘肃":
		return GAN_SU, nil
	case province == "青海":
		return QING_HAI, nil
	case province == "陕西":
		return SH_AN_XI, nil
	case province == "宁夏":
		return NING_XIA, nil
	case province == "河南":
		return HE_NAN, nil
	case province == "山东":
		return SHAN_DONG, nil
	case province == "山西":
		return SHAN_XI, nil
	case province == "安徽":
		return AN_HUI, nil
	case province == "湖北":
		return HU_BEI, nil
	case province == "湖南":
		return HU_NAN, nil
	case province == "江苏":
		return JIANG_SU, nil
	case province == "四川":
		return SI_CHUAN, nil
	case province == "贵州":
		return GUI_ZHOU, nil
	case province == "云南":
		return YUN_NAN, nil
	case province == "广西":
		return GUANG_XI, nil
	case province == "西藏":
		return XI_ZANG, nil
	case province == "江西":
		return JIANG_XI, nil
	case province == "广东":
		return GUANG_DONG, nil
	case province == "福建":
		return FU_JIAN, nil
	case province == "台湾":
		return TAI_WAN, nil
	case province == "海南":
		return HAI_NAN, nil
	case province == "香港":
		return XIANG_GANG, nil
	case province == "澳门":
		return AO_MEN, nil
	}

	var p Province
	return p, errors.Errorf("not a valid provice, %q", p)
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
		return "未知的省份"
	}
}

func CheckProvinceFuncName() string {
	return "province"
}
