package notify

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"time"
)

// Platform interface abstracts OS-specific behavior.
type Platform interface {
	OpenBrowser(url string) error
	Notify(title, message, url string) error
}

// Get current platform implementation
func GetPlatform() (Platform, error) {
	switch runtime.GOOS {
	case "windows":
		return &windowsPlatform{}, nil
	case "darwin":
		return &darwinPlatform{}, nil
	case "linux":
		return &linuxPlatform{}, nil
	default:
		return nil, fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
}

// ===== Windows Implementation =====

type windowsPlatform struct{}

func (p *windowsPlatform) OpenBrowser(url string) error {
	if url == "" {
		return errors.New("url must not be empty")
	}
	return exec.Command("cmd", "/c", "start", url).Start()
}

func (p *windowsPlatform) Notify(title, message, url string) error {
	ts := timestamp()
	return exec.Command("powershell", "-Command",
		"New-BurntToastNotification -Text '"+title+"', '"+message+" "+url+" - Time: "+ts+"'").Run()
}

// ===== macOS Implementation =====

type darwinPlatform struct{}

func (p *darwinPlatform) OpenBrowser(url string) error {
	if url == "" {
		return errors.New("url must not be empty")
	}
	return exec.Command("open", url).Start()
}

func (p *darwinPlatform) Notify(title, message, url string) error {
	ts := timestamp()
	script := fmt.Sprintf(`display notification "%s %s - Time: %s" with title "%s" sound name "default"`, message, url, ts, title)
	return exec.Command("osascript", "-e", script).Run()
}

// ===== Linux Implementation =====

type linuxPlatform struct{}

func (p *linuxPlatform) OpenBrowser(url string) error {
	if url == "" {
		return errors.New("url must not be empty")
	}
	return exec.Command("xdg-open", url).Start()
}

func (p *linuxPlatform) Notify(title, message, url string) error {
	ts := timestamp()
	return exec.Command("notify-send", title, message+"\n"+url+"\n - Time: "+ts).Run()
}

// ====== Utility ======

func timestamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
