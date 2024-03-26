import random
import os
import time
from queue import LifoQueue
import sys


sys.setrecursionlimit(5000000)

n = 111 # maze size
stack = LifoQueue()


def create_grid(n):
    grid = [[0 for _ in range(n)] for _ in range(n)]
    for row in range(len(grid)):
        for i in range(len(grid[row])):
            if row % 2 == 0:
                grid[row][i] = '#'
            elif i % 2 == 0:
                grid[row][i] = '#'
    return grid


def print_grid(grid):
    # os.system('cls')
    for row in grid:
        for i in row:
            print(i, end=' ')
        print()


def save_maze(grid):
    with open('test.txt', 'w') as out:
        for row in grid:
            for i in row:
                out.write(i)
            out.write('\n')

def get_neighbours(curr, grid):
    global n
    y_curr = curr[0]
    x_curr = curr[1]
    unvisited = []

    if not y_curr - 2 < 0 and grid[y_curr - 2][x_curr] == 0:
        unvisited.append([y_curr - 2, x_curr])

    if not y_curr + 2 == n and grid[y_curr + 2][x_curr] == 0:
        unvisited.append([y_curr + 2, x_curr])

    if not x_curr - 2 < 0 and grid[y_curr][x_curr - 2] == 0:
        unvisited.append([y_curr, x_curr - 2])

    if not x_curr + 2 == n and grid[y_curr][x_curr + 2] == 0:
        unvisited.append([y_curr, x_curr + 2])

    if len(unvisited) > 0:
        return unvisited[random.randint(0, len(unvisited) - 1)]
    else:
        return


def run(curr, grid):
    global stack

    y_curr = curr[0]
    x_curr = curr[1]
    grid[y_curr][x_curr] = ' '
    # print_grid(grid)
    # time.sleep(0.05)

    next = get_neighbours(curr, grid)


    if next:

        stack.put([y_curr, x_curr])

        y_next = next[0]
        x_next = next[1]

        if y_curr == y_next:
            if x_next > x_curr:
                grid[y_curr][x_curr + 1] = ' '
            else:
                grid[y_curr][x_curr - 1] = ' '
        else:
            if y_next > y_curr:
                grid[y_curr + 1][x_curr] = ' '
            else:
                grid[y_curr - 1][x_curr] = ' '

        run(next, grid)
    elif not stack.empty():
        run(stack.get(), grid)


def main():
    start = time.time()
    global n
    grid = create_grid(n)
    # print_grid(grid)

    run([1, 1], grid)


    save_maze(grid)

    elapsed = time.time() - start
    print_grid(grid)
    print(f"Program took {elapsed:.5f} seconds to complete.")

main()



