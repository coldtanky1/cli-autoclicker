package main

import (
	"os"
	"os/exec"
	"strconv"
	"time"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	hook "github.com/robotn/gohook"
)

type ClickerState struct {
	CPS        int
	isClicking bool
}

func main() {
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)

	letters := putils.LettersFromString("CLI Autoclicker")
	pterm.DefaultBigText.WithLetters(letters).Render()

	pterm.DefaultBasicText.Printfln("Welcome to %s.", pterm.Blue("CLI Autoclicker"))
	pterm.DefaultBasicText.Println("To toggle the autoclicker on/off use 'Ctrl + Y'.")

	result, _ := pterm.DefaultInteractiveTextInput.Show("Enter desired CPS")
	cps, err := strconv.Atoi(result)
	if err != nil {
		logger.Fatal("Could not convert string to int. Exiting.")
		os.Exit(1)
	}

	clicker := ClickerState{
		CPS:        cps,
		isClicking: false,
	}

	go clicker.StartListener()
	clicker.StartClicker()
}

func (c *ClickerState) ToggleClicker() {
	c.isClicking = !c.isClicking
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)
	logger.Info("Clicker toggled", logger.Args("active", c.isClicking))
}

func (c *ClickerState) StartListener() {
	logger := pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)
	logger.Info("Shortcut listener started!")

	hook.Register(hook.KeyDown, []string{"ctrl", "y"}, func(e hook.Event) {
		c.ToggleClicker()
	})

	s := hook.Start()
	<-hook.Process(s)
}

func (c *ClickerState) StartClicker() {
	yms := CPS2Ms(c.CPS)

	for {
		if c.isClicking {
			cmd := exec.Command("ydotool", "click", "0xC0")
			_ = cmd.Run()
			time.Sleep(time.Duration(yms) * time.Millisecond)
		} else {
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func CPS2Ms(CPS int) int {
	return 1000 / (2 * CPS)
}
