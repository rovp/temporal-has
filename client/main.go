package main

import (
    "context"
    "log"

    "go.temporal.io/sdk/client"
    "github.com/rovp/temporal/internal/workflow"
)

func main() {
    c, err := client.Dial(client.Options{})
    if err != nil {
        log.Fatalln("Unable to create Temporal client", err)
    }
    defer c.Close()

    workflowOptions := client.StartWorkflowOptions{
        ID:        "payment_workflow",
        TaskQueue: "HMS_TASK_QUEUE",
    }

    we, err := c.ExecuteWorkflow(context.Background(), workflowOptions, workflow.PaymentWorkflow, "patient123")
    if err != nil {
        log.Fatalln("Unable to execute workflow", err)
    }

    log.Println("Started workflow", "WorkflowID", we.GetID(), "RunID", we.GetRunID())
}
