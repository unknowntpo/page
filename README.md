# page: Dcard Interview 2023 Assignment

[![Test](https://github.com/unknowntpo/page/actions/workflows/main.yml/badge.svg?event=push)](https://github.com/unknowntpo/page/actions/workflows/main.yml)

## Getting Started

See help messages:

```
Usage:
  help          print this help message
  mock/gen      generate mock $(IFASE) implementation against interface inside internal/domain, e.g. make mock/gen IFASE=PageUsecase
  proto/gen     generate code from grpc proto
  redis/setup   set up development environment
  redis/flush   wipe out data in redis
  redis/down    delete redis container
  test          run unit tests
  build         build the server binary
  run/server    build and run the server binary
  run/client    build and run the client binary
```

Install required package for `connect-go`:
See [official website](https://connect.build/docs/go/getting-started#install-tools)
```
$ go install github.com/bufbuild/buf/cmd/buf@latest
$ go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
$ go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
$ go install github.com/bufbuild/connect-go/cmd/protoc-gen-connect-go@latest
```

Start redis-stack container:
```
$ make redis/setup
````

Build and run binary:

```
$ make run/server
```

Run example client binary:

```
$ make run/client
```

Run the tests:

```
$ make test
```

or the run specific tests:

```
$ make test VERBOSE=1 TESTPKG=./page/repo/redis FOCUS="SetPage.*related data.*"
```

## Choice of Database

:question: **Why not PostgreSQL ?**

PostgreSQL implements *Multi-version Concurrency Control (MVCC)* by *Snapshot Isolation (SI)*, and for the deleted row, `t_xmax` field will be marked, and when it comes to VACUUM, this deleted row (dead tuple) will be cleaned. If we delete data frequently, there will be a lot of dead tuples in heap page. This will cause `Index Scan` require more disk IO because the actual data is spreaded across multiple pages.

Although we can use some tricks like:
1. Using online clustering tool e.g. [`pg_repack`](https://reorg.github.io/pg_repack/) to reorganize table
2. Put data with similar expired time under same table, and Drop the table if all rows are expired.

But I think this will increase complexity.

See [The Internals of PostgreSQL - Chapter 5: Concurrency Control ](https://www.interdb.jp/pg/pgsql05.html#_5.5.) for more information about Snapshot Isolation in PostgreSQL.

> 👍 **TimescaleDB** address this problem by directly drop chunks (PostgreSQL Table) to avoid vacuuming problem. See this talk for more information [Rearchitecting a SQL Database for Time-Series Data | TimescaleDB](https://youtu.be/eQKbbCg0NqE?t=867)

:question: **Why I choose Redis ?**

- It can delete expired key automatically
- Insertion is faster than PostgreSQL

## Choice of gRPC package

- Simplicity: To be honest, it's because I can't generate correct gRPC .go file with Google gRPC package. The command is too hard to use
- Ecosystem: Can use Golang's `net/http` ecosystem, and for testing, we can use `net/http/httptest` package

## Redis Storage Design

### Goals:
> 1. Make sure that every operation is 1 RTT (Round-trip time)
> 2. Make sure that related data structure is stored in same Node

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


`NewList` initializes the list with `listMeta`, `pageList`.

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

`GetHead` gets the head for head field of HashMap `listMeta`, but if head inside it is expired,
we use ZREMRANGEBYSCORE to remove expired pageKey entry in sorted set `pageList`, then set the new head to `listMeta.head`.

  See [ZREMRANGEBYSCORE command](https://redis.io/commands/zremrangebyscore/) for more information.

#### GetPage

> - **Time Complexity**: O(1)
> - **RTT**: 1

`GetPage` gets the page by `pageKey`.

### Drawback of this design
> - TTL can not be adjusted because the linked list will broken.
> - Redis is in-memory database, and we don't have a way to swap unused object into disk file, this will require lots of Redis node, which will be very expensive

## TODO:
### Switch to ScyllaDB

Pros:
- Cloud native, high availability, don't need to worry about sharding things.
- Can use TTL with [Time Window Compaction Strategy (TWCS)](https://docs.scylladb.com/stable/kb/ttl-facts.html) to automatically whole expired SSTable.
- Can leverage disk storage, in contract, Redis can only use memory to store data.
- Can use CQL to query data, compared to Redis, this approach is more natural for me.

### Use `Packer`, `Terraform` to deploy services to cloud

### K6 Pressure Test

### Use structured logging package like [zerolog](https://github.com/rs/zerolog)
