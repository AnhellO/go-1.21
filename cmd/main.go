package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"testing/slogtest"
	"time"

	"log/slog"
)

func main() {
	// TYPE INFERENCE
	value := 42
	fmt.Println(genericFunc(value)) // La inferencia de tipos detecta automáticamente el tipo de valor

	// LOOP VARIABLE CAPTURE (Use GOEXPERIMENT=loopvar flag)
	var out []*int
	for i := 0; i < 3; i++ {
		out = append(out, &i)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
	fmt.Println("Addresses:", out[0], out[1], out[2])

	for i := 1; i <= 10; i++ {
		go func() {
			fmt.Println(i)
		}()
	}

	time.Sleep(1 * time.Second)

	// PKG log/slog
	slog.Info("hello, world")
	slog.Info("hello, world", "user", os.Getenv("USER"))

	logger := slog.Default()
	logger.Info("hello, world", "user", os.Getenv("USER"))

	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))
	logger.Info("hello, world", "user", os.Getenv("USER"))

	logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	logger.Info("hello, world", "user", os.Getenv("USER"))

	// PKG testing/slogtest
	var buf bytes.Buffer
	h := slog.NewJSONHandler(&buf, nil)

	results := func() []map[string]any {
		var ms []map[string]any
		for _, line := range bytes.Split(buf.Bytes(), []byte{'\n'}) {
			if len(line) == 0 {
				continue
			}
			var m map[string]any
			if err := json.Unmarshal(line, &m); err != nil {
				panic(err) // In a real test, use t.Fatal.
			}
			ms = append(ms, m)
		}
		return ms
	}
	err := slogtest.TestHandler(h, results)
	if err != nil {
		log.Fatal(err)
	}
}

func genericFunc[T any](x T) T {
	// Implementación genérica de la función
	return x
}
