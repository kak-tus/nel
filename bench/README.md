## Benchmarks

In Go code you usually use structures with some Pointer fields like this

```
type pointerSmall struct {
	Field000 *string
	Field001 *string
	Field002 *string
	Field003 *string
	Field004 *string
	Field005 *string
}
```

Let's try compare this struct benchmark with benchmark for struct with simple Values

```
type valueSmall struct {
	Field000 string
	Field001 string
	Field002 string
	Field003 string
	Field004 string
	Field005 string
}
```

Struct with Pointers processed with zero allocation, as fields with Values. And we see here better processing time for struct with Pointers.

Small notice: here we can see two Go mechanics, that can explain this results.

First - is inlining, that do replace of function call to the body of this function. In this case inlining allow to don't do escape Pointer to heap. As result - no allocations. All next benchmarks will be doing with enabled inlining, as in usual Go code.

Second - is copying structs. Spended time for copy struct get worse processing time per operation for struct with Values.

```
BenchmarkPointerSmall-8   	1000000000	         0.295 ns/op	       0 B/op	       0 allocs/op
BenchmarkValueSmall-8     	184702404	         6.51 ns/op	       0 B/op	       0 allocs/op
```

Let's do some function calls with structs. Structs with Pointers still processed with zero allocations and with better time, then structs with Values.

```
BenchmarkPointerSmallChain-8   	1000000000	         0.297 ns/op	       0 B/op	       0 allocs/op
BenchmarkValueSmallChain-8     	59185880	        20.3 ns/op	       0 B/op	       0 allocs/op
```

We can use JSONs for input and output in our service. Try to do marshalling and unmarshalling with jsoniter. Here situation changed. Struct with Values got lower allocations, memory usage, lower processing time.

```
BenchmarkPointerSmallJSON-8   	   49522	     23724 ns/op	   14122 B/op	      28 allocs/op
BenchmarkValueSmallJSON-8     	   52234	     22806 ns/op	   14011 B/op	      15 allocs/op
```

Let's try to improve json speed with easyjson. Better mostly all results for both structs, except slightly bigger memory usage per operations.

```
BenchmarkPointerSmallEasyJSON-8   	   64482	     17815 ns/op	   14591 B/op	      21 allocs/op
BenchmarkValueSmallEasyJSON-8     	   63136	     17537 ns/op	   14444 B/op	      14 allocs/op
```

**Intermediate conclusion**: if your code processing pipeline is to produce some value, chains it thru some function call and got some result - sometimes better to use struct with Pointers. But if your processing is to do some conversions (marshalling, unmarshalling) with structs - prefer to use struct with Values.

Go next. Sometimes struct growths bigger.

```
type pointerBig struct {
	Field000 *string
	...
	Field999 *string
}
```

```
type valueBig struct {
	Field000 string
	...
	Field999 string
}
```

Try to do benchmark for this struct. Here we see, that struct with Values left zero allocations and got bigger processing time (this is normal, because struct is bigger). And struct with Pointers lost advantage - non zero allocations, much worst processing time and memory usage.

```
BenchmarkPointerBig-8   	   36787	     32243 ns/op	   24192 B/op	    1001 allocs/op
BenchmarkValueBig-8     	  721375	      1613 ns/op	       0 B/op	       0 allocs/op
```

Try function calls chain. Not changed anything in struct with Pointers and slightly added processing time for struct with Values (still lower then struct with Pointers).

```
BenchmarkPointerBigChain-8   	   36607	     31709 ns/op	   24192 B/op	    1001 allocs/op
BenchmarkValueBigChain-8     	  351693	      3216 ns/op	       0 B/op	       0 allocs/op
```

Try do marshal and unmarshal. Struct with Values better in all.

```
BenchmarkPointerBigJSON-8   	     250	   4640020 ns/op	 5326593 B/op	    4024 allocs/op
BenchmarkValueBigJSON-8     	     270	   4289834 ns/op	 4110721 B/op	    2015 allocs/op
```

Try to improve result with easyjson. Struct with Values better in all. Also better, then jsoniter.

```
BenchmarkPointerBigEasyJSON-8   	     364	   3204100 ns/op	 2357440 B/op	    3066 allocs/op
BenchmarkValueBigEasyJSON-8     	     380	   3058639 ns/op	 2302248 B/op	    1063 allocs/op
```

**Final conclusion**: Don't do optimisations at first code stage - prefer to use struct with Values, then struct with Pointers. And only when perfomance tuning needed - review your processing pipeline and try switch to struct with Pointers on hot places. Prefer to use codegens (easyjson and others) then native runtime processing - they get better results in mostly cases.
