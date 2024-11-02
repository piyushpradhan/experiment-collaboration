package types

type Message struct {
	Sender string `json:"sender"`
	Color string `json:"color"`
	Name string `json:"username"`
	X float64 `json:"x"`
	Y float64 `json:"y"`
}