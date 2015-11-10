package tests

import (
	"fmt"
	"testing"
	"time"
)

//go test -v time_test.go
func TestTime(t *testing.T) {
	fmt.Println(time.Now())
}
