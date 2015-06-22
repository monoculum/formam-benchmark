formam-benchmark
================

Benchmark about formam package (compared with [ajg/form](https://github.com/ajg/form) and [gorilla/schema](https://github.com/gorilla/schema) and [built-in/json](http://golang.org/pkg/encoding/json/))

In a Mac i7 2,8Ghz 4 cores, 8 GB 1067Mhz DDR3

### test 1

**NOTE**: `gorilla/schema` not support map...

```
BenchmarkAJGFormTest1	   20000	     64840 ns/op	    5886 B/op	     116 allocs/op
BenchmarkFormamTest1	   50000	     28037 ns/op	    2027 B/op	      86 allocs/op
BenchmarkJSONTest1	      100000	     23509 ns/op	    1761 B/op	      34 allocs/op
```

### test 2

```
BenchmarkAJGFormTest2	   50000	     35664 ns/op	    3156 B/op	      66 allocs/op
BenchmarkSchemaTest2	   50000	     30770 ns/op	    3504 B/op	      80 allocs/op
BenchmarkFormamTest2	  100000	     12305 ns/op	     844 B/op	      43 allocs/op
BenchmarkJSONTest2	      100000	     13712 ns/op	     848 B/op	      17 allocs/op
```

Conclusion
----------

Formam is twice faster than [ajg/form](https://github.com/ajg/form) and [gorilla/schema](https://github.com/gorilla/schema), and equal to or slightly slower than [built-in/json](http://golang.org/pkg/encoding/json/).
Furthermore, Formam allocates in memory twice less than [ajg/form](https://github.com/ajg/form) and [gorilla/schema](https://github.com/gorilla/schema), and slightly more than [built-in/json](http://golang.org/pkg/encoding/json/).
