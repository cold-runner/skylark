// Code generated by "stringer -type Grade -linecomment"; DO NOT EDIT.

package consts

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[FRESHMAN-0]
	_ = x[SOPHOMORE-1]
	_ = x[JUNIOR-2]
	_ = x[SENIOR-3]
	_ = x[GRADUATES-4]
}

const _Grade_name = "大一大二大三大四毕业生"

var _Grade_index = [...]uint8{0, 6, 12, 18, 24, 33}

func (i Grade) String() string {
	if i < 0 || i >= Grade(len(_Grade_index)-1) {
		return "Grade(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Grade_name[_Grade_index[i]:_Grade_index[i+1]]
}
