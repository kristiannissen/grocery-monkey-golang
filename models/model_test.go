package models

import (
    "testing"
)

var tests []struct {
    name, got, want string
}{
    {"GetUser", "Kitty", "Hello Kitty"}
}

// Setup
func init() {
    
}

func TestGetUser(t *testing.T) {
    for _, tt := range tests {
        if m.GetUser(tt.got) != tt.want {
            t.Errorf("Test %s: Got %s - want %s", tt.name, tt.got, tt.want))
        }
    }
}
