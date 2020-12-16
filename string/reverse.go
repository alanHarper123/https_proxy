package string

func Reverse(s string) string{
  data := []rune(s)
  for i:=0; i<len(data)/2;i++{
    j:=len(data)-1-i
    data[i],data[j] = data[j], data[i]
  }
  return string(data)
}
