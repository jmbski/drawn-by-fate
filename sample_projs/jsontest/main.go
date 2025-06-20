package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type MyEnum string
type IntEnum int

const (
	EnumOne   MyEnum = "EnumOne"
	EnumTwo   MyEnum = "EnumTwo"
	EnumThree MyEnum = "EnumThree"
)

const (
	IntEnumA IntEnum = iota
	IntEnumB
	IntEnumC
)

type Config struct {
	Mode    MyEnum  `json:"mode"`
	IntEnum IntEnum `json:"int_enum"`
}

func (e *MyEnum) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	switch MyEnum(s) {
	case EnumOne, EnumTwo, EnumThree:
		*e = MyEnum(s)
		return nil
	default:
		return fmt.Errorf("invalid MyEnum value: %q", s)
	}
}

func main() {
	jsonStr := `{"mode":"EnumTwo", "int_enum": "IntEnumA"}`

	var cfg Config
	if err := json.Unmarshal([]byte(jsonStr), &cfg); err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg.Mode, cfg.IntEnum) // Output: EnumTwo

}
