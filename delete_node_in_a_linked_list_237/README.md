# 237. Delete Node in a Linked List

https://leetcode.com/problems/delete-node-in-a-linked-list/

## Difficulty:

Easy

## Description

Write a function to delete a node (except the tail) in a singly linked list, 
given only access to that node.

Given linked list -- `head = [4,5,1,9]`, which looks like following:
![linked list](https://assets.leetcode.com/uploads/2018/12/28/237_example.png)

Example 1:
```
Input: head = [4,5,1,9], node = 5
Output: [4,1,9]
Explanation: You are given the second node with value 5, the linked list should 
become 4 -> 1 -> 9 after calling your function.
```

Example 2:
```
Input: head = [4,5,1,9], node = 1
Output: [4,5,9]
Explanation: You are given the third node with value 1, the linked list should 
become 4 -> 5 -> 9 after calling your function.
```
