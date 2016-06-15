formam-benchmark
================

Benchmark about formam package (compared with [ajg/form](https://github.com/ajg/form) and [gorilla/schema](https://github.com/gorilla/schema) and [built-in/json](http://golang.org/pkg/encoding/json/))

In a macbook air i5 1.7GHz, 4 GB 1333MHz DDR3 and Go 1.7beta1

### test 1

**NOTE**: `gorilla/schema` not support map... `formam` has more features than rest to decoder, so has more work.

```
BenchmarkAJGFormTest1-4   	   30000	     46746 ns/op	    5868 B/op	     114 allocs/op
BenchmarkFormamTest1-4    	   50000	     38204 ns/op	    2294 B/op	     104 allocs/op
BenchmarkFormTest1-4      	   50000	     34370 ns/op	    6776 B/op	     165 allocs/op
BenchmarkJSONTest1-4      	  100000	     18688 ns/op	    1696 B/op	      32 allocs/op
```

### test 2

```
BenchmarkAJGFormTest2-4   	  100000	     20980 ns/op	    3152 B/op	      66 allocs/op
BenchmarkSchemaTest2-4    	  100000	     20826 ns/op	    4025 B/op	      79 allocs/op
BenchmarkFormamTest2-4    	  100000	     16675 ns/op	     933 B/op	      54 allocs/op
BenchmarkFormTest2-4      	  200000	      9780 ns/op	    2504 B/op	      45 allocs/op
BenchmarkJSONTest2-4      	  200000	     10488 ns/op	     848 B/op	      17 allocs/op
```

Conclusion
----------

Formam is twice faster than [ajg/form](https://github.com/ajg/form) and [gorilla/schema](https://github.com/gorilla/schema), and equal to or slightly slower than [built-in/json](http://golang.org/pkg/encoding/json/).
Furthermore, Formam allocates in memory twice less than [ajg/form](https://github.com/ajg/form) and [gorilla/schema](https://github.com/gorilla/schema), and slightly more than [built-in/json](http://golang.org/pkg/encoding/json/).
