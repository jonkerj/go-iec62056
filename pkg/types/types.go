package types

type Telegram struct {
	Identification string
	Objects        []Object
	Checksum       string
}

type Object struct {
	ID        string
	Value     Value
	Timestamp *string
}

type Value struct {
	Value *string
	Unit  *string
}
