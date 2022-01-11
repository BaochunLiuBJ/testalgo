package weektwo

import (
	"crypto/sha256"
	"strconv"
)

type direction int

const (
	north direction = iota
	east
	south
	west
)

func robotSim(commands []int, obstacles [][]int) int {
	curX, curY := 0, 0
	curDir := north
	var obstacleMap = make(map[string]struct{})

	for _, o := range obstacles {
		obstacleMap[calculateHash(o[0], o[1])] = struct{}{}
	}
	maxDistance := 0

	for _, command := range commands {
		if command == -1 {
			curDir = (curDir + 1) % 4
		} else if command == -2 {
			curDir = (curDir + 3) % 4
		} else {
			for i := 0; i < command; i++ {
				nextX, nextY := nextCoordination(curX, curY, curDir)
				h := calculateHash(nextX, nextY)
				_, ok :=  obstacleMap[h]
				if ok {
					break
				} else {
					curX = nextX
					curY = nextY
					if curX * curX + curY * curY > maxDistance {
						maxDistance = curX * curX + curY * curY
					}
				}
			}
		}
	}

	return maxDistance

}

func calculateHash(x int, y int) string {
	s := strconv.Itoa(x) + "," + strconv.Itoa(y)
	return string(sha256.New().Sum([]byte(s)))
}

func nextCoordination(x int, y int, dir direction) (int, int) {
	switch dir{
	case north:
		return x, y + 1 
	case east:
		return x+1, y
	case south:
		return x, y - 1
	case west:
		return x -1, y 	
	}
	return x, y 
}
