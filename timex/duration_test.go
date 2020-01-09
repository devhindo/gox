package timex

import (
	"fmt"
	"testing"
	"time"
)

// ExampleRound shows how to use the Round() function.
func ExampleRound() {
	ds := []time.Duration{
		time.Hour + time.Second + 123*time.Millisecond, // 1h0m1.123s
		time.Hour + time.Second + time.Microsecond,     // 1h0m1.000001s
		123456789 * time.Nanosecond,                    // 123.456789ms
		123456 * time.Nanosecond,                       // 123.456µs
		123 * time.Nanosecond,                          // 123ns
	}

	fmt.Println("Duration      |0 digits      |1 digit       |2 digits      |3 digits")
	fmt.Println("---------------------------------------------------------------------------")
	for _, d := range ds {
		fmt.Printf("%-14v|", d)
		for digits := 0; digits <= 3; digits++ {
			fmt.Printf("%-14v|", Round(d, digits))
		}
		fmt.Println()
	}

	// Output:
	// Duration      |0 digits      |1 digit       |2 digits      |3 digits
	// ---------------------------------------------------------------------------
	// 1h0m1.123s    |1h0m1s        |1h0m1.1s      |1h0m1.12s     |1h0m1.123s    |
	// 1h0m1.000001s |1h0m1s        |1h0m1s        |1h0m1s        |1h0m1s        |
	// 123.456789ms  |123ms         |123.5ms       |123.46ms      |123.457ms     |
	// 123.456µs     |123µs         |123.5µs       |123.46µs      |123.456µs     |
	// 123ns         |123ns         |123ns         |123ns         |123ns         |
}

func TestRound(t *testing.T) {
	cases := []struct {
		name      string // Name of the test case
		d         time.Duration
		formatted []string // Expected results when using String() on output with index used as digits input
	}{
		{
			"no-fraction-sec",
			time.Hour + time.Second + time.Microsecond, // 1h0m1.000001s
			[]string{"1h0m1s", "1h0m1s", "1h0m1s", "1h0m1s"},
		},
		{
			"no-fraction-ns",
			123 * time.Nanosecond, // 123ns
			[]string{"123ns", "123ns", "123ns", "123ns", "123ns"},
		},
		{
			"s-fraction",
			time.Hour + time.Second + 123*time.Millisecond, // 1h0m1.123s
			[]string{"1h0m1s", "1h0m1.1s", "1h0m1.12s", "1h0m1.123s", "1h0m1.123s", "1h0m1.123s"},
		},
		{
			"ms-fraction",
			123456789 * time.Nanosecond, // 123.456789ms
			[]string{"123ms", "123.5ms", "123.46ms", "123.457ms", "123.4568ms", "123.45679ms", "123.456789ms", "123.456789ms"},
		},
		{
			"µs-fraction",
			123456 * time.Nanosecond, // 123.456µs
			[]string{"123µs", "123.5µs", "123.46µs", "123.456µs", "123.456µs", "123.456µs", "123.456µs", "123.456µs", "123.456µs"},
		},
	}

	for _, c := range cases {
		for digits, formatted := range c.formatted {
			if got := Round(c.d, digits).String(); got != formatted {
				t.Errorf("[%s (digits:%d)] Expected: %v, got: %v", c.name, digits, formatted, got)
			}
		}
	}

	// Test negative digits:
	exp := "123ms"
	if got := Round(123456789*time.Nanosecond, -1).String(); got != exp {
		t.Errorf("[negative-digits] Expected: %v, got: %v", exp, got)
	}
}
