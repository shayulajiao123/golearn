package main

import (
	"fmt"
	"time"
)

func main(){
	var tz string
	tz = "+08:00:00.000000"
	t, err := time.Parse("-07:00:00", tz)
	if err!= nil {
		fmt.Println(err)
		fmt.Println(t)
	}
}

