import (
	"log"
	"net/http"
)

func helloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", helloFunc)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
