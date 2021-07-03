### 文档上的注意点⚠️
1. Should always time.Time, not *time.Time 
2. 大部分时间是 concurrency-safe 的, 除了 GobDecode, UnmarshalBinary, UnmarshalJSON
```
A Time value can be used by multiple goroutines simultaneously except
that the methods GobDecode, UnmarshalBinary, UnmarshalJSON and
UnmarshalText are not concurrency-safe.
```
3. 零值, a simple way of detecting a time that has not been initialized explicitly.
```
The zero value of type Time is January 1, year 1, 00:00:00.000000000 UTC.
```

4. See the “Monotonic Clocks” section in the package documentation for details.

5. 什么是 "Monotonic Clocks" 和 “wall clock”

6. prefer t.Equal(u), // to t == u

7. wall and ext encode the wall time seconds, wall time nanoseconds,