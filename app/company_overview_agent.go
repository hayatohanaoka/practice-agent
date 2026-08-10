package main

import (
	"google.golang.org/adk/v2/agent"
	"google.golang.org/adk/v2/agent/llmagent"
	"google.golang.org/adk/v2/model"
	"google.golang.org/adk/v2/tool"
)

func companyOverviewAgent(llm model.LLM, tools []tool.Tool) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "company_overview_agent",
		Model:       llm,
		Description: "企業の概要を調査するエージェント",
		Instruction: "企業の概要を調査するエージェントです。",
		Tools:       tools,
	})
}
