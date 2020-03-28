# 289. Game Of Life

https://leetcode.com/problems/game-of-life/

# Description

According to the Wikipedia's article: "The Game of Life, also known simply as Life, 
is a cellular automaton devised by the British mathematician John Horton Conway in 1970."

Given a board with m by n cells, each cell has an initial state live (1) or dead (0). 
Each cell interacts with its eight neighbors (horizontal, vertical, diagonal) using the 
following four rules (taken from the above Wikipedia article):

1. Any live cell with fewer than two live neighbors dies, as if caused by under-population.
2. Any live cell with two or three live neighbors lives on to the next generation.
3. Any live cell with more than three live neighbors dies, as if by over-population..
4. Any dead cell with exactly three live neighbors becomes a live cell, as if by reproduction.

Write a function to compute the next state (after one update) of the board given its 
current state. The next state is created by applying the above rules simultaneously to 
every cell in the current state, where births and deaths occur simultaneously.

Example:
```
Input: 
[
  [0,1,0],
  [0,0,1],
  [1,1,1],
  [0,0,0]
]
Output: 
[
  [0,0,0],
  [1,0,1],
  [0,1,1],
  [0,1,0]
]
```

Follow up:

Could you solve it in-place? 
Remember that the board needs to be updated at the same time: You cannot update 
some cells first and then use their updated values to update other cells.

In this question, we represent the board using a 2D array. In principle, 
the board is infinite, which would cause problems when the active area encroaches 
the border of the array. How would you address these problems?
