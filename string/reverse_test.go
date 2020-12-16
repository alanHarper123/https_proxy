package string
import (
  "testing"
)
func Test(t *testing.T) {
  var tests = []struct{
    s, want string
  }{
    {"十三分","分三十"},
    {"123456","654321"},
  }
  for _,c := range tests {
    got:=Reverse(c.s)
    if got !=c.want {
      t.Errorf("Reverse(%q) should be %q, but get %q",c.s,c.want,got)
    }
  }
}
