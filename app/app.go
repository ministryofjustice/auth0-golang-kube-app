package app

import (
	"encoding/gob"
	"github.com/gorilla/sessions"
	"os"
	"math"
)

var (
	Store *sessions.FilesystemStore
)

func Init() error {
	Store = sessions.NewFilesystemStore(os.TempDir(), []byte("something-very-secret"))
	Store.MaxLength(math.MaxInt64)
	gob.Register(map[string]interface{}{})
	return nil
}
