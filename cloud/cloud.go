package cloud

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("handler", Handler)
}

func Handler(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Capy is here")
		return
	}
	if d.Name == "" {
		fmt.Fprint(w, "Capy is here")
		return
	}
	fmt.Fprintf(w, "Hello, %s!", html.EscapeString(d.Name))
}
