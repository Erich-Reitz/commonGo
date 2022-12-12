package advent 


func MinOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min > i {
			min = i
		}
	}

	return min
}

func MaxOf(vars ...int) int {
	min := vars[0]

	for _, i := range vars {
		if min < i {
			min = i
		}
	}

	return min
}

func AbsInt(x int) int {
	return AbsDiffInt(x, 0)
 }
 
 func AbsDiffInt(x, y int) int {
	if x < y {
	   return y - x
	}
	return x - y
 }
 
 func AbsDiffUint(x, y uint) uint {
	if x < y {
	   return y - x
	}
	return x - y
 }
 
