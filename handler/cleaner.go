package handler

import (
	"fmt"
	"time"
)

func CleanTTl() {
	for {
		err := dbInstance.TTLControl()
		if err != nil {
			fmt.Println(err)
		}
		time.Sleep(time.Second)
	}
}
