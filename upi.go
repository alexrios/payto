package main

import (
	"errors"
	"math/big"
	"net/url"
)

/**
Name:  upi

   Description:  Unified Payment Interface (UPI).  The path is an
      account alias.  The amount and receiver-name options are mandatory
      for this payment target.  Limitations on the length and character
      set of option values are defined by the implementation of the
      handler.  Language tags and internationalization of options are
      not supported.

   Example:
      payto://upi/alice@example.com?receiver-name=Alice&amount=INR:200
*/

type UPI struct {
	authority    string
	accountAlias string
	amount       big.Float
	receiver     string

	//optional fields
	sender      string
	message     string
	instruction string
}

func NewUPI(accountAlias, receiver string, amount string) (UPI, error) {
	floatAmount, succeeded := new(big.Float).SetString(amount)
	if !succeeded {
		return UPI{}, errors.New("")
	}
	return UPI{
		authority:    UnifiedPaymentInterface,
		accountAlias: accountAlias,
		receiver:     receiver,
		amount:       *floatAmount,
	}, nil
}

func (U UPI) Amount() string {
	return U.amount.String()
}

func (U UPI) ReceiverName() string {
	return U.receiver
}

func (U UPI) SenderName() string {
	return U.sender
}

func (U UPI) Message() string {
	return U.message
}

func (U UPI) Instruction() string {
	return U.instruction
}

func (U UPI) URL() string {
	values := url.Values{}
	values.Set("receiver-name", U.receiver)
	values.Set("amount", U.amount.String())
	if U.sender != "" {
		values.Set("sender-name", U.sender)
	}
	encodeValues := values.Encode()
	u := &url.URL{
		Scheme:   "payto",
		Host:     UnifiedPaymentInterface,
		Path:     U.accountAlias,
		RawQuery: encodeValues,
	}
	return u.String()
}
