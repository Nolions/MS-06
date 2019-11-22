# Network

主要參考[Writing TCP scanner in Go]("https://developer20.com/tcp-scanner-in-go/")，透過實作TCP Scanner來練習go land網路相關實作。

`net` 是Go Lang原生所提供的package，支援TCP、UDP或HTTP、FTP等網路相關操作。

[Writing TCP scanner in Go]("https://developer20.com/tcp-scanner-in-go/")一文中透過goroutine實作並發，用以加快scanning時的速度，但goroutine的數量會因為相關環境不同而有數量上限制，故透過chan方式來實作goroutine pool
，用以限制可同時啟動goroutine數目。

主要就是透過當channel滿的時候會阻塞，直到channel有空時阻塞才會取消的概念，下面為透過chan 控制goroutine數目sample code。

```
var ch = make(chan int, 2)
ch <- 1
go func() {
	.
    .
    .
	<-ch
}()
```

## REFERENCE

1. [Writing TCP scanner in Go]("https://developer20.com/tcp-scanner-in-go/")

