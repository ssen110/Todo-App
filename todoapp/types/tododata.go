package types

type TODODETAILS struct {
	Id   int `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	Completed bool `json:"completed"`

}