package todo

type Todo struct {
	Id    string `json:"id" bson:"_id,omitempty"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}
