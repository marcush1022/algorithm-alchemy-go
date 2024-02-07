package backtracking

/**
Given two integers n and k, return an array of all the integers of length n where the difference between every two consecutive digits is k.
You may return the answer in any order.

Note that the integers should not have leading zeros. Integers as 02 and 043 are not allowed.


Example 1:

Input: n = 3, k = 7
Output: [181,292,707,818,929]
Explanation: Note that 070 is not a valid number, because it has leading zeroes.

Example 2:

Input: n = 2, k = 1
Output: [10,12,21,23,32,34,43,45,54,56,65,67,76,78,87,89,98]

Constraints:

2 <= n <= 9
0 <= k <= 9
*/

func numsSameConsecDiff(n int, k int) []int {
	results := &[]int{}
	for i := 1; i <= 9; i++ {
		searchNumsSameConsec(n-1, k, i, results)
	}
	return *results
}

func searchNumsSameConsec(numbersCount, difference, currNumber int, results *[]int) {
	//fmt.Println(">>>>> numbersCount, currNumber", numbersCount, currNumber)
	if numbersCount == 0 {
		*results = append(*results, currNumber)
		return
	}

	lastDigit := currNumber % 10
	if lastDigit >= difference {
		// 8,9
		searchNumsSameConsec(numbersCount-1, difference, currNumber*10+lastDigit-difference, results)
	}
	if difference > 0 && lastDigit+difference < 10 {
		// 0,1,2,3,4,5,6
		searchNumsSameConsec(numbersCount-1, difference, currNumber*10+lastDigit+difference, results)
	}
}
