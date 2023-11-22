type ProductionLine = {
    ID:     string,
    Name:   string,
    Status: ProductionLineStatus,
}

type ProductionLineStatus = "Staring" | "Running" | "Stopping" | "Stopped" | "Standby"