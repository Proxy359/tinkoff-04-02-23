/*Василий — фермер в двумерном государстве под названием Флатландия.
У Василия есть 4n единиц забора, являющихся отрезками длины 1. Единицы забора разрешается располагать только параллельно осям O_xOx
​и O_yOy, при этом крайние точки забора должны оказаться в целочисленных точках.

Василий хочет построить ровно k загонов для овечек, имеющих размеры 1 × 1. Каждый загон обязан иметь форму квадрата.
Какое минимальное количество овечек могут вмещать k загонов, построенных Василием? Какое максимальное количество овечек могут вмещать k загонов, построенных Василием?

Формат входных данных
Первая строка содержит число n — количество четверок однометровых кусочков забора, которыми располагает Василий.
Вторая строка содержит число k — количество загонов, которое хочет построить Василий. Гарантируется, что выполнена цепочка неравенств
11 \leqslant   k \leqslant  n\leqslant  10^9.1⩽k⩽n⩽10
9
 .

Формат выходных данных
Через пробел выведите минимальную и максимальную суммарные площади загонов, которые могут иметь k загонов, сооруженных Василием.*/

package main

import (
	"fmt"
)

func main() {
	var material, points int
	fmt.Scan(&material, &points)

	if points == 1 {
		fmt.Println(material*material, material*material)
		return
	}

	min := makeMin(material, points)
	max := makeMax(material, points)
	fmt.Println(min, max)

}

func makeMax(material, points int) int {
	prev := material - points + 1
	sum := prev*prev + points - 1
	return sum
}

func makeMin(material, points int) int {
	sum := 0

	for i := 0; i < points; i++ {
		minus := material / (points - i)
		sum += minus * minus
		material = material - minus
	}
	return sum
}
