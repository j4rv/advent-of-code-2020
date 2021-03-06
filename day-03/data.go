package main

const openCell = '.'
const treeCell = '#'

const data = `..#...##...###.........#..#..#.
#.###........#..##.#......#...#
#.#.###..#.#..#.#............#.
.##............#......#...#.#..
..#..#.....##..##..##..........
...#...........###.#.##........
....#.#...#..#..##............#
....#....##...##..##........#..
.#..#..#....#...#..##.....##...
.#.###..#......####........##..
..#...###....#......#.....##.##
..#...#.......#......#..##....#
#...##....#.#..#.......#....#..
.#......#..#...........#....##.
.##.#......#.#.#.....##........
.....#.................#.#..#.#
....#..#........##......#..#.#.
..#...#..##.......#..##...#..#.
..#.......#.............#.#....
.#.................#.........#.
..#..#.#.#.#............##.#..#
.#.#.##.#.....#.....#..#......#
..#.#..#.#..........##........#
.........#...#.....#.#...#####.
##..#.....##.##........#...##..
.#.....#....##.#..#....##...##.
.##.....#.#....#.#.....#......#
.....#..#.##.....#.#....#.#..##
#......##..##....##...###..#...
.......#..#...........#......#.
#...#......#........#..#.......
##..#.....##.....#...#...#....#
.###..##..#.#........#..#.#....
#.#...#...#......##........#.#.
......#....#.#........##...#..#
.#.....#..#.#.....#......##....
.....#.....#.#.#....###.....#..
#.......##.#......#.#.#....###.
.......#..#..#...#.#.##........
.#......##..#.........###..#...
....#..##.......##.###...###...
.##............#..#.##.....#.##
..##.#.......##....#.......##.#
#..###............#.#...#...#.#
...##.#.#.#..#.##........#.#...
.#.....#...##.#..###..##.##...#
..............#.#.#.........#..
.....#...........#.#...#....#..
.....#...##.##.#....#.###..#...
#..###.........#......#.#.#....
.....#..#...##...###.#....#....
#..........#.#.#....#..#......#
###...................#......#.
........####......#.#..........
.......#.....#...#.......#...#.
.....#.....................#...
...#.#...#...#...........#.....
..#.........#...#....###..#....
.....#.#..##......##........##.
..#.............#............#.
.#....##.......#..#............
.#............#.#..#.##....#...
.####...##.#....#.....#...#....
##..#....#.#.#...........#..#.#
...#..#...........#...#..#....#
.....##.....#..#...#.........#.
...##........#....##........#..
.##.#...#...#..#.....#....##.#.
#.#...#.#.#.#.#..#....#....#...
#..#.#...#..#........#....#.#..
....#.#.....#......##...#....#.
.###.##...#....##.#...###..#..#
###..#...##..#......#.........#
..#.#......##.......#.....#...#
..#...#........#.........#.#...
#....#..#.........###.#......##
...#..#....#...#.......##.#.#..
....#.......#....###...##.#....
..#.....#.#.....###..#####....#
##......#....#.....##..#..#...#
#...........#..#..#....#....#..
.#...#.##.#.#.#....#......#..#.
.......#.#....#....#...#.#..#..
..#.#..#.##..##...##..#..#.....
...##.##.................#.#...
.....#...#......##.#....#.....#
......#..##.#..#.#.........#...
.............##.#......#.......
..#.#.....#...#.#.....#..#.....
.........#..#.#......#..#......
#..#.#.##..........#.##......##
......#.......#.....#..#.#...#.
.#....#....#.#.....#.......#...
#..#..##..................#....
............#...........##.#...
####...#..##.#....#.##..#......
#...#...#.....#.#...#.#........
.......##.........#.....#.....#
.....#...........##......#.####
.##....#.#.##......###.#.##....
........#.####.......#.#...#...
.#.#...##.#.#.#.........##.....
....#............###.##....#...
...##........##.#...#....#..###
..#.#.........#....##.#........
..#..##..#...##..#.##...#.....#
.#......#..#..#.........#......
..#........##.#......#.....#...
.##.......#....#.#......#......
#........#....#.####...#.#.....
##......#.............#....###.
..#....####.#.#.#.#...##......#
#.#.#.....#...#.......#...##...
........#...#....#..#......#.#.
#..#...#.#.##.....#.#....#....#
#...#....#......#.........##.##
..##.#..##............#........
#.........####.........#.......
#.##.........#..##....#.#.#.#..
.###......##..#.#.....#.#...##.
...#.........#.#...##.##....#..
#..#......#....##..#.#...#...##
...#.......#.#.#.....#..##...#.
....................###........
#...........###......#.#...##.#
.................#.#...###....#
...#..###..#.##...#..#....##...
###..#..#.#...#.....#.#.......#
.........#..#.#.....#.........#
.##..#.........#.#.....##......
.....#..........#.#.##....#....
........#.##.....#...##...#....
#.#.#...#......##....#.###.....
.#.##.....##.....#....#.##.#.##
.#...#.....##.#.##....#.....#.#
...#.....#........#............
##...####..#....#..##...#......
#.......#...#.#...#........#..#
......#.....#....#..#..#.#.....
..............##.....#.##....##
.#..##.........###..#..........
......#......#............#..#.
#.....###...###..........#.....
...###...........#....##..#....
.....................#...#.##..
###....#.#....#...#....#.#..##.
..#.............#.#....#..#.#..
.......#..###....#...#...#...#.
.##..#....##..............##...
...#...#..#..#.##.#....##...#..
#..#....##......#....###..####.
.##...#.#....#..#..#....#....#.
.#.##..#..#.........#.#......##
#..#.................#.....#...
..#............#........#...#..
##.##.......#.#....#..#....##.#
..##...#.#.....#......#........
......#.##.........##...#......
......##.#......#.##....##.#..#
.#.#......####.#.#.#.#..#......
..#.#....#...###....#.#...###..
.#.#...#....##..###.#..#.......
..#.....##..#............#..#.#
.#...#....#.....#....#..#.#.#..
..#....#..#......##...##......#
....#.......#.##.#.........#..#
#............#...##.....#..##..
......#..#..........#.#..#.....
...........#.#..#...##.#...#..#
.........###..#......##.###.#..
.....#....#......#...#....##...
..#.......#..#.#.#......#......
.#....#.....#.#.#.##...#....##.
....#.##.##.......#..#.....#...
.#.....#......#.......#..#.....
....###.....##.....##..#.#...#.
#.......#.#....#.#.#....#......
#...#..#.#......#...##.#.......
....##.##....#..#.......#..#.#.
.#.##.#.#..#.....#.#.......##..
..#..#..#..#.###...............
#.#......##....##.#.#.....#.#..
..##...#.........#.#..#.##..#..
.........####...#.....##.#.....
..#...................#.###....
..#.....#..#....#..#...........
.....###.#.........#.#.........
#..#..#........#..#......#..#..
###..##...#.......#........#.#.
.#.#.#.###.#............##.....
#..............#......#....#.##
.#...#..###...###.#..#.#.......
.###....##.#.......#.#.....###.
.##.....#.#....................
#..#.....#.....#...#.....#..##.
#.#....##..#......##..#...##..#
...........#....#..#.##.##....#
......#.......##....#.#....#.#.
###..#.#..........#.......#.#..
..#.#..##....##............##..
..#.....#..#....###............
.#...#...##...#..#..#..#.#....#
...#....#........#.............
#.#......#.#.....##..........#.
....#..#...............#...##..
........#..#....#..#..#..##....
....#......#.#.#...#.......#.##
.#.....#.#.#........###....#...
.#..#.......#...........#...#..
#.#.#####..#......#...#.#.###..
...##...#.#.....#..#...#...#...
..#....#.....#..#....#.#.....#.
....#.......#.....#........###.
..##..........#...###.......#..
#.#.##..#........##...#.#......
....##...#......#..........#.#.
#.......#..#.##.............#..
......#..........#...#....#...#
#.#.....#.##.#.#.............##
#...#.........##.##......#.##..
.........##.....#....#...##..#.
#.#...##.#...#.....##...#..#..#
......##.#.....#.#.....#.##....
....#.............#...#.......#
.#......##...#.#...#.##........
...#..#..#...........#..#..#.#.
.#...#...........#.#.##....#...
..#...#...#.#..#....#..#..#....
..###..##..#..#.........#.#..#.
....#.##.#...#.......#...#.....
.#.#.................#.......#.
.#..#.....#.##...#.#.....#...#.
.#.......#...#....#.......#....
###....##....#..#...#.#..##....
.........#.#.#.#...###.......##
....##.......#......#......##..
......##.###.#..#...#.#.#.#....
.#.###.#.#......#.#.......##...
.#.....##..#.#.....#...#.##....
.#..##.#.....#........###...#.#
.......#.#...........#........#
..........#...##..##.......#.#.
...#..#..#...#....#.#......#.#.
....#...#..#....#....#.#.##....
...#.#...##...#...##..#........
..#........#...##.#...##.#.#.#.
...##.#..#.......###.#.#.#.....
..##......#.#.#.........###.#..
.......#.#...##...#.#.......#.#
.#....##..#..#....#..#...#.....
.#....#.#.......#..#..##.#....#
#.#..##..#.#............#....#.
##..#....#.##.#....#...#..##...
.###.#.#..#...##........###.##.
...........#..#...#..#.#..###..
.#.#.#...#....#...##.##........
.....###.........#......#####..
#.#.#.....#....#..#...#.#.#.#..
..##.....#..#..#.#.#...#....#.#
......#.##..##..#.#.#.......#..
...#.#..###.........#........#.
......#.##..####...#...#..#...#
#.......#.........#....#....###
#......#...#........#.##....###
.#.#..............#...#...###.#
.#....####...#..##.###.....#...
.......#......#..#...#..##.#...
...........#.......#...........
..............#...#.#.#.#...##.
.###.....##.#.....#..........#.
##.##......#....##..#.....###.#
#.......#...##...#....##...#...
##.#.##...#....#....#....#.....
.....####...........#.#......#.
......#...#....#..#......#.....
...#...##..........#.......##.#
.#....#..........#.####........
...##...#..#...##........##..#.
.........##....#...##..#.##.#..
##.#.....#.......#.....#.......
#..#....#.##.#........#........
#.#...#...##........#.#.....###
....#...................#.#.#..
.......#..#.#...#....#.##.#....
....##...###.#.#.##...#...#....
.#....#....#...##.#......#...#.
............##..#.#.#........#.
...#....#.....#......#........#
...#.#.....#.##.....#....#...#.
.....#..##.......#.##.......#.#
........##................#....
....#..###...##.#..#...#......#
.#.#.......#.......#....##.#..#
..#........#............#......
..##.......#..#..#....#....#..#
#...###.......#.##...#.........
.....#...#...#..##..#....#..#..
.##.#..#...##.........###.#....
..#.#..#...#...####.#...#.#.#.#
#....#..###.....#......#.##..##
##......#...##...###......#.##.
...........#.....##...#...#...#
..#..#.#.....#..#.....###...#..
.............#..........###...#
....##............#....###.##.#
..##.#..##.....#.#.........#.#.
....#.#...........####.........
.##.###.##.#.#......#.##.#.#...
.....##.........#..#.......#...
...........#.........#....###..
...#.#..#..........#.....#..#..
.#..###.......##........#.#....
.#...###.....#..#.#..#...#.##..
##...###.#.#....#......#...#..#
....#.......#..#..##..#.#......
#.#......#.##..#......#..#....#
....#..#..#.....#.#......#..#..
..#...###......##.............#
..#....####...##.#...##.#......
.....#.......###...............
.......#...#.#.......#.#.##.###
.#.#...#.....#...##.........#..
..#..........#..#.........##...`
