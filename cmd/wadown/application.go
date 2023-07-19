package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/mdp/qrterminal"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/types/events"
)

const PERM_DEFAULT = 0o755

type Application struct {
	client *whatsmeow.Client
	config Config
}

func (app *Application) Connect() error {
	store.SetOSInfo("Wadown Image Downloader", [3]uint32{0, 0, 1})
	if app.client.Store.ID != nil { // If already connected
		if err := app.client.Connect(); err != nil {
			return err
		}
		return nil
	}

	qrChan, _ := app.client.GetQRChannel(context.Background())
	if err := app.client.Connect(); err != nil {
		return err
	}
	for evt := range qrChan {
		if evt.Event == "code" {
			qrterminal.GenerateHalfBlock(evt.Code, qrterminal.L, os.Stdout)
		} else {
			fmt.Println("Login event:", evt.Event)
		}
	}
	return nil
}

func (app *Application) HandleEvent(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		switch v.Info.MediaType {
		case "image":
			dirname := filepath.Join("images", app.config.GetUserAlias(v.Info.Sender.User))
			os.MkdirAll(dirname, PERM_DEFAULT)
			filename := filepath.Join(dirname, fmt.Sprintf("%s.jpg", v.Info.ID))

			mediaBytes, err := app.client.Download(v.Message.GetImageMessage())
			if err != nil {
				log.Printf("Failed to download image '%s': %s\n", filename, err)
				return
			}
			err = os.WriteFile(filename, mediaBytes, PERM_DEFAULT)
			if err != nil {
				log.Printf("Failed to download image '%s': %s\n", filename, err)
				return
			}
		}
	}
}
