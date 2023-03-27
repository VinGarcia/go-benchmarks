package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"testing"
)

type SomeStruct struct {
	Name string
	Age  int
}

// Used to make sure the compiler dos not try to optimize the benchmarks below:
var ageCache int

func BenchmarkMapKeys(b *testing.B) {
	file, err := os.CreateTemp("", "*")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(file.Name())

	stringMap := map[string]*SomeStruct{}
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			key := fmt.Sprintf("keyI%d|keyJ%d|keyIJ%d", i, j, i+j)
			stringMap[key] = &SomeStruct{
				Name: key,
				Age:  i + j,
			}
		}
	}

	b.Run("using strings", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			i := rand.Intn(15)
			j := rand.Intn(15)
			key := fmt.Sprintf("keyI%d|keyJ%d|keyIJ%d", i, j, i+j)
			ageCache = stringMap[key].Age
		}
	})

	// This is done just to force Go not to optimize the benchmarks
	// by ignoring the contents of the map:
	fmt.Fprintln(file, ageCache)

	// Otimization so that we don't have to convert integers to string inside the benchmark:
	var intToStr = make([]string, 30)
	for i := 0; i < 30; i++ {
		intToStr[i] = strconv.Itoa(i)
	}

	b.Run("using strings optimized", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			i := rand.Intn(15)
			j := rand.Intn(15)
			key := "keyI" + intToStr[i] + "|keyJ" + intToStr[j] + "|keyIJ" + intToStr[i+j]
			ageCache = stringMap[key].Age
		}
	})

	// This is done just to force Go not to optimize the benchmarks
	// by ignoring the contents of the map:
	fmt.Fprintln(file, ageCache)

	type keyType struct {
		KeyI  int
		KeyJ  int
		KeyIJ int
	}
	structMap := map[keyType]*SomeStruct{}
	for i := 0; i < 15; i++ {
		for j := 0; j < 15; j++ {
			key := keyType{
				KeyI:  i,
				KeyJ:  j,
				KeyIJ: i + j,
			}
			structMap[key] = &SomeStruct{
				Name: fmt.Sprintf("keyI%d|keyJ%d|keyIJ%d", i, j, i+j),
				Age:  i + j,
			}
		}
	}

	b.Run("using structs", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			i := rand.Intn(15)
			j := rand.Intn(15)
			key := keyType{
				KeyI:  i,
				KeyJ:  j,
				KeyIJ: i + j,
			}
			ageCache = structMap[key].Age
		}
	})

	// This is done just to force Go not to optimize the benchmarks
	// by ignoring the contents of the map:
	fmt.Fprintln(file, ageCache)
}
