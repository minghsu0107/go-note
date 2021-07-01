package flyweight

import (
	"sync"
	"testing"
)

const parCount = 100000

func ExampleFlyweight() {
	viewer := NewImageViewer("image1.png")
	viewer.Display()
	// Output:
	// Display: image data image1.png
}

func TestFlyweight(t *testing.T) {
	viewer1 := NewImageViewer("image1.png")
	viewer2 := NewImageViewer("image1.png")

	if viewer1.ImageFlyweight != viewer2.ImageFlyweight {
		t.Fail()
	}
}

func TestParallelSingleton(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(parCount)
	instances := [parCount]*ImageViewer{}
	for i := 0; i < parCount; i++ {
		go func(index int) {
			instances[index] = NewImageViewer("image1.png")
			wg.Done()
		}(i)
	}
	wg.Wait()
	for i := 1; i < parCount; i++ {
		if instances[i].ImageFlyweight != instances[i-1].ImageFlyweight {
			t.Fatal("instance is not equal")
		}
	}
}
