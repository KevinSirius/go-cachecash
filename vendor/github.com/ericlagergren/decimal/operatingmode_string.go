// Code generated by "stringer -type OperatingMode"; DO NOT EDIT.

package decimal

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[GDA-0]
	_ = x[Go-1]
}

const _OperatingMode_name = "GDAGo"

var _OperatingMode_index = [...]uint8{0, 3, 5}

func (i OperatingMode) String() string {
	if i >= OperatingMode(len(_OperatingMode_index)-1) {
		return "OperatingMode(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OperatingMode_name[_OperatingMode_index[i]:_OperatingMode_index[i+1]]
}
