package main

import (
    "context"
    "fmt"
    "net/http"
    "os"
    "os/signal"
    "sync"
    "syscall"
    "time"
)

const addr = ":80"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprint(w, "New Http Server")
    })
    //ʹ��Ĭ��·�ɴ��� http server
    srv := http.Server{
        Addr:    addr,
        Handler: http.DefaultServeMux,
    }
    //ʹ��WaitGroupͬ��Goroutine
    var wg sync.WaitGroup
    exit := make(chan os.Signal)
    //���� Ctrl+C �ź�
    signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-exit
        wg.Add(1)
        //ʹ��context����srv.Shutdown�ĳ�ʱʱ��
        ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
        defer cancel()
        err := srv.Shutdown(ctx)
        if err != nil {
            fmt.Println(err)
        }
        wg.Done()
    }()

    fmt.Println("listening at " + addr)
    err := srv.ListenAndServe()

    fmt.Println("waiting for the remaining connections to finish...")
    wg.Wait()
    if err != nil && err != http.ErrServerClosed {
        panic(err)
    }
    fmt.Println("gracefully shutdown the http server...")
}
