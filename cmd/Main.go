package main

import "fmt"

const (
	countofinputs      = 7
	countOfNodesLayer1 = 10
	countOfNodesLayer2 = 10
	groupSize          = 40
	variant            = 3
)

func main() {
	mindeviation := uint32(4000)
	mindevall := uint64(27000000)
	deviationtest := uint32(0)
	ts := loadtimeslots()

	ml := GenerateMinModel(ts)
	filltimeslots(&ml, ts, 0, groupSize)
	deviation := calculateAll(&ml)
	fmt.Printf("model was made, initial deviation = %d\n", deviation)
	for j := 0; j < 155; j++ {
		var devall uint64
		for i := range ts {
			if i > len(ts)-60 {
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
			SaveModel(&ml, "t")
			mindeviation = deviationtest
			fmt.Printf("!!! test answer = %d\n", testInput(&ml))
		} else if devall < mindevall {
			SaveModel(&ml, "a")
			fmt.Printf("!!! test answer = %d\n", testInput(&ml))
			mindevall = devall
		}
	}

}
