// Copyright (c) 2018 Kevin Kragenbrink, II
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package bot

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/kkragenbrink/slate/config"
)

// SlateCommand describes a command to be registered with slate.
type SlateCommand interface {
	// Name describes the string used to route input to this command
	Name() string

	// Synopsis describes what the command will do
	Synopsis() string

	// Usage describes the helpfile and usage for this command
	Usage() string

	// Execute is the handler which will be run when this command is called
	Execute(context.Context, []string, DiscordSession, *discordgo.MessageCreate)
}

// DiscordFactory describes a factory function used to create a discord session.
// 		cfg: The slate configuration object.
type DiscordFactory func(cfg *config.Config) (DiscordSession, error)

// DiscordSession describes the functions required to work with discord
type DiscordSession interface {
	// AddHandler adds a function handler to Discord.
	AddHandler(interface{}) func()

	// ChannelMessageSend allows a message to be sent to a channel on Discord
	// todo: Remove the hard reference to discordgo
	ChannelMessageSend(string, string) (*discordgo.Message, error)

	// Open establishes the websocket connection to Discord.
	Open() error

	// Close closes the websocket connection to Discord.
	Close() error
}
