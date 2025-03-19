package generate

import (
	"errors"
	"strconv"
	"sync"
	"time"
)

type UniqueIDGenerator struct {
	mu       sync.Mutex
	epoch    time.Time
	nodeID   int64
	sequence int64
	lastTime int64
}

type ID int64

const (
	Epoch         int64 = 1735664400000 // 2025-01-01 00:00:00 UTC
	TimestampBits       = 39
	CustomBits          = 1
	NodeBits            = 10
	SequenceBits        = 13

	NodeMask     = (1 << NodeBits) - 1
	SequenceMask = (1 << SequenceBits) - 1

	NodeShift      = SequenceBits
	CustomShift    = NodeBits + SequenceBits
	TimestampShift = CustomBits + NodeBits + SequenceBits

	CustomValue = 1
)

func NewUniqueIDGenerator(nodeID int64) (*UniqueIDGenerator, error) {
	if nodeID < 0 || nodeID > NodeMask {
		return nil, errors.New("NodeID must be between 0 and " + strconv.FormatInt(NodeMask, 10))
	}

	curTime := time.Now()
	return &UniqueIDGenerator{
		epoch:  curTime.Add(time.Unix(Epoch/1000, (Epoch%1000)*1000000).Sub(curTime)),
		nodeID: nodeID,
	}, nil
}

func (n *UniqueIDGenerator) GenerateID() ID {
	n.mu.Lock()
	defer n.mu.Unlock()

	now := time.Since(n.epoch).Milliseconds()

	if now == n.lastTime {
		n.sequence = (n.sequence + 1) & SequenceMask
		if n.sequence == 0 {
			for now <= n.lastTime {
				now = time.Since(n.epoch).Milliseconds()
			}
		}
	} else {
		n.sequence = 0
	}

	n.lastTime = now

	id := ID((now << TimestampShift) | (CustomValue << CustomShift) | (n.nodeID << NodeShift) | n.sequence)

	return id
}

// Int64 returns an int64 of the Generator ID
func (f ID) Int64() int64 {
	return int64(f)
}

// ParseInt64 converts an int64 into a Generator ID
func ParseInt64(id int64) ID {
	return ID(id)
}

// Base2 returns a string base2 of the Generator ID
func (f ID) Base2() string {
	return strconv.FormatInt(int64(f), 2)
}

// ParseBase2 converts a Base2 string into a Generator ID
func ParseBase2(id string) (ID, error) {
	i, err := strconv.ParseInt(id, 2, 64)
	return ID(i), err
}

// String returns a string of the Generator ID
func (f ID) String() string {
	return strconv.FormatInt(int64(f), 10)
}

// ParseString converts a string into a Generator ID
func ParseString(id string) (ID, error) {
	i, err := strconv.ParseInt(id, 10, 64)
	return ID(i), err

}
