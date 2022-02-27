package main

import "fmt"

func fillArray(arr [10]int) [10]int {

	for i := 0; i < len(arr); i++ {
		arr[i] = i
		p := &arr[i]
		fmt.Println(arr, arr[i], p)
	}
	fmt.Println(arr)
	return arr

}

func main() {
	var array = [3]string{ //создание массива из 3 элементов
		"hello",
		"Worls",
		"!!!",
	}

	fmt.Println(array)
	fmt.Println(array[0])
	fmt.Println(array[1])
	fmt.Println(array[2])

	fmt.Println(len(array)) //размер массива

	for index, item := range array {
		fmt.Println("индекс", index, "значение", item)
	}

	array2 := [7]int64{}
	fmt.Println(len(array2))
	var j int64 = 0
	for i := 0; i <= len(array2)-1; i++ {

		if i == 2 {

			continue
		} else if i == 5 {
			break
		}
		array2[i] = j
		j++
	}
	fmt.Println(array2)

	//null array

	var arr [10]int
	fmt.Println(arr)
	fillArray(arr)
	fmt.Println(arr)
	var arr2 [10]int
	arr2 = fillArray(arr)
	fmt.Println(arr2)
}
