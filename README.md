# HTML Pretty Print for Go (golang)
Package hpp (github.com/Joker/hpp) is a HTML formatter for Go.


example:
```html
<!DOCTYPE html><html lang="en"><head><title>Pug</title><script type="text/javascript">if (foo) {
    bar(1 + 5)
}
</script></head><body><h1>Pug - template engine</h1><div id="container" class="col">
<p>You are amazing</p><form><br>First name:<input type="text" name="firstname"><br>Last name:
<input type="text" name="lastname"></form><p>Pug is a terse and simple templating
language with a <b>strong</b> focus on 
performance and powerful features.</p></div></body></html>
```
becomes
```html
<!DOCTYPE html>
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
```


### Installation

```sh
$ go get github.com/Joker/hpp
```