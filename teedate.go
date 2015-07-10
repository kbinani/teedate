package main

import "os"
import flag "github.com/ogier/pflag"
import "time"
import strftime "github.com/jehiah/go-strftime"

func main() {
	flagAppend := flag.BoolP("append", "a", false, "append to file")
	timeFormat := flag.String("format", "%Y-%m-%d %H:%M:%S", "time format, compatible with strftime(3)")
	flag.Parse()

	var outStreams []*os.File = []*os.File{os.Stdout}
	for i := 0; i < len(flag.Args()); i++ {
		if *flagAppend {
			f, _ := os.OpenFile(flag.Args()[i], os.O_WRONLY | os.O_CREATE | os.O_APPEND, 0644)
			outStreams = append(outStreams, f)
		} else {
			f, _ := os.OpenFile(flag.Args()[i], os.O_WRONLY | os.O_CREATE | os.O_TRUNC, 0644)
			outStreams = append(outStreams, f)
		}
	}

	const kBufferSize = 1
	inputBuffer := make([]byte, kBufferSize)
	outputBuffer := make([]byte, kBufferSize)
	var outputBufferSize int = 0

	kSeparator := []byte{' '}
	const kLF byte = '\x0a'

	AppendBuffer := func(buffer []byte) {
		var remaining int = len(buffer)
		var offset int = 0
		for remaining > 0 {
			amount := remaining
			if outputBufferSize + amount > kBufferSize {
				amount = kBufferSize - outputBufferSize
			}
			for i := 0; i < amount; i++ {
				outputBuffer[outputBufferSize] = buffer[i + offset]
				outputBufferSize++
			}
			if outputBufferSize == kBufferSize {
				for _, stream := range outStreams {
					stream.Write(outputBuffer)
				}
				outputBufferSize = 0
			}
			remaining -= amount
			offset += amount
		}
	}

	PrintTime := func() {
		AppendBuffer([]byte(strftime.Format(*timeFormat, time.Now())))
		AppendBuffer(kSeparator)
	}

	var printTime bool = true

	for true {
		size, _ := os.Stdin.Read(inputBuffer)

		if size == 0 {
			break;
		}

		for i := 0; i < size; i++ {
			if printTime {
				PrintTime()
				printTime = false
			}

			if inputBuffer[i] == kLF {
				printTime = true
			}
			AppendBuffer(inputBuffer[i:i + 1])
		}
	}

	for _, stream := range outStreams {
		stream.Write(outputBuffer[0:outputBufferSize])
		stream.Close()
	}
}
