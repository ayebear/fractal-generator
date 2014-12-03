Assignment 6
============

Due Dec 11
----------

Fractals, as discussed in class, are a form of mathematical recursion. You will now take the opportunity to write a program that will create an image of the Mandelbrot set, possibly the most famous fractal.

Your program should take 2 parameters: the resolution in pixels per unit, and the number of iterations that your algorithm performs on each point. Your program should generate a PNG file in the same directory, in black and white. Do not print in color. Do not use a format other than PNG.

Note that the Mandelbrot set can be calculated in parallel for increased performance. A truly outstanding solution would take advantage of this fact.

```
	a6 100 5
```
![alt tag](https://dl.dropboxusercontent.com/u/10528991/assignment6_images/m20-25.png)

```
	a6 20 25
```
![alt tag](https://dl.dropboxusercontent.com/u/10528991/assignment6_images/m100-5.png)

```
	a6 600 100
```
![alt tag](https://dl.dropboxusercontent.com/u/10528991/assignment6_images/m600-100.png)
