package migration

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db, err := sql.Open("sqlite3", "db/app.db")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`
CREATE TABLE IF NOT EXISTS users (
	id integer not null primary key AUTOINCREMENT,
    username varchar(255) not null,
    password varchar(255) not null,
    role varchar(255) not null,
    loggedin boolean not null,
    created_at Time.time,
    updated_at Time.time
);
CREATE TABLE IF NOT EXISTS campaign (
  id integer not null primary key AUTOINCREMENT,
  title varchar(255) not null,
  targetAmount integer not null,
  currentAmount integer not null,
  description varchar(255) not null,
  timeStart Time.time,
  timeEnd Time.time  
);
CREATE TABLE IF NOT EXISTS transaction (
  	  id integer not null primary key AUTOINCREMENT,
  	  amount integer not null,
  	  status boolean not null
);


`)
	if err != nil {
		panic(err)
	}
}
