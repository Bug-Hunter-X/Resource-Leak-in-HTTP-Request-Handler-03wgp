func handleRequest(w http.ResponseWriter, r *http.Request) {
    // ... other code ...
    defer r.Body.Close()
    // ... more code ...
    // Handle any errors during the request processing, including closing the request body if needed. 
    if err != nil {
        http.Error(w, "Internal Server Error", http.StatusInternalServerError)
        return
    }
} 