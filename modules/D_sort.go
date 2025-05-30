package modules

func SortBySalary() {
	var i, j, maxIdx int
	var temp JobListing

	for i = 0; i < jobCount-1; i++ {
		maxIdx = i

		for j = i + 1; j < jobCount; j++ {
			if jobListings[j].Salary > jobListings[maxIdx].Salary {
				maxIdx = j
			}
		}
		temp = jobListings[i]
		jobListings[i] = jobListings[maxIdx]
		jobListings[maxIdx] = temp
	}
}

// func SortByReleveance() {
// 	var i, j int
// 	var key JobListing

// 	for i = 0; i < jobCount; i++ {
// 		key = jobListings[i]

// 		for j = i - 1; j >= 0 && jobListings[j].Relevance < key.Relevance; j++ {
// 			jobListings[j+1] = jobListings[j]
// 		}
// 		jobListings[j+1] = key
// 	}
// }
