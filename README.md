# Overlapping Rectangle Problem in Go

This is reimplementation of previous project https://github.com/miqdude/rectangleOverlap.

## How to use the program
- Use command go run main.go -file <your_input_JSON_file>
- Make sure your input JSON format is as following
```
{
    "rects": [
        {"x": 100,"y": 100,"w": 250,"h": 80},
        .
        .
        .
        .
    ]
}
```
- See the input example in input.json

## Expected Output
```
Input:
         0 : Rectangle at ( 100 , 100 ), w= 450 , h= 280
         1 : Rectangle at ( 120 , 200 ), w= 490 , h= 550
         2 : Rectangle at ( 140 , 160 ), w= 530 , h= 420
         3 : Rectangle at ( 160 , 140 ), w= 670 , h= 470
Intersections:
        0: Between 1 and 3 at ( 140, 160 ), w= 210, h= 20.
        1: Between 2 and 3 at ( 140, 200 ), w= 230, h= 60.
        2: Between 1 and 4 at ( 160, 140 ), w= 190, h= 40.
        3: Between 2 and 4 at ( 160, 200 ), w= 210, h= 130.
        4: Between 1,3, and 4 at ( 160, 160 ), w= 190, h= 20.
        5: Between 2,3, and 4 at ( 160, 200 ), w= 210, h= 60.
        6: Between 3 and 4 at ( 160, 160 ), w= 230, h= 100.
```
