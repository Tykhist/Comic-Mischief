package main

import "fmt"

func main() {
  // Prints comic book info based on certain parameters
  var publisher, writer, artist, title string
  var year, pageNumber int
  var grade float32

  title = "Mr. GoToSleep"
  writer = "Tracey Hatchet"
  artist = "Jewel Tampson"
  publisher = "DizzyBooks Publishing Inc."
  year = 1997
  pageNumber = 14
  grade = 6.5

  fmt.Println(title, "written by", writer, "drawn by", artist, "published in", year, "by", publisher, "with", pageNumber, "pages, with a condition grade of", grade)

  fmt.Println()

  title = "Epic Vol. 1"
  writer = "Ryan N. Shawn"
  artist = "Phoebe Paperclips"
  publisher = "DizzyBooks Publishing Inc."
  year = 2013
  pageNumber = 160
  grade = 9.0

  fmt.Println(title, "written by", writer, "drawn by", artist, "published in", year, "by", publisher, "with", pageNumber, "pages, with a condition grade of", grade)
}

/*definition
My goals for the project are to:
1. Add the variables into the function
2. Setting the variables to be user input
*/
