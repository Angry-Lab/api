package helpers

import (
	"io"
	"log"
)

func Close(c io.Closer) {
	err := c.Close()
	if err != nil {
		log.Printf("closer returns error: %s", err)
	}
}
