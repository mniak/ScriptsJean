package main

import (
	"os"
	"os/signal"
	"syscall"

	_ "github.com/mattn/go-sqlite3"
	"github.com/samber/lo"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	walog "go.mau.fi/whatsmeow/util/log"
)

func main() {
	dbLog := walog.Stdout("Database", "DEBUG", true)
	clientLog := walog.Stdout("Client", "DEBUG", true)

	app := Application{}
	app.config = lo.Must(LoadConfig())

	container := lo.Must(sqlstore.New("sqlite3", "file:wadown.db?_foreign_keys=on", dbLog))
	deviceStore := lo.Must(container.GetFirstDevice())

	app.client = whatsmeow.NewClient(deviceStore, clientLog)
	app.client.AddEventHandler(app.HandleEvent)

	lo.Must0(app.Connect())
	defer app.client.Disconnect()

	// Listen to Ctrl+C (you can also do something else that prevents the program from exiting)
	sigterm := make(chan os.Signal)
	signal.Notify(sigterm, os.Interrupt, syscall.SIGTERM)
	<-sigterm
}
