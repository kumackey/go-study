package main

import (
	"database/sql"
	"embed"
	"fmt"
	"github.com/gocraft/web"
	"github.com/jhillyerd/enmime"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"io"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"path"
	"regexp"
	"strings"
	"time"
)

func main() {
	defineFunc()
}

// smtp
func smtpFunc() {
	smtpHost := "my-mail-server:25"
	smtpAuth := smtp.PlainAuth(
		"example.com",
		"example-user",
		"example-password",
		"auth.example.com",
	)

	sender := enmime.NewSMTP(smtpHost, smtpAuth)
	master := enmime.Builder().
		From("宗慎太郎", "taro@example.com").
		Subject("件名").
		Text([]byte("本文")).
		HTML([]byte("<b>本文</b>")).
		AddFileAttachment("document.pdf")

	msg := master.To("宛先花子", "hanako@example.com")
	err := msg.Send(sender)
	if err != nil {
		log.Fatal(err)
	}
}

// 　define
func defineFunc() {
	t := template.Must(template.New("").ParseGlob("web5/*.tmpl"))
	err := t.ExecuteTemplate(os.Stdout, "index", "これは本文です。")
	if err != nil {
		log.Fatal(err)
	}
}

func templateParseFiles() {
	t := template.Must(template.ParseFiles("web5/tmpl.html", "web5/tmpljs.html"))
	data := struct {
		Value string
	}{
		Value: "Hello",
	}
	err := t.Execute(os.Stdout, data)
	if err != nil {
		log.Fatal(err)
	}
}

func templateFuncMap() {
	t := template.New("").Funcs(template.FuncMap{
		"FormatDatetime": func(format string, d time.Time) string {
			if d.IsZero() {
				return ""
			}
			return d.Format(format)
		}})
	tmpl := `{{FormatDatetime "2006/01/02" .}}`
	t = template.Must(t.Parse(tmpl))
	err := t.Execute(os.Stdout, time.Now())
	if err != nil {
		log.Fatal(err)
	}
}

//go:embed tmpljs.html
var tmplJS string

func templateJSFunc() {
	t := template.Must(template.New("").Parse(tmplJS))
	err := t.Execute(os.Stdout, template.JS(`alert("<script>1</script>")`))
	if err != nil {
		log.Fatal(err)
	}
}

//go:embed tmpl.html
var tmplHtml string

// template.HTMLの使用
func templateHTMLFunc() {
	t := template.Must(template.New("").Parse(tmplHtml))
	err := t.Execute(os.Stdout, template.HTML(`<b>HTML</b>`))
	if err != nil {
		log.Fatal(err)
	}
}

type Employee struct {
	Name string
}
type Company struct {
	Employees []Employee
}

//go:embed company.txt
var f embed.FS

func companyExecuteFunc() {
	tmpl, err := template.ParseFS(f, "company.txt")
	if err != nil {
		log.Fatal(err)
	}
	company := Company{
		Employees: []Employee{
			{Name: "Alice"},
			{Name: "Bob"},
			{Name: ""},
		},
	}
	err = tmpl.Execute(os.Stdout, company)
	if err != nil {
		log.Fatal(err)
	}
}

// if
func ifFunc() {
	tmpl := `{{if gt .Age 20}}
{{.Name}} is older than 20
{{else}}
{{.Name}} is not older than 20
{{end}}`
	t := template.Must(template.New("").Parse(tmpl))
	user := User{Name: "Taro", Age: 19}
	err := t.Execute(os.Stdout, user)
	if err != nil {
		log.Fatal(err)
	}
}

// range/end
func rangeFunc() {
	tmpl := `{{range .}}
<p>{{.}}</p>{{end}}
{{index . 1}}`
	t := template.Must(template.New("").Parse(tmpl))
	values := []string{"a", "b", "c"}
	err := t.Execute(os.Stdout, values)
	if err != nil {
		log.Fatal(err)
	}
}

// html/templateの例
func templateFunc() {
	tmpl := `{{.Name}}`
	t := template.Must(template.New("").Parse(tmpl))
	user := User{Name: "Taro"}
	err := t.Execute(os.Stdout, user)
	if err != nil {
		log.Fatal(err)
	}

}

type User struct {
	Name string
	Age  int
}

// オリジナルのcontent-typeを使いたい場合
func fileServeMyContentTypeFunc() {
	const prefix = "/public/"
	fileserver := http.StripPrefix(prefix, http.FileServer(http.Dir("./best4/static")))
	http.HandleFunc(prefix, func(w http.ResponseWriter, r *http.Request) {
		if typ, ok := mimeMap[path.Ext(r.URL.Path)]; ok {
			w.Header().Set("Content-Type", typ)
		}
		fileserver.ServeHTTP(w, r)
	})
	http.ListenAndServe(":8080", nil)
}

var mimeMap = map[string]string{
	".xls":  "application/vnd.ms-excel",
	".xlsx": "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet",
	".ppt":  "application/vnd.ms-powerpoint",
	".doc":  "application/msword",
	".docx": "application/vnd.openxmlformats-officedocument.wordprocessingml.document",
}

// fileのserve
func fileServeFunc() {
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("./best4/static"))))
	http.ListenAndServe(":8080", nil)
}

type routerParam map[string]string
type routerFunc func(routerParam, http.ResponseWriter, *http.Request)
type routerItem struct {
	method  string
	matcher *regexp.Regexp
	fnc     routerFunc
}
type router struct {
	items []routerItem
}

func (rt *router) GET(prefix string, fnc routerFunc) {
	rt.items = append(rt.items, routerItem{
		method:  http.MethodGet,
		matcher: regexp.MustCompile("^" + prefix + "$"),
		fnc:     fnc,
	})
}

func (rt *router) POST(prefix string, fnc routerFunc) {
	rt.items = append(rt.items, routerItem{
		method:  http.MethodPost,
		matcher: regexp.MustCompile("^" + prefix + "$"),
		fnc:     fnc,
	})
}

func (rt *router) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, v := range rt.items {
		if v.method == r.Method && v.matcher.MatchString(r.RequestURI) {
			match := v.matcher.FindStringSubmatch(r.RequestURI)
			params := make(map[string]string)
			for i, name := range v.matcher.SubexpNames() {
				params[name] = match[i]
			}
			v.fnc(params, w, r)
			return
		}
	}
	http.NotFound(w, r)
}

func webfunc() {
	router := web.New(AppContext{}).
		Middleware(web.LoggerMiddleware).
		Middleware(web.ShowErrorsMiddleware).
		Middleware((*AppContext).SetHelloCount).
		Get("/", (*AppContext).SayHello)

	http.ListenAndServe(":8080", router)
}

type AppContext struct {
	HelloCount int
}

func (c *AppContext) SetHelloCount(rw web.ResponseWriter, req *web.Request, next web.NextMiddlewareFunc) {
	c.HelloCount = 3
	next(rw, req)
}

func (c *AppContext) SayHello(rw web.ResponseWriter, req *web.Request) {
	fmt.Fprint(rw, strings.Repeat("Hello ", c.HelloCount), "World!")
}

func routerfunc() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/hello/:name", Hello)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Hello World!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

// method value
type I int

func (i I) Add(n int) I {
	return i + I(n)
}

func methodValueFunc() {
	i := I(1)
	n := i.Add(1).Add(2)
	fmt.Println(n)

	add := n.Add
	fmt.Println(add(3))

	fmt.Printf("%T\n", n.Add)
	fmt.Printf("%T\n", I.Add)

	fmt.Println(I.Add(n, 4))
}

func NewMyContext() *MyContext {
	return &MyContext{}
}

type MyContext struct {
	db *sql.DB
}

func (m *MyContext) handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World! db")
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World!")
}

func myHandler2(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		fmt.Fprint(w, "GET Hello World!")
	default:
		f, err := os.Open("./web5/content.txt")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer f.Close()
		io.Copy(w, f)
	}
}
