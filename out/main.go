package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

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
	Token string `json:"token,omitempty"`
}

type Params struct {
	Channel     string `json:"channel,omitempty"`
	ChannelFile string `json:"channel_file,omitempty"`
	Title       string `json:"title,omitempty"`
	TitleFile   string `json:"title_file,omitempty"`
	Message     string `json:"message,omitempty"`
	MessageFile string `json:"message_file,omitempty"`
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

	inputs := os.Args[1]
	params, err := ReasonAboutParams(inputs, payload.Params)
	if err != nil {
		panic(err)
	}

	discord, err := discordgo.New("Bot " + payload.Source.Token)
	if err != nil {
		panic(err)
	}

	if err = discord.Open(); err != nil {
		panic(err)
	}
	defer discord.Close()

	embed := &discordgo.MessageEmbed{
		Author:      &discordgo.MessageEmbedAuthor{},
		Title:       params.Title,
		Description: params.Message,
	}

	_, err = discord.ChannelMessageSendEmbed(params.Channel, embed)
	if err != nil {
		panic(err)
	}

	output, err := json.Marshal(Output{})
	if err != nil {
		panic(err)
	}
	fmt.Print(string(output))
}

func ReasonAboutParams(inputs string, params Params) (Params, error) {
	if params.ChannelFile != "" {
		file := filepath.Join(inputs, params.ChannelFile)
		contents, err := ioutil.ReadFile(file)
		if err != nil {
			return params, err
		}
		params.Channel = string(contents)
	}

	if params.TitleFile != "" {
		file := filepath.Join(inputs, params.TitleFile)
		contents, err := ioutil.ReadFile(file)
		if err != nil {
			return params, err
		}
		params.Title = string(contents)
	}

	if params.MessageFile != "" {
		file := filepath.Join(inputs, params.MessageFile)
		contents, err := ioutil.ReadFile(file)
		if err != nil {
			return params, err
		}
		params.Message = string(contents)
	}

	return params, nil
}
