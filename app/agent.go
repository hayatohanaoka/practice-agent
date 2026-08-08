package main

import (
	"google.golang.org/adk/v2/agent"
	"google.golang.org/adk/v2/agent/llmagent"
	"google.golang.org/adk/v2/model"
	"google.golang.org/adk/v2/tool"
)

func newTimeAgent(llm model.LLM, tools []tool.Tool) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "hello_time_agent",
		Model:       llm,
		Description: "Tells the current time in a specified city.",
		Instruction: "You are a helpful assistant that tells the current time in a city.",
		Tools:       tools,
	})
}

