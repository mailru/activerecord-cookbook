package serializer

func SearchQueryMarshal(data map[string]string) ([]string, error) {
	ret := make([]string, 0, 2*len(data))
	for k, v := range data {
		ret = append(ret, k, v)
	}
	return ret, nil
}
