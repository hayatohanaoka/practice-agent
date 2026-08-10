package main

import (
	"context"
	"log"
	"os"

	"google.golang.org/genai"

	"google.golang.org/adk/v2/agent"
	"google.golang.org/adk/v2/agent/workflowagent"
	"google.golang.org/adk/v2/cmd/launcher"
	"google.golang.org/adk/v2/cmd/launcher/full"
	"google.golang.org/adk/v2/model/gemini"
	"google.golang.org/adk/v2/session"
	"google.golang.org/adk/v2/tool"
	"google.golang.org/adk/v2/tool/geminitool"
	"google.golang.org/adk/v2/workflow"
)

func main() {
	ctx := context.Background()

	model, err := gemini.NewModel(ctx, "gemini-2.5-flash", &genai.ClientConfig{
		APIKey: os.Getenv("GOOGLE_API_KEY"),
	})
	if err != nil {
		log.Fatalf("Failed to create model: %v", err)
	}

	companyOverviewAgent, err := companyOverviewAgent(model, []tool.Tool{geminitool.GoogleSearch{}})
	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	comoanyOverviewNode, err := workflow.NewAgentNode(companyOverviewAgent, workflow.NodeConfig{
		RetryConfig: workflow.DefaultRetryConfig(),
	})

	companyReviewAgent, err := companyReviewAgent(model, []tool.Tool{geminitool.GoogleSearch{}})
	if err != nil {
		log.Fatalf("Failed to create agent: %v", err)
	}

	companyReviewNode, err := workflow.NewAgentNode(companyReviewAgent, workflow.NodeConfig{
		RetryConfig: workflow.DefaultRetryConfig(),
	})

	gatherNode := workflow.NewJoinNode("gather")
	edgeBuilder := workflow.NewEdgeBuilder()
	edgeBuilder.AddFanOut(workflow.Start, comoanyOverviewNode, companyReviewNode)
	edgeBuilder.AddFanIn(gatherNode, comoanyOverviewNode, companyReviewNode)

	workflowAgent, err := workflowagent.New(workflowagent.Config{
		Name:        "company_search_workflow",
		Description: "企業情報の調査ワークフロー",
		Edges:       edgeBuilder.Build(),
	})
	if err != nil {
		log.Fatalf("Failed to create workflow: %v", err)
	}

	config := &launcher.Config{
		AgentLoader:    agent.NewSingleLoader(workflowAgent),
		SessionService: session.InMemoryService(),
	}

	args := os.Args[1:]
	if len(args) == 0 {
		args = []string{"web", "api"}
	}

	l := full.NewLauncher()
	if err := l.Execute(ctx, config, args); err != nil {
		log.Fatalf("Failed to launch launcher: %v", err)
	}
}
