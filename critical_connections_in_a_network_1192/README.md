# 1192. Critical Connections in a Network

https://leetcode.com/problems/critical-connections-in-a-network/

# Description

There are n servers numbered from 0 to n-1 connected by undirected 
server-to-server connections forming a network where `connections[i] = [a, b]` 
represents a connection between servers a and b. Any server can reach any 
other server directly or indirectly through the network.

A critical connection is a connection that, if removed, 
will make some server unable to reach some other server.

Return all critical connections in the network in any order.

Example 1:
![graph](https://assets.leetcode.com/uploads/2019/09/03/1537_ex1_2.png)

```
Input: n = 4, connections = [[0,1],[1,2],[2,0],[1,3]]
Output: [[1,3]]
Explanation: [[3,1]] is also accepted.
```

Constraints:
- 1 <= n <= 10^5
- n-1 <= connections.length <= 10^5
- connections[i][0] != connections[i][1]
- There are no repeated connections.
