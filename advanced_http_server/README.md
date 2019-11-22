# advanced_http_server

進階Http Server練習

主要餐參考這邊文章[Fun with Concurrency in Golang]("https://www.alexsears.com/2019/10/fun-with-concurrency-in-golang/")，差別在於原文中是使用GO Lang原生的`net/http`函數處理http server，而這邊則是使用了Gin這個Library

1. 透過並發方式同時啟動多個http server( `helloWord`, `hellowordName`, `echo`)
2. gracefully(優雅退出)
3. 當程式接收到系統發出`SIGINT`&`SIGTERM`通知是才會進行關閉
4. 當啟動三個http server有其中一個發生錯誤時，就會將全部啟動http server關閉

## Server & Router List

| router | method | Server            |
| ------ | ------ | ----------------- |
| /      | GET    | Hello Word Server |
| /      | GET    | Hello Name Server |
| /      | POST   | echo Server       |

### Example:

***Hello Word Server***

```
http :7000 

HTTP/1.1 303 See Other
Content-Length: 9
Content-Type: text/plain; charset=utf-8
Date: Fri, 22 Nov 2019 03:35:35 GMT

HelloWord
```

***Hello Name Server*** 

```
$ http :7001 name==Noions

HTTP/1.1 200 OK
Content-Length: 12
Content-Type: text/plain; charset=utf-8
Date: Fri, 22 Nov 2019 03:49:19 GMT

Hello Noions
```

***echo Server***

```
$ hhttp POST :7002 name=Alex job="Software Engineer" coffee=please
 
HTTP/1.1 200 OK
Content-Length: 64
Content-Type: text/plain; charset=utf-8
Date: Fri, 22 Nov 2019 03:49:45 GMT

{
    "coffee": "please",
    "job": "Software Engineer",
    "name": "Alex"
}
```

## REFERENCE

1. [Fun with Concurrency in Golang]("https://www.alexsears.com/2019/10/fun-with-concurrency-in-golang/")
2. [funwithconcurrency repository of gitlab]("https://gitlab.com/searsaw/funwithconcurrency")