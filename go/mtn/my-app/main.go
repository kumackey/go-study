package main

import (
	"context"
	"embed"
	_ "embed"
	"fmt"
	"github.com/labstack/echo/v4"
	"io"
	"io/fs"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"sort"
	"sync"
	"time"
)

func main() {
	argsFunc("a", "b", "c")
}

//

// 可変調引数
func argsFunc(args ...string) {
	for _, arg := range args {
		fmt.Println(arg)
	}
}

// embedでdirectory
//
//go:embed static
var local embed.FS

func embedDirFunc2() {
	e := echo.New()
	e.GET("/", echo.WrapHandler(http.FileServer(http.FS(local))))
	e.Logger.Fatal(e.Start(":8989"))
}

func embedDirFunc() {
	entries, err := local.ReadDir("static")
	if err != nil {
		log.Fatal(err)
	}
	for _, entry := range entries {
		in, err := local.Open(path.Join("static", entry.Name()))
		if err != nil {
			log.Fatal(err)
		}
		out, err := os.Create("embed-" + path.Base(entry.Name()))
		if err != nil {
			log.Fatal(err)
		}
		io.Copy(out, in)
		out.Close()
		in.Close()
		log.Println("wrote", "embed-"+path.Base(entry.Name()))
	}
}

// embedのstringの例
//
//go:embed static/message.txt
var message string

func messageFunc() {
	fmt.Println(message)
}

//go:embed static/logo.png
var contents []byte

// embedの例
func embedFunc() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.Blob(http.StatusOK, "image/png", contents)
	})
	e.Logger.Fatal(e.Start(":8989"))
}

// contextによる中断
func contextFunc() {
	var wg sync.WaitGroup
	wg.Add(1)
	ctx, cancel := context.WithCancel(context.Background())
	go f(ctx, &wg)

	time.Sleep(3 * time.Second)
	cancel()
	wg.Wait()
}

func f(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			// 何か処理
		}
		fmt.Println("goroutine: 処理")
		time.Sleep(1 * time.Second)
	}
}

// os.Openの例
func osOpenFunc() error {
	f, err := os.Open("test.txt")
	if err != nil {
		return fmt.Errorf("os.Open: %v", err)
	}
	defer f.Close()

	err = os.Mkdir("test", 0755)
	if err != nil {
		return fmt.Errorf("os.Mkdir: %v", err)
	}
	defer os.RemoveAll("test")

	return nil
}

// WalkDirの例
func walkDirFunc() {
	files := []string{}
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	err = filepath.WalkDir(cwd, func(path string, info os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			if info.Name()[0] == '.' {
				return fs.SkipDir
			}
			return nil
		}
		files = append(files, path)
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(files)
	fmt.Println(len(files))
}

// filepathパッケージの例
func filepathFunc() {
	fmt.Println(filepath.Base("C:/path/to/file.txt"))     // file.txt
	fmt.Println(filepath.Dir("C:/path/to/file.txt"))      // C:/path/to
	fmt.Println(filepath.Clean(`C:/path/to/../file.txt`)) // C:/path/file.txt
	fmt.Println(filepath.Ext("C:/path/to/file.txt"))      // .txt
	fmt.Println(filepath.IsAbs(`C:\path\to\file.txt`))    // true
	fmt.Println(filepath.IsAbs(`.\file.txt`))             // false
	fmt.Println(filepath.Join("C:/path/to", "file.txt"))  // C:/path/to/file.txt
	absolute, err := filepath.Abs("../file.txt")
	if err != nil {
		println(absolute)
	}

	absolute, err = filepath.Rel(`C:\path`, `C:\path\to\file.txt`)
	if err != nil {
		println(absolute)
	}

}

// time.Duration
func timeDurationFunc() {
	d, err := time.ParseDuration("3s")
	if err != nil {
		log.Fatal(err)
	}
	// 型と値を、[1]を使って出力
	fmt.Printf("%[1]T: %[1]v\n", d)

	d, err = time.ParseDuration("4m")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%[1]T: %[1]v\n", d)

	d, err = time.ParseDuration("5h")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%[1]T: %[1]v\n", d)
}

// log.Fatalの例
func logFatalFunc() {
	f, err := os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	log.Println("app started")
	log.Fatalln("fatal error")

	// 到達しない
	log.Println("app finished")
}

// log package
func logFunc() {
	f, err := os.OpenFile("test.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0600)
	if err != nil {
		panic(err)
	}
	log.SetOutput(f)
	log.Println("app started2")
}

// F Stringerインターフェースの使用
type F struct {
	Name string
	Age  int
}

func (f *F) String() string {
	return fmt.Sprintf("NAME=%q,AGE=%d", f.Name, f.Age)
}

// %vと%+Vと%#vの違い
func printFunc() {
	// 構造体
	t := &struct {
		a int
		b float64
		c string
	}{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)

}

type serverParam struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

func callBuilder() {
	logger := log.New(os.Stdout, "server: ", log.LstdFlags)
	srv := NewBuilder("localhost", 8080).
		TimeOut(2 * time.Second).
		Logger(logger).Build()
	srv.Start()
}

func NewBuilder(host string, port int) *serverParam {
	return &serverParam{host: host, port: port}
}

func (b *serverParam) TimeOut(timeout time.Duration) *serverParam {
	b.timeout = timeout
	return b
}

func (b *serverParam) Logger(logger *log.Logger) *serverParam {
	b.logger = logger
	return b
}

func (b *serverParam) Build() *BServer {
	return &BServer{param: *b}
}

func (b *BServer) Start() error {
	if b.param.logger != nil {
		b.param.logger.Printf("start server on %s", b.param.host)
	}

	fmt.Println(b.param.host)
	return nil
}

// Builder pattern
type BServer struct {
	param serverParam
}

// serverの呼び出し
func callServer() {
	logger := log.New(os.Stdout, "server: ", log.LstdFlags)
	srv := NewServer("localhost", 8080, WithTimeOut(5*time.Second), WithLogger(logger))
	srv.Start()
}

type Server struct {
	host    string
	port    int
	timeout time.Duration
	logger  *log.Logger
}

func NewServer(host string, port int, options ...Option) *Server {
	srv := &Server{host: host, port: port}
	for _, option := range options {
		option(srv)
	}
	return srv
}

func WithTimeOut(timeout time.Duration) Option {
	return func(s *Server) {
		s.timeout = timeout
	}
}

func WithLogger(logger *log.Logger) Option {
	return func(s *Server) {
		s.logger = logger
	}
}

type Option func(*Server)

func (s *Server) Start() error {
	fmt.Printf("start server on %s:%d\n", s.host, s.port)
	return nil
}

// recover
func recoverFunc() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Printf("%[1]T: %[1]s\n", err)
		}
	}()
	var a [2]int
	n := 2
	println(a[n])
}

func useChannel() {
	var s string
	ch := make(chan string)
	go server(ch)

	s = <-ch
	fmt.Println(s)
	s = <-ch
	fmt.Println(s)
	s = <-ch
	fmt.Println(s)
}

// 　channelの例
func server(ch chan string) {
	defer close(ch)
	ch <- "hello"
	ch <- "world"
	ch <- "!"
}

// race conditionが発生する
func raceCondition() {
	n := 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1000; i++ {
			mu.Lock()
			n++
			mu.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println(n)
}

// ループ内のGoroutineで無名関数に渡す
func gorutineLoop2() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}

// ループ内のGoroutine
func gorutineLoop() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		v := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(v)
		}()
	}
	wg.Wait()
}

// gorutineのキャプチャのタイミング
func gorutineCapture() {
	msg := "hi"
	go func() {
		sendMsg(msg)
	}()
	msg = "hello"

	sendMsg(msg)
	time.Sleep(time.Second)
	sendMsg(msg)
	time.Sleep(time.Second)
}

func sendMsg(msg string) {
	fmt.Println(msg)
}

// deferに無名関数を渡した場合の挙動の違い
func deferFunc2() {
	n := 1
	defer func() {
		fmt.Println(n)
	}()
	n = 2
}

// deferに無名関数を渡した場合の挙動の違い
func deferFunc3() {
	n := 1
	defer fmt.Println(n)
	n = 2
}

// deferによる後処理
func deferFunc() {
	defer fmt.Println("6")
	defer fmt.Println("5")
	defer fmt.Println("4")
	fmt.Println("1")
	fmt.Println("2")
	fmt.Println("3")
}

type Speaker interface {
	Speak() error
}

func DoSpeak(s Speaker) error {
	return s.Speak()
}

type Dog struct{}
type Cat struct{}

func (d Dog) Speak() error {
	fmt.Println("わんわん")
	return nil
}

func (c *Cat) Speak() error {
	fmt.Println("にゃーにゃー")
	return nil
}

// 型アサーション
func typeAssertion() {
	var x interface{} = 3
	i := x.(int)
	fmt.Println(i)
	x = "test"
	s := x.(string)
	fmt.Println(s)
	m := x.(int) // panic
	fmt.Println(m)
}

// 型アサーションのpanicを回避する
func typeAssertion2() {
	var x interface{} = 3
	if i, isInt := x.(int); !isInt {
		fmt.Println("not int")
	} else {
		fmt.Println(i)
	}
}

// 型スイッチ
func typeSwitch(x interface{}) {
	switch t := x.(type) {
	case int, int32, int64:
		fmt.Println("int", t)
	case string:
		fmt.Println("string", t)
	default:
		fmt.Println("default")
	}
}

// reflectパッケージによるtype switch
func typeSwitch2(x interface{}) {
	rt := reflect.TypeOf(x)
	switch rt.Kind() {
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("int", rt)
	case reflect.String:
		fmt.Println("string", rt)
	default:
		fmt.Println("default")
	}
}

// 文字列を書き換える
func overwriteString() {
	s := "こんにちは"
	b := []rune(s)
	println(b)
	b[0] = 'c'
	s2 := string(b)
	fmt.Println(s2)
}

// 固定長の配列をスライスに変換する
func convertArrayToSlice() {
	a := [3]int{1, 2, 3}
	fmt.Printf("%T\n", a)
	s := a[:]
	fmt.Println(s)
	fmt.Printf("%T\n", s)
}

// 複数行のテキスト
func multipleLineText() {
	var content = `Hello
World`
	fmt.Println(content)
}

// mapのfor-rangeは毎回順序が変わる
func mapRange() {
	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	for key, value := range m {
		fmt.Println(key, value)
	}
}

// もし順番通りに実行したいなら、先にkeyを取得してからfor-rangeを回す
func mapRange2() {
	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	keys := make([]string, 0, len(m))
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	for _, key := range keys {
		fmt.Println(key, m[key])
	}
}

// mapは存在しないキーを指定すると、値ゼロの型がかえる
func mapZeroValue() {
	m := map[string]int{
		"Mike":  20,
		"Nancy": 24,
		"Messi": 30,
	}
	v, ok := m["zoo"]
	fmt.Println(v, ok)
}

// ポインタの実体に代入
func pointer() {
	v := 1
	p := &v
	*p = 100
	fmt.Println(v)
}

type User struct {
	Name string
}

func (u User) Add(n string) {
	u.Name += n
}
func (u *User) AddP(n string) {
	u.Name += n
}

// ユーザ名を追加しようとする
func addUserName() {
	u := User{Name: "Taro"}
	u.Add("Yamada")
	fmt.Println(u.Name)
	u.AddP("Yamada")
	fmt.Println(u.Name)
}

func pointer2() {
	user := new(User)
	fmt.Printf("%T\n", user)
	user = &User{}
	fmt.Printf("%T\n", user)
}
