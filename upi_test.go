package main

import (
	"reflect"
	"testing"
)

func TestNewUPI(t *testing.T) {
	tests := []struct {
		name      string
		UPI       UPI
		wantedURL string
	}{
		{
			name:      "should create a valid UPI wth the mandatory options ",
			UPI:       createNewUPI(t, "my-lil-acc@payto.com", "the-receiver", "123.39"),
			wantedURL: "payto://upi/my-lil-acc@payto.com?amount=123.39&receiver-name=the-receiver",
		},
		{
			name:      "should create a valid UPI wth the mandatory options and sender ",
			UPI:       createNewUPI(t, "my-lil-acc@payto.com", "the-receiver", "123.39", UPISender("the-sender")),
			wantedURL: "payto://upi/my-lil-acc@payto.com?amount=123.39&receiver-name=the-receiver&sender-name=the-sender",
		},
		{
			name:      "should create a valid UPI wth the mandatory options and message ",
			UPI:       createNewUPI(t, "my-lil-acc@payto.com", "the-receiver", "123.39", UPIMessage("the-message")),
			wantedURL: "payto://upi/my-lil-acc@payto.com?amount=123.39&message=the-message&receiver-name=the-receiver",
		},
		{
			name:      "should create a valid UPI wth the mandatory options and sender ",
			UPI:       createNewUPI(t, "my-lil-acc@payto.com", "the-receiver", "123.39", UPIInstruction("the-instruction")),
			wantedURL: "payto://upi/my-lil-acc@payto.com?amount=123.39&instruction=the-instruction&receiver-name=the-receiver",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.UPI.URL(), tt.wantedURL) {
				t.Errorf("NewUPI() got = %v, want %v", tt.UPI.URL(), tt.wantedURL)
			}
		})
	}
}

func createNewUPI(t *testing.T, accountAlias, receiver, amount string, options ...UPIOptional) UPI {
	t.Helper()
	upi, err := NewUPI(accountAlias, receiver, amount, options...)
	if err != nil {
		t.Fatal(err)
	}
	return upi
}
