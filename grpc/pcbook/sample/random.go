package sample

import (
	"math/rand"
	"pcbook/pb"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomBool() bool {
	return rand.Intn(2) == 0
}

func randomFromSet[T any](set ...T) T {
	setLen := len(set)
	if setLen == 0 {
		var result T
		return result
	}
	return set[rand.Intn(setLen)]
}

func randomFloat(min float64, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	var layout pb.Keyboard_Layout
	switch rand.Intn(3) {
	case 0:
		layout = pb.Keyboard_QWERTY
	case 1:
		layout = pb.Keyboard_QWERTZ
	case 2:
		layout = pb.Keyboard_AZERTY
	}
	return layout
}

func randomCPUBrand() string {
	return randomFromSet("Intel", "AMD")
}

func randomCPUName(brand string) string {
	switch brand {
	case "Intel":
		return randomFromSet(
			"Xeon E-2286M",
			"Core i9-9980HK",
			"Core i7-9750H",
			"Core i5-9400F",
			"Core i3-1005G1",
		)
	case "AMD":
		return randomFromSet(
			"Ryzen 7 PRO 2700U",
			"Ryzen 5 PRO 3500U",
			"Ryzen 3 PRO 3200GE",
		)
	}
	return ""
}

func randomGPUBrand() string {
	return randomFromSet("NVidia", "AMD")
}

func randomGPUName(brand string) string {
	switch brand {
	case "NVidia":
		return randomFromSet(
			"RTX 2060",
			"RTX 2070",
			"GTX 1660-Ti",
			"GTX 1070",
		)
	case "AMD":
		return randomFromSet(
			"RX 590",
			"RX 580",
			"RX 5700-XT",
			"RX Vega-56",
		)
	}
	return ""
}

func NewRAMinGigabytes(set ...int) *pb.Memory {
	return &pb.Memory{
		Value: uint64(randomFromSet(set...)),
		Unit:  pb.Memory_GIGABYTE,
	}
}

func randomScreenResolution() *pb.Screen_Resolution {
	height := randomFromSet(1024, 1280, 1920, 2560, 3840)
	width := height / 16 * 9

	return &pb.Screen_Resolution{
		Width:  uint32(height),
		Height: uint32(width),
	}
}

func randomScreenPanel() pb.Screen_Panel {
	return randomFromSet(pb.Screen_IPS, pb.Screen_OLED)
}

func randomLaptopBrand() string {
	return randomFromSet("Apple", "Dell", "Lenovo")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Apple":
		return randomFromSet(
			"Macbook Air",
			"Macbook Pro",
		)
	case "Dell":
		return randomFromSet(
			"Lalitude",
			"Vostro",
			"XPS",
			"Alienware",
		)
	case "Lenovo":
		return randomFromSet(
			"Thinkpad X1",
			"Thinkpad P1",
			"Thinkpad P53",
		)
	}
	return ""
}
