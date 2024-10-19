package app

import (
	"VoGPro/GoPkg/pkg/mutex"
	"fmt"
)

func RunApp() {
	m := mutex.NewMutex(3)

	for i := 0; i < 3; i++ {
		go func(i int) {
			defer m.Unlock()
			fmt.Printf("Горутина %d: Привет, мир!\n", i)
		}(i)
	}
	m.Wait()
}
