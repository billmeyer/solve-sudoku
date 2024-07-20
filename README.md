# Puzzle 8: You Won't Want to Play Sudoku Again

Program to solve a Sudoku puzzle.

Adopted original program from **Programming For The Puzzled** by Prof. Srini Devadas rewritten in [**Go**](https://go.dev/).

See https://ocw.mit.edu/courses/6-s095-programming-for-the-puzzled-january-iap-2018/pages/puzzle-8-you-wont-want-to-play-sudoku-again/ for original Python code.

Run using:

```bash
go run cmd/solve-sudoku.go
```

Output:

```bash
Starting Puzzle
===============
[0 0 6] [7 8 0] [0 0 0]
[0 8 5] [0 0 4] [0 0 0]
[0 0 0] [0 0 0] [8 0 9]

[6 0 0] [3 0 2] [4 0 0]
[3 5 0] [0 1 0] [0 2 7]
[0 0 1] [6 0 7] [0 0 5]

[9 0 2] [0 0 0] [0 0 0]
[0 0 0] [4 0 0] [2 1 0]
[0 0 0] [0 7 1] [9 0 0]

Solution
===============
[1 9 6] [7 8 3] [5 4 2]
[7 8 5] [9 2 4] [1 3 6]
[2 4 3] [1 6 5] [8 7 9]

[6 7 9] [3 5 2] [4 8 1]
[3 5 4] [8 1 9] [6 2 7]
[8 2 1] [6 4 7] [3 9 5]

[9 1 2] [5 3 8] [7 6 4]
[5 3 7] [4 9 6] [2 1 8]
[4 6 8] [2 7 1] [9 5 3]
Backtracks = 1211
cpu time: 163.292µs
```
