package lift

type Lift struct {
	currentFloor          int
	maximumWeightCapacity float64
}

func (l Lift) Move(weight float64, targetFloor int) string {

	if l.isWeightOverload(weight) {
		return "stop"
	}

	return l.directState(targetFloor)
}

func (l Lift) isWeightOverload(weight float64) bool {
	return weight > l.maximumWeightCapacity
}

func (l Lift) directState(targetFloor int) string {
	if l.currentFloor == targetFloor {
		return "stop"
	}

	if l.currentFloor < targetFloor {
		return "up"
	}

	return "down"
}
