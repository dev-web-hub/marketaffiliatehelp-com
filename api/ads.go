package api
import "net/http"

func AdsHandler(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("Advertiser page — $599 / $499 per month"))
}
