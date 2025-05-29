# 🌐 Cross-Platform Browser & Notification Package for Go

A simple and elegant Go package to handle:
- ✅ Opening URLs in the default browser
- 🔔 Sending native desktop notifications

---

## 🖥️ Supported Operating Systems

| Platform | Browser Support | Notification Support |
|----------|------------------|-----------------------|
| 🪟 Windows | ✅ | ✅ (via PowerShell + BurntToast) |
| 🍎 macOS   | ✅ | ✅ (via `osascript`)             |
| 🐧 Linux   | ✅ | ✅ (via `notify-send`)           |

---

## 📦 Installation

### 📁 Use `go get`
```bash
go get github.com/saravanan611/notiFly
```


## 🚀 Quick Start

```go
package main

import (
	"fmt"

	"github.com/saravanan611/notiFly"
)

func main() {
	p, err := notify.GetPlatform()
	if err != nil {
		panic(err)
	}

	url := "https://example.com"
	title := "Hello"
	message := "This is a test notification."

	// 🌐 Open in browser
	if err := p.OpenBrowser(url); err != nil {
		fmt.Println("❌ Browser error:", err)
	}

	// 🔔 Send desktop notification
	if err := p.Notify(title, message, url); err != nil {
		fmt.Println("❌ Notification error:", err)
	}
}

```

---

## 📘 API Reference

### 🔹 `GetPlatform() (Platform, error)`

Returns a `Platform` interface based on the detected OS.

### 🔹 `OpenBrowser(url string) error`

Opens the provided URL using the system’s default browser.

### 🔹 `Notify(title, message, url string) error`

Sends a system-native notification with a title, message, and clickable link (optional).

---

## 🧪 Example Output

### ✅ Windows (BurntToast)

```
🔔 Hello
This is a test notification. https://example.com - Time: 2025-05-24 17:30:00
```

### ✅ macOS

System notification:

> Title: Hello
> Body: This is a test notification. [https://example.com](https://example.com) - Time: 2025-05-24 17:30:00

### ✅ Linux

A `notify-send` popup with the title and message.

---

## ⚙️ Requirements

Make sure required tools are installed for your platform:

| OS      | Requirement                 | Install Command                   |
| ------- | --------------------------- | --------------------------------- |
| Windows | PowerShell + BurntToast     | `Install-Module -Name BurntToast` |
| macOS   | `osascript` (pre-installed) | *(nothing to install)*            |
| Linux   | `notify-send`               | `sudo apt install libnotify-bin`  |


