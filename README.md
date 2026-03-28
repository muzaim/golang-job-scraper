# 🚀 Golang Job Scraper + Telegram Notifier

Simple automation tool built with Go to scrape job listings and send real-time notifications to Telegram.

---

## ✨ Features

* 🔎 Scrape job listings from websites
* ⚡ Keyword-based filtering
* 🔔 Telegram notification
* 🧠 In-memory deduplication (no spam)
* ⏱️ Auto run with scheduler

---

## 🛠️ Tech Stack

* Go (Golang)
* net/http (HTTP client)
* goquery (HTML parsing)
* Telegram Bot API

---

## 📂 Project Structure

```
job-scraper/
├── main.go
├── scraper/
├── parser/
├── notifier/
├── storage/
```

---

## ⚙️ Setup

### 1. Clone repo

```
git clone https://github.com/your-username/job-scraper.git
cd job-scraper
```

---

### 2. Install dependencies

```
go mod tidy
```

---

### 3. Create `.env`

```
BOT_TOKEN=your_telegram_bot_token
CHAT_ID=your_chat_id
KEYWORD=golang
URL=https://remoteok.com/remote-dev-jobs
```

---

## 🤖 Setup Telegram Bot

1. Open Telegram
2. Search **BotFather**
3. Run:

```
/start
/newbot
```

4. Copy BOT TOKEN

Get your Chat ID:

```
https://api.telegram.org/bot<YOUR_TOKEN>/getUpdates
```

---

## ▶️ Run

```
go run main.go
```

---

## 📸 Example Output

```
🔥 Job Found!

Title: Senior Backend Engineer
Link: https://remoteok.com/xxx
```

---

## 🚧 Future Improvements

* [ ] Multi-keyword support
* [ ] Persistent storage (SQLite/Redis)
* [ ] Concurrency (goroutines + worker pool)
* [ ] CLI support (flags)
* [ ] Docker support

---

## 💡 Why this project?

This project demonstrates:

* Working with HTTP requests
* Web scraping & parsing
* Automation workflow
* Integration with external APIs (Telegram)
* Basic concurrency potential in Go

---

## 🧑‍💻 Author

Built as a portfolio project to practice Golang beyond CRUD APIs.

---

## ⭐️ Support

If you like this project, consider giving it a star!
