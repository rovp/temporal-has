package notification

import (
    "context"
    "fmt"
)

func SendReceipt(ctx context.Context, patientID string) error {
    fmt.Println("Sending receipt to", patientID)
    // Implement receipt sending logic here
    return nil
}
