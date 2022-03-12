package models

import (
    "testing"
)

// Setup
func init() {
    
}

func TestGetUser(t *testing.T) {
    tests := []struct {
        name, got, want string
    }{
        {"GetUser", "Kitty", "Kitty"},
    }

    m := Model{}

    for _, tt := range tests {
        if m.GetUser(tt.got).NickName != tt.want {
            t.Errorf("Test %s: Got %s - want %s", tt.name, m.GetUser(tt.got).NickName, tt.want)
        }
    }
}
