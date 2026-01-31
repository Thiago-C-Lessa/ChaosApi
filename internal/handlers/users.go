package handlers

import (
	"net/http"
)

func UsersHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`[{"id":1,"name":"Alice"},{"id":2,"name":"Bob"}]`))
}
