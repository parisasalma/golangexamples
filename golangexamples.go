package golangexamples
import "github.com/ehteshamz/greetings"

func ConcatSlice(SliceToConcat[]byte) string {
	var s string
	for _,value:=range SliceToConcat{
		s=s+string(value)+"_"
	}
	return s
}
func Encrypt(sliceToEncrypt* []byte,ceaserCount int)  {
	tmp := *sliceToEncrypt
	for i,x:=range tmp{
    ascii := int(x)
		ascii=ascii+ceaserCount
		tmp[i]=byte(ascii)
	}
}
func EZGreetings(name string) string {
	return string(greetings.PrintGreetings(name))
}
