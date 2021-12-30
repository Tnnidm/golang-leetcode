package parkingSystem

type ParkingSystem struct {
	park []int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	ps := ParkingSystem{}
	ps.park = []int{big, medium, small}
	return ps
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.park[carType-1] != 0 {
		this.park[carType-1]--
		return true
	}
	return false
}

/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
