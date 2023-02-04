/*Каждый раз, возвращаясь с работы, Василий вынужден отстоять большую очередь, чтобы сесть в маршрутку.
В очереди перед Василием находится n человек. Когда человек заходит в маршрутку, он идет к месту, находящемуся как можно ближе к середине. Если таких мест несколько, человек идет к месту с меньшим номером. Например, если в маршрутке всего 6 мест и места с номерами 3, 4 заняты, то следующий пассажир сядет на место под номером 2.
Когда в маршрутке кончаются места, она уезжает, а на ее место приезжает новая. Василий по очереди записывал, на какие места садятся пассажиры. Что записал Василий?


Формат входных данных
Единственная строка содержит числа n (1⩽n⩽5∗10^5)(1⩽n⩽5∗105)  и m (1⩽m⩽5∗10^5)(1⩽m⩽5∗105) — количество пассажиров, стоящих в очереди, и вместимость маршрутки соответственно.

Формат выходных данных
В i-йi−й строке выведите, на какое место сел i-йi−й пассажир.*/

package main

import (
	"fmt"
)

func main() {
	var passengerNumber, capacity int
	fmt.Scan(&passengerNumber, &capacity)
	someMap := []int{}
	for i := 0; i < capacity; i++ {
		someMap = append(someMap, (i + 1))
	}

	numsCp := makeCp(someMap)
	for i := 0; i < passengerNumber; i++ {
		var avg int
		if len(numsCp)%2 == 1 {
			avg = len(numsCp) / 2
		} else {
			avg = len(numsCp)/2 - 1
		}

		fmt.Println(numsCp[avg])
		numsCp = removeByIndex(numsCp, avg)
		if len(numsCp) == 0 {
			numsCp = makeCp(someMap)
		}
	}
}

func removeByIndex(array []int, index int) []int {
	return append(array[:index], array[index+1:]...)
}

func makeCp(someSlice []int) []int {
	numsCp := make([]int, len(someSlice))

	copy(numsCp, someSlice)
	return numsCp
}
