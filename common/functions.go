package common

func GetEndpoint(base string, endpoint string) string {
	return "/" + base + "/" + endpoint
}

func Filter[T any](items []*T, predicate func(*T) bool) []*T{
	var results []*T

	for _, item := range items{
		if predicate(item){
			results = append(results, item)
		}
	}

	return results
}
