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

	// 初始化并回收 costwhere 采集
	ctx, cw := costwhere.Init(ctx)
	defer func() {
		stacks, err := cw.EndWithJSON() // 以 JSON 格式对采集结果进行输出
		if err != nil {
			log.Printf("%+v", err)
			return
		}

		err = os.WriteFile("costwhere.json", stacks, 0644) // 将采集结果保存到文件（或输出到日志）
		if err != nil {
			log.Printf("%+v", err)
		}
	}()

	F(ctx)
}

func F(ctx context.Context) {
	defer costwhere.Mark(ctx)() // 在任何需要进行耗时统计的地方复制粘贴此代码

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

## output

![](./flamegraph.svg)