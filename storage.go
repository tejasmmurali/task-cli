package main

import (
	"encoding/json"
	"os"
)

const filePath = "task.json"

// load reads tasks from the JSON file. Returns an empty list if the file doesn't exist yet.
func load() (TaskList, error) {
	data, err := os.ReadFile(filePath)
	if os.IsNotExist(err) {
		return TaskList{}, nil
	}

	if err != nil {
		return nil, err
	}

	var tasks TaskList
	err = json.Unmarshal(data, &tasks)
	return tasks, err
}

func save(tasks TaskList) error {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return os.WriteFile(filePath, data, 0644)
}
