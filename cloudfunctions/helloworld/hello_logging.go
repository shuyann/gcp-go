package helloworld

import (
	"fmt"
	"log"
	"net/http"
)

func HelloLogging(w http.ResponseWriter, r *http.Request) {
	log.Println("This is stderr")
	fmt.Println("This is stdout")

	fmt.Println(`{"message": "This has ERROR severity", "severity":"error"`)
}
