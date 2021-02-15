package main

import (
	"encoding/json"
	"fmt"
)

type Movie struct {
	Title  string   `json:"title"`
	Year   int      `json:"year"`
	Price  int      `json:"rmb"`
	Actors []string `json:"actors"`
}

func main() {
	movie := Movie{"喜剧之王", 2000, 10, []string{"zhouxingchi", "zhangbozhi"}}
	//encode
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("Marshal error")
		return
	}
	fmt.Printf("str=%s\n", jsonStr)

	//decode
	//{"title":"喜剧之王","year":2000,"rmb":10,"actors":["zhouxingchi","zhangbozhi"]}
	//myMoive:=Movie{}
	var myMoive Movie
	err = json.Unmarshal(jsonStr, &myMoive)

	if err != nil {
		fmt.Println("Unmarshal error")
		return
	}

	fmt.Println("myMoive=", myMoive)

}
