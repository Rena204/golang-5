
package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

# Using a JSON file for data persistence
# This is a basic example of how to serialize and deserialize data in Go

package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name string
	Age  int
}

func main() {
	user := User{"Alice", 30}

	file, _ := json.MarshalIndent(user, "", " ")

	_ = os.WriteFile("user.json", file, 0644)

	byteValue, _ := os.ReadFile("user.json")

	var user2 User

	json.Unmarshal(byteValue, &user2)

	fmt.Println(user2.Name)
	fmt.Println(user2.Age)
}

# Setting up a proxy with nginx
# This is a basic example of an nginx configuration file

http {
    upstream backend {
        server localhost:8080;
        server localhost:8081;
    }

    server {
        listen 80;

        location / {
            proxy_pass http://backend;
        }
    }
}

# Testing the application
# This is a basic example of a unit test in Go

package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	request, _ := http.NewRequest(http.MethodGet, "", nil)
	response := httptest.NewRecorder()

	handler(response, request)

	got := response.Body.String()
	want := "Hello, !"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

