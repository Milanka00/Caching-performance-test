package main

import (
    "fmt"
    "net/http"
    "os"
    "strconv"
    "sync"
    "time"
)

var payload []byte
var once sync.Once

func main() {
    
    generatePayload()

   
    http.HandleFunc("/nocache", func(w http.ResponseWriter, r *http.Request) {
        NoCacheHandler(w, r)
    })
    http.HandleFunc("/publiccache", func(w http.ResponseWriter, r *http.Request) {
        PublicCacheHandler(w, r)
    })
    http.HandleFunc("/privatecache", func(w http.ResponseWriter, r *http.Request) {
        PrivateCacheHandler(w, r)
    })
    http.HandleFunc("/getresponse", func(w http.ResponseWriter, r *http.Request) {
        getresponseWithoutHeaders(w, r)
    })

    // Start server
    fmt.Println("Server is listening on port 8081...")
    http.ListenAndServe(":8081", nil)
}


func NoCacheHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Cache-Control", "no-store")
    w.Write(payload)
    
}


func PublicCacheHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Cache-Control", "public, max-age=60")
    w.Write(payload)
    // additionalContent := []byte("cached as public")
    // w.Write(additionalContent)
    
}


func PrivateCacheHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Cache-Control", "private, max-age=3600")
    w.Write(payload)
    
}

func getresponseWithoutHeaders(w http.ResponseWriter, r *http.Request) {
    w.Write(payload)
    
}

func sleepBeforeRespond() {
    sleepTimeStr := os.Getenv("SLEEP_TIME")
    sleepTime, err := strconv.Atoi(sleepTimeStr)
    if err != nil {
        sleepTime = 15
    }
    time.Sleep(time.Duration(sleepTime) * time.Second)
}

func generatePayload() {
    once.Do(func() {
        payload = make([]byte, 100)
        for i := range payload {
            payload[i] = 'x'
        }
    })
}
