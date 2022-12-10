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

func absInt(x int) int {
	return absDiffInt(x, 0)
 }
 
 func absDiffInt(x, y int) int {
	if x < y {
	   return y - x
	}
	return x - y
 }
 
 func absDiffUint(x, y uint) uint {
	if x < y {
	   return y - x
	}
	return x - y
 }
 