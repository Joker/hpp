package hpp

import (
	"fmt"
	"strings"
	"testing"
)

func TestPrint(t *testing.T) {
	input := `<!DOCTYPE html><html lang="en"><head><title>Pug</title><script type="text/javascript">if (foo) {
				bar(1 + 5)
			}
			</script></head><body><h1>Pug - template engine</h1><div id="container" class="col">
            <p>You are amazing</p><form><br>First name:<input type="text" name="firstname"><br>Last name:
            <input type="text" name="lastname"></form><p>Pug is a terse and simple templating
			language with a <b>strong</b> focus on 
			performance and powerful features.</p></div></body></html>`

	expect := `<!DOCTYPE html>
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

	output := string(Print(strings.NewReader(input)))

	if expect != output {
		t.Errorf("\n------ expect:\n%s\n------ output:\n%s", expect, output)
	} else {
		fmt.Print(output)
	}
}
