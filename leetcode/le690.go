package main

import "fmt"

type Employee struct {
     Id int
     Importance int
     Subordinates []int
 }

func getImportance(employees []*Employee, id int) int {
 	maps := make(map[int]int)
	importence :=0
 	var getall func(int)
 	for i,e :=range employees{
 		maps[e.Id] = i
	}
 	getall = func(i int)  {
 		e := employees[maps[i]]
 		importence += e.Importance
 		for _,v := range e.Subordinates {
 			getall(v)
		}
	}
	getall(1)
 	return importence
}

func main() {
	fmt.Println(getImportance([]*Employee{
		&Employee{Id: 1,Importance: 2,Subordinates: []int{2,3}},
			&Employee{Id: 2,Importance: 2,Subordinates: []int{}},
		&Employee{Id: 3,Importance: 2,Subordinates: []int{}},
		},
		1))
}