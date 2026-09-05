package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "os"
    "runtime"
    "time"
)

type SystemReport struct {
    Status      string    `json:"status"`
    Timestamp   time.Time `json:"timestamp"`
    GoVersion   string    `json:"go_version"`
    OS          string    `json:"os"`
    Architecture string   `json:"architecture"`
    Hostname    string    `json:"hostname"`
    CPUs        int       `json:"cpu_count"`
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
    hostname, _ := os.Hostname()
    report := SystemReport{
        Status:       "ONLINE & OPERATIONAL",
        Timestamp:    time.Now(),
        GoVersion:    runtime.Version(),
        OS:           runtime.GOOS,
        Architecture: runtime.GOARCH,
        Hostname:     hostname,
        CPUs:         runtime.NumCPU(),
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(report)
    fmt.Println("[LOG] Health check probe executed successfully.")
}

func main() {
    fmt.Println("==================================================")
    fmt.Println("   🚀 INITIALIZING ENTERPRISE GOLANG MICROSERVICE ")
    fmt.Println("==================================================")
    port := ":8080"
    http.HandleFunc("/health", healthHandler)
    
    fmt.Printf("[INFO] Server listening on port %s...\n", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        fmt.Printf("[ERROR] Server failed: %v\n", err)
    }
}
