package workflow

import (
    "time"

    "go.temporal.io/sdk/workflow"
    "github.com/rovp/temporal/internal/cart"
    "github.com/rovp/temporal/internal/billing"
    "github.com/rovp/temporal/internal/payment"
    "github.com/rovp/temporal/internal/notification"
)

func PaymentWorkflow(ctx workflow.Context, patientID string) error {
    ao := workflow.ActivityOptions{
        StartToCloseTimeout: time.Second * 10,
    }
    ctx = workflow.WithActivityOptions(ctx, ao)

    if err := workflow.ExecuteActivity(ctx, cart.ValidateCart, patientID).Get(ctx, nil); err != nil {
        return err
    }

    if err := workflow.ExecuteActivity(ctx, billing.GenerateBill, patientID).Get(ctx, nil); err != nil {
        return err
    }

    if err := workflow.ExecuteActivity(ctx, payment.ProcessPayment, patientID).Get(ctx, nil); err != nil {
        // Compensation: cancel bill
        _ = workflow.ExecuteActivity(ctx, billing.CancelBill, patientID).Get(ctx, nil)
        return err
    }

    if err := workflow.ExecuteActivity(ctx, notification.SendReceipt, patientID).Get(ctx, nil); err != nil {
        return err
    }

    return nil
}
