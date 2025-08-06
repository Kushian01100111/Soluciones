package main

func finishMergeInt(i int, mergedArr []Basket, indexArr int, arr []Basket) []Basket {
	for j := i; j < len(mergedArr) && indexArr < len(arr); j++ {
		mergedArr[j] = arr[indexArr]
		indexArr++
	}
	return mergedArr
}

func mergeInt(arr1 []Basket, arr2 []Basket, criterion string) []Basket {
	mergedArr := make([]Basket, len(arr1)+len(arr2))
	indexArr1, indexArr2 := 0, 0
	comp := getComparator(criterion)

	for i := 0; i < len(mergedArr) && indexArr1 < len(arr1) && indexArr2 < len(arr2); i++ {
		if comp(arr1[indexArr1], arr2[indexArr2]) {
			mergedArr[i] = arr1[indexArr1]
			indexArr1++
		} else if comp(arr2[indexArr2], arr1[indexArr1]) {
			mergedArr[i] = arr2[indexArr2]
			indexArr2++
		} else {
			// En caso de igualdad, preferimos arr1 (izquierda)
			mergedArr[i] = arr1[indexArr1]
			indexArr1++
		}

		if indexArr1 == len(arr1) {
			mergedArr = finishMergeInt(i+1, mergedArr, indexArr2, arr2)
			break
		}
		if indexArr2 == len(arr2) {
			mergedArr = finishMergeInt(i+1, mergedArr, indexArr1, arr1)
			break
		}
	}

	return mergedArr
}

func mergeSortInt(arr []Basket, criterion string) []Basket {
	if len(arr) <= 1 {
		return arr
	}

	if len(arr) == 2 && criterion == "primero" {
		if arr[0].Val > arr[1].Val {
			return []Basket{arr[1], arr[0]}
		}
		return arr
	} else if len(arr) == 2 && criterion == "segundo" {
		if arr[0].Index > arr[1].Index {
			return []Basket{arr[1], arr[0]}
		}
		return arr
	}

	middle := len(arr) / 2
	left := mergeSortInt(arr[:middle], criterion)
	right := mergeSortInt(arr[middle:], criterion)
	return mergeInt(left, right, criterion)
}
