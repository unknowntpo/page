# page
Dcard Interview Repo

[![Test](https://github.com/unknowntpo/page/actions/workflows/main.yml/badge.svg?event=push)](https://github.com/unknowntpo/page/actions/workflows/main.yml)

## Redis Storage Design

### Goals:
1. Make sure that every operation is 1 RTT (Round-trip time)
2. Make sure that related data structure is stored in

First one is achieved by using lua script, it's convenient because every lua script is ensured to be executed atomicly, there's no way to interrupt this, so we don't use something like `MULTI`, `EXEC`, `WATCH`, which is hard to use them within 1 RTT. 

The Second one can be achieved by using Hash Tag 
see [Redis cluster specification](https://redis.io/docs/reference/cluster-spec/) for more information.

### Implementation Details
#### Data Structures

We have three redis data structure to store the list data:

**listMeta**:
  - Key: `listMeta:<listKey>:<userID>`
  - Purpose: store the metadata of a list, like `head`, `tail`

**pageList**
  - Key: `pageList:{<listKey>:<userID>}`
  - Purpose: store the score of `pageKey`, the score will be the expire time of `pageKey`

**RedisJSON Object**
  - Key = `page:{<listKey>:<userID>}`

It stores the actual data, the TTL will be set to the object, so we don't need to delete expired data manually. 
The reason I choose RedisJSON because in linked-list, we need to update page.next to new pageKey,
if I use lua module like [cjson](https://github.com/mpx/lua-cjson), there will be very slow if content in a page is very large.
The reason is that `cjson` stores data in text format it means that , so if we need to change `.next` pointer frequently, it needs to decode and encode the text, which is costly.




Design:
- Goals
- DataStructures
- Detail
- - NewList
- - Setpage
- - GetHead
- - GetPage

#### API Design

#### NewList:

Time Complexity: O(1)

NewList initialize the list with `listMeta`, `pageList`

#### SetPage:

Time Complexity: O(1)

SetPage need to do this in a single lua script, and given page should content the `.next` field, which have the pagKey of next candidate page after setting this page.

## TODO: ScyllaDB

Pros:
- Cloud native, high availability, don't need to worry about sharding things.
- Can use TTL with Time Window Compaction Strategy (TWCS) to Drop whole expired SSTable.

# Reference:
- [Time to Live (TTL) and Compaction¶](https://docs.scylladb.com/stable/kb/ttl-facts.html)

