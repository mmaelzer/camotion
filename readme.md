camotion
========

A motion detection library written in Go.

## Usage

#### func  Blended

```go
func Blended(img1, img2 image.Image, threshold int) image.Image
```
Returns a black and white image where white pixels denote a change between the
two images that is greater than the given threshold

#### func  Motion

```go
func Motion(img1, img2 image.Image, minChange, threshold int) bool
```
Motion will determine if motion occurs between two images given a minChange
value which is a 1-100 value representing percentage of image changed and a
threshold value which sets a below this value is noise. A reasonable threshold
default for jpegs (and potentially other formats) is 2500.
