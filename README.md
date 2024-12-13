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
		stacks, err := cw.EndWithJSON()
		if err != nil {
			log.Printf("%+v", err)
			return
		}

		log.Println(string(stacks)) // output JSON
	}()

	F(ctx)
}

func F(ctx context.Context) {
	defer costwhere.Mark(&ctx, "F")()

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