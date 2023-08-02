// Code generated by "stringer -type=Status status.go"; DO NOT EDIT.

package moduletest

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Pending-0]
	_ = x[Skip-1]
	_ = x[Pass-2]
	_ = x[Fail-3]
	_ = x[Error-4]
}

const _Status_name = "PendingSkipPassFailError"

var _Status_index = [...]uint8{0, 7, 11, 15, 19, 24}

func (i Status) String() string {
	if i < 0 || i >= Status(len(_Status_index)-1) {
		return "Status(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Status_name[_Status_index[i]:_Status_index[i+1]]
}
