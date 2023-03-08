package main

func rotate(movements int, dir int, list *[]int) {
	if movements > len(*list) {
		movements -= len(*list)
	}
	switch dir {
	case 0:
		*list = append((*list)[movements:], (*list)[:movements]...)
	case 1:
		*list = append((*list)[len(*list)-movements:], (*list)[:len(*list)-movements]...)
	default:
		println("The dir variable must be 0 for left or 1 for right rotation")
	}
}

func main() {
	print("List before rotation: ")
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	for i := 0; i < len(nums); i++ {
		print(nums[i], " ")
	}
	rotate(2, 1, &nums)
	print("\nList after rotation: ")
	for i := 0; i < len(nums); i++ {
		print(nums[i], " ")
	}
}
