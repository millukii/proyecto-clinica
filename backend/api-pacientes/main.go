package main

import (
	"api-pacientes/internal/model"
	"fmt"
)

func main() {

	session := model.NewMockSession()
	documents, _ := session.DB("test").C("other_test").GetMyDocuments()

	fmt.Println(documents)
}
