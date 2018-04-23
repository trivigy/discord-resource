package main

import (
	"fmt"
	"os"
	"bufio"
	"encoding/json"

	"github.com/bwmarrin/discordgo"
)

type Datum struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

type Version struct {
	Ref string `json:"ref,omitempty"`
}

type Output struct {
	Version  Version `json:"version,omitempty"`
	Metadata []Datum `json:"metadata,omitempty"`
}

type Source struct {
	Token    string `json:"token,omitempty"`
	Insecure bool   `json:"insecure,omitempty"`
}

type Params struct {
	Channel string `json:"channel,omitempty"`
	Message string `json:"message,omitempty"`
}

type Payload struct {
	Source Source `json:"source,omitempty"`
	Params Params `json:"params,omitempty"`
}

func main() {
	stat, err := os.Stdin.Stat()
	if err != nil {
		panic(err)
	}

	if stat.Mode()&os.ModeNamedPipe == 0 {
		panic("stdin is empty")
	}

	var payload Payload
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if err := json.Unmarshal(scanner.Bytes(), &payload); err != nil {
			panic(err)
		}
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	discord, err := discordgo.New("Bot " + payload.Source.Token)
	if err != nil {
		panic(err)
	}

	_, err = discord.ChannelMessageSend(
		payload.Params.Channel,
		payload.Params.Message,
	)
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	output, err := json.Marshal(Output{})
	if err != nil {
		panic(err)
	}
	fmt.Print(string(output))
}
