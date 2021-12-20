package repl

import (
	"bufio"
	"fmt"
	"github.com/inspiration33/interpreter-go/lexer"
	"github.com/inspiration33/interpreter-go/token"
	"io"
)

// REPL: Read(읽기) Eval(평가) Print(출력) Loop(반복)
/**
* 입력을 읽고(Read),
* 인터프리터에 보내 평가(Eval)하고,
* 인터프리터의 결과물을 출력하고(Print)
* 이런 동작을 반복(Loop)하기 때문에 REPL (READ, Eval, Print, Loop)이다.
 */

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
	}

	line := scanner.Text()
	l := lexer.New(line)

	for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
		fmt.Fprintf(out, "%+v\n", tok)
	}
}
