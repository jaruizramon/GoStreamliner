package main

import (
	"Gostreamliner/actions"
	"fmt"
)

func main() {

	actions.GetAdbCoords()
	fmt.Println("FINISHED!")

	//rawContent, _ := ioutil.ReadFile("executerMP.csv")
	//ioContent := strings.NewReader(string(rawContent))
	//
	//df := dataframe.ReadCSV(ioContent)
	//
	//appendedRefList := []int{}
	//appendedEpicList := []string{}
	//
	//idx := 0
	//for idx != df.Nrow()-1 {
	//
	//	epic := df.Elem(idx, 0).String()
	//	if len(epic) != 0 || epic != "NaN" {
	//		appendedEpicList = append(appendedEpicList, epic)
	//	}
	//
	//	cavity := df.Elem(idx, 1).String()
	//	text := df.Elem(idx, 2).String()
	//
	//	refno, err := df.Elem(idx, 5).Int()
	//	if err != nil {
	//		refno = 0
	//	}
	//	if refno != 0 {
	//		appendedRefList = append(appendedRefList, refno)
	//
	//	}
	//
	//	comment := df.Elem(idx, 4).String()
	//
	//	sleep := df.Elem(idx, 5).Float()
	//	if sleep < 0 {
	//		sleep = 0
	//	}
	//
	//	fmt.Printf("%d. {%s , %s , %s , %d , %s , %f} \n", idx, epic, cavity, text, refno, comment, sleep)
	//	idx = idx + 1
	//}
	//
	//fmt.Println(appendedRefList)
	//fmt.Println(appendedEpicList)

}
