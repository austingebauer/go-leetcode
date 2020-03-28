# 141. Linked List Cycle

https://leetcode.com/problems/linked-list-cycle/

## Difficulty:

Easy

## Description

Given a linked list, determine if it has a cycle in it.

To represent a cycle in the given linked list, 
we use an integer pos which represents the position (0-indexed) 
in the linked list where tail connects to. If pos is -1, then 
there is no cycle in the linked list.

Example 1:
![example-1](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist.png)
```
Input: head = [3,2,0,-4], pos = 1
Output: true
Explanation: There is a cycle in the linked list, where tail connects to the second node.
```

Example 2:
![example-2](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test2.png)
```
Input: head = [1,2], pos = 0
Output: true
Explanation: There is a cycle in the linked list, where tail connects to the first node.
```

Example 3:
![example-3](https://assets.leetcode.com/uploads/2018/12/07/circularlinkedlist_test3.png)
```
Input: head = [1], pos = -1
Output: false
Explanation: There is no cycle in the linked list.
```

Follow up:
- Can you solve it using O(1) (i.e. constant) memory?
