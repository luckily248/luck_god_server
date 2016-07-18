package model

type Step struct {
	description string
	root        *Step
	pStep       *Step
	nStep       *Step
}
