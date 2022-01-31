package leet207

import (
	"container/list"
)

func canFinish(numCourses int, prerequisites [][]int) bool {
	c := initCourses(numCourses, prerequisites)
	return c.check()
}

type courses struct {
	toMap      map[int][]int
	inDegree   map[int]int
	queue      *list.List
	numCourses int
}

func initCourses(numCourses int, precourses [][]int) *courses {
	c := &courses{}
	c.inDegree = make(map[int]int)
	c.toMap = make(map[int][]int)
	c.queue = list.New()

	for i := 0; i < numCourses; i++ {
		c.inDegree[i] = 0
	}

	for _, edge := range precourses {
		pre := edge[1]
		end := edge[0]
		if c.toMap[pre] == nil {
			c.toMap[pre] = make([]int, 0)
		}
		c.toMap[pre] = append(c.toMap[pre], end)
		c.inDegree[end]++

	}
	return c
}

func (c courses) check() bool {
	courseSlice := make([]int, 0)
	for course, degree := range c.inDegree {
		if degree == 0 {
			c.queue.PushBack(course)
			courseSlice = append(courseSlice, course)
		}
	}

	for c.queue.Len() != 0 {
		element := c.queue.Front()
		course := element.Value.(int)
		courseSlice = append(courseSlice, course)
		c.queue.Remove(element)
		for _, cc := range c.toMap[course] {
			c.inDegree[cc]--
			if c.inDegree[cc] == 0 {
				c.queue.PushBack(cc)
			}
		}
	}

	result := true
	for _, d := range c.inDegree {
		if d != 0 {
			result = false
		}
	}

	return result
}
