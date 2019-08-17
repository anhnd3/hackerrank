package bonappetit

import "fmt"

// BonAppetit ...
func BonAppetit(bill []int32, k int32, b int32) {
	var actual int32
	length := len(bill)
	for index := 0; index < length; index++ {
		if int32(index) == k {
			continue
		}
		actual += bill[index]
	}
	actual /= 2
	if actual == b {
		fmt.Println("Bon Appetit")
		return
	}
	fmt.Println(b - actual)
}
