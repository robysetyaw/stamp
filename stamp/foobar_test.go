package stamp

import (
	"reflect"
	"testing"
)

func TestStamp(t *testing.T) {
	runTestStamp(t, "Should return []string{'1'} when input = 1", 1, []string{"1"})
	runTestStamp(t, "Should return []string{'1', '4', 'Foo'} when input = 6", 6, []string{"Foo", "4", "1"})
	runTestStamp(t, "Should return []string{'1','4', 'Foo',  '8', 'Foo', 'Bar', 'Foo','14', 'FooBar', '16'}) when input = 16", 16, []string{"16", "FooBar", "14", "Foo", "Bar", "Foo", "8", "Foo", "4", "1"})
}

func runTestStamp(t *testing.T, description string, input int, expectation []string) {
	t.Run(description, func(t *testing.T) {
		got := Foobar(input)
		if !reflect.DeepEqual(got, expectation) {
			t.Errorf("Failed %s, FooBar(%v) = %s, want %v", description, input, got, expectation)
		}
	})
}
