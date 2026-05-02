package main

type FilterOptions struct {
	MatchCodes []int	// -mc: HTTP Statuscode muốn giữ
	FilterCodes  []int	// -fc: HTTP statuscode muốn bỏ
	FilterSizes []int64 // -fs: các content-length muốn bỏ
}

func (f *FilterOptions)IsValid(result Result) bool {
	if containInt(f.FilterCodes, result.StatusCode){
		return false
	}

	if containInt64(f.FilterSizes, result.ContentLength){
		return false
	}

	if len(f.MatchCodes) > 0 {
		if !containInt(f.MatchCodes, result.StatusCode){
			return false
		}
	}

	return true
}


func containInt(slice []int, val int) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func containInt64(slice []int64, val int64) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}