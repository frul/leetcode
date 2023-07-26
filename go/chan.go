// impotrs

func main() {
    sum := 0
    var sync.WaitGroup wg
    c := make(chan int, 1000)

    for i := 0; i < 100; i++ {
       wg.Add(1)
       go foo(c, done)
    }
    
    go func() {
        wg.Wait()
        close(c)
    }
    
    for r := range c {
        sum += r
    }
    
    fmt.Println(sum)
}

func foo(c chan int, wg *sync.WaitGroup ) {
    defer wg.Done()
    r := rand.Int()
    for i := 0; i < r; i++ {
        c <- r
    }
}