package wpp

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/boombuler/barcode/qr"
	_ "github.com/mattn/go-sqlite3" // Importa o driver SQLite3
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/store/sqlstore"
	waLog "go.mau.fi/whatsmeow/util/log"
)

func RunClient(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Encerrando servidor.")
			return
		default:
			dbLog := waLog.Stdout("Database", "DEBUG", true)
			container, err := sqlstore.New(context.Background(), "sqlite3", "file:./database.db?_foreign_keys=on", dbLog)
			if err != nil {
				panic(err)
			}
			deviceStore, err := container.GetFirstDevice(context.Background())
			if err != nil {
				panic(err)
			}
			clientLog := waLog.Stdout("Client", "DEBUG", true)
			client := whatsmeow.NewClient(deviceStore, clientLog)
			ConnClient = client
			// client.AddEventHandler(EventHandler)
			if client.Store.ID == nil {
				// No ID stored, new login
				qrChan, _ := client.GetQRChannel(context.Background())
				err = client.Connect()
				if err != nil {
					panic(err)
				}
				for evt := range qrChan {
					if evt.Event == "code" {
						// Gerar o QR Code a partir da string
						qrCode, err := qr.Encode(evt.Code, qr.L, qr.Auto)
						if err != nil {
							panic(err)
						}
						CodQrCode = qrCode
					}
				}
			} else {
				// Already logged in, just connect
				err = client.Connect()
				if err != nil {
					panic(err)
				}
			}
			// Listen to Ctrl+C
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt, syscall.SIGTERM)
			<-c

			client.Disconnect()
		}
	}
}
