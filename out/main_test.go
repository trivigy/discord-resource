package main

import (
	"io/ioutil"
	"os"
	"testing"
)

func TestFileOverrides(t *testing.T) {
	setup()
	// _file fields should always win
	sourceParams := Params{
		Channel:     "some-channel",
		ChannelFile: "channel-file",
		Title:       "some-title",
		TitleFile:   "title-file",
		Message:     "some message",
		MessageFile: "message-file",
	}
	expectedParams := Params{
		Channel:     "channel-name",
		ChannelFile: "channel-file",
		Title:       "message title",
		TitleFile:   "title-file",
		Message:     "message body",
		MessageFile: "message-file",
	}

	givenParams, _ := ReasonAboutParams(sourceParams)

	if givenParams != expectedParams {
		t.Fail()
	}

	cleanup()
}

func TestParamsOnly(t *testing.T) {
	// If only non-_file params are set they should
	// remain unchanged
	sourceParams := Params{
		Channel: "channel-name",
		Title:   "message title",
		Message: "message body",
	}
	expectedParams := Params{
		Channel: "channel-name",
		Title:   "message title",
		Message: "message body",
	}

	givenParams, _ := ReasonAboutParams(sourceParams)

	if givenParams != expectedParams {
		t.Fail()
	}
}

func setup() {
	ioutil.WriteFile("channel-file", []byte("channel-name"), 0644)
	ioutil.WriteFile("title-file", []byte("message title"), 0644)
	ioutil.WriteFile("message-file", []byte("message body"), 0644)
}

func cleanup() {
	os.Remove("channel-file")
	os.Remove("title-file")
	os.Remove("message-file")
}
