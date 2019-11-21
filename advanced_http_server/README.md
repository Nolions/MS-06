# advanced_http_server

進階Http Server練習

主要餐參考這邊文章[Fun with Concurrency in Golang]("https://www.alexsears.com/2019/10/fun-with-concurrency-in-golang/")，差別在於原文中是使用GO Lang原生的`net/http`函數處理http server，而這邊則是使用了Gin這個Library

1. 透過並發方式同時啟動多個http server
2. gracefully(優雅退出)
3. 當程式接收到系統發出`SIGINT`&`SIGTERM`通知是才會進行關閉
4. 當啟動三個http server有其中一個發生錯誤時，就會將全部啟動http server關閉


## REFERENCE

1. [Fun with Concurrency in Golang]("https://www.alexsears.com/2019/10/fun-with-concurrency-in-golang/")
2. [funwithconcurrency repository of gitlab]("https://gitlab.com/searsaw/funwithconcurrency")