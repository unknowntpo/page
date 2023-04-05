# page
Dcard Interview Repo

## Redis Storage Design

We have three redis data structure to store the list data:

- listMeta `listMeta:<listKey>:<userID>` to store the metadata of a list, like `head`, `tail`, `nextCandidate`   
- pageList `pageList:<listKey>:<userID>` to store the score of `pageKey`, the score will be the expire time of `pageKey`
- redis String with key = `pageKey`, and value is actual data, the TTL will be set, so we don't need to delete expired data manually. 

### NewList:
- Initialize `listMeta:<listKey>:<userID>`, with `listMeta:<listKey>:<userID>.head == ""`, `pageMeta.tail == ""`, `nextCandidate = <pageKey>` 

```
127.0.0.1:6379> hgetall testList-meta:33
1) "head"
2) ""
3) "tail"
4) ""
5) "nextCandidate"
6) ""
127.0.0.1:6379>
```

### SetPage:

SetPage need to do this in a single lua script, and given page should content the `.next` field, which have the pagKey of next candidate page after setting this page.

- Append it the sorted set `pageList:<listKey>:<userID>`
- Get candidate pageKey from `listMeta:<listKey>:<userID>.nextCandidate`, e.g. `page:43147719-0af6-4701-8c7b-0f63d95677d1`
- Call `SET <pageKey>` to set data with 1 Day TTL, e.g. `SET page:43147719-0af6-4701-8c7b-0f63d95677d1 <actual-data> EX 86400`
- Concate the linked-list by modifying `tail`, of `listMeta:<listKey>:<userID>`, and `nextCandidate` field should be set to the one inside given page's next field


NOTE: SetPage will also set `listMeta:<listKey>:<userID>.nextCandidate` to next pageKey and the `next` field of current page.

So If I have a page like:
```
{
	"content": [...],
    "next": "page:64c60df7-45a3-4ad8-9ee5-7a4848ff016e"
}
```

listMeta `listMeta:<listKey>:<userID>.nextCandidate` have:

```

```





```
zrange testList:33 0 +inf byscore
1) "page:43147719-0af6-4701-8c7b-0f63d95677d1"
```

### GetHead:

We can get head directly from `<listKey>-meta:<userID>`, but head may not exist because of expiration.
So after obtain the pageKey from `listMeta:<listKey>:<userID>`, we also need to use `EXISTS <head>` to do the re-check,
if it doesn't exist, then we need to get the oldest pageKey from `pageList:<listKey>:<userID>` by 

```
ZREMRANGEBYSCORE `pageList:<listKey>:<userID>`, "0", <current_timestamp-TTL>
```

Then get the real head by calling
```
ZRANGE <listKey>-SS:<userID> 0 +inf BYSCORE LIMIT 0 1
```

finally, set it back to `<listKey>-meta:<userID>.head`, and return the head.

### GetPage:
Call `GET <pageKey>`, if it doesn't exist, it can have two scenarios
- Page expired
- Page doesn't exist
Since we can't distinguish between these scenarios, so just return `NotFound` error.

Time complexity: `O(1)`

## TODO: ScyllaDB

Pros:
- Cloud native, high availability, don't need to worry about sharding things.
- Can use TTL with Time Window Compaction Strategy (TWCS) to Drop whole expired SSTable.

# Reference:
- [Time to Live (TTL) and Compaction¶](https://docs.scylladb.com/stable/kb/ttl-facts.html)

