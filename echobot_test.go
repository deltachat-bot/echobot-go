package main

import (
	"testing"

	"github.com/deltachat-bot/deltabot-cli-go/botcli"
	"github.com/deltachat/deltachat-rpc-client-go/acfactory"
	"github.com/stretchr/testify/assert"
)

func TestEcho(t *testing.T) {
	var err error

	bot := acfactory.OnlineBot()
	defer acfactory.StopRpc(bot)
	user := acfactory.OnlineAccount()
	defer acfactory.StopRpc(user)

	bot.OnNewMsg(echo)
	go bot.Run() //nolint:errcheck

	chatWithBot, err := acfactory.CreateChat(user, bot.Account)
	assert.Nil(t, err)

	_, err = chatWithBot.SendText("hi")
	assert.Nil(t, err)

	msg, err := acfactory.NextMsg(user)
	assert.Nil(t, err)
	assert.Equal(t, "hi", msg.Text)
}

func TestOnBotInit(t *testing.T) {
	var err error

	bot := acfactory.OnlineBot()
	defer acfactory.StopRpc(bot)
	user := acfactory.OnlineAccount()
	defer acfactory.StopRpc(user)

	onBotInit(nil, bot, nil, nil)
	go bot.Run() //nolint:errcheck

	chatWithBot, err := acfactory.CreateChat(user, bot.Account)
	assert.Nil(t, err)

	_, err = chatWithBot.SendText("hi")
	assert.Nil(t, err)

	msg, err := acfactory.NextMsg(user)
	assert.Nil(t, err)
	assert.Equal(t, "hi", msg.Text)
}

func TestOnBotStart(t *testing.T) {
	bot := acfactory.OnlineBot()
	defer acfactory.StopRpc(bot)

	cli := botcli.New("testbot")
	onBotStart(cli, bot, nil, nil)
}

func TestOnInfoCmd(t *testing.T) {
	bot := acfactory.OnlineBot()
	defer acfactory.StopRpc(bot)

	onInfoCmd(nil, bot, nil, nil)
}
