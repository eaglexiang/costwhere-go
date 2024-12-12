# costwhere

show where your project cost time

## install

tool install

```bash
go install github.com/eaglexiang/costwhere-go/cmd/costwhere

```

```go
func main() {
	ctx := context.Background()
	ctx, cw := costwhere.Init(ctx, "main")
	defer func() {
		stacks := cw.End()
		buf, _ := json.Marshal(stacks)
        fmt.Println(string(buf)) // output JSON
	}()

	F(ctx)
}

func F(ctx context.Context) {
	ctx, end := costwhere.Begin(ctx, "F")
	defer end()
	time.Sleep(100 * time.Millisecond)
}

```

dependency install

```bash
yay -S flamegraph # archlinux

```

> [FlameGraph](https://github.com/brendangregg/FlameGraph)

## use

```bash
costwhere -if="./data.json" -of="./flamegraph.svg"
# output flamegraph.svg

```