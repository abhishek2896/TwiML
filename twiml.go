package TwiML

import (
	"encoding/xml"
	"log"
)

type ResponseElement struct {
	XMLName  xml.Name      `xml:"Response"`
	Contents []interface{} `xml:",xmlvalue"`
}

func (element ResponseElement) String() string {
	bytes, err := xml.Marshal(element)
	if err != nil {
		log.Println(err.Error())
	}
	return string(bytes)
}

type GatherElement struct {
	Contents []interface{} `xml:",innerxml"`

	Input *string `xml:"input,attr"`

	Action *string `xml:"action,attr"`

	Method *string `xml:"method,attr"`

	Timeout *int `xml:"timeout,attr"`

	FinishOnKey *string `xml:"finishOnKey,attr"`

	NumDigits *int `xml:"numDigits,attr"`

	PartialResultCallback *string `xml:"partialResultCallback,attr"`

	PartialResultCallbackMethod *string `xml:"partialResultCallbackMethod,attr"`

	Language *string `xml:"language,attr"`

	Hints *string `xml:"hints,attr"`

	ProfanityFilter *bool `xml:"profanityFilter,attr"`

	SpeechTimeout *string `xml:"speechTimeout,attr"`

	XMLName xml.Name `xml:"Gather"`
}

func (e GatherElement) SetInput(value string) GatherElement {
	e.Input = &value
	return e
}

func (e GatherElement) SetAction(value string) GatherElement {
	e.Action = &value
	return e
}

func (e GatherElement) SetMethod(value string) GatherElement {
	e.Method = &value
	return e
}

func (e GatherElement) SetTimeout(value int) GatherElement {
	e.Timeout = &value
	return e
}

func (e GatherElement) SetFinishOnKey(value string) GatherElement {
	e.FinishOnKey = &value
	return e
}

func (e GatherElement) SetNumDigits(value int) GatherElement {
	e.NumDigits = &value
	return e
}

func (e GatherElement) SetPartialResultCallback(value string) GatherElement {
	e.PartialResultCallback = &value
	return e
}

func (e GatherElement) SetPartialResultCallbackMethod(value string) GatherElement {
	e.PartialResultCallbackMethod = &value
	return e
}

func (e GatherElement) SetLanguage(value string) GatherElement {
	e.Language = &value
	return e
}

func (e GatherElement) SetHints(value string) GatherElement {
	e.Hints = &value
	return e
}

func (e GatherElement) SetProfanityFilter(value bool) GatherElement {
	e.ProfanityFilter = &value
	return e
}

func (e GatherElement) SetSpeechTimeout(value string) GatherElement {
	e.SpeechTimeout = &value
	return e
}

func (e GatherElement) SetContents(value []interface{}) GatherElement {
	e.Contents = value
	return e
}

type SayElement struct {
	Contents string `xml:",innerxml"`

	Voice *string `xml:"voice,attr"`

	Language *string `xml:"language,attr"`

	Loop *int `xml:"loop,attr"`

	XMLName xml.Name `xml:"Say"`
}

func (e SayElement) SetVoice(value string) SayElement {
	e.Voice = &value
	return e
}

func (e SayElement) SetLanguage(value string) SayElement {
	e.Language = &value
	return e
}

func (e SayElement) SetLoop(value int) SayElement {
	e.Loop = &value
	return e
}

func (e SayElement) SetContents(value string) SayElement {
	e.Contents = value
	return e
}

type PauseElement struct {
	Length *int `xml:"length,attr"`

	XMLName xml.Name `xml:"Pause"`
}

func (e PauseElement) SetLength(value int) PauseElement {
	e.Length = &value
	return e
}

type PlayElement struct {
	Contents string `xml:",innerxml"`

	Loop *int `xml:"loop,attr"`

	Digits *string `xml:"digits,attr"`

	XMLName xml.Name `xml:"Play"`
}

func (e PlayElement) SetLoop(value int) PlayElement {
	e.Loop = &value
	return e
}

func (e PlayElement) SetDigits(value string) PlayElement {
	e.Digits = &value
	return e
}

func (e PlayElement) SetContents(value string) PlayElement {
	e.Contents = value
	return e
}

type RecordElement struct {
	Action *string `xml:"action,attr"`

	Method *string `xml:"method,attr"`

	Timeout *int `xml:"timeout,attr"`

	FinishOnKey *string `xml:"finishOnKey,attr"`

	MaxLength *int `xml:"maxLength,attr"`

	PlayBeep *bool `xml:"playBeep,attr"`

	Trim *string `xml:"trim,attr"`

	RecordingStatusCallback *string `xml:"recordingStatusCallback,attr"`

	RecordingStatusCallbackMethod *string `xml:"recordingStatusCallbackMethod,attr"`

	Transcribe *bool `xml:"transcribe,attr"`

	TranscribeCallback *string `xml:"transcribeCallback,attr"`

	XMLName xml.Name `xml:"Record"`
}

func (e RecordElement) SetAction(value string) RecordElement {
	e.Action = &value
	return e
}

func (e RecordElement) SetMethod(value string) RecordElement {
	e.Method = &value
	return e
}

func (e RecordElement) SetTimeout(value int) RecordElement {
	e.Timeout = &value
	return e
}

func (e RecordElement) SetFinishOnKey(value string) RecordElement {
	e.FinishOnKey = &value
	return e
}

func (e RecordElement) SetMaxLength(value int) RecordElement {
	e.MaxLength = &value
	return e
}

func (e RecordElement) SetPlayBeep(value bool) RecordElement {
	e.PlayBeep = &value
	return e
}

func (e RecordElement) SetTrim(value string) RecordElement {
	e.Trim = &value
	return e
}

func (e RecordElement) SetRecordingStatusCallback(value string) RecordElement {
	e.RecordingStatusCallback = &value
	return e
}

func (e RecordElement) SetRecordingStatusCallbackMethod(value string) RecordElement {
	e.RecordingStatusCallbackMethod = &value
	return e
}

func (e RecordElement) SetTranscribe(value bool) RecordElement {
	e.Transcribe = &value
	return e
}

func (e RecordElement) SetTranscribeCallback(value string) RecordElement {
	e.TranscribeCallback = &value
	return e
}

type SmsElement struct {
	To *string `xml:"to,attr"`

	From *string `xml:"from,attr"`

	Action *string `xml:"action,attr"`

	Method *string `xml:"method,attr"`

	StatusCallback *string `xml:"statusCallback,attr"`

	XMLName xml.Name `xml:"Sms"`
}

func (e SmsElement) SetTo(value string) SmsElement {
	e.To = &value
	return e
}

func (e SmsElement) SetFrom(value string) SmsElement {
	e.From = &value
	return e
}

func (e SmsElement) SetAction(value string) SmsElement {
	e.Action = &value
	return e
}

func (e SmsElement) SetMethod(value string) SmsElement {
	e.Method = &value
	return e
}

func (e SmsElement) SetStatusCallback(value string) SmsElement {
	e.StatusCallback = &value
	return e
}

type HangupElement struct {
	XMLName xml.Name `xml:"Hangup"`
}

type LeaveElement struct {
	XMLName xml.Name `xml:"Leave"`
}

type RejectElement struct {
	Reason  *string  `xml:"reason,attr"`
	XMLName xml.Name `xml:"Reject"`
}

func (e RejectElement) SetReason(value string) RejectElement {
	e.Reason = &value
	return e
}

type RedirectElement struct {
	Contents string   `xml:",innerxml"`
	Method   *string  `xml:"method,attr"`
	XMLName  xml.Name `xml:"Redirect"`
}

func (e RedirectElement) SetMethod(value string) RedirectElement {
	e.Method = &value
	return e
}

func (e RedirectElement) SetContents(value string) RedirectElement {
	e.Contents = value
	return e
}

type EnqueueElement struct {
	Contents      string   `xml:",innerxml"`
	Action        *string  `xml:"action,attr"`
	WaitUrl       *string  `xml:"waitUrl,attr"`
	WaitUrlMethod *string  `xml:"waitUrlMethod,attr"`
	Method        *string  `xml:"method,attr"`
	WorkflowSid   *string  `xml:"workflowSid,attr"`
	XMLName       xml.Name `xml:"Enqueue"`
}

func (e EnqueueElement) SetContents(value string) EnqueueElement {
	e.Contents = value
	return e
}

func (e EnqueueElement) SetAction(value string) EnqueueElement {
	e.Action = &value
	return e
}

func (e EnqueueElement) SetWaitUrl(value string) EnqueueElement {
	e.WaitUrl = &value
	return e
}

func (e EnqueueElement) SetWaitUrlMethod(value string) EnqueueElement {
	e.WaitUrlMethod = &value
	return e
}

func (e EnqueueElement) SetMethod(value string) EnqueueElement {
	e.Method = &value
	return e
}

func (e EnqueueElement) SetWorkflowSid(value string) EnqueueElement {
	e.WorkflowSid = &value
	return e
}

type DialElement struct {
	Contents                      []interface{} `xml:",innerxml"`
	Action                        *string       `xml:"action,attr"`
	Method                        *string       `xml:"method,attr"`
	Timeout                       *int          `xml:"timeout,attr"`
	HangupOnStar                  *bool         `xml:"hangupOnStar,attr"`
	TimeLimit                     *int          `xml:"timeLimit,attr"`
	CallerId                      *string       `xml:"callerId,attr"`
	Record                        *string       `xml:"record,attr"`
	Trim                          *string       `xml:"trim,attr"`
	RecordingStatusCallback       *string       `xml:"recordingStatusCallback,attr"`
	RecordingStatusCallbackMethod *string       `xml:"recordingStatusCallbackMethod,attr"`
	RecordingStatusCallbackEvent  *string       `xml:"recordingStatusCallbackEvent,attr"`
	AnswerOnBridge                *bool         `xml:"answerOnBridge,attr"`
	RingTone                      *string       `xml:"ringTone,attr"`
	XMLName                       xml.Name      `xml:"Dial"`
}

func (e DialElement) SetContents(value []interface{}) DialElement {
	e.Contents = value
	return e
}

func (e DialElement) SetAction(value string) DialElement {
	e.Action = &value
	return e
}

func (e DialElement) SetMethod(value string) DialElement {
	e.Method = &value
	return e
}

func (e DialElement) SetTimeout(value int) DialElement {
	e.Timeout = &value
	return e
}

func (e DialElement) SetHangupOnStar(value bool) DialElement {
	e.HangupOnStar = &value
	return e
}

func (e DialElement) SetTimeLimit(value int) DialElement {
	e.TimeLimit = &value
	return e
}

func (e DialElement) SetCallerId(value string) DialElement {
	e.CallerId = &value
	return e
}

func (e DialElement) SetRecord(value string) DialElement {
	e.Record = &value
	return e
}

func (e DialElement) SetTrim(value string) DialElement {
	e.Trim = &value
	return e
}

func (e DialElement) SetRecordingStatusCallback(value string) DialElement {
	e.RecordingStatusCallback = &value
	return e
}

func (e DialElement) SetRecordingStatusCallbackMethod(value string) DialElement {
	e.RecordingStatusCallbackMethod = &value
	return e
}

func (e DialElement) SetRecordingStatusCallbackEvent(value string) DialElement {
	e.RecordingStatusCallbackEvent = &value
	return e
}

func (e DialElement) SetAnswerOnBridge(value bool) DialElement {
	e.AnswerOnBridge = &value
	return e
}

func (e DialElement) SetRingTone(value string) DialElement {
	e.RingTone = &value
	return e
}

type ClientElement struct {
	Contents             string   `xml:",innerxml"`
	Url                  *string  `xml:"url,attr"`
	Method               *string  `xml:"method,attr"`
	StatusCallback       *string  `xml:"statusCallback,attr"`
	StatusCallbackEvent  *string  `xml:"statusCallbackEvent,attr"`
	StatusCallbackMethod *string  `xml:"statusCallbackMethod,attr"`
	XMLName              xml.Name `xml:"Client"`
}

func (e ClientElement) SetContents(value string) ClientElement {
	e.Contents = value
	return e
}

func (e ClientElement) SetUrl(value string) ClientElement {
	e.Url = &value
	return e
}

func (e ClientElement) SetMethod(value string) ClientElement {
	e.Method = &value
	return e
}

func (e ClientElement) SetStatusCallback(value string) ClientElement {
	e.StatusCallback = &value
	return e
}

func (e ClientElement) SetStatusCallbackMethod(value string) ClientElement {
	e.StatusCallbackMethod = &value
	return e
}

func (e ClientElement) SetStatusCallbackEvent(value string) ClientElement {
	e.StatusCallbackEvent = &value
	return e
}

type ConferenceElement struct {
	Contents                      string   `xml:",innerxml"`
	Muted                         *bool    `xml:"muted,attr"`
	Beep                          *bool    `xml:"beep,attr"`
	StartConferenceOnEnter        *bool    `xml:"startConferenceOnEnter,attr"`
	EndConferenceOnExit           *bool    `xml:"endConferenceOnExit,attr"`
	WaitUrl                       *string  `xml:"waitUrl,attr"`
	WaitMethod                    *string  `xml:"waitMethod,attr"`
	MaxParticipants               *int     `xml:"maxParticipants,attr"`
	Record                        *string  `xml:"record,attr"`
	Region                        *string  `xml:"region,attr"`
	Trim                          *string  `xml:"trim,attr"`
	Coach                         *string  `xml:"coach,attr"`
	StatusCallbackEvent           *string  `xml:"statusCallbackEvent,attr"`
	StatusCallback                *string  `xml:"statusCallback,attr"`
	StatusCallbackMethod          *string  `xml:"statusCallbackMethod,attr"`
	RecordingStatusCallback       *string  `xml:"recordingStatusCallback,attr"`
	RecordingStatusCallbackMethod *string  `xml:"recordingStatusCallbackMethod,attr"`
	RecordingStatusCallbackEvent  *string  `xml:"recordingStatusCallbackEvent,attr"`
	EventCallbackUrl              *string  `xml:"eventCallbackUrl,attr"`
	XMLName                       xml.Name `xml:"Conference"`
}

func (e ConferenceElement) SetContents(value string) ConferenceElement {
	e.Contents = value
	return e
}

func (e ConferenceElement) SetMuted(value bool) ConferenceElement {
	e.Muted = &value
	return e
}

func (e ConferenceElement) SetBeep(value bool) ConferenceElement {
	e.Beep = &value
	return e
}

func (e ConferenceElement) SetStartConferenceOnEnter(value bool) ConferenceElement {
	e.StartConferenceOnEnter = &value
	return e
}

func (e ConferenceElement) SetEndConferenceOnExit(value bool) ConferenceElement {
	e.EndConferenceOnExit = &value
	return e
}

func (e ConferenceElement) SetWaitUrl(value string) ConferenceElement {
	e.WaitUrl = &value
	return e
}

func (e ConferenceElement) SetWaitMethod(value string) ConferenceElement {
	e.WaitMethod = &value
	return e
}

func (e ConferenceElement) SetMaxParticipants(value int) ConferenceElement {
	e.MaxParticipants = &value
	return e
}
func (e ConferenceElement) SetRecord(value string) ConferenceElement {
	e.Record = &value
	return e
}

func (e ConferenceElement) SetRegion(value string) ConferenceElement {
	e.Region = &value
	return e
}

func (e ConferenceElement) SetTrim(value string) ConferenceElement {
	e.Trim = &value
	return e
}

func (e ConferenceElement) SetCoach(value string) ConferenceElement {
	e.Coach = &value
	return e
}

func (e ConferenceElement) SetStatusCallback(value string) ConferenceElement {
	e.StatusCallback = &value
	return e
}

func (e ConferenceElement) SetStatusCallbackEvent(value string) ConferenceElement {
	e.StatusCallbackEvent = &value
	return e
}

func (e ConferenceElement) SetStatusCallbackMethod(value string) ConferenceElement {
	e.StatusCallbackMethod = &value
	return e
}

func (e ConferenceElement) SetRecordingStatusCallback(value string) ConferenceElement {
	e.RecordingStatusCallback = &value
	return e
}

func (e ConferenceElement) SetRecordingStatusCallbackMethod(value string) ConferenceElement {
	e.RecordingStatusCallbackMethod = &value
	return e
}

func (e ConferenceElement) SetRecordingStatusCallbackEvent(value string) ConferenceElement {
	e.RecordingStatusCallbackEvent = &value
	return e
}

func (e ConferenceElement) SetEventCallbackUrl(value string) ConferenceElement {
	e.EventCallbackUrl = &value
	return e
}

type NumberElement struct {
	Contents             string   `xml:",innerxml"`
	SendDigits           *string  `xml:"sendDigits,attr"`
	Url                  *string  `xml:"url,attr"`
	Method               *string  `xml:"method,attr"`
	StatusCallback       *string  `xml:"statusCallback,attr"`
	StatusCallbackEvent  *string  `xml:"statusCallbackEvent,attr"`
	StatusCallbackMethod *string  `xml:"statusCallbackMethod,attr"`
	XMLName              xml.Name `xml:"Number"`
}

func (e NumberElement) SetContents(value string) NumberElement {
	e.Contents = value
	return e
}

func (e NumberElement) SetSendDigits(value string) NumberElement {
	e.SendDigits = &value
	return e
}

func (e NumberElement) SetUrl(value string) NumberElement {
	e.Url = &value
	return e
}

func (e NumberElement) SetMethod(value string) NumberElement {
	e.Method = &value
	return e
}

func (e NumberElement) SetStatusCallback(value string) NumberElement {
	e.StatusCallback = &value
	return e
}

func (e NumberElement) SetStatusCallbackMethod(value string) NumberElement {
	e.StatusCallbackMethod = &value
	return e
}

func (e NumberElement) SetStatusCallbackEvent(value string) NumberElement {
	e.StatusCallbackEvent = &value
	return e
}

type QueueElement struct {
	Contents            string   `xml:",innerxml"`
	Url                 *string  `xml:"url,attr"`
	Method              *string  `xml:"method,attr"`
	ReservationSid      *string  `xml:"reservationSid,attr"`
	PostWorkActivitySid *string  `xml:"postWorkActivitySid,attr"`
	XMLName             xml.Name `xml:"Queue"`
}

func (e QueueElement) SetContents(value string) QueueElement {
	e.Contents = value
	return e
}

func (e QueueElement) SetUrl(value string) QueueElement {
	e.Url = &value
	return e
}

func (e QueueElement) SetMethod(value string) QueueElement {
	e.Method = &value
	return e
}

func (e QueueElement) SetReservationSid(value string) QueueElement {
	e.ReservationSid = &value
	return e
}

func (e QueueElement) SetPostWorkActivitySid(value string) QueueElement {
	e.PostWorkActivitySid = &value
	return e
}

type SimElement struct {
	Contents string   `xml:",innerxml"`
	XMLName  xml.Name `xml:"Sim"`
}

func (e SimElement) SetContents(value string) SimElement {
	e.Contents = value
	return e
}

type SipElement struct {
	Contents             string   `xml:",innerxml"`
	Username             *string  `xml:"username,attr"`
	Password             *string  `xml:"password,attr"`
	Url                  *string  `xml:"url,attr"`
	Method               *string  `xml:"method,attr"`
	StatusCallback       *string  `xml:"statusCallback,attr"`
	StatusCallbackEvent  *string  `xml:"statusCallbackEvent,attr"`
	StatusCallbackMethod *string  `xml:"statusCallbackMethod,attr"`
	XMLName              xml.Name `xml:"Sip"`
}

func (e SipElement) SetContents(value string) SipElement {
	e.Contents = value
	return e
}

func (e SipElement) SetUrl(value string) SipElement {
	e.Url = &value
	return e
}

func (e SipElement) SetUsername(value string) SipElement {
	e.Username = &value
	return e
}

func (e SipElement) SetPassword(value string) SipElement {
	e.Password = &value
	return e
}

func (e SipElement) SetMethod(value string) SipElement {
	e.Method = &value
	return e
}

func (e SipElement) SetStatusCallback(value string) SipElement {
	e.StatusCallback = &value
	return e
}

func (e SipElement) SetStatusCallbackMethod(value string) SipElement {
	e.StatusCallbackMethod = &value
	return e
}

func (e SipElement) SetStatusCallbackEvent(value string) SipElement {
	e.StatusCallbackEvent = &value
	return e
}

type MessageElement struct {
	Contents string `xml:",innerxml"`

	To *string `xml:"to,attr"`

	From *string `xml:"from,attr"`

	Action *string `xml:"action,attr"`

	Method *string `xml:"method,attr"`

	XMLName xml.Name `xml:"Message"`
}

func (e MessageElement) SetContents(value string) MessageElement {
	e.Contents = value
	return e
}

func (e MessageElement) SetTo(value string) MessageElement {
	e.To = &value
	return e
}

func (e MessageElement) SetFrom(value string) MessageElement {
	e.From = &value
	return e
}

func (e MessageElement) SetAction(value string) MessageElement {
	e.Action = &value
	return e
}

func (e MessageElement) SetMethod(value string) MessageElement {
	e.Method = &value
	return e
}

type RecieveElement struct {
	Action *string `xml:"action,attr"`

	Method *string `xml:"method,attr"`

	MediaType *string `xml:"mediaType,attr"`

	StoreMedia *string `xml:"storeMedia,attr"`

	XMLName xml.Name `xml:"Recieve"`
}

func (e RecieveElement) SetAction(value string) RecieveElement {
	e.Action = &value
	return e
}

func (e RecieveElement) SetMethod(value string) RecieveElement {
	e.Method = &value
	return e
}

func (e RecieveElement) SetMediaType(value string) RecieveElement {
	e.MediaType = &value
	return e
}

func (e RecieveElement) SetStoreMedia(value string) RecieveElement {
	e.StoreMedia = &value
	return e
}
