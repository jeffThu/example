// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type Character struct {
	ID         string     `json:"id"`
	Name       string     `json:"name"`
	IsHero     bool       `json:"isHero"`
	CliqueType CliqueType `json:"cliqueType"`
}

type CharacterInput struct {
	Name       string     `json:"name"`
	ID         *string    `json:"id,omitempty"`
	IsHero     *bool      `json:"isHero,omitempty"`
	CliqueType CliqueType `json:"cliqueType"`
}

type Mutation struct {
}

type Query struct {
}

type CliqueType string

const (
	// People who are elite with parents having money
	CliqueTypeKooks CliqueType = "KOOKS"
	// People who desperate to move up the social ladder to become new versions of themselves and establish new beginnings
	CliqueTypePogues CliqueType = "POGUES"
)

var AllCliqueType = []CliqueType{
	CliqueTypeKooks,
	CliqueTypePogues,
}

func (e CliqueType) IsValid() bool {
	switch e {
	case CliqueTypeKooks, CliqueTypePogues:
		return true
	}
	return false
}

func (e CliqueType) String() string {
	return string(e)
}

func (e *CliqueType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CliqueType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CliqueType", str)
	}
	return nil
}

func (e CliqueType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
