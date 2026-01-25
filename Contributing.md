# Contributing to Go Practice

Thanks for your interest in contributing!
This repository is meant to be a **hands-on learning space for Go**, covering multiple topics with small, focused tasks.

Contributions are welcome and encouraged.

---

## What you can contribute

### 1. Implement missing tasks

Some tasks already have a description (`.md`) but **no Go code yet**.

You can:

* Read the task description
* Add a `main.go` implementation
* Keep the solution simple and idiomatic

---

### 2. Add new Go topics

This repo is **not limited to channels**.

Future topics include (but are not limited to):

* Goroutines & concurrency patterns
* Context & cancellation
* Error handling
* Interfaces & composition
* Memory & performance
* Standard library deep-dives
* Testing & benchmarking

You can propose or add:

* New topic folders
* New levels
* New tasks with descriptions

---

### 3. Improve existing tasks

* Refactor code for clarity
* Fix bugs or edge cases
* Improve comments or explanations
* Make concurrency behavior clearer

---

## Project guidelines

* One task = one focused concept
* Prefer clarity over cleverness
* Avoid unnecessary abstractions
* No `time.Sleep` for correctness (allowed only for learning/demo)
* Keep tasks runnable with:

  ```bash
  go run main.go
  ```

---

## Adding a new task

1. Create a new task folder:

   ```text
   topic/level-x/task-y/
   ```
2. Add:

    * `main.go`
    * (optional) `README.md` or `.md` task description
3. Keep the task small and self-contained

---

## Code style

* Use idiomatic Go
* Follow `gofmt`
* Prefer standard library solutions
* Explicit > implicit

---

## How to contribute

1. Fork the repository
2. Create a new branch
3. Make your changes
4. Open a pull request with a clear description

---

## Notes

This repository is primarily educational.
Perfect code is less important than **clear intent and learning value**.
