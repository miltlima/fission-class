package main

import (
   "net/http"
)

func Handler(w http.ResponseWriter, *http.Request) {
    msg := "Hello World"
    w.Write([]byte(msg))
} 
