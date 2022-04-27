package hutex

import (
    "sync"
)

type Hutex struct {
    Lock     sync.RWMutex
    Children []*Hutex
}
