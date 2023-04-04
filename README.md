# page
Dcard Interview Repo

## Redis Storage Design

### NewList:
- Initialize `listMeta`, with `pageMeta.head == nil`, `pageMeta.tail == nil` 
- 

### GetPage:
Call `GET <pageKey>`, if it doesn't exist, it can have two scenarios
- Page expired
- Page doesn't exist
Since we can't distinguish between these scenarios, so just return `NotFound` error.

Time complexity: `O(1)`
### SetPage:

- Append it the sorted set `pageList:<userID>`
- Call `Set page:pageKey` to set data with 1 Day TTL
- Concate the linked-list by modifying `tail`, `nextCandidate` of `ListMeta:<userID>`

## TODO: ScyllaDB

Pros:
- Cloud native, high availability, don't need to worry about sharding things.
- 

