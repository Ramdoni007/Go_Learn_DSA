package Go_learning_DSA

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	testCases := []struct {
		n    int
		want string
	}{
		{1, "1\n"},
		{2, "1, 2\n"},
		{3, "1, 2, Fizz\n"},
		{4, "1, 2, Fizz, 4\n"},
		{5, "1, 2, Fizz, 4, Buzz\n"},
		{6, "1, 2, Fizz, 4, Buzz, Fizz\n"},
		{7, "1, 2, Fizz, 4, Buzz, Fizz, 7\n"},
		{8, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8\n"},
		{9, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz\n"},
		{10, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz\n"},
		{11, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11\n"},
		{12, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz\n"},
		{13, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13\n"},
		{14, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14\n"},
		{15, "1, 2, Fizz, 4, Buzz, Fizz, 7, 8, Fizz, Buzz, 11, Fizz, 13, 14, Fizz Buzz\n"},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf(""), func(t *testing.T) {
			testStdout, writer, err := os.Pipe()
			if err != nil {
				t.Fatalf("os.Pipe() err = %v; want  %v", err, nil)
			}
			osStdout := os.Stdout // keep backup of the real stdout
			os.Stdout = writer
			defer func() {
				os.Stdout = osStdout
			}()
			FizzBuzz(tc.n)
			writer.Close()

			var buf bytes.Buffer
			io.Copy(&buf, testStdout)
			got := buf.String()
			if got != tc.want {
				t.Fatalf("FizzBuzz(%d) = %q; want %q", tc.n, got, tc.want)
			}

		})
	}

}
