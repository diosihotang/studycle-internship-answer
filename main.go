package main

import "fmt"

func main() {
	array := []int{3, 4, 12, 1}

	array = BubbleSort(array)

	fmt.Printf("Sorted Array : %+v \n", array)

	sum := 0
	times := 0

	for _, v := range array {
		sum += v

		if times == 0 {
			times = v
		}

		times *= v
	}

	fmt.Printf("Nilai rata-rata : %f \n", float64(sum)/float64(len(array)))
	fmt.Printf("Semua bilangan di kalikan : %d \n", times)

	// bila total bilangan genap, nilai tengah merupakan jumlah dari kedua bilangan tengah di bagi 2
	if len(array)%2 == 0 {
		fmt.Printf("Nilai Tengah : %f \n", (float64(array[len(array)/2])+float64(array[(len(array)/2)-1]))/2)
	} else {
		fmt.Printf("Nilai Tengah : %d \n", array[(len(array)-1)/2])
	}
}

func BubbleSort(array []int) []int {
	for i := 0; i < len(array)-1; i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}
	return array
}
