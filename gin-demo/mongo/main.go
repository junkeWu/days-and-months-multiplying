package main

type Student struct {
	Name string `json:"name,omitempty"`
	Age  int    `json:"age,omitempty"`
}

func main() {
}

func (s *Student) InsertOne() {
	mog.Database("")
}
