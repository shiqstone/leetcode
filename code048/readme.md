48. Rotate Image
---
You are given an `n*n` matrix reprenting a Image.

Rotate the image by 90 degree (clockwise).

Note:
You have to rotate the image **in-place**, which mean you have to modify the input 2D matrix directly, DO NOT allocate another 2D matrix and do the rotate.

Example 1:
```
Given the input matrix = 
[  
    [1, 2, 3],
    [4, 5, 6],
    [7, 8, 9],
]

rotate the input matrix in-place such that it becames:
[
    [7, 4, 1],
    [8, 5, 2],
    [9, 6, 3],
]
```

Example 2:
```
input = [
    [5, 1, 9, 11],
    [2, 4, 8, 10],
    [13, 3, 6, 7],
    [15, 14, 12, 16],
]

Output = [
    [15, 13, 2, 5],
    [14, 3, 4, 1],
    [12, 6, 8, 9],
    [16, 12, 14, 15],
]
```