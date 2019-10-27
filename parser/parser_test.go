package parser

import (
	"price_calculation/token"
	"reflect"
	"testing"
)

func TestParseTokens(t *testing.T) {
	cases := []struct {
		input  string
		expect []token.Token
	}{
		{"10 11 + 100 * / < = ? a b", []token.Token{
			token.NewIntegerToken("10"),
			token.NewIntegerToken("11"),
			&token.PlusToken{},
			token.NewIntegerToken("100"),
			&token.MulToken{},
			&token.DivToken{},
			&token.LtToken{},
			&token.EqToken{},
			&token.LogicToken{},
			token.NewVariableToken("a"),
			token.NewVariableToken("b"),
		}},
	}

	parser := Parser{}
	for _, testCase := range cases {
		err, result := parser.Parse(testCase.input)
		if err != nil {
			t.Errorf("Parse error for input \"%q\": %q", testCase.input, err)
		}
		if !reflect.DeepEqual(testCase.expect, result) {
			t.Errorf("Expect %q. Got %q", testCase.expect, result)
		}
	}
}
