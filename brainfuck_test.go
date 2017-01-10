package brainfuck

import (
	"strings"
	"testing"
)

func TestExec(t *testing.T) {
	r, err := Exec("++++++++++[>+++++++>++++++++++>+++>+<<<<-]>++.>+.+++++++..+++.>++.<<+++++++++++++++.>.+++.------.--------.>+.>.")
	if err != nil {
		t.Errorf("Error occurred while executing: %v", err)
	}

	if strings.TrimSpace(string(r)) != "Hello World!" {
		t.Errorf("Returned string (%s) should be `Hello World!`", string(r))
	}

	_, err = Exec("+[->,----------")
	if err == nil {
		t.Errorf("Error `Couldn't find loop's end` is expected")
	}

	_, err = Exec("+->],----------")
	if err == nil {
		t.Errorf("Error `Couldn't find loop's beginning` is expected")
	}
}
