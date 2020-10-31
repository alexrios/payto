package main

import (
	"errors"
	"math/big"
	"net/url"
)

/**
Name:  bic

   Description:  Business Identifier Code (BIC).  The path consists of
      just a BIC.  This is used for wire transfers between banks.  The
      registry for BICs is provided by the Society for Worldwide
      Interbank Financial Telecommunication (SWIFT).  The path does not
      allow specifying a bank account number.  Limitations on the length
      and character set of option values are defined by the
      implementation of the handler.  Language tagging and
      internationalization of options are not supported.

   Example:
      payto://bic/SOGEDEFFXXX
*/

type BIC struct {
	swift     string
	authority string
	amount    big.Float
	receiver  string
	//Optionals
	sender      string
	message     string
	instruction string
}

type BICOptional func(*BIC)

func BICSender(sender string) BICOptional {
	return func(bic *BIC) {
		bic.sender = sender
	}
}

func BICMessage(message string) BICOptional {
	return func(bic *BIC) {
		bic.message = message
	}
}

func BICInstruction(instruction string) BICOptional {
	return func(bic *BIC) {
		bic.instruction = instruction
	}
}

func NewBIC(swiftCode, receiver string, amount string, options ...BICOptional) (BIC, error) {
	floatAmount, succeeded := new(big.Float).SetString(amount)
	if !succeeded {
		return BIC{}, errors.New("")
	}
	bic := BIC{
		authority:   BusinessIdentifierCode,
		swift:       swiftCode,
		receiver:    receiver,
		amount:      *floatAmount,
		sender:      "",
		message:     "",
		instruction: "",
	}

	for _, option := range options {
		option(&bic)
	}

	return bic, nil
}

func (U BIC) Amount() string {
	return U.amount.String()
}

func (U BIC) ReceiverName() string {
	return U.receiver
}

func (U BIC) SenderName() string {
	return U.sender
}

func (U BIC) Message() string {
	return U.message
}

func (U BIC) Instruction() string {
	return U.instruction
}

func (U BIC) URL() string {
	values := url.Values{}
	values.Set("receiver-name", U.receiver)
	values.Set("amount", U.amount.String())
	if U.sender != "" {
		values.Set("sender-name", U.sender)
	}
	if U.instruction != "" {
		values.Set("instruction", U.instruction)
	}
	if U.message != "" {
		values.Set("message", U.message)
	}
	encodeValues := values.Encode()
	u := &url.URL{
		Scheme:   "payto",
		Host:     BusinessIdentifierCode,
		Path:     U.swift,
		RawQuery: encodeValues,
	}
	return u.String()
}
