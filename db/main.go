package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var (
	cmd, userName, depart string
	id                    int
)

type sqLite struct {
	db  *sql.DB
	err error
}

func main() {
	d := new(sqLite)
	d.db, d.err = sql.Open("sqlite3", "./foo.db")
	if d.err != nil {
		log.Fatal(d.err)
	}

	d.creatorTable()

	flag.StringVar(&cmd, "c", "o", "Read cmd")
	flag.IntVar(&id, "id", -1, "User uid")
	flag.StringVar(&userName, "u", "", "User Name")
	flag.StringVar(&depart, "d", "", "Depart Name")

	flag.Parse()

	switch cmd {
	case "q":
		fmt.Println("Query user uid is", id, "......")
		d.query(id)
		break
	case "i":
		fmt.Println("Insert user ......")
		fmt.Println("User Name:", userName)
		fmt.Println("Depart:", depart)
		d.insert(userName, depart)
		break
	case "u":
		fmt.Println("Update user ......")
		fmt.Println("id:", id)
		fmt.Println("User Name:", userName)
		fmt.Println("Depart:", depart)
		d.update(id, userName, depart)
		break
	case "d":
		fmt.Println("Delete user's uid is", id, "......")
		d.delete(id)
		break
	}

}

func (d sqLite) delete(id int) {
	stmt, err := d.db.Prepare("DELETE FROM userinfo WHERE uid=?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(affect)
}

func (d sqLite) update(id int, userName, depart string) {
	stmt, err := d.db.Prepare("UPDATE userinfo SET username=?, departname=? WHERE uid=?")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(userName, depart, id)
	if err != nil {
		log.Fatal(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(affect)
}

func (d sqLite) insert(userName, depart string) int64 {
	stmt, err := d.db.Prepare("INSERT INTO userinfo (username, departname, created) values(?,?,?)")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(userName, depart, time.Now().Format("2006-01-02"))
	if err != nil {
		log.Fatal(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return id
}

func (d sqLite) query(id int) {
	rows, err := d.db.Query("SELECT * FROM userinfo WHERE uid=?", id)
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created time.Time
		err = rows.Scan(&uid, &username, &department, &created)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
}

func (d sqLite) creatorTable() {
	sqlTable := `
    CREATE TABLE IF NOT EXISTS userinfo(
        uid INTEGER PRIMARY KEY AUTOINCREMENT,
        username VARCHAR(64) NULL,
        departname VARCHAR(64) NULL,
        created DATE NULL
	);
	`

	d.db.Exec(sqlTable)
}
