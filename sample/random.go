package sample

import (
	"math/rand"

	"github.com/google/uuid"
	"github.com/manh737/go-grpc/protos"
)

func randomKeyboardLayout() protos.Keyboard_Layout {
	switch rand.Intn(4) {
	case 0:
		return protos.Keyboard_UNKNOWN
	case 1:
		return protos.Keyboard_QWERTY
	case 2:
		return protos.Keyboard_QWERTZ
	default:
		return protos.Keyboard_AZERTY
	}
}

func randomMemoryUnit() protos.Memory_Unit {
	switch rand.Intn(7) {
	case 0:
		return protos.Memory_UNKNOWN
	case 1:
		return protos.Memory_BIT
	case 2:
		return protos.Memory_BYTE
	case 3:
		return protos.Memory_KILOBYTE
	case 4:
		return protos.Memory_MEGABYTE
	case 5:
		return protos.Memory_GIGABYTE
	default:
		return protos.Memory_TERABYTE
	}
}

func randomScreenPanel() protos.Screen_Panel {
	switch rand.Intn(3) {
	case 0:
		return protos.Screen_UNKNOWN
	case 1:
		return protos.Screen_IPS
	default:
		return protos.Screen_OLED
	}
}

func randomStorageDriver() protos.Storage_Driver {
	switch rand.Intn(3) {
	case 0:
		return protos.Storage_UNKNOWN
	case 1:
		return protos.Storage_HDD
	default:
		return protos.Storage_SSD
	}
}

func randomBoolean() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomCPUName() string {
	return randomStringFromSet("i3", "i5", "i7", "i9", "Ryzen 3", "Ryzen 5", "Ryzen 7", "Ryzen 9")
}

func randomGPUBrand() string {
	return randomStringFromSet("Intel", "AMD", "Apple", "Qualcomm", "Nvidia")
}

func randomGPUName() string {
	return randomStringFromSet("HD Graphics", "UHD Graphics", "Radeon", "GeForce", "Quadro", "Tesla")
}

func randomLaptopBrand() string {
	return randomStringFromSet("Apple", "Dell", "Lenovo", "HP", "Asus", "Acer", "Microsoft", "Google")
}

func randomLaptopName() string {
	return randomStringFromSet("Macbook Pro", "Thinkpad X1", "Yoga 720", "Ideapad 320", "Inspiron 15", "Latitude 7480", "Alienware 17", "Zenbook", "Chromebook", "Spectre x360")
}

func randomInt(min, max int) int {
	return min + rand.Intn((max-min)+1)
}

func randomFloat64(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randomStringFromSet(a ...string) string {
	len_set := len(a)
	if len_set == 0 {
		return ""
	}
	return a[rand.Intn(len_set)]
}

func randomID() string {
	return uuid.New().String()
}
