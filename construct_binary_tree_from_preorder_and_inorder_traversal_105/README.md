# 105. Construct Binary Tree from Preorder and Inorder Traversal

https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

## Difficulty:

Medium

## Description

Given preorder and inorder traversal of a tree, construct the binary tree.

Note:
- You may assume that duplicates do not exist in the tree.

For example, given
```
preorder = [3,9,20,15,7]
inorder = [9,3,15,20,7]
Return the following binary tree:

    3
   / \
  9  20
    /  \
   15   7
```

## Solution

This is a tough problem and takes close attention to understand. 

Key is that preorder starts with root, then left, then right.

The inorder for some value splits the tree into left and right subtrees.

This image really helps to visualize the how the algorithm comes together:  
![algo](https://leetcode.com/uploads/files/1486248260436-screenshot-2017-02-04-17.44.08.png) 
