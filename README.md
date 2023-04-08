# page: Dcard Interview 2023 Assignment

[![Test](https://github.com/unknowntpo/page/actions/workflows/main.yml/badge.svg?event=push)](https://github.com/unknowntpo/page/actions/workflows/main.yml)

## Getting Started

> TODO

## Choice of Database

:question: Why not PostgreSQL ?

## Choice of GRPC package

## Redis Storage Design

### Goals:
> 1. Make sure that every operation is 1 RTT (Round-trip time)
> 2. Make sure that related data structure is stored in

First one is achieved by using lua script, it's convenient because every lua script is ensured to be executed atomically, there's no way to interrupt this, so we don't use something like `MULTI`, `EXEC`, `WATCH`, which is hard to use them within 1 RTT.

The Second one can be achieved by using Hash Tag 
see [Redis cluster specification](https://redis.io/docs/reference/cluster-spec/) for more information.

### Implementation Details
#### Data Structures

We have three redis data structure to store the list data:

**listMeta**:
- **Key**: `listMeta:{<listKey>:<userID>}`
- **Purpose**: store the metadata of a list, like `head`, `tail`

**pageList**
  - **Key**: `pageList:{<listKey>:<userID>}`
  - **Purpose**: store the score of `pageKey`, the score will be the expire time of `pageKey`

**RedisJSON Object**
  - **Key**: `page:{<listKey>:<userID>}`
  - **Purpose**: It stores the actual data, the TTL will be set to the object, so we don't need to delete expired data manually. 
The reason I choose RedisJSON because in linked-list, we need to update page.next to new pageKey,
if I use lua module like [cjson](https://github.com/mpx/lua-cjson), there will be very slow if content in a page is very large.
The reason is that `cjson` stores data in text format it means that , so if we need to change `.next` pointer frequently, it needs to decode and encode the text, which is costly.

#### API Design

#### NewList:

> - **Time Complexity**: O(1)
> - **RTT**: 1


`NewList` initializes the list with `listMeta`, `pageList`

#### SetPage:

> - **Time Complexity**: O(1)
> - **RTT**: 1

`SetPage` do the following things:
- Find the tail recorded in `listMeta.tail`, modify `.next` field of tail `RedisJSON` object point to itself.
- Add new `pageKey` with expire time to sorted set (`pageList`)
- Add new `RedisJSON` object

#### GetHead

> - **Time Complexity**:
>   - Best: O(1) if head haven't exired
>   - Worst: O(log(N)+M) with N being the number of elements in the sorted set and M the number of expired elements removed by the operation.
> - **RTT**: 1


  See [ZREMRANGEBYSCORE command](https://redis.io/commands/zremrangebyscore/) for more information.

#### GetPage

> - **Time Complexity**: O(1)
> - **RTT**: 1

### Drawback of this design
> - TTL can not be adjusted because the linked list will broken.

## TODO: ScyllaDB

Pros:
- Cloud native, high availability, don't need to worry about sharding things.
- Can use TTL with Time Window Compaction Strategy (TWCS) to Drop whole expired SSTable.

# Reference:
- [Time to Live (TTL) and Compaction¶](https://docs.scylladb.com/stable/kb/ttl-facts.html)

