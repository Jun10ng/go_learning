package main

import (
	"fmt"
	"github.com/bitly/go-simplejson"
)

func main() {
	m, _ := simplejson.NewJson([]byte(raw))
	k := m.Get("movie_info")
	//ej := &simplejson.Json{}
	if k.Interface() == nil {
		fmt.Println("te")
	}
	fmt.Println(k)

	fmt.Println(m.GetPath("fulfillment_result", "movie").Get("booking_number"))
}
