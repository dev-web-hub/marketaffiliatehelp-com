package api
import "net/http"

func DigestHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Digest endpoint — will render weekly posts"))
}
