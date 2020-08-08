package normalizer_test

import (
	"fmt"
	"reflect"
	// "strings"
	"testing"
	// "unicode"

	// "golang.org/x/text/transform"
	// "golang.org/x/text/unicode/norm"

	"github.com/sugarme/tokenizer/normalizer"
)

func TestNormalized_NewNormalizedFrom(t *testing.T) {
	gotN := normalizer.NewNormalizedFrom("élégant")
	gotN.NFD()

	want := []normalizer.Alignment{
		{0, 1},
		{0, 1},
		{1, 2},
		{2, 3},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
	}
	got := gotN.Get().Alignments

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want: %v\n", want)
		t.Errorf("Got: %v\n", got)
	}
}

// Unchanged: Remove accents - Mark, nonspacing (Mn)
func TestNormalized_RemoveAccents(t *testing.T) {
	gotN := normalizer.NewNormalizedFrom("élégant")
	gotN.RemoveAccents()

	want := []normalizer.Alignment{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{5, 6},
		{6, 7},
	}
	got := gotN.Get().Alignments

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want: %v\n", want)
		t.Errorf("Got: %v\n", got)
	}
}

// Removed Chars
func TestNormalized_Filter(t *testing.T) {
	gotN := normalizer.NewNormalizedFrom("élégant")

	gotN.Filter('n')

	want := []normalizer.Alignment{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{6, 7},
	}
	got := gotN.Get().Alignments

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want: %v\n", want)
		t.Errorf("Got: %v\n", got)
	}
}

// Mixed addition and removal
func TestNormalized_Mixed(t *testing.T) {
	gotN := normalizer.NewNormalizedFrom("élégant")

	gotN.RemoveAccents()
	gotN.Filter('n')

	want := []normalizer.Alignment{
		{0, 1},
		{1, 2},
		{2, 3},
		{3, 4},
		{4, 5},
		{6, 7},
	}
	got := gotN.Get().Alignments

	if !reflect.DeepEqual(want, got) {
		t.Errorf("Want: %v\n", want)
		t.Errorf("Got: %v\n", got)
	}
}

// Range and Conversion
func TestNormalized_RangeConversion(t *testing.T) {
	gotN := normalizer.NewNormalizedFrom(`    __Hello__   `)

	gotN.Filter(' ')
	fmt.Printf("Original string after filtering: '%v'\n", gotN.GetOriginal())
	gotN.Lowercase()
	fmt.Printf("Original string after lowercase: '%v'\n", gotN.GetOriginal())

	got1, err := gotN.RangeOriginal([]int{6, 7, 8, 9, 10, 11})
	if err != nil {
		t.Fatal(err)
	}
	want1 := "Hello"
	if !reflect.DeepEqual(want1, got1) {
		t.Errorf("Want: %v\n", want1)
		t.Errorf("Got: %v\n", got1)
	}

	got2, err := gotN.Range([]int{6, 7, 8, 9, 10, 11})
	if err != nil {
		t.Fatal(err)
	}
	want2 := "hello"
	if !reflect.DeepEqual(want1, got1) {
		t.Errorf("Want: %v\n", want2)
		t.Errorf("Got: %v\n", got2)
	}

}
