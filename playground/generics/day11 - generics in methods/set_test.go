package set

import "testing"

func TestSet(t *testing.T) {
	//preparing mock data
	mockIntA := Set[int]{Items: map[int]struct{}{1: {}, 3: {}, 4: {}, 5: {}}}
	mockIntB := Set[int]{Items: map[int]struct{}{2: {}, 3: {}, 4: {}, 6: {}}}
	mockIntC := Set[int]{Items: map[int]struct{}{3: {}, 4: {}}}

	mockStringA := Set[string]{Items: map[string]struct{}{"aaa": {}, "bbb": {}, "ccc": {}, "ddd": {}}}
	mockStringB := Set[string]{Items: map[string]struct{}{"AAA": {}, "bbb": {}, "ccc": {}, "DDD": {}}}
	mockStringC := Set[string]{Items: map[string]struct{}{"bbb": {}, "ccc": {}}}

	mockRuneA := Set[rune]{Items: map[rune]struct{}{'A': {}, 'B': {}, 'C': {}, 'D': {}}}
	mockRuneB := Set[rune]{Items: map[rune]struct{}{'a': {}, 'B': {}, 'C': {}, 'd': {}}}
	mockRuneC := Set[rune]{Items: map[rune]struct{}{'B': {}, 'C': {}}}

	//starting tests

	//Contains
	if !mockIntA.Contains(4) || !mockStringA.Contains("ccc") || !mockRuneA.Contains('C') {
		t.Error("Contains-method doesn't work correclty with EXISTING items")
	}
	if mockIntA.Contains(100) || mockStringA.Contains("fgh") || mockRuneA.Contains('Z') {
		t.Error("Contains-method doesn't work correclty with NON-EXISTING items")
	}

	//IsEmpty
	if mockIntA.IsEmpty() || mockRuneA.IsEmpty() || mockStringA.IsEmpty() {
		t.Error("IsEmpty-method doesn't work correctly with non-empty structures")
	}
	//Len
	if mockIntA.Len() != 4 || mockRuneA.Len() != 4 || mockStringA.Len() != 4 {
		t.Error("Len-method doesn't work properly")
	}
	//Add
	mockIntA.Add(11)
	mockStringA.Add("eee")
	mockRuneA.Add('E')
	if mockIntA.Len() != 5 || mockRuneA.Len() != 5 || mockStringA.Len() != 5 {
		t.Error("Len-method doesn't work properly after Add-method")
	}
	//Remove
	mockIntA.Remove(11)
	mockStringA.Remove("eee")
	mockRuneA.Remove('E')
	if mockIntA.Len() != 4 || mockRuneA.Len() != 4 || mockStringA.Len() != 4 {
		t.Error("Len-method doesn't work properly after Remove-method")
	}
	if mockIntA.Contains(11) || mockStringA.Contains("eee") || mockRuneA.Contains('E') {
		t.Error("Contains-method returned TRUE for the items removed by Remove-method")
	}
	//ToSlice
	sliceIntsA := mockIntA.ToSlice()
	sliceStringsA := mockStringA.ToSlice()
	sliceRunesA := mockRuneA.ToSlice()
	if mockIntA.Len() != len(sliceIntsA) || mockStringA.Len() != len(sliceStringsA) || mockRuneA.Len() != len(sliceRunesA) {
		t.Error("ToSlice-method returned different number than the actual size of the original map")
	} else {
		for _, i := range sliceIntsA {
			if _, ok := mockIntA.Items[i]; !ok {
				t.Error("ToSlice-method returned INTs-slice with element, non-existing in the original map")
			}
		}
		for _, i := range sliceStringsA {
			if _, ok := mockStringA.Items[i]; !ok {
				t.Error("ToSlice-method returned STRINGs-slice with element, non-existing in the original map")
			}
		}
		for _, i := range sliceRunesA {
			if _, ok := mockRuneA.Items[i]; !ok {
				t.Error("ToSlice-method returned RUNEs-slice with element, non-existing in the original map")
			}
		}
	}
	//Union
	unionInts := mockIntA.Union(&mockIntB)
	unionStrings := mockStringA.Union(&mockStringB)
	unionRunes := mockRuneA.Union(&mockRuneB)
	for i := range unionInts.Items {
		if !mockIntA.Contains(i) && !mockIntB.Contains(i) {
			t.Error("Union-method returned INTs-structure with item non-existing in source-structures")
			break
		}
	}
	for i := range unionStrings.Items {
		if !mockStringA.Contains(i) && !mockStringB.Contains(i) {
			t.Error("Union-method returned STRINGs-structure with item non-existing in source-structures")
			break
		}
	}
	for i := range unionRunes.Items {
		if !mockRuneA.Contains(i) && !mockRuneB.Contains(i) {
			t.Error("Union-method returned RUNEs-structure with item non-existing in source-structures")
			break
		}
	}
	//Intersection
	intersectionInts := mockIntA.Intersection(&mockIntB)
	intersectionStrings := mockStringA.Intersection(&mockStringB)
	intersectionRunes := mockRuneA.Intersection(&mockRuneB)
	for i := range intersectionInts.Items {
		if !mockIntA.Contains(i) || !mockIntB.Contains(i) {
			t.Error("Intersection-method returned INTs-structure with item non-existing in original structures")
			break
		}
	}
	for i := range intersectionStrings.Items {
		if !mockStringA.Contains(i) || !mockStringB.Contains(i) {
			t.Error("Intersection-method returned STRINGs-structure with item non-existing in original structures")
			break
		}
	}
	for i := range intersectionRunes.Items {
		if !mockRuneA.Contains(i) || !mockRuneB.Contains(i) {
			t.Error("Intersection-method returned RUNEs-structure with item non-existing in original structures")
			break
		}
	}
	//Difference
	differenceInts := mockIntA.Difference(&mockIntB)
	differenceStrings := mockStringA.Difference(&mockStringB)
	differenceRunes := mockRuneA.Difference(&mockRuneB)
	for i := range differenceInts.Items {
		if !mockIntA.Contains(i) || mockIntB.Contains(i) {
			t.Error("Difference-method returned INTs-structure with item non-existing in original structure A, or existing in strcture B")
			break
		}
	}
	for i := range differenceStrings.Items {
		if !mockStringA.Contains(i) || mockStringB.Contains(i) {
			t.Error("Difference-method returned STRINGs-structure with item non-existing in original structure A, or existing in strcture B")
			break
		}
	}
	for i := range differenceRunes.Items {
		if !mockRuneA.Contains(i) || mockRuneB.Contains(i) {
			t.Error("Difference-method returned RUNEs-structure with item non-existing in original structure A, or existing in strcture B")
			break
		}
	}
	//IsSubset
	if !mockIntC.IsSubset(&mockIntA) || mockIntB.IsSubset(&mockIntA) {
		t.Error("IsSubset-method returned 'C is not a subset of A' or 'B is a subset of A' for INTs")
	}
	if !mockStringC.IsSubset(&mockStringA) || mockStringB.IsSubset(&mockStringA) {
		t.Error("IsSubset-method returned 'C is not a subset of A' or 'B is a subset of A' for STRINGs")
	}
	if !mockRuneC.IsSubset(&mockRuneA) || mockRuneB.IsSubset(&mockRuneA) {
		t.Error("IsSubset-method returned 'C is not a subset of A' or 'B is a subset of A' for RUNEs")
	}
	//Clear
	mockIntC.Clear()
	mockRuneC.Clear()
	mockStringC.Clear()
	if mockIntC.Len() != 00 || mockStringC.Len() != 0 || mockRuneC.Len() != 0 {
		t.Error("Clear-function didn't clear input-structure")
	}
}
