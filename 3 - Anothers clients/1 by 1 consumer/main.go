package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/dixonwille/wmenu"
	"github.com/segmentio/kafka-go"
	"github.com/sethvargo/go-password/password"
)

type PatentUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Patent   string `json:"patent"`
	Password string `json:"password"`
}

func readRegister() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "CG1",
		Topic:    "register-event",
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	m, err := r.FetchMessage(ctx)
	if err != nil {
		return
	}

	if err := r.CommitMessages(ctx, m); err != nil {
		log.Fatal("failed to commit messages:", err)
	}

	User := PatentUser{}
	json.Unmarshal(m.Value, &User)

	pass, err := password.Generate(15, 5, 5, false, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(User.Email)
	fmt.Println(User.Name)
	fmt.Println(pass)
}

func main() {
	actFunc := func(opts []wmenu.Opt) error {
		if opts[0].ID == 0 { // If yes, do this
			readRegister()
		}
		if opts[0].ID == 1 { // If no, do this
			fmt.Println("Adios!")
			return nil
		}
		return nil
	}

	for {
		menu := wmenu.NewMenu("Quieres consumir un registro?")
		menu.Action(actFunc)
		menu.IsYesNo(0)
		err := menu.Run()
		if err != nil {
			log.Fatal(err)
		}
	}

}
