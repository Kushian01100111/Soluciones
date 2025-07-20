package main

import (
	"strings"
)

func finishMerge(i int, mergedArr []string, indexArr int, arr []string) []string {
	for j := i; j < len(mergedArr) && indexArr < len(arr); j++ {
		mergedArr[j] = arr[indexArr]
		indexArr++
	}
	return mergedArr
}

func merge(arr1 []string, arr2 []string) []string {
	mergedArr := make([]string, (len(arr1) + len(arr2)))

	indexArr1, indexArr2 := 0, 0
	for i := 0; i < len(mergedArr) && indexArr1 < len(arr1) && indexArr2 < len(arr2); i++ {
		if arr1[indexArr1] <= arr2[indexArr2] {
			mergedArr[i] = arr1[indexArr1]
			indexArr1++
			if indexArr1 == len(arr1) {
				mergedArr = finishMerge(i+1, mergedArr, indexArr2, arr2)
				break
			}
		} else if arr1[indexArr1] > arr2[indexArr2] {
			mergedArr[i] = arr2[indexArr2]
			indexArr2++
			if indexArr2 == len(arr2) {
				mergedArr = finishMerge(i+1, mergedArr, indexArr1, arr1)
				break
			}
		}
	}
	return mergedArr
}

func mergeSort(folder []string) []string {
	if len(folder) <= 1 {
		return folder
	}

	if len(folder) == 2 {
		if folder[0] > folder[1] {
			return []string{folder[1], folder[0]}
		}
		return folder
	}

	middle := len(folder) / 2
	left := mergeSort(folder[:middle])
	right := mergeSort(folder[middle:])
	return merge(left, right)
}

func removeSubfolders(folder []string) []string {
	sortedFolder := mergeSort(folder)
	mainFolders := make([]string, 0)

	currentPrefix := "-"
	for i := 0; i < len(sortedFolder); i++ {
		comparation := sortedFolder[i] == currentPrefix || strings.HasPrefix(sortedFolder[i], currentPrefix+"/")
		if !comparation {
			currentPrefix = sortedFolder[i]
			mainFolders = append(mainFolders, sortedFolder[i])
		}
	}

	return mainFolders
}
