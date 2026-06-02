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

	result, _ := pterm.DefaultInteractiveTextInput.Show("Enter desired CPS")
	cps, err := strconv.Atoi(result)
	if err != nil {
		logger.Fatal("Could not convert string to int. Exiting.")
		os.Exit(1)
	}
	pterm.DefaultBasicText.Printfln("Got it! %d is now your desired CPS.", cps)

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
	logger.Info("Global listeners started!")
	logger.Info("Press Ctrl+Y to toggle clicker")
	logger.Info("Press Ctrl+Shift+V to change CPS")

	hook.Register(hook.KeyDown, []string{"ctrl", "y"}, func(e hook.Event) {
		c.ToggleClicker()
	})

	hook.Register(hook.KeyDown, []string{"ctrl", "shift", "v"}, func(e hook.Event) {
		if c.isClicking {
			c.isClicking = false

			result, _ := pterm.DefaultInteractiveTextInput.Show("Enter desired CPS")
			cps, err := strconv.Atoi(result)
			if err != nil {
				logger.Fatal("Could not convert string to int. Exiting.")
				os.Exit(1)
			}
			logger.Info("Got it! " + strconv.Itoa(cps) + " is now your new desired CPS.")
			c.ChangeCPS(cps)
		} else {

			result, _ := pterm.DefaultInteractiveTextInput.Show("Enter desired CPS")
			cps, err := strconv.Atoi(result)
			if err != nil {
				logger.Fatal("Could not convert string to int. Exiting.")
				os.Exit(1)
			}
			logger.Info("Got it! " + strconv.Itoa(cps) + " is now your new desired CPS.")
			c.ChangeCPS(cps)
		}
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

func (c *ClickerState) ChangeCPS(newCPS int) {
	c.CPS = newCPS
}

func CPS2Ms(CPS int) int {
	return 1000 / CPS
}
