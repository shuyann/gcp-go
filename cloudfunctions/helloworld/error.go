package helloworld

import (
	"fmt"
	"net/http"
	"os"
)

func HTTPError(w http.ResponseWriter, r *http.Request) {
	fmt.Println("An error occurred (stdout)")
	fmt.Fprintln(os.Stderr, "An error occurred (stderr)")

	panic("An error occurred (panic)")
}
