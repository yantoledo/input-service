package message

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageWithoutText(t *testing.T) {
	message := NewMessage()
	message.Text = ""
	message.Type = "Text"

	err := message.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "Invalid Message", err.Error())
}

func TestMessageWithoutType(t *testing.T) {
	message := NewMessage()
	message.Text = "Texting"
	message.Type = ""

	err := message.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "Invalid Message", err.Error())
}
