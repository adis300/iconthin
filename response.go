package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func sendResponse(w http.ResponseWriter, response interface{}) {
	b, err := json.Marshal(response)
	if err == nil {
		w.Write(b)
	} else {
		fmt.Printf("ERROR:JSON marshaling error:%v", err)
		w.Write([]byte(`{"error":"JSON marshaling error"}`))
	}
}
