package main

import "net/url"

/**
  Applications MUST accept URIs with options in any order.  The
   "amount" option MUST NOT occur more than once.  Other options MAY be
   allowed multiple times, with further restrictions depending on the
   payment target type.  The following options SHOULD be understood by
   every payment target type.

   amount:  The amount to transfer.  The format MUST be:

        amount = currency ":" unit [ "." fraction ]
        currency = 1*ALPHA
        unit = 1*(DIGIT / ",")
        fraction = 1*(DIGIT / ",")

      If a 3-letter 'currency' is used, it MUST be an [ISO4217]
      alphabetic code.  A payment target type MAY define semantics
      beyond ISO 4217 for currency codes that are not 3 characters.  The
      'unit' value MUST be smaller than 2^53.  If present, the
      'fraction' MUST consist of no more than 8 decimal digits.  The use
      of commas is optional for readability, and they MUST be ignored.

   receiver-name:  Name of the entity that receives the payment
      (creditor).  The value of this option MAY be subject to lossy
      conversion, modification, and truncation (for example, due to line
      wrapping or character set conversion).

   sender-name:  Name of the entity that makes the payment (debtor).
      The value of this option MAY be subject to lossy conversion,
      modification, and truncation (for example, due to line wrapping or
      character set conversion).

   message:  A short message to identify the purpose of the payment.
      The value of this option MAY be subject to lossy conversion,
      modification, and truncation (for example, due to line wrapping or
      character set conversion).

   instruction:  A short message giving payment reconciliation
      instructions to the recipient.  An instruction that follows the
      character set and length limitation defined by the respective
      payment target type SHOULD NOT be subject to lossy conversion.
*/

const (
	InternationalBankAccountNumber = "iban"
	AutomatedClearingHouse         = "ach"
	UnifiedPaymentInterface        = "upi"
	Bitcoin                        = "bitcoin"
	InterledgerProtocol            = "ilp"
	Void                           = "void"
)

type Payto interface {
	Amount() string
	// receiver-name (creditor)
	ReceiverName() string
	// sender-name (debitor)
	SenderName() string
	// A short message to identify the purpose of the payment.
	Message() string
	// A short message giving payment reconciliation instructions to the recipient.
	Instruction() string
	/**
	  payto-URI = "payto://" authority path-abempty [ "?" opts ]
	  opts = opt *( "&" opt )
	  opt-name = generic-opt / authority-specific-opt
	  opt-value = *pchar
	  opt = opt-name "=" opt-value
	  generic-opt = "amount" / "receiver-name" / "sender-name" /
	                "message" / "instruction"
	  authority-specific-opt = ALPHA *( ALPHA / DIGIT / "-" / "." )
	  authority = ALPHA *( ALPHA / DIGIT / "-" / "." )
	*/
	URL() url.URL
}
