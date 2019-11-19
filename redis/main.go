package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gofrs/uuid"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
	"time"
)

type event struct {
	UUID     string
	Event    string
	DeleteAt time.Time
	CreateAt time.Time
	DeadLine time.Time
}

var (
	client                  *redis.Client
	cmd                     int
	uid, incident, deadLine string
)

const (
	CMD_CREATOR int = 0
	CMD_GET     int = 1
	CMD_REMOVE  int = 2
)

func main() {
	flag.IntVar(&cmd, "c", 0, "model")
	flag.StringVar(&uid, "uid", "", "UUID")
	flag.StringVar(&incident, "e", "", "event")
	flag.StringVar(&deadLine, "t", "", "deadline")
	flag.Parse()

	client = connect()
	defer client.Close()

	u := event{}

	switch cmd {
	case CMD_CREATOR:
		layout := "2006-01-02"
		t, err := time.Parse(layout, deadLine)
		if err != nil {
			fmt.Println(err)
		}

		u = event{
			UUID:     uuid.Must(uuid.NewV4()).String(),
			Event:    incident,
			DeleteAt: time.Time{},
			CreateAt: time.Now(),
			DeadLine: t,
		}
		u.creator()
		break
	case CMD_GET:
		u.UUID = uid
		u.get()
		fmt.Println(u)
		break
	case CMD_REMOVE:
		u.UUID = uid
		u.remove()
		break
	}
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error:", err)
	}
}

func connect() *redis.Client {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		log.Fatal("error:", err)
	}

	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_ADDRESS"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       db,                          // use default DB
	})
}

func (e *event) creator() {
	b, err := json.Marshal(&e)
	if err != nil {
		log.Fatalln("error:", err)
	}

	client.Set(e.UUID, string(b), 0)
}

func (e *event) get() {
	s, err := client.Get(e.UUID).Result()
	if err != nil {
		log.Fatalln("error:", err)
	}

	err = json.Unmarshal([]byte(s), e)
	if err != nil {
		log.Fatalln("error:", err)
	}
}

func (e *event) remove() {
	client.Set(e.UUID, "", 1*time.Second)
}
