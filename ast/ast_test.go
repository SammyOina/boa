package ast

import (
	"testing"

	"github.com/sammyoina/boa/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&AssignStatement{
				Token: token.Token{
					Type:    token.IDENTIFIER,
					Literal: "x",
				},
				Name: &Identifier{
					Token: token.Token{
						Type:    token.IDENTIFIER,
						Literal: "x",
					},
					Value: "x",
				},
				Value: &Identifier{
					Token: token.Token{
						Type:    token.IDENTIFIER,
						Literal: "y",
					},
					Value: "y",
				},
			},
		},
	}
	if program.String() != "x = y" {
		t.Errorf("program>String() wrong got %q expected x = y", program.String())
	}
}
