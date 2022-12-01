package connection

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var Conn *pgx.Conn

func ConnectDB() {

	// variable nampung URI for Db connection bentukny string
	// format "postgres://username:password@localhost:5432/database"
	uri := "postgres://arifth:593001@localhost:5432/courses"

	// var err buat nampung error dr connection variable

	var err error

	// open connection to database,then baru cek error konek atau tidak
	Conn, err = pgx.Connect(context.Background(), uri)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Ga konek bossque %v \n", err)
		// hentikan app golang
		os.Exit(1)
	} else {
		fmt.Println("sukses konnek database ")

	}

}
