package main

import (
	"ekyu.moe/soda/i18n"

	"fmt"
	"os"

	surveyTerminal "gopkg.in/AlecAivazis/survey.v1/terminal"
)

func gracefulFatal(err error) {
	if err == surveyTerminal.InterruptErr {
		// 直接退
		os.Exit(2)
	}
	colorRed.Printf("\n%s：\n%s\n", i18n.EXCEPTION_OCCURRED, err)
	gracefulError(err)
	colorDim.Println(i18n.PRESS_ENTER_TO_EXIT)
	fmt.Scanln()
	os.Exit(2)
}

func gracefulError(err error) {
	if err == surveyTerminal.InterruptErr {
		// 不算错误，不处理
		fmt.Println()
		return
	}
	colorRed.Printf("\n%s：\n%s\n", i18n.EXCEPTION_OCCURRED, err)
}
