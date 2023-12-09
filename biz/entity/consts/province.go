package consts

import "github.com/pkg/errors"

//go:generate stringer -type Province -linecomment

// Province 省份
type Province int

const (
	ZHE_JIANG      Province = iota // 浙江
	SHANG_HAI                      // 上海
	BEI_JING                       // 北京
	TIAN_JIN                       // 天津
	CHONG_QING                     // 重庆
	HEI_LONG_JIANG                 // 黑龙江
	JI_LIN                         // 吉林
	LIAO_NING                      // 辽宁
	NEI_MENG_GU                    // 内蒙古
	HE_BEI                         // 河北
	XIN_JIANG                      // 新疆
	GAN_SU                         // 甘肃
	QING_HAI                       // 青海
	SH_AN_XI                       // 陕西
	NING_XIA                       // 宁夏
	HE_NAN                         // 河南
	SHAN_DONG                      // 山东
	SHAN_XI                        // 山西
	AN_HUI                         // 安徽
	HU_BEI                         // 湖北
	HU_NAN                         // 湖南
	JIANG_SU                       // 江苏
	SI_CHUAN                       // 四川
	GUI_ZHOU                       // 贵州
	YUN_NAN                        // 云南
	GUANG_XI                       // 广西
	XI_ZANG                        // 西藏
	JIANG_XI                       // 江西
	GUANG_DONG                     // 广东
	FU_JIAN                        // 福建
	TAI_WAN                        // 台湾
	HAI_NAN                        // 海南
	XIANG_GANG                     // 香港
	AO_MEN                         // 澳门
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
	return p, errors.Errorf("not a valid province, %q", p)
}

func CheckProvinceFuncName() string {
	return "province"
}
