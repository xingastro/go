package main

import (
	"bytes"
	"io"
	"testing"
)

func algOne(data []byte, find []byte, repl []byte, output *bytes.Buffer) {
	// Use a bytes Buffer to provide a stream to process.
	input := bytes.NewBuffer(data)

	// The number of bytes we are looking for.
	size := len(find)

	// Declare the buffers we need to process the stream.
	buf := make([]byte, size)
	end := size - 1

	// Read in an inintal number of bytes we need to get started.
	if n, err := io.ReadFull(input, buf[:end]); err != nil {
		output.Write(buf[:n])
		return
	}

	for {
		// Read in one byte from the input stream.
		if _, err := io.ReadFull(input, buf[end:]); err != nil {
			// Flush the reset of the bytes we have.
			output.Write(buf[:end])
			return
		}

		// If we have a match, replace the bytes.
		if bytes.Compare(buf, find) == 0 {
			output.Write(repl)

			// Read a new initial number of bytes.
			if n, err := io.ReadFull(input, buf[:end]); err != nil {
				output.Write(buf[:n])
				return
			}

			continue
		}

		// Write the front byte since it has been compared.
		output.WriteByte(buf[0])

		// Slice that front byte out.
		copy(buf, buf[1:])
	}
}

func BenchmarkAlgorithmOne(b *testing.B) {
	var outout bytes.Buffer
	in := assemebleInputStream()
	find := []byte("elvis")
	repl := []byte("Elvis")

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		output.Reset()
		algOne(in, find, repl, &output)
	}
}
