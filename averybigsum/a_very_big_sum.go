package averybigsum

// AVeryBigSum Complete the aVeryBigSum function below.
func AVeryBigSum(ar []int64) int64 {
	var sum int64
	for _, v := range ar {
		sum += v
	}
	return sum
}
