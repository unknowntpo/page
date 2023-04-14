package domain

import (
	"fmt"
	"time"

	"github.com/oklog/ulid/v2"
	"golang.org/x/exp/rand"
)

// getRedisPartitionKey generates given redis partition key by
// userID, listKey, and have format {<listKey>:<userID>}
// NOTE: This format can not be changed, or every key will be migrated to other Redis Node.
func getRedisPartitionKey(listKey ListKey, userID int64) string {
	return fmt.Sprintf("{%s:%d}", listKey, userID)
}

// GeneratePageKeyByListKeyUserID generates given redis key by
// userID, listKey, and have format page:{<listKey>:<userID>}:<ulid>
func GeneratePageKeyByListKeyUserID(listKey ListKey, userID int64, now time.Time) PageKey {
	entropy := ulid.Monotonic(rand.New(rand.NewSource(uint64(now.UnixNano()))), 0)
	ms := ulid.Timestamp(now)
	ulid, err := ulid.New(ms, entropy)
	if err != nil {
		panic(err)
	}
	return PageKey(fmt.Sprintf("page:%s:%s", getRedisPartitionKey(listKey, userID), ulid.String()))
}

// GenerateListKeyByUserID generates pageList (sorted set) redis key by
// listKey, userID, and have format pageList:{<listKey>:<userID>}
func GenerateListKeyByUserID(listKey ListKey, userID int64) ListKey {
	return ListKey(fmt.Sprintf("pageList:%s", getRedisPartitionKey(listKey, userID)))
}

// GenerateListKeyByUserID generates pageList (sorted set) redis key by
// listKey, userID, and have format pageList:{<listKey>:<userID>}
func GenerateListMetaKeyByUserID(listKey ListKey, userID int64) PageMetaKey {
	return PageMetaKey(fmt.Sprintf("listMeta:%s", getRedisPartitionKey(listKey, userID)))
}
