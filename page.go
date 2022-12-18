package linkding

type Page[T any] struct {
	Count    int64  `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Items    []T    `json:"results"`
}
