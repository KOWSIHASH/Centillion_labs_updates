package main

import "fmt"

func main() {

	var ages = [3]int{20, 25, 30}

	names := [4]string{"deep", "marko", "web", "bowser"}
	names[1] = "ksksk"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	
	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	
	rangeOne := names[1:4] 
	rangeTwo := names[2:]  
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)
	fmt.Printf("the type of rangeOne is %T \n", rangeOne)

	rangeOne = append(rangeOne, "kupdu")
	fmt.Println(rangeOne)

}