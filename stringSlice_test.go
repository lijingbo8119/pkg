package pkg_test

import (
	"fmt"
	"github.com/lijingbo8119/pkg"
	"testing"
)

func TestStringSlice_Length(t *testing.T) {
	s := pkg.NewStringSlice("a", "b", "c", "d")
	fmt.Println(s)
}
