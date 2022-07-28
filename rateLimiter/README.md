## Sliding Window Rate Limiter

A Sliding Window Rate Limiter implementation. The idea of sliding window is to keep track of all requests timestamp and calculate whether the counter exceeds in the past fixed period when a request arrives.


![SlidingWindowExample](https://habrastorage.org/webt/nv/8a/9z/nv8a9zz0_ycg7sm4htfair2fk44.gif)

### Usage

```go
func main() {
	rl := rateLimiter.New(5, time.Second)
	for i := 1; i <= 10; i++ {
		if rl.Take() {
			fmt.Printf("event %d passed\n", i)
		} else {
			fmt.Printf("event %d rejected\n", i)
		}
	}
}
```

