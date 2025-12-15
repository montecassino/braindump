type CircuitState = "OPEN" | "CLOSED" | "HALF_OPEN";

class CircuitBreaker {
    private openTimestamp: number = 0;

    private resetTime: number; // milliseconds
    private currentFailureCount: number = 0;
    private failureCountThreshold: number;
    private state: CircuitState = "CLOSED";

    constructor(resetTime: number, failureCountThreshold: number) {
        this.resetTime = resetTime;
        this.failureCountThreshold = failureCountThreshold;
    }

    public async execute<T>(serviceCall: () => Promise<T>): Promise<T> {
        if (this.state === "OPEN") {
            const timeSinceOpen = Date.now() - this.openTimestamp;
            if (timeSinceOpen > this.resetTime) {
                this.state = "HALF_OPEN";
            } else {
                throw new Error("Service is unavailable");
            }
        }

        try {
            const result = await serviceCall();

            if (this.state === "HALF_OPEN") {
                this.state = "CLOSED";
            }
            this.currentFailureCount = 0;

            return result;

        } catch (error) {
            if (this.state === "HALF_OPEN") {
                this.state = "OPEN";
                this.openTimestamp = Date.now();
            } else if (this.state === "CLOSED") {
                this.currentFailureCount++;
                if (this.currentFailureCount >= this.failureCountThreshold) {
                    this.state = "OPEN";
                    this.openTimestamp = Date.now();
                }
            }
            
            throw error; 
        }
    }
}