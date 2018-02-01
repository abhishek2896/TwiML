package TwiML

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleXML(t *testing.T) {
	assert.Equal(t, "<Response><Message>TwiML®</Message></Response>", ResponseElement{
		Contents: []interface{}{
			new(MessageElement).SetContents("TwiML®"),
		},
	}.String())
}

func TestComplexXML(t *testing.T) {
	assert.Equal(t, "<Response><Message to=\"123\" from=\"321\">TwiML®</Message><Record action=\"action\" method=\"POST\" timeout=\"10\" finishOnKey=\"key\" maxLength=\"100\" playBeep=\"true\" recordingStatusCallback=\"url\" recordingStatusCallbackMethod=\"POST\"></Record><Dial action=\"action\" method=\"POST\" timeout=\"10\" hangupOnStar=\"true\" timeLimit=\"10\" recordingStatusCallback=\"url\" recordingStatusCallbackMethod=\"POST\"><Number sendDigits=\"123\">123</Number></Dial><Pause length=\"10\"></Pause><Say voice=\"MAN\" language=\"en\" loop=\"10\">text</Say><Gather action=\"action\" timeout=\"10\" finishOnKey=\"#\" numDigits=\"10\" speechTimeout=\"10\"></Gather><Conference muted=\"true\" startConferenceOnEnter=\"true\" endConferenceOnExit=\"true\" maxParticipants=\"10\" record=\"record-from-start\" statusCallback=\"url\" statusCallbackMethod=\"method\">name</Conference><Redirect method=\"POST\">url</Redirect><Play loop=\"10\">url</Play><Hangup></Hangup></Response>", ResponseElement{
		Contents: []interface{}{
			new(MessageElement).SetTo("123").SetFrom("321").SetContents("TwiML®"),
			new(RecordElement).SetRecordingStatusCallbackMethod("POST").SetRecordingStatusCallback("url").SetAction("action").SetFinishOnKey("key").SetPlayBeep(true).SetTimeout(10).SetMaxLength(100).SetMethod("POST"),
			new(DialElement).SetContents([]interface{}{
				new(NumberElement).SetContents("123").SetSendDigits("123"),
			}).SetTimeLimit(10).SetMethod("POST").SetHangupOnStar(true).SetRecordingStatusCallbackMethod("POST").SetRecordingStatusCallback("url").SetAction("action").SetTimeout(10),
			new(PauseElement).SetLength(10),
			new(SayElement).SetContents("text").SetLanguage("en").SetLoop(10).SetVoice("MAN"),
			new(GatherElement).SetSpeechTimeout("10").SetFinishOnKey("#").SetAction("action").SetNumDigits(10).SetTimeout(10).SetContents([]interface{}{}),
			new(ConferenceElement).SetStatusCallback("url").SetStatusCallbackMethod("method").SetContents("name").SetEndConferenceOnExit(true).SetMaxParticipants(10).SetMuted(true).SetRecord("record-from-start").SetStartConferenceOnEnter(true),
			new(RedirectElement).SetMethod("POST").SetContents("url"),
			new(PlayElement).SetContents("url").SetLoop(10),
			new(HangupElement),
		},
	}.String())
}
