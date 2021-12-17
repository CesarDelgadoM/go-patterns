package builder

import "encoding/json"

//JSONMessageBuilder is concrete builder
type JSONMessageBuilder struct {
	Message Message
}

//SetRecipient asigna el receptor del mensaje
func (b *JSONMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	b.Message.Recipient = recipient
	return b
}

func (b *JSONMessageBuilder) SetMessage(message string) MessageBuilder {
	b.Message.Text = message
	return b
}

func (b *JSONMessageBuilder) Build() (*MessageRepresented, error) {
	data, err := json.Marshal(b.Message)
	if err != nil {
		return nil, err
	}
	return &MessageRepresented{
		Body:   data,
		Format: "JSON",
	}, nil
}
