package main

import (
	"reflect"
	"testing"
)

func TestNewBIC(t *testing.T) {
	tests := []struct {
		name      string
		BIC       BIC
		wantedURL string
	}{
		{
			name:      "should create a valid BIC wth the mandatory options ",
			BIC:       createNewBIC(t, "SOGEDEFFXXX", "the-receiver", "123.39"),
			wantedURL: "payto://bic/SOGEDEFFXXX?amount=123.39&receiver-name=the-receiver",
		},
		{
			name:      "should create a valid BIC wth the mandatory options and sender ",
			BIC:       createNewBIC(t, "SOGEDEFFXXX", "the-receiver", "123.39", BICSender("the-sender")),
			wantedURL: "payto://bic/SOGEDEFFXXX?amount=123.39&receiver-name=the-receiver&sender-name=the-sender",
		},
		{
			name:      "should create a valid BIC wth the mandatory options and message ",
			BIC:       createNewBIC(t, "SOGEDEFFXXX", "the-receiver", "123.39", BICMessage("the-message")),
			wantedURL: "payto://bic/SOGEDEFFXXX?amount=123.39&message=the-message&receiver-name=the-receiver",
		},
		{
			name:      "should create a valid BIC wth the mandatory options and sender ",
			BIC:       createNewBIC(t, "SOGEDEFFXXX", "the-receiver", "123.39", BICInstruction("the-instruction")),
			wantedURL: "payto://bic/SOGEDEFFXXX?amount=123.39&instruction=the-instruction&receiver-name=the-receiver",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(tt.BIC.URL(), tt.wantedURL) {
				t.Errorf("NewBIC() got = %v, want %v", tt.BIC.URL(), tt.wantedURL)
			}
		})
	}
}

func createNewBIC(t *testing.T, swift, receiver, amount string, options ...BICOptional) BIC {
	t.Helper()
	bic, err := NewBIC(swift, receiver, amount, options...)
	if err != nil {
		t.Fatal(err)
	}
	return bic
}
