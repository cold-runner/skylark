package sms

type Sms interface {
	SendToSingle(phone string, paramSet []string) error
	SendToMultiple(phoneSet, paramSet []string) error
}
