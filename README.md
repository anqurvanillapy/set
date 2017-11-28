# set

A missing `set` toy in Go.

Again, **TOY ALERT**!

## Install

```bash
$ go get -u github.com/anqurvanillapy/set
```

## Usage

```go
import "github.com/anqurvanillapy/set"

func main() {
    us := set.NewUSet()
    us.Add(42) //=> false
    us.Has(42) //=> true
    us.Add(42) //=> true
    us.Delete(42)
}
```

## License

MIT
