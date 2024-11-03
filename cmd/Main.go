package main

import "fmt"

const (
	countofinputs      = 7
	countOfNodesLayer1 = 10
	countOfNodesLayer2 = 10
	groupSize          = 20
)

func main() {
	mindeviation := uint32(2000)
	ts := loadtimeslots()
	ml := generateModel()
	deviationtest := uint32(0)
	filltimeslots(&ml, ts, 0, groupSize)
	deviation := calculateAll(&ml)
	fmt.Printf("model was made, initial deviation = %d\n", deviation)
	for j := 0; j < 155; j++ {
		var devall uint64
		for i := range ts {
			if i > len(ts)-50 {
				break
			}
			filltimeslots(&ml, ts, i, groupSize)
			learn(&ml, 2)
			deviation = calculateAll(&ml)
			devall += uint64(deviation)
			//fmt.Printf("dv = %d\n ", deviation)
		}
		for i := range ts {
			if i < len(ts)-40 {
				continue
			}
			filltimeslots(&ml, ts, i, groupSize)
			deviationtest = calculateAll(&ml)
			break
			//fmt.Printf("dv = %d\n ", deviation)
		}

		fmt.Printf("************************ gen %d deviation = %d avg = %d test = %d ************************\n", j, devall, devall/uint64(len(ts)), deviationtest)
		if deviationtest < mindeviation {
			SaveModel(&ml)
			mindeviation = deviationtest
		}
	}

}
