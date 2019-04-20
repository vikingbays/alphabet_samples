package module

import (
	"fmt"
	"math"
	"math/rand"
)

type User1 struct {
	Usrid   int
	Name    string
	Nanjing bool
	Money   float64
}

func GetUser1List(min int, max int) []interface{} {

	var list = make([]interface{}, 0, 1)
	counter := 0
	for i := min; i < max; i++ {
		user1 := User1{Usrid: i, Name: fmt.Sprintf("Ikkk_%d", i), Nanjing: true, Money: math.Ceil(rand.Float64()*10000) / 100}
		list = append(list, user1)
		counter = counter + 1
		if counter == 30 {
			break
		}
	}
	return list
}
