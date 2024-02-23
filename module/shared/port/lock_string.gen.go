// Code generated by "stringer -type=Lock -trimprefix=Lock -output=lock_string.gen.go"; DO NOT EDIT.

package sharedport

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[LockNone-0]
	_ = x[LockShared-1]
	_ = x[LockExclusive-2]
}

const _Lock_name = "NoneSharedExclusive"

var _Lock_index = [...]uint8{0, 4, 10, 19}

func (i Lock) String() string {
	if i < 0 || i >= Lock(len(_Lock_index)-1) {
		return "Lock(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Lock_name[_Lock_index[i]:_Lock_index[i+1]]
}