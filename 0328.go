package main 

import "fmt"
import "math/rand"
import "time"

func main() {
    rand.Seed(time.Now().Unix())
    
    for i := 1; i <= 3; i++ {
        fmt.Printf("%d回目のおみくじ結果: ", i)
        
        number := rand.Intn(6)
        switch number {
            case 0:
                fmt.Println("凶です")
            case 1, 2:
                fmt.Println("吉です")
            case 3, 4:
                fmt.Println("中吉です")
            case 5:
                fmt.Println("大吉です")
        }
    }
}



package main

func main() {
    // 数値の7を出力してください
    println(7)
    
    // 9に3足した値を出力してください
    println(9 + 3)
    
    // 「9 + 3」を文字列として出力してください   
    println("9 + 3")
    
}

package main

func main() {
    // 12を3で割った値を出力してください
    println(12 / 3)
    
    // 3に6を掛けた値を出力してください
    println(3 * 6)
    
    // 8を3で割った時の余りを出力してください
    println(8 % 3)
}



package main

func main() {
    //「こんにちは世界」となるように文字列を連結して出力してください
    println("こんにちは" + "世界")
    
    // 文字列"38"と"19"を連結して出力してください
    println("38" + "19")
    
    // 数値38に19を足して出力してください
    println(38 + 19)

}


package main

func main() {
    // 変数messageを定義し、「Hello, 世界」を代入してください
    var message string = "Hello, 世界"
    
    // 変数messageの値を出力してください
    println(message)
    
}


package main

func main () {
    var message string = "Hello, 世界"

    // 変数messageの値を、文字列"Hello, Go"に更新してください
	message = "Hello, Go"
	
	println(message)
}


package main

func main() {
    // 変数messageを省略して定義し、"Hello, 世界"を代入してください
    message := "Hello, 世界"
    
    // 変数numberを省略して定義し、100を代入してください
    number := 100
    
    // 変数messageと変数numberを同時に出力してください
    println(message, number)
    
}

package main

func main() {
    // s1とs2の定義がまだされていなく、出力できないので消してください
    
    
    s1 := "Hello"
    s2 := "世界"
    
    // 変数定義から代入に変更してください
    s1 = "こんにちは"

    // 変数s1はString型なので、消してください
    
    
    // 変数s1とs2を同時に出力してください
    println(s1, s2)
    
}

package main

func main() {
    n := 100
    // nから10を引いてnに代入しましょう
    n -= 10
    
    println(n)
    
    
    m := 10
    // mに1を足してmに代入しましょう
    m++
    
    println(m)
}
