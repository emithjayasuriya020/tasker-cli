package main

import "fmt"

func addTask(title string) {
	tasks, _ := loadTasks()
	task := Task{
		ID:    len(tasks) + 1,
		Title: title,
	}
	tasks = append(tasks, task)
	saveTasks(tasks)
	fmt.Printf("✅ Added: %s\n", title)
}

func listTasks() {
	tasks, _ := loadTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks yet!")
		return
	}
	for _, t := range tasks {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		fmt.Printf("%s #%d: %s\n", status, t.ID, t.Title)
	}
}

func markDone(id int) {
	tasks, _ := loadTasks()
	for i, t := range tasks {
		if t.ID == id {
			tasks[i].Done = true
			saveTasks(tasks)
			fmt.Printf("✅ Marked done: %s\n", t.Title)
			return
		}
	}
	fmt.Println("Task not found")
}

func deleteTask(id int) {
	tasks, _ := loadTasks()
	newTasks := []Task{}
	for _, t := range tasks {
		if t.ID != id {
			newTasks = append(newTasks, t)
		}
	}
	saveTasks(newTasks)
	fmt.Printf("🗑️ Deleted task #%d\n", id)
}

func printHelp() {
	fmt.Println(`
Tasker - A simple task manager

Commands:
	add <title>		Add a new task
	list			List all tasks
	done <id>		Mark task as donw
	delete <id>		Delete a task
		`)
}
