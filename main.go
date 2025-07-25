package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"runtime"
	"sort"
	"strings"
	"text/tabwriter"
	"time"

	"github.com/zihxs/gorm-pgx-benchmarks/helper"

	"github.com/zihxs/gorm-pgx-benchmarks/bench"

	_ "github.com/lib/pq"
)

const VERSION = "v1.0.2"

var defaultBenchmarkNames = []string{"pgx_pool", "gorm"}

type ListOpts []string

func (opts *ListOpts) String() string {
	return fmt.Sprint(*opts)
}

func (opts *ListOpts) Set(value string) error {
	if value != "all" && !strings.Contains(" "+strings.Join(defaultBenchmarkNames, " ")+" ", " "+value+" ") {
		return fmt.Errorf("wrong run name %s", value)
	}

	*opts = append(*opts, value)
	return nil
}

func (opts ListOpts) Shuffle() {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < len(opts); i++ {
		a := rd.Intn(len(opts))
		b := rd.Intn(len(opts))
		opts[a], opts[b] = opts[b], opts[a]
	}
}

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var orms ListOpts
	var all bool

	flag.StringVar(&helper.OrmSource, "source", "host=localhost user=postgres password=postgres dbname=test sslmode=disable", "postgres dsn source")
	flag.Var(&orms, "orm", "orm name: all, "+strings.Join(defaultBenchmarkNames, ", "))
	flag.IntVar(&helper.OrmMaxIdle, "max_idle", 200, "max idle conns")
	flag.IntVar(&helper.OrmMaxConn, "max_conn", 200, "max open conns")
	flag.BoolVar(&helper.DebugMode, "debug", false, "Enable debug mode (also prints not-sorted results of ORMs)")
	version := flag.Bool("version", false, "prints current roxy version")
	flag.Parse()

	if *version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	if len(orms) == 0 {
		all = true
	} else {
		for _, n := range orms {
			if n == "all" {
				all = true
			}
		}
	}

	if all {
		orms = defaultBenchmarkNames
	}
	orms.Shuffle()

	helper.Errors = make(map[string]map[string]string, 0)
	for _, name := range defaultBenchmarkNames {
		helper.Errors[name] = make(map[string]string, 0)
	}

	runBenchmarks(orms)
}

func runBenchmarks(orms ListOpts) {
	benchmarks := map[string]helper.ORMInterface{
		"gorm":     bench.CreateGorm(),
		"pgx_pool": bench.CreatePgxPool(),
	}

	table := new(tabwriter.Writer)
	table.Init(os.Stdout, 0, 8, 2, '\t', tabwriter.AlignRight)

	reports := make(map[string]helper.BenchmarkReport, 0)
	i := 0
	for _, n := range orms {
		orm := benchmarks[n]
		if orm == nil {
			panic(fmt.Sprintf("Unknown ORM: %s", n))
		}

		res, err := helper.RunBenchmarks(orm, reports)
		if err != nil {
			panic(fmt.Sprintf("An error occured while running the benchmarks: %v", err))
		}

		if helper.DebugMode {
			if i != 0 {
				_, _ = fmt.Fprint(table, "\n")
			}
			_, _ = fmt.Fprintf(table, "%s Benchmark Results:\n", n)
			for _, result := range res.Results {
				if result.ErrorMsg == "" {
					_, _ = fmt.Fprintf(table, "%s:\t%d\t%d ns/op\t%d B/op\t%d allocs/op\n", result.Method, result.N, result.NsPerOp, result.MemBytes, result.MemAllocs)
				} else {
					_, _ = fmt.Fprintf(table, "%s:\t%s\n", result.Method, result.ErrorMsg)
				}
			}
			_ = table.Flush()
		}
		i++
	}

	for _, v := range reports {
		sort.Sort(v)
	}

	if helper.DebugMode {
		_, _ = fmt.Fprint(table, "\n")
	}
	_, _ = fmt.Fprintf(table, "Reports:\n\n")

	i = 1
	for method, report := range reports {
		_, _ = fmt.Fprintf(table, "%s\n", method)
		for _, result := range report {
			if result.ErrorMsg == "" {
				_, _ = fmt.Fprintf(table, "%s:\t%d\t%d ns/op\t%d B/op\t%d allocs/op\n", result.Name, result.N, result.NsPerOp, result.MemBytes, result.MemAllocs)
			} else {
				_, _ = fmt.Fprintf(table, "%s:\t%s\n", result.Name, result.ErrorMsg)
			}
		}

		if i != len(reports) {
			_, _ = fmt.Fprintf(table, "\n")
		}
		i++
	}
	_ = table.Flush()
}
