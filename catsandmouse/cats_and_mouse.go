package catsandmouse

// CatAndMouse ...
func CatAndMouse(x int32, y int32, z int32) string {
	distanceA := x - z
	if distanceA < 0 {
		distanceA *= -1
	}

	distanceB := y - z
	if distanceB < 0 {
		distanceB *= -1
	}

	if distanceA > distanceB {
		return "Cat B"
	} else if distanceB > distanceA {
		return "Cat A"
	} else {
		return "Mouse C"
	}
}
