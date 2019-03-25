package hpp

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

var input = `<!DOCTYPE html><html lang="en"><head><title>Pug</title><script type="text/javascript">if (foo) {
            bar(1 + 5)
        }
        </script></head><body><h1>Pug - template engine</h1><div id="container" class="col">
        <p>You are amazing</p><form><br>First name:<input type="text" name="firstname"><br>Last name:
        <input type="text" name="lastname"></form><p>Pug is a terse and simple templating
        language with a <b>strong</b> focus on 
        performance and powerful features.</p></div></body></html>`

var expect = `<!DOCTYPE html>
<html lang="en">
    <head>
        <title>Pug</title>
        <script type="text/javascript">
            if (foo) {
                bar(1 + 5)
            } 
        </script>
    </head>
    <body>
        <h1>Pug - template engine</h1>
        <div id="container" class="col">
            <p>You are amazing</p>
            <form>
                <br>First name:
                <input type="text" name="firstname">
                <br>Last name: 
                <input type="text" name="lastname">
            </form>
            <p>
                Pug is a terse and simple templating
                language with a <b>strong</b> focus on 
                performance and powerful features.
            </p>
        </div>
    </body>
</html>
`

func TestFormat(t *testing.T) {
	bf := new(bytes.Buffer)
	Format(strings.NewReader(input), bf)
	output := bf.String()
	if expect != strings.TrimLeft(output, "\n\r\t ") {
		t.Errorf("\n------ expect:\n%s\n------ output:\n%s", expect, output)
	}
}

func TestPrint(t *testing.T) {
	output := string(Print(strings.NewReader(input)))
	if expect != output {
		t.Errorf("\n------ expect:\n%s\n------ output:\n%s", expect, output)
	}
}

func TestPrPrint(t *testing.T) {
	output := PrPrint(input)
	if expect != output {
		t.Errorf("\n------ expect:\n%s\n------ output:\n%s", expect, output)
	}

	fmt.Print(output)
}
