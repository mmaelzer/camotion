// Package camotion provides functions for determining if motion
// occurred between two images
package camotion

import (
	"image"
	"image/color"
	"math"
)

func pixels(img image.Image) int {
	b := img.Bounds()
	return b.Dx() * b.Dy()
}

// MotionWithStep will call Motion() with a step value that determines
// how many pixels to skip between comparisons. By default, Motion has step
// of 1, which means that every pixel will be analyzed to determine if
// motion is detected. Adding a step value provides a way to dial down
// processing requirements in a linear fashion to reduce the CPU burden
// of motion detection at the cost of precision.
func MotionWithStep(img1, img2 image.Image, minChange, threshold, step int) bool {
	if !img1.Bounds().Eq(img2.Bounds()) {
		return false
	}

	psize := float64(pixels(img1))
	mc := (psize / float64(step*step)) * float64(float64(minChange)/100.0)
	change := float64(0)

	b := img1.Bounds()
	for x := b.Min.X; x <= b.Max.X; x += step {
		for y := b.Min.Y; y <= b.Max.Y; y += step {
			avg1 := average(img1.At(x, y))
			avg2 := average(img2.At(x, y))
			diff := math.Abs(avg1 - avg2)
			if diff > float64(threshold) {
				change++
				if change >= mc {
					return true
				}
			}
		}
	}
	return false
}

// Motion will determine if motion occurs between two images
// given a minChange value which is a 1-100 value representing
// percentage of image changed and a threshold value which
// sets a below this value is noise. A reasonable threshold default for
// jpegs (and potentially other formats) is 2500.
func Motion(img1, img2 image.Image, minChange, threshold int) bool {
	return MotionWithStep(img1, img2, minChange, threshold, 1)
}

// Blended returns a black and white image where white pixels denote
// a change between the two images that is greater than the given
// threshold
func Blended(img1, img2 image.Image, threshold int) image.Image {
	t := float64(threshold)

	blend := image.NewNRGBA(img1.Bounds())
	if !img1.Bounds().Eq(img2.Bounds()) {
		return blend
	}

	b := img1.Bounds()
	for x := b.Min.X; x <= b.Max.X; x++ {
		for y := b.Min.Y; y <= b.Max.Y; y++ {
			avg1 := average(img1.At(x, y))
			avg2 := average(img2.At(x, y))
			diff := math.Abs(avg1 - avg2)
			var c color.Color
			if diff > t {
				c = color.White
			} else {
				c = color.Black
			}
			blend.Set(x, y, c)
		}
	}
	return blend
}

func average(c color.Color) float64 {
	r, g, b, _ := c.RGBA()
	avg := float64((r + g + b) / 3)
	return avg
}
