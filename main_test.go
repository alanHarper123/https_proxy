package main
import "github.com/alanHarper123/https_proxy/string"
import "testing"
func Test(t *testing.T) {
  if string.Reverse("12345")!="54321"{
    t.Error("gg")
  }

}
