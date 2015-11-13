package main

// http://qiita.com/masahikoofjoyto/items/f60188f4252e455541d4#app%E3%82%B5%E3%83%BC%E3%83%90
// http://qiita.com/yuya_takeyama/items/26ccfd37d9ff32bca4c7
// http://qiita.com/najeira/items/47539ab346fa0c00dc62
// http://qiita.com/masahikoofjoyto/items/f60188f4252e455541d4

import (
	"fmt"
	"github.com/guregu/kami"
	"golang.org/x/net/context"
	"net/http"
	"os"
	"os/signal"
	"runtime"
	"syscall"
)

func work(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	name := kami.Param(ctx, "name")
	fmt.Println("**********request has come**************")
	fmt.Println(name)
}

func main() {
	runtime.GOMAXPROCS(30)

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt)
	signal.Notify(sig, syscall.SIGTERM)

	ctx := context.Background()
	kami.Context = ctx

	kami.Post("/work", work)

	go func() {
		kami.Serve()
	}()

	<-sig

}
