package main

import (
	"time"
)

type Task struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Done      bool      `json:"done"`
	CreatedAT time.Time `json:"created_at"`
}

type TaskList []Task

func (tl TaskList) NextID() int {
	max := 0
	for _, t := range tl {
		if t.ID > max {
			max = t.ID
		}
	}
	return max + 1
}

func (tl TaskList) FindIndex(id int) int {
	for i, t := range tl {
		if t.ID == id {
			return i
		}
	}
	return -1
}
