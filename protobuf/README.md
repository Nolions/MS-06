# protobuf

## 產生go lang用的Proto文件

    指令
    protoc --go_out=. <protobuf file>
    EX:
    protoc --go_out=. ./protobuf/*.proto> //編譯./protobuf下所有副檔名為proto的檔案
