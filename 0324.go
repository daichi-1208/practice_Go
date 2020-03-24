package main

import (
	"fmt"
	"net/http"
	"time"
)

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World! %s", time.Now())
}

func main() {
	http.HandleFunc("/", greet)
	http.ListenAndServe(":8080", nil)
}


package main

import "fmt"

// 引数なし関数
func sayHi() {
    fmt.Println("hi!")
}

// 引数あり関数
func sayName(name string) {
    fmt.Println(name)
}

// return関数
func getMessage(name string) string {
    msg := "hi! my name is " + name
    return msg
}

// 名前付きreturn関数
// 関数内で使った変数名を返す
func getHelloMessage(name string) (msg string) {
    msg = "Hello " + name
    return
}

// メイン関数
func main() {
    sayHi()
    sayName("gcfuji")
    fmt.Println(getMessage("gcfuji"))
    fmt.Println(getHelloMessage("Gemcook"))
}

package main

import "fmt"

// 複数の返り値を返す事ができる
func swap(a int, b int) (int, int) {
    return b, a
}

func main() {
    fmt.Println(swap(5, 2))

    // 関数もデータ型なので、変数に代入可能
    // その際は関数名はいらない
    f := func(a int, b int) (int, int) {
        return b, a
    }

    fmt.Println(f(3, 8))

    // 即時関数的な事も可能
    func(msg string) {
        fmt.Println(msg)
    }("Fujimoto")
}



