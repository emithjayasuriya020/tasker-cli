# tasker-cli ⚡

A lightweight command-line task manager built with Go. Add, list, complete, and delete tasks right from your terminal.

## Features

- ✅ Add tasks
- 📋 List all tasks
- ✔️ Mark tasks as done
- 🗑️ Delete tasks
- 💾 Persistent storage using JSON

## Installation

### Prerequisites

- [Go](https://go.dev/dl/) 1.18 or higher

### Build from source

```bash
git clone https://github.com/your-username/tasker-cli.git
cd tasker-cli
go build -o tasker.exe   # Windows
go build -o tasker       # macOS / Linux
```

## Usage

```bash
# Add a task
.\tasker.exe add "Learn Go"

# List all tasks
.\tasker.exe list

# Mark a task as done
.\tasker.exe done 1

# Delete a task
.\tasker.exe delete 1
```

### Example

```
$ .\tasker.exe add "Learn Go"
✅ Added: Learn Go

$ .\tasker.exe add "Build CLI tool"
✅ Added: Build CLI tool

$ .\tasker.exe list
[ ] #1: Learn Go
[ ] #2: Build CLI tool

$ .\tasker.exe done 1
✅ Marked done: Learn Go

$ .\tasker.exe list
[x] #1: Learn Go
[ ] #2: Build CLI tool

$ .\tasker.exe delete 2
🗑️ Deleted task #2
```

## Project Structure

```
tasker/
├── main.go        # Entry point and CLI argument handling
├── task.go        # Task struct, JSON load/save
├── commands.go    # add, list, done, delete logic
├── tasks.json     # Auto-generated task storage
└── go.mod
```

## How It Works

Tasks are stored locally in a `tasks.json` file in the same directory as the binary. The file is created automatically when you add your first task.

## What I Learned

This project was built as a hands-on introduction to Go, covering:

- Structs and slices
- Error handling
- JSON encoding/decoding
- File I/O
- CLI argument parsing with `os.Args`
- Building binaries with `go build`

## License

MIT
