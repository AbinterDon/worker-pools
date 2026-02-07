# Worker Pools Practice (Go)

[中文版 (Chinese Version)](README_zh.md)

---

### What is a Worker Pool?
A **Worker Pool** is a design pattern used in concurrent programming to manage a large number of tasks using a fixed number of workers (goroutines). This approach prevents the system from being overwhelmed by too many concurrent processes, which could exhaust memory or CPU resources.

### How it Works
The pattern typically involves three components:
1.  **Job Queue**: A channel where tasks are submitted.
2.  **Workers**: A fixed set of goroutines that pull jobs from the queue, process them, and potentially return results.
3.  **Result Queue**: A channel where workers send the outcomes of their work.

### Benefits
-   **Resource Control**: Limits the number of active goroutines.
-   **Performance**: Reduces the overhead of creating/destroying many goroutines.
-   **Stability**: Prevents resource exhaustion under heavy load.

### This Project
This is a simple Go project demonstrating the Worker Pool pattern. It uses:
-   `sync.WaitGroup` to coordinate completion.
-   Buffered channels to manage jobs and results.

### Concurrency Paradigms Comparison
| Paradigm | Resource Control | Performance | Complexity | Use Case |
| :--- | :--- | :--- | :--- | :--- |
| **Synchronous** | Excellent | Poor (Slow) | Low | Simple scripts, linear logic |
| **Unrestricted** | Poor (Risk) | Peak (Fast) | Medium | Small number of short tasks |
| **Worker Pool** | **Excellent** | **High/Stable** | **High** | **High load, limited external resources** |

> [!IMPORTANT]
> **The Resource Bottleneck**: Unrestricted concurrency can crash external resources (Databases, APIs). Worker Pools act as a "Throttle" to ensure high performance without overwhelming the infrastructure.

---

### Real-world Scenarios

#### 1. Synchronous
- **Scenario**: A tool to back up 5 local configuration files.
- **Why**: The files are few, the operation is fast, and the setup cost of concurrency outweighs the benefits.

#### 2. Unrestricted Concurrency (`go func()`)
- **Scenario**: Sending an email notification after a user signs up.
- **Why**: The task is independent, low frequency, and doesn't share a limited pool of resources.

#### 3. Worker Pool (Controlled Concurrency)
- **Scenario**: Resizing 10,000 product images for an e-commerce site.
- **Why**: CPU/Memory is limited. Opening 10,000 goroutines to process heavy images at once would crash the server. A pool of 8 or 16 workers ensures steady progress.
