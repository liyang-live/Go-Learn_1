package main

func main() {

	//p:=models.Person{Id:10,Name:"Test",Age:10,Sex:1}

	//fmt.Println(p.Name)

	router := initRouter()
	router.Run(":10080")
}
