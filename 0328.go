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