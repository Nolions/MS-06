# Go parameters
GOCMD=go

GORUN=$(GOCMD) run

# 執行 json範例
run-json:
	$(GORUN) json/main.go

# 執行 protobuf範例
run-protobuf:
	$(GORUN) protobuf/main.go

# 執行 互斥鎖範例
run-Mutex:
	$(GORUN) mutex/main.go

# 執行 讀寫鎖範例
run-RWMutex:
	$(GORUN) rw_mutex/main.go


# 執行 Simple Http Server 範例
run-simpleHttpServer:
	$(GORUN) http_server/main.go

# 執行 DB 範例
ID=1
NAME=User Name
DEPART=Depart Name

# QUERY
run-DB_QUERY:
	$(GORUN) db/main.go -c=q -id=${ID}

# INSERT
run-DB_INSERT:
	$(GORUN) db/main.go -c=i -u=${NAME} -d=${DEPART}

# UPDATE
run-DB_UPDATE:
	$(GORUN) db/main.go -c=u -id=${ID} -u=${NAME} -d=${DEPART}

# DELETE
run-DB_DELETE:
	$(GORUN) db/main.go -c=d -id=${ID}

# 執行 redis讀/寫 範例
UID=911b62af-e283-45d2-90f7-c01c809a7029
EVENT=test
DEADLINE=2019-12-31

run-Redis_Write:
	$(GORUN) redis/main.go -c=0 -e=${EVENT} -t=${DEADLINE}

run-Redis_Read:
	$(GORUN) redis/main.go -c=1 -uid=${UID}

run-Redis_Remove:
	$(GORUN) redis/main.go -c=2 -uid=${UID}

run filepath_walk:
	$(GORUN) filepath_walk/main.go -path=/Users/nolions/workspace/MS-06/scanDir_File
