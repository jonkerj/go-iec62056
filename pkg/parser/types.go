package parser

type Telegram struct {
	Identification string
	Objects        []Object
	Checksum       string
}

type Object struct {
	ID        ID
	Value     Value
	Timestamp *string
}

type Value struct {
	Value *string
	Unit  *string
}

type ID struct {
	A *uint8
	B *uint8
	C *uint8
	D *uint8
	E *uint8
	F *uint8
}
