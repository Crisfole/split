package word

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
	"testing"
)

var testStrings []string = []string{
	`In Düsseldorf I took my hat off. But I can't put it back on.`,
	"On a $50,000 mortgage of 30 years at 8 percent, the monthly payment would be $366.88.",
	"\"We beat some pretty good teams to get here,\" Slocum said.",
	"Well, we couldn't have this predictable, cliche-ridden, \"Touched by an Angel\" (a show creator John Masius worked on) wanna-be if she didn't.",
	"I cannot cannot work under these conditions!",
	"The company spent $30,000,000 last year.",
	"The company spent 40.75% of its income last year.",
	"He arrived at 3:00 pm.",
	"I bought these items: books, pencils, and pens.",
	"Though there were 150, 100 of them were old.",
	"There were 300,000, but that wasn't enough.",
	"Good muffins cost $3.88\nin New York.  Please buy me\ntwo of them.\n\nThanks.",
	"Alas, it has not rained today. When, do you think, will it rain again?",
}

func TestTokenize(t *testing.T) {
	for _, str := range testStrings {
		test(t, str)
	}
}

func test(t *testing.T, str string) {
	tokenized := strings.TrimSpace(Tokenize(str))
	pyTokenized := strings.TrimSpace(pyTokenize(str))
	fmt.Println(str)
	fmt.Println("\t", tokenized)
	fmt.Println("\t", pyTokenized)
	if tokenized != pyTokenized {
		t.Errorf("Failed On: %s\nGo:\t%s\tLen: %d\nPython:\t%s\tLen: %d\n\n", str, tokenized, len(tokenized), pyTokenized, len(pyTokenized))
	}
}

func pyTokenize(str string) string {
	cmd := exec.Command("python", "word_tokenize.py")
	cmd.Stdin = strings.NewReader(str)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		fmt.Errorf(err.Error())
		return ""
	}
	return out.String()
}
