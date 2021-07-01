package flyweight

import (
	"fmt"
	"sync"
)

type myMap struct {
	sync.Map
}

type ImageFlyweightFactory struct {
	//maps map[string]*ImageFlyweight
	maps myMap
}

var imageFactory *ImageFlyweightFactory
var once sync.Once

func GetImageFlyweightFactory() *ImageFlyweightFactory {
	once.Do(func() {
		imageFactory = &ImageFlyweightFactory{
			//maps: make(map[string]*ImageFlyweight),
		}
	})

	return imageFactory
}

func (f *ImageFlyweightFactory) Get(filename string) *ImageFlyweight {
	/*
		image := f.maps[filename]
		if image == nil {
			image = NewImageFlyweight(filename)
			f.maps[filename] = image
		}

		return image
	*/
	if image, ok := f.maps.Load(filename); ok {
		return image.(*ImageFlyweight) // type assertion
	}
	image, _ := f.maps.LoadOrStore(filename, NewImageFlyweight(filename))
	return image.(*ImageFlyweight) // type assertion
}

type ImageFlyweight struct {
	data string
}

func NewImageFlyweight(filename string) *ImageFlyweight {
	// Load image file
	data := fmt.Sprintf("image data %s", filename)
	return &ImageFlyweight{
		data: data,
	}
}

func (i *ImageFlyweight) Data() string {
	return i.data
}

type ImageViewer struct {
	*ImageFlyweight
}

func NewImageViewer(filename string) *ImageViewer {
	image := GetImageFlyweightFactory().Get(filename)
	return &ImageViewer{
		ImageFlyweight: image,
	}
}

func (i *ImageViewer) Display() {
	fmt.Printf("Display: %s\n", i.Data())
}
