package hosts

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

func TestGetUserHomeDir(t *testing.T) {
	fmt.Println(GetUserHomeDir())
}
//system32\drivers\etc
func TestGetWinSystemDir(t *testing.T) {

	pp := filepath.Join(getWinSystemDir(),"system32","drivers","etc","hosts")
	body, _ := ioutil.ReadFile(pp)
	fmt.Println(string(body))
}