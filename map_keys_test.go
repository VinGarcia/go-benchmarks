package main

import (
	"fmt"
	"math/rand"
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

	fmt.Println(ageCache)

	b.Run("using strings optimized", func(b *testing.B) {
		for n := 0; n < b.N; n++ {
			i := rand.Intn(15)
			j := rand.Intn(15)
			key := "keyI" + strconv.Itoa(i) + "|keyJ" + strconv.Itoa(j) + "|keyIJ" + strconv.Itoa(i+j)
			ageCache = stringMap[key].Age
		}
	})

	fmt.Println(ageCache)

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

	fmt.Println(ageCache)
}
