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
	f := firego.New("https://shining-heat-7158.firebaseio.com/chatlog")

	if r.Method == "GET" {
		var v map[string]interface{}
		if err := f.Value(&v); err != nil {
			c.Errorf("%s", err.Error())
			fmt.Fprintf(w, "%s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "%s", v)
	}

	if r.Method == "POST" {
		v := map[string]string{"name": "gae", "text": r.FormValue("text")}
		pushedFirego, err := f.Push(v)
		if err != nil {
			c.Errorf("%s", err.Error())
			fmt.Fprintf(w, "%s", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var result map[string]string
		if err := pushedFirego.Value(&result); err != nil {
			c.Errorf("%s", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "%s", err)
			return
		}

		fmt.Fprintf(w, "%s: %s\n", pushedFirego, result)
	}

}
