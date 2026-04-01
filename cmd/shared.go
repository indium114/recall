package cmd

import (
	"encoding/json"
	"github.com/fatih/color"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

type Task struct {
	Name  string `json:"name"`
	Prio  int    `json:"prio"`
	State int    `json:"state"`
	ID    int    `json:"id"`
}

type XP struct {
	XP          int   `json:"XP"`
	LastChecked int64 `json:"last_checked"`
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
	ensureFile(xpPath, []byte(`{"XP":0,"last_checked":0}`))
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

func applyDailyPenalty() {
	xp := loadXP()

	now := time.Now().Unix()
	last := xp.LastChecked

	if last == 0 {
		xp.LastChecked = now
		_ = saveXP(xp)
		return
	}

	days := int((now - last) / 86400)
	if days <= 0 {
		return
	}

	tasks := loadTasks()

	pending := 0
	for _, t := range tasks {
		if t.State == 0 {
			pending++
		}
	}

	if pending == 0 {
		xp.LastChecked = now
		_ = saveXP(xp)
		return
	}

	penalty := days * pending * 5
	xp.XP -= penalty
	xp.LastChecked = now

	_ = saveXP(xp)

	color.Red("󰓑 %d tasks pending, penalty of -%d XP applied", pending, penalty)
}
