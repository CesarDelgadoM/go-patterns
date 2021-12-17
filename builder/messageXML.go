package builder

import "encoding/xml"

type XMLMessageBuilder struct {
	message Message
}

func (b *XMLMessageBuilder) SetRecipient(recipient string) MessageBuilder {
	b.message.Recipient = recipient
	return b
}

func (b *XMLMessageBuilder) SetMessage(message string) MessageBuilder {
	b.message.Text = message
	return b
}

func (b *XMLMessageBuilder) Build() (*MessageRepresented, error) {
	data, err := xml.Marshal(b.message)
	if err != nil {
		return nil, err
	}
	return &MessageRepresented{
		Body:   data,
		Format: "XML",
	}, nil
}
