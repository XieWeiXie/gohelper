# gist

> 收集代码

```text
var wg sync.WaitGroup
const max = 1000
for i := 1; i < max; i++ {
    wg.Add(1)
    go func() {
        defer wg.Done()
        IOCounters()
    }()
}
wg.Wait()
```