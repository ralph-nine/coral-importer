package importer

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"os"
	"sync"

	"github.com/mailru/easyjson"
)

// Importer manages importing documents into MongoDB after a validation process.
type Importer struct {
	docs       chan easyjson.Marshaler
	collection string
}

// Import will insert the document onto the importer to get processed.
func (i *Importer) Import(doc easyjson.Marshaler) error {
	if doc == nil {
		return errors.New("can't import nil")
	}

	i.docs <- doc

	return nil
}

// Start should be called as a part of a new goroutine to process importing
// tasks for new documents coming in.
func (i *Importer) Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	f, err := os.Create(fmt.Sprintf("%s.import.json", i.collection))
	if err != nil {
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer w.Flush()

	for {
		select {
		case <-ctx.Done():
			close(i.docs)
			return
		case doc := <-i.docs:
			if doc == nil {
				return
			}

			bytes, err := easyjson.Marshal(doc)
			if err != nil {
				return
			}

			w.Write(bytes)
			w.WriteString("\n")
		}
	}
}

func (i *Importer) Done() {
	close(i.docs)
}

// New returns a configured Importer.
func New(collection string) *Importer {
	return &Importer{
		collection: collection,
		docs:       make(chan easyjson.Marshaler, 100),
	}
}
