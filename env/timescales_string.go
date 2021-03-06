// Code generated by "stringer -type=TimeScales"; DO NOT EDIT.

package env

import (
	"errors"
	"strconv"
)

var _ = errors.New("dummy error")

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Event-0]
	_ = x[Trial-1]
	_ = x[Tick-2]
	_ = x[Sequence-3]
	_ = x[Block-4]
	_ = x[Condition-5]
	_ = x[Epoch-6]
	_ = x[Run-7]
	_ = x[Expt-8]
	_ = x[Scene-9]
	_ = x[Episode-10]
	_ = x[TimeScalesN-11]
}

const _TimeScales_name = "EventTrialTickSequenceBlockConditionEpochRunExptSceneEpisodeTimeScalesN"

var _TimeScales_index = [...]uint8{0, 5, 10, 14, 22, 27, 36, 41, 44, 48, 53, 60, 71}

func (i TimeScales) String() string {
	if i < 0 || i >= TimeScales(len(_TimeScales_index)-1) {
		return "TimeScales(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _TimeScales_name[_TimeScales_index[i]:_TimeScales_index[i+1]]
}

func (i *TimeScales) FromString(s string) error {
	for j := 0; j < len(_TimeScales_index)-1; j++ {
		if s == _TimeScales_name[_TimeScales_index[j]:_TimeScales_index[j+1]] {
			*i = TimeScales(j)
			return nil
		}
	}
	return errors.New("String: " + s + " is not a valid option for type: TimeScales")
}
