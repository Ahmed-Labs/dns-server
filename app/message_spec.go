package main

type QuestionType uint16

const (
	QTYPE_A     QuestionType = 1  // a host address
	QTYPE_NS    QuestionType = 2  // an authoritative name server
	QTYPE_MD    QuestionType = 3  // a mail destination (Obsolete - use MX)
	QTYPE_MF    QuestionType = 4  // a mail forwarder (Obsolete - use MX)
	QTYPE_CNAME QuestionType = 5  // the canonical name for an alias
	QTYPE_SOA   QuestionType = 6  // marks the start of a zone of authority
	QTYPE_MB    QuestionType = 7  // a mailbox domain name (EXPERIMENTAL)
	QTYPE_MG    QuestionType = 8  // a mail group member (EXPERIMENTAL)
	QTYPE_MR    QuestionType = 9  // a mail rename domain name (EXPERIMENTAL)
	QTYPE_NULL  QuestionType = 10 // a null RR (EXPERIMENTAL)
	QTYPE_WKS   QuestionType = 11 // a well-known service description
	QTYPE_PTR   QuestionType = 12 // a domain name pointer
	QTYPE_HINFO QuestionType = 13 // host information
	QTYPE_MINFO QuestionType = 14 // mailbox or mail list information
	QTYPE_MX    QuestionType = 15 // mail exchange
	QTYPE_TXT   QuestionType = 16 // text strings
)

type QuestionClass uint16

const (
	QCLASS_IN  QuestionClass = 1   // Internet
	QCLASS_CS  QuestionClass = 2   // CSNET (obsolete)
	QCLASS_CH  QuestionClass = 3   // CHAOS
	QCLASS_HS  QuestionClass = 4   // Hesiod
	QCLASS_ANY QuestionClass = 255 // Any class
)
