package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/keybase/go-keybase-chat-bot/kbchat"
)

var kbc *kbchat.API

func fail(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(3)
}

func send(channel kbchat.Channel, body string) {
  if _, err := kbc.SendMessage(channel, body); err != nil {
    fail("error echo'ing message: %s", err.Error())
  }
}

func main() {
	var kbLoc string
	var err error

	flag.StringVar(&kbLoc, "keybase", "keybase", "the location of the Keybase app")
	flag.Parse()

	if kbc, err = kbchat.Start(kbchat.RunOptions{KeybaseLocation: kbLoc}); err != nil {
		fail("Error creating API: %s", err.Error())
	}

  sub, err := kbc.ListenForNewTextMessages()
	if err != nil {
		fail("Error listening: %s", err.Error())
	}

	fmt.Println("Tarantula running.")

	for {
		msg, err := sub.Read()
		if err != nil {
			fail("failed to read message: %s", err.Error())
		}

		if msg.Message.Content.Type != "text" { // If the message isnt text
			continue
		}

		args := strings.Split(msg.Message.Content.Text.Body, " ")
    switch body := args[0]; body {
  	case "!ping":
      send(msg.Message.Channel, "Pong")
		case "!lookup":
			send(msg.Message.Channel, Lookup(args[1]))
  	}
	}
}
