package main

import (
    "log"

    "go.temporal.io/sdk/client"
    "go.temporal.io/sdk/worker"
    "github.com/rovp/temporal/internal/workflow"
    "github.com/rovp/temporal/internal/cart"
    "github.com/rovp/temporal/internal/billing"
    "github.com/rovp/temporal/internal/payment"
    "github.com/rovp/temporal/internal/notification"
)

func main() {
    c, err := client.Dial(client.Options{})
    if err != nil {
        log.Fatalln("Unable to create Temporal client", err)
    }
    defer c.Close()

    w := worker.New(c, "HMS_TASK_QUEUE", worker.Options{})

    w.RegisterWorkflow(workflow.PaymentWorkflow)
    w.RegisterActivity(cart.ValidateCart)
    w.RegisterActivity(billing.GenerateBill)
    w.RegisterActivity(billing.CancelBill)
    w.RegisterActivity(payment.ProcessPayment)
    w.RegisterActivity(notification.SendReceipt)

    err = w.Run(worker.InterruptCh())
    if err != nil {
        log.Fatalln("Unable to start worker", err)
    }
}

