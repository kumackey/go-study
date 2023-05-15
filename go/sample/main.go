package main

// 英小文字のみからなる文字列Sが与えられます。Sの各文字を英大文字に変換して得られる文字列Tを返す関数
func convert(s string) string {
	var t string
	for _, r := range s {
		t += string(r - 32)
	}
	return t
}

func main() {
	s := "hello"
	t := convert(s)
	println(t)
}
