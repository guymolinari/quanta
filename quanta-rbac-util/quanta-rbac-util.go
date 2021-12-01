package main

import (
	"github.com/disney/quanta/client"
	"github.com/disney/quanta/core"
	"github.com/disney/quanta/rbac"
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
)

// Variables to identify the build
var (
	Version string
	Build   string
)

// Exit Codes
const (
	Success = 0
)

// Main strct defines command line arguments variables and various global meta-data associated with record loads.
type Main struct {
	UserID string
	Port   int
}

// NewMain allocates a new pointer to Main struct with empty record counter
func NewMain() *Main {
	return &Main{}
}

func main() {

	app := kingpin.New(os.Args[0], "Quanta RBAC Utility").DefaultEnvars()
	app.Version("Version: " + Version + "\nBuild: " + Build)

	userID := app.Arg("user-id", "User ID for SystemAdmin grant.").Required().String()
	port := app.Arg("port", "Port number for service").Default("4000").Int32()
	environment := app.Flag("env", "Environment [DEV, QA, STG, VAL, PROD]").Default("DEV").String()
	core.InitLogging("WARN", *environment, "RBAC Utility", Version, "Quanta")

	kingpin.MustParse(app.Parse(os.Args[1:]))

	main := NewMain()
	main.UserID = *userID
	main.Port = int(*port)

	log.Printf("User ID %v.\n", main.UserID)
	log.Printf("Service port %d.\n", main.Port)

	conn := quanta.NewDefaultConnection()
	conn.ServicePort = main.Port
	conn.Quorum = 3
	if err := conn.Connect(nil); err != nil {
		log.Fatal(err)
	}
	store := quanta.NewKVStore(conn)
	ctx, err2 := rbac.NewAuthContext(store, main.UserID, true)
	if err2 != nil {
		log.Fatal(err2)
	}
	err3 := ctx.GrantRole(rbac.SystemAdmin, main.UserID, "", true)
	if err3 != nil {
		log.Fatal(err3)
	}
	log.Printf("Success!\n")
}