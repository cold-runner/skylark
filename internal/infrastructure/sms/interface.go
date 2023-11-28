package sms

import "context"

type Sms interface {
	SendToSingle(c context.Context, phone string, paramSet []string) error
	SendToMultiple(c context.Context, phoneSet, paramSet []string) error
}
