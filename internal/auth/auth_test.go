package auth

import (
    "testing"
    "net/http"
)

func TestGetAPIKey(t *testing.T){
    apiKey := "8213adsa21a213ax"
    headers := http.Header{
        "Authorization": []string{"ApiKe " + apiKey},
    }
    got,err := GetAPIKey(headers)
    if err !=nil{
        t.Fatalf("Unexpected error calling GetAPIKey, %v",err)
    }
    want := apiKey
    if want != got {
        t.Fatalf("expected: %v, got %v",want,got)
    }
}
