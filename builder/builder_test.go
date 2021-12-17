package builder

import "testing"

func TestSender_BuildMessage_JSON(t *testing.T) {
	json := &JSONMessageBuilder{}
	sender := &Sender{}
	sender.SetBuilder(json)
	msg, err := sender.BuildMessage("Gopher", "Hello world!")
	if err != nil {
		t.Fatalf("BuidMessage(): no se esperaba error, se obtuvo: %v", err)
	}
	t.Log(string(msg.Body))
}

func TestSender_BuildMessage_XML(t *testing.T) {
	xmlmb := &XMLMessageBuilder{}
	sender := &Sender{}
	sender.SetBuilder(xmlmb)
	msg, err := sender.BuildMessage("Gopher", "Hello world!")
	if err != nil {
		t.Fatalf("BuidMessage(): no se esperaba error, se obtuvo: %v", err)
	}
	t.Log(string(msg.Body))
}
