package utils

/*
func main() {
	slice1 := []string{"apple", "banana", "orange"}
	slice2 := []string{"banana", "grape", "apple"}

	merged := mergeAndRemoveDuplicates(slice1, slice2)

	fmt.Println("Merged and deduplicated slice:", merged)
}
*/

func MergeAndRemoveDuplicates(slice1 []string, slice2 []string) []string {
	result := []string{}

	// 遍历第一个切片，并将每个字符串添加到结果切片中（如果它还不存在于结果切片中）
	for _, str := range slice1 {
		if !contains(result, str) {
			result = append(result, str)
		}
	}

	// 遍历第二个切片，检查每个字符串是否已存在于结果切片中，如果不存在，则将其添加到结果切片中
	for _, str := range slice2 {
		if !contains(result, str) {
			result = append(result, str)
		}
	}

	return result
}

// 去除 A 中 与 B 重复的字符串
func RemoveDuplicates(A, B []string) []string {
	result := []string{}

	for _, str := range A {
		if !contains(B, str) {
			result = append(result, str)
		}
	}

	return result
}

// 检查切片中是否包含指定的字符串
func contains(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}
