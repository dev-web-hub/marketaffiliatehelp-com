package api
import "net/http"

func SignupHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Signup endpoint — future subscriber storage"))
}
