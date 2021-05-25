package main

import (
	"math"
)

func collides(c1, c2 Circle) bool {
	dist := math.Sqrt(math.Pow(c2.center.x-c1.center.x, 2) +
		math.Pow(c2.center.y-c1.center.y, 2))

	return dist <= c1.radius+c2.radius
}

func checkForCollions() error {
	// fmt.Println("checking for collisions")
	for i := 0; i < len(Objects)-1; i++ {
		for j := i + 1; j < len(Objects); j++ {
			for _, c1 := range Objects[i].collisions {
				for _, c2 := range Objects[j].collisions {
					if collides(c1, c2) && Objects[i].isActive && Objects[j].isActive {
						err := Objects[i].Collision(Objects[j])
						if err != nil {
							return err
						}
						err = Objects[j].Collision(Objects[i])
						if err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil

}
