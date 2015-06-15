package firebase_playground

import (
	"appengine"
	"fmt"
	"github.com/CloudCom/firego"
	"net/http"
)

func init() {
	http.HandleFunc("/hello", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	f := firego.New("https://shining-heat-7158.firebaseio.com/")

	var v map[string]interface{}
	if err := f.Value(&v); err != nil {
		c.Errorf("%v", err.Error())
	}
	fmt.Fprintf(w, "%s", v)
}
