RandomPick

A simple function which takes a map of weighted options and picks one at random. This might be useful for a lottery or raffle where participants might have multiple tickets.

```go
rand.Seed(time.Now().UnixNano())

m := map[string]int{
    "Bob": 900,
    "Steve": 2000,
    "Kyle": 300,
}

fmt.Println(randompick.Result(m))
```