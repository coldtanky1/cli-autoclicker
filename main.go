package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-vgo/robotgo"
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	hook "github.com/robotn/gohook"
)

var Debug = false

type ClickerState struct {
	CPS        int
	isClicking bool
	MStoSleep  time.Duration
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
	clicker.CPS2Ms()

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
	logger.Info("Press Ctrl+Shift+; to change CPS")

	hook.Register(hook.KeyDown, []string{"ctrl", "y"}, func(e hook.Event) {
		c.ToggleClicker()
	})

	hook.Register(hook.KeyDown, []string{"ctrl", "shift", ";"}, func(e hook.Event) {
		if c.isClicking {
			c.isClicking = false

			changeCPSResult, _ := pterm.DefaultInteractiveTextInput.Show("Enter desired CPS")
			newCps, err := strconv.Atoi(changeCPSResult)
			if err != nil {
				logger.Fatal("Could not convert string to int. Exiting.")
				os.Exit(1)
			}
			logger.Info("Got it! " + strconv.Itoa(newCps) + " is now your new desired CPS.")
			c.CPS = newCps
			c.CPS2Ms()
		} else {

			changeCPSResult, _ := pterm.DefaultInteractiveTextInput.Show("Enter desired CPS")
			newCps, err := strconv.Atoi(changeCPSResult)
			if err != nil {
				logger.Fatal("Could not convert string to int. Exiting.")
				os.Exit(1)
			}
			logger.Info("Got it! " + strconv.Itoa(newCps) + " is now your new desired CPS.")
			c.CPS = newCps
			c.CPS2Ms()
		}
	})

	s := hook.Start()
	<-hook.Process(s)
}

func (c *ClickerState) StartClicker() {
	for {
		if c.isClicking {
			DebugPrintf("c.MStoSleep = %d", c.MStoSleep)
			DebugPrintf("c.CPS = %d", c.CPS)
			robotgo.Click()
			time.Sleep(c.MStoSleep * time.Millisecond)
		} else {
			time.Sleep(50 * time.Millisecond)
		}
	}
}

func DebugPrintf(format string, v ...any) {
	if Debug {
		log.Printf(format, v)
	} else {
		return
	}
}

func (c *ClickerState) CPS2Ms() {
	c.MStoSleep = time.Duration(1000 / c.CPS)
}
