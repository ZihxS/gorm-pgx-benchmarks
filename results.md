### Raw Performance Benchmarks
```
Read
gorm:        52472  ns/op      4749 B/op        96 allocs/op
pgxpool:     33164  ns/op       963 B/op        19 allocs/op

ReadSlice
gorm:        196773 ns/op     44654 B/op      2190 allocs/op
pgxpool:      83255 ns/op     30381 B/op       513 allocs/op

Insert
gorm:        122017 ns/op      6148 B/op        97 allocs/op
pgxpool:      97497 ns/op       298 B/op         8 allocs/op

InsertMulti
gorm:        2784810 ns/op   253598 B/op      5216 allocs/op
pgxpool:      435200 ns/op    49439 B/op        39 allocs/op

Update
gorm:        123075 ns/op      5829 B/op        89 allocs/op
pgxpool:     106403 ns/op       306 B/op         8 allocs/op
```

### Performance Benchmarks (Table)

| Operation | Metric | GORM | pgxpool |
|----|----|----|----|
| **Read** | Nanoseconds per Operation (ns/op) | 52.472 | **33.164** |
|    | Bytes per Operation (B/op) | 4.749 | **963** |
|    | Allocations per Operation (allocs/op) | 96 | **19** |
| **ReadSlice** | Nanoseconds per Operation (ns/op) | 196.773 | **83.255** |
|    | Bytes per Operation (B/op) | 44.654 | **30.381** |
|    | Allocations per Operation (allocs/op) | 2.190 | **513** |
| **Insert** | Nanoseconds per Operation (ns/op) | 122.017 | **97.497** |
|    | Bytes per Operation (B/op) | 6.148 | **298** |
|    | Allocations per Operation (allocs/op) | 97 | **8** |
| **Insert Multiple** **(100 Rows)** | Nanoseconds per Operation (ns/op) | 2.784.810 | **435.200** |
|    | Bytes per Operation (B/op) | 253.598 | **49.439** |
|    | Allocations per Operation (allocs/op) | 5.216 | **39** |
| **Update** | Nanoseconds per Operation (ns/op) | 123.075 | **106.403** |
|    | Bytes per Operation (B/op) | 5.829 | **306** |
|    | Allocations per Operation (allocs/op) | 89 | **8** |

### Total Memory Usage and Allocations for 100.000 Operations

---

| Operation | Metric | GORM | pgxpool |
|----|----|----|----|
| **Read** | Total Bytes (MB) | 474 MB | 96 MB |
|    | Total Memory Allocations | 9.600.000 | 1.900.000 |
| **ReadSlice** | Total Bytes (MB) | 4.465 MB (4.47 GB) | 3.038 MB (3.04 GB) |
|    | Total Memory Allocations | 219.000.000 | 51.300.000 |
| **Insert** | Total Bytes (MB) | 614 MB | 29 MB |
|    | Total Memory Allocations | 9.700.000 | 800.000 |
| **Insert Multiple** **(100 Rows)** | Total Bytes (MB) | 25.359 MB (25.36 GB) | 4.943 MB (4.94 GB) |
|    | Total Memory Allocations | 521.600.000 | 3.900.000 |
| **Update** | Total Bytes (MB) | 582.90 MB | 30.60 MB |
|    | Total Memory Allocations | 8.900.000 | 800.000 |
