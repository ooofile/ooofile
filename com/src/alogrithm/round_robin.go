package alogrithm

/**
 * Round Robin
 */
func SimpleSchedule(index *int, total int) {
	*index = (*index + 1) % total
}

/**
 * Weighted Round Robin
 */
func WeightSchedule(index, p_cw *int, total, gcd, max_weight int, getWeight func(index int) int) {
	i := *index
	cw := *p_cw
	for {
		i = (i + 1) % total
		if i == 0 {
			cw = cw - gcd
			if cw <= 0 {
				cw = max_weight
				if cw == 0 { //if max weight is 0
					i = *index
					break
				}
			}
		}
		weight := getWeight(i)
		if weight >= cw {
			break
		}
	}
	*p_cw = cw
	*index = i
}
