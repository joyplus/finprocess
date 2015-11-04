package tests
import (
	"time"
	"fmt"
	"testing"
)


//go test -v time_test.go
func TestTime(t *testing.T){
	fmt.Println(time.Now().Unix())
}