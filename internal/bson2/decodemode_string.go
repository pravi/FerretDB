// Code generated by "stringer -linecomment -type decodeMode"; DO NOT EDIT.

package bson2

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[decodeShallow-1]
	_ = x[decodeDeep-2]
	_ = x[decodeCheckOnly-3]
}

const _decodeMode_name = "decodeShallowdecodeDeepdecodeCheckOnly"

var _decodeMode_index = [...]uint8{0, 13, 23, 38}

func (i decodeMode) String() string {
	i -= 1
	if i < 0 || i >= decodeMode(len(_decodeMode_index)-1) {
		return "decodeMode(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _decodeMode_name[_decodeMode_index[i]:_decodeMode_index[i+1]]
}