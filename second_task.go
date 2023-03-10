/*Изучая древние надписи, историк Василий обнаружил, что все они состоят из n символов, каждый из которых является нулем или единицей.
Как всем известно, чтобы найти количество лет между датами появления двух надписей, необходимо взять все позиции (нумеруются с единицы), в которых строки отличаются, и сложить. Например, строки «011» и «110» отличаются в позициях 1 и 3, поэтому количество лет между их написанием составляет 4 года.
Хоть надписи, над которыми работает Василий, достаточно велики, в них присутствуют достаточно большие блоки из подряд идущих одинаковых символов. Чтобы с ними было удобнее рабо- тать, Василий выписал две последовательности — в одной из них содержатся символы, находящиеся в блоках, а во второй — размеры самих блоков. Например, надпись «111001111» Василий может закодировать как ({1, 0, 1, 1}, {3, 2, 1, 3}) (хотя это далеко не единственный способ кодирования).
Василий очень боится допустить вычислительную ошибку, поэтому обратился за помощью к вам. По последовательностям, кодирующим две надписи, определите, сколько лет прошло между их появлением.

Формат входных данных
Первая строка содержит число n (1 \leqslant   n \leqslant   2 · 10^5)(1⩽n⩽2⋅105) — количество блоков в первой надписи.

Вторая строка содержит символы s_1, s_2, . . . , s_n (s_i ∈ {0, 1})s1,s2,...,sn(si∈0,1), записанные через пробел, где s_isi— символ, которому равны все символы в i-м блоке первой надписи.
Третья строка содержит числа a_1, a_2, . . . , a_n (1 \leqslant   a_i \leqslant   10^9),a1,a2,...,an(1⩽ai⩽109), где a_iai — количество символов в i-м блоке первой надписи.
Следующие три строки содержат информацию о второй надписи в том же формате. Гарантируется, что надписи содержат одинаковое количество символов, не превосходящее 110^9109


Формат выходных данных
Выведите количество лет, прошедшее между появлением этих надписей.*/

package main

import "fmt"

func main() {
	var num int
	fmt.Scan(&num)

	firstSlice := firstGive(num)
	secondSlice := firstGive(num)
	bigFirstSlice := []int{}

	for i := 0; i < num; i++ {
		bigFirstSlice = makeGig(bigFirstSlice, firstSlice[i], secondSlice[i])
	}

	fmt.Scan(&num)

	firstSlice = firstGive(num)
	secondSlice = firstGive(num)
	bigSecondSlice := []int{}
	for i := 0; i < num; i++ {
		bigSecondSlice = makeGig(bigSecondSlice, firstSlice[i], secondSlice[i])
	}

	counter := 0
	for i := 0; i < len(bigFirstSlice); i++ {
		if bigFirstSlice[i] != bigSecondSlice[i] {
			counter += i + 1
		}
	}
	fmt.Println(counter)
}

func makeGig(someSlice []int, someIntToAdd, num int) []int {
	for i := 0; i < num; i++ {
		someSlice = append(someSlice, someIntToAdd)
	}
	return someSlice
}

func firstGive(num int) []int {
	var cNum int
	someSlice := []int{}

	for i := 0; i < num; i++ {
		fmt.Scan(&cNum)
		someSlice = append(someSlice, cNum)
	}

	return someSlice
}
