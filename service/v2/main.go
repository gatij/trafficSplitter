package main

import (
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "strings"
)

type User struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    Version string `json:"version"`
    Email   string `json:"email"`   // Additional field for v2
}

func main() {
    // Register HTTP handler
    http.HandleFunc("/users/", handleGetUser)

    fmt.Println("Starting server for V2 on :8082...")
    log.Fatal(http.ListenAndServe(":8082", nil))
}

func handleGetUser(w http.ResponseWriter, r *http.Request) {
    // Extract user ID from the URL
    pathParts := strings.Split(r.URL.Path, "/")

    if len(pathParts) != 3 {
        http.Error(w, "Invalid path. Use /users/{id}", http.StatusBadRequest)
        return
    }

    // Assuming the user ID is the third part of the path
    userIDFromUrl := pathParts[2]

    // Convert string ID to int
    userID, err := strconv.Atoi(userIDFromUrl)
    if err != nil {
        http.Error(w, "Invalid user ID. Must be a number.", http.StatusBadRequest)
        return
    }

    // User Response
    // Note the differences from v1: Name is "John Doe" and we added an email field
    user := User{
        ID:      userID,
        Name:    "John Doe",
        Version: "v2",
        Email:   "john@example.com",
    }

    // Set response headers and write the user data as JSON
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(user)

    fmt.Println("Handled request for user ID from V2 for:", userID)
}