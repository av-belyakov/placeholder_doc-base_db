package jsonprocessing

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/av-belyakov/placeholder_doc-base_db/internal/supportingfunctions"
)

var (
	dataByte []byte

	err error
)

func TestMain(m *testing.M) {
	dataByte, err = os.ReadFile("../test_json/case_39100.json")
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(m.Run())
}

func TestGetElementFromJSON(t *testing.T) {
	var num int
	for range 3 {
		data, err := supportingfunctions.GetElementsFromJSON(t.Context(), dataByte)
		if err != nil {
			fmt.Println("ERROR:", err)
		}

		fmt.Println("COUNT elements:", len(data))
		num++
	}

	assert.Equal(t, num, 3)
}

/*
func BenchmarkGetElementFromJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		data, err := supportingfunctions.GetElementsFromJSON(b.Context(), dataByte)
		if err != nil {
			fmt.Println("ERROR:", err)
		}

		fmt.Println("COUNT elements:", len(data))
	}
}
*/
