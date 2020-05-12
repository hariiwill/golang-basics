package main

import (
	"fmt"
	"sort"
)

/*func sort(nums ...int) []int {

	/*for i := 0; i < len(nums); i++ {

		for j := i + 1; j < len(nums); j++ {

			if nums[j] > nums[i] {

				temp := nums[j]
				nums[j] = nums[i]
				nums[i] = temp
			}

		}

	}

	//sort.Ints(nums)
	return nums

}*/

func main() {

	// fmt.Println("Sort numbers in ascending order")
	// nums := []int{45, 223, 76, 109, 87, 31, 567, 90, 132, 198}
	// sort.Ints(nums)
	// //num := sort(nums...)
	// fmt.Println(num)
	// s := []int{4, 2, 3, 1}
	// sort.Ints(s)
	// fmt.Println(s)
	list := []string{"vihaan", "subhi", "hari", "mano", "mani"}
	sort.Strings(list)
	fmt.Println(list)
}
