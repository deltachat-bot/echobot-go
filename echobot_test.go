package main

import (
	"testing"

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
