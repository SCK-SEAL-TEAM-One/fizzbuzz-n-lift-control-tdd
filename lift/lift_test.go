package lift

import "testing"

var gotoFloor = 3
var weight = 70.00
var maximumWeightCapacityValue = 1000.00

func Test_current_1_not_overload_goto_floor_3_should_up_and_stop_at_floor_3(t *testing.T) {
	lift := Lift{currentFloor: 1, maximumWeightCapacity: maximumWeightCapacityValue}
	expectedState := "up"

	actualState := lift.Move(weight, gotoFloor)

	isEqual(expectedState, actualState, t)
}

func Test_current_2_not_overload_goto_floor_3_should_up_and_stop_at_floor_3(t *testing.T) {
	lift := Lift{currentFloor: 2, maximumWeightCapacity: maximumWeightCapacityValue}
	expectedState := "up"

	actualState := lift.Move(weight, gotoFloor)

	isEqual(expectedState, actualState, t)
}

func Test_current_3_not_overload_goto_floor_3_should_stop_at_floor_3(t *testing.T) {
	lift := Lift{currentFloor: 3, maximumWeightCapacity: maximumWeightCapacityValue}
	expectedState := "stop"

	actualState := lift.Move(weight, gotoFloor)

	isEqual(expectedState, actualState, t)

}

func isEqual(expectedState string, actualState string, t *testing.T) {
	if expectedState != actualState {
		t.Error("expected : ", expectedState, "but it is ", actualState)
	}
}
