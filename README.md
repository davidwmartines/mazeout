# mazeout
### Text-base maze puzzle game

The object is to guide Moze ☻ through a maze-like 2D world, to a final destination point ⚑.  Moze moves one point at a time in one of four possible directions (left, right, up, down).  You provide guidance by entering the direction to move, and Moze will move and reports the new location (in Row/Column coordinates) ...Or not, if your provided direction is not possible because of a wall or obstacle.

The maze-like structure is generated differently for each game (well, mostly), using a cellular-automation algorithm (Conway's Life) with a psuedo-random seed pattern and psuedo-randomly selected number of generations. The activated cells from the output of the life sequence are used to create the obstacles (i.e. the inner "walls")  on the maze.  The idea is that maze should be different each time the game is played, yet still be in a somewhat playable or useful configuration and not consisting simply of totally scattered, psuedo-randomly placed walls.

Also, you can ask for **help** a limited number of times during the game causing a graphical representation of the maze and Moze's current location to be displayed.

### Example of game play:
```
Help me through the maze!
The maze is a 0-indexed grid of 10 rows by 30 columns, with 123 wall points.
I am starting at row 0, column 1, and I need to get to row 9, column 28.
Tell me which way to go - Left (l), Right (r), Up (u), or Down (d).
Hint: You can also ask for help (h) to view the maze, up to 10 times.

I am at row 0, column 1.
which way? h
 █☻████████████████████████████
 █               ██           █
 █             ██ █    ██     █
 █      ██   ██   █   █  █    █
 █   █   ██      █    ███     █
 █  ██  ██      █     █ █     █
 █ ██   █    ███  ██  ██      █
 █  █████    █    ██          █
 █    █           ██          █
 ████████████████████████████⚑█
You have used 1 of 10 available helps.
I am at row 0, column 1.
which way? d
Moving down
I am at row 1, column 1.
which way? d
Moving down
I am at row 2, column 1.
which way? d
Moving down
I am at row 3, column 1.
which way? r
Moving right
I am at row 3, column 2.
which way? r
Moving right
I am at row 3, column 3.
which way? r
Moving right
I am at row 3, column 4.
which way? r
Moving right
I am at row 3, column 5.
which way? h
 █ ████████████████████████████
 █               ██           █
 █             ██ █    ██     █
 █    ☻ ██   ██   █   █  █    █
 █   █   ██      █    ███     █
 █  ██  ██      █     █ █     █
 █ ██   █    ███  ██  ██      █
 █  █████    █    ██          █
 █    █           ██          █
 ████████████████████████████⚑█
You have used 2 of 10 available helps.
I am at row 3, column 5.
which way? u
Moving up
I am at row 2, column 5.
which way? r
Moving right
I am at row 2, column 6.
which way? r
Moving right
I am at row 2, column 7.
which way? r
Moving right
I am at row 2, column 8.
which way? d
Can't go down into wall!
which way? l
Moving left
I am at row 2, column 7.
which way? d
Can't go down into wall!
which way? r
Moving right
I am at row 2, column 8.
which way? r
Moving right
I am at row 2, column 9.
which way? d
Moving down
I am at row 3, column 9.
which way? d
Can't go down into wall!
which way? d
Can't go down into wall!
which way? h
 █ ████████████████████████████
 █               ██           █
 █             ██ █    ██     █
 █      ██☻  ██   █   █  █    █
 █   █   ██      █    ███     █
 █  ██  ██      █     █ █     █
 █ ██   █    ███  ██  ██      █
 █  █████    █    ██          █
 █    █           ██          █
 ████████████████████████████⚑█
You have used 3 of 10 available helps.
I am at row 3, column 9.
which way? r
Moving right
I am at row 3, column 10.
which way? d
Moving down
I am at row 4, column 10.
which way? d
Moving down
I am at row 5, column 10.
which way? d
Moving down
I am at row 6, column 10.
which way? d
Moving down
I am at row 7, column 10.
which way? d
Moving down
I am at row 8, column 10.
which way? d
Can't go down into wall!
which way? h
 █ ████████████████████████████
 █               ██           █
 █             ██ █    ██     █
 █      ██   ██   █   █  █    █
 █   █   ██      █    ███     █
 █  ██  ██      █     █ █     █
 █ ██   █    ███  ██  ██      █
 █  █████    █    ██          █
 █    █    ☻      ██          █
 ████████████████████████████⚑█
You have used 4 of 10 available helps.
I am at row 8, column 10.
which way? r
Moving right
I am at row 8, column 11.
which way? r
Moving right
I am at row 8, column 12.
which way? r
Moving right
I am at row 8, column 13.
which way? r
Moving right
I am at row 8, column 14.
which way? r
Moving right
I am at row 8, column 15.
which way? r
Moving right
I am at row 8, column 16.
which way? r
Can't go right into wall!
which way? u
Moving up
I am at row 7, column 16.
which way? u
Moving up
I am at row 6, column 16.
which way? u
Moving up
I am at row 5, column 16.
which way? u
Can't go up into wall!
which way? h
 █ ████████████████████████████
 █               ██           █
 █             ██ █    ██     █
 █      ██   ██   █   █  █    █
 █   █   ██      █    ███     █
 █  ██  ██      █☻    █ █     █
 █ ██   █    ███  ██  ██      █
 █  █████    █    ██          █
 █    █           ██          █
 ████████████████████████████⚑█
You have used 5 of 10 available helps.
I am at row 5, column 16.
which way? r
Moving right
I am at row 5, column 17.
which way? r
Moving right
I am at row 5, column 18.
which way? r
Moving right
I am at row 5, column 19.
which way? d
Moving down
I am at row 6, column 19.
which way? d
Moving down
I am at row 7, column 19.
which way? d
Moving down
I am at row 8, column 19.
which way? d
Can't go down into wall!
which way? r
Moving right
I am at row 8, column 20.
which way? r
Moving right
I am at row 8, column 21.
which way? r
Moving right
I am at row 8, column 22.
which way? r
Moving right
I am at row 8, column 23.
which way? r
Moving right
I am at row 8, column 24.
which way? r
Moving right
I am at row 8, column 25.
which way? r
Moving right
I am at row 8, column 26.
which way? r
Moving right
I am at row 8, column 27.
which way? r
Moving right
I am at row 8, column 28.
which way? r
Can't go right into wall!
which way? h
 █ ████████████████████████████
 █               ██           █
 █             ██ █    ██     █
 █      ██   ██   █   █  █    █
 █   █   ██      █    ███     █
 █  ██  ██      █     █ █     █
 █ ██   █    ███  ██  ██      █
 █  █████    █    ██          █
 █    █           ██         ☻█
 ████████████████████████████⚑█
You have used 6 of 10 available helps.
I am at row 8, column 28.
which way? d
Moving down
Yay! I made it out. Thank you!
```