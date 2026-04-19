# 🎯 Number Guessing Game (Go)

A simple command-line Number Guessing Game built with Go, where the player tries to guess a randomly generated number between 1 and 100 within a limited number of attempts based on difficulty level.

This project is based on the roadmap project:
👉 https://roadmap.sh/projects/number-guessing-game

---

## 🚀 Features

- 🎮 3 difficulty levels:
  - Easy (10 chances)
  - Medium (5 chances)
  - Hard (3 chances)
- 🎲 Random number between 1 and 100
- ⌨️ Robust CLI input handling using `bufio`
- 📊 Tracks number of attempts
- 🧠 Clean and modular Go structure (package-based design)

---

## 📁 Project Structure

```

number-guessing-game/
├── main.go
├── game/
│ ├── game.go
│ ├── difficulty.go
│ ├── input.go
│ └── messages.go

````

---

## 🛠️ How to Run

### 1. Clone the repository
```bash
git clone https://github.com/rizkyfauziilmi/number-guessing-game.git
cd number-guessing-game
````

### 2. Run the game

```bash
go run main.go
```

---

## 🎯 How to Play

1. Start the program
2. Choose difficulty:

   * `1` = Easy
   * `2` = Medium
   * `3` = Hard
3. Guess a number between 1–100
4. You will get hints:

   * Too high → "lower"
   * Too low → "higher"
5. Win by guessing correctly before chances run out

---

## 🧠 What You Learn

This project helps you understand:

* Go project structure (packages)
* Structs and methods
* Pointer usage in Go
* Input handling with `bufio`
* Random number generation
* Game loop logic

---

## 📌 Reference

Project challenge:
[https://roadmap.sh/projects/number-guessing-game](https://roadmap.sh/projects/number-guessing-game)