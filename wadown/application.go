package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/mdp/qrterminal"
	"github.com/pkg/errors"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store"
	"go.mau.fi/whatsmeow/types"
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

			imgBytes, err := app.client.Download(v.Message.GetImageMessage())
			if err != nil {
				log.Printf("Failed to download image '%s': %s\n", v.Info.ID, err)
				return
			}
			err = app.SaveImage(v.Info, imgBytes)
			if err != nil {
				log.Printf("Failed to store image '%s': %s\n", v.Info.ID, err)
				return
			}
		}
	}
}

func (app *Application) SaveImage(info types.MessageInfo, imgBytes []byte) error {
	dirname := filepath.Join("images", app.config.GetUserAlias(info.Timestamp.Format(time.DateOnly)))
	err := os.MkdirAll(dirname, PERM_DEFAULT)
	if err != nil {
		return errors.WithMessage(err, fmt.Sprintf("could not create directory '%s'", dirname))
	}
	filename := filepath.Join(dirname, fmt.Sprintf("%s.jpg", info.ID))

	err = os.WriteFile(filename, imgBytes, PERM_DEFAULT)
	if err != nil {
		return err
	}
	return nil
}
