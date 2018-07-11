package main

import (
  "log"
  "net/http"
)

func main() {
  fs := http.FileServer(http.Dir("static"))
  http.Handle("/", fs)

  log.Println("Run me as Admin, i want to host port 9002")
  log.Println("Try the following to test me: curl  http://localhost:9002/dhcpstatus.xml")
  log.Println("Listening...")
  log.Fatal(http.ListenAndServe(":9002", nil))
}
