package times

import (
	"os"
	"testing"
)

func BenchmarkGet(t *testing.B) {
	fileTest(t, func(f *os.File) {
		fi, err := os.Stat(f.Name())
		if err != nil {
			t.Error(err)
		}

		for t.Loop() {
			Get(fi)
		}
	})
	t.ReportAllocs()
}

func BenchmarkStatFile(t *testing.B) {
	fileTest(t, func(f *os.File) {
		for t.Loop() {
			StatFile(f)
		}
	})
	t.ReportAllocs()
}

func BenchmarkStat(t *testing.B) {
	fileTest(t, func(f *os.File) {
		for t.Loop() {
			Stat(f.Name())
		}
	})
	t.ReportAllocs()
}

func BenchmarkLstat(t *testing.B) {
	fileTest(t, func(f *os.File) {
		for t.Loop() {
			Lstat(f.Name())
		}
	})
	t.ReportAllocs()
}

func BenchmarkOsStat(t *testing.B) {
	fileTest(t, func(f *os.File) {
		for t.Loop() {
			os.Stat(f.Name())
		}
	})
	t.ReportAllocs()
}

func BenchmarkOsLstat(t *testing.B) {
	fileTest(t, func(f *os.File) {
		for t.Loop() {
			os.Lstat(f.Name())
		}
	})
	t.ReportAllocs()
}
