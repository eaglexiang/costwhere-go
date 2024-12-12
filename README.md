# costwhere

show where your project cost time

## install

tool install

```bash
go install github.com/eaglexiang/costwhere-go/cmd/costwhere

```

dependency install

```bash
yay -S flamegraph # archlinux

```

> [FlameGraph](https://github.com/brendangregg/FlameGraph)

## use

```bash
costwhere if="./data.json"
# output flamegraph.svg

```