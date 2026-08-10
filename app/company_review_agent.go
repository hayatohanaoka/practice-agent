package main

import (
	"google.golang.org/adk/v2/agent"
	"google.golang.org/adk/v2/agent/llmagent"
	"google.golang.org/adk/v2/model"
	"google.golang.org/adk/v2/tool"
)

func companyReviewAgent(llm model.LLM, tools []tool.Tool) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "company_review_agent",
		Model:       llm,
		Description: "企業のレビューを調査するエージェント",
		Instruction: "企業のレビューを調査するエージェントです。",
		Tools:       tools,
	})
}
