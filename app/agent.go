package main

import (
	"google.golang.org/adk/v2/agent"
	"google.golang.org/adk/v2/agent/llmagent"
	"google.golang.org/adk/v2/model"
	"google.golang.org/adk/v2/tool"
)

func searchAgent(llm model.LLM, tools []tool.Tool) (agent.Agent, error) {
	return llmagent.New(llmagent.Config{
		Name:        "search_agent",
		Model:       llm,
		Description: "ユーザーが入力した質問に、Web検索のデータをもとに回答するエージェント",
		Instruction: "自身の知識で答えるのではなく、Web検索のデータをもとに回答するエージェントです。",
		Tools:       tools,
	})
}
