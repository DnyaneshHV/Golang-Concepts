package main

import "fmt"

type Image interface {

   ColorModel() color.Model  // whet is the color mode of the image

  Bounds() Rectangle  // what will the height and width of the image

  At(x, y int) color.Color  // it asks like position of the color tile.

}