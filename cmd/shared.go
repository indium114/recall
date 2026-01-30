package cmd

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strconv"
)

type Task struct {
	Name  string `json:"name"`
	Prio  int    `json:"prio"`
	State bool   `json:"state"`
	ID    int    `json:"id"`
}

type XP struct {
	XP int `json:"XP"`
}

var (
	homeDir, _ = os.UserHomeDir()
	todoPath   = filepath.Join(homeDir, ".recall")
	xpPath     = filepath.Join(homeDir, ".recall_xp")
)

func ensureFile(path string, content []byte) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		_ = os.WriteFile(path, content, 0644)
	}
}

func initStorage() {
	ensureFile(todoPath, []byte("[]"))
	ensureFile(xpPath, []byte(`{"XP":0}`))
}

func loadTasks() []Task {
	data, err := os.ReadFile(todoPath)
	if err != nil {
		return []Task{}
	}

	var tasks []Task
	_ = json.Unmarshal(data, &tasks)
	return tasks
}

func saveTasks(tasks []Task) error {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	return os.WriteFile(todoPath, data, 0644)
}

func loadXP() XP {
	data, err := os.ReadFile(xpPath)
	if err != nil {
		return XP{XP: 0}
	}

	var xp XP
	_ = json.Unmarshal(data, &xp)
	return xp
}

func saveXP(xp XP) error {
	data, _ := json.MarshalIndent(xp, "", "  ")
	return os.WriteFile(xpPath, data, 0644)
}

func increaseXP(amount int) {
	xp := loadXP()
	xp.XP += amount
	_ = saveXP(xp)
}

func decreaseXP(amount int) {
	xp := loadXP()
	xp.XP -= amount
	_ = saveXP(xp)
}

func atoi(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
