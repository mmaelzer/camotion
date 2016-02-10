# camotion
[![godoc reference](https://godoc.org/github.com/mmaelzer/camotion?status.png)](https://godoc.org/github.com/mmaelzer/camotion)
[![go report card](https://goreportcard.com/badge/github.com/mmaelzer/camotion)](https://goreportcard.com/badge/github.com/mmaelzer/camotion)

A motion detection library written in Go.

```
import "github.com/mmaelzer/camotion"
```

## Usage

#### func  Blended

```go
func Blended(img1, img2 image.Image, threshold int) image.Image
```
Blended eturns a black and white image where white pixels denote a change
between the two images that is greater than the given threshold

#### func  Motion

```go
func Motion(img1, img2 image.Image, minChange, threshold int) bool
```
Motion will determine if motion occurs between two images given a minChange
value which is a 1-100 value representing percentage of image changed and a
threshold value which sets a below this value is noise. A reasonable threshold
default for jpegs (and potentially other formats) is 2500.

#### func  MotionWithStep

```go
func MotionWithStep(img1, img2 image.Image, minChange, threshold, step int) bool
```
MotionWithStep will call Motion() with a step value that determines how many
pixels to skip between comparisons. By default, Motion has step of 1, which
means that every pixel will be analyzed to determine if motion is detected.
Adding a step value provides a way to dial down processing requirements in a
linear fashion to reduce the CPU burden of motion detection at the cost of
precision.
