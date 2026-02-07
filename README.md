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
- **Scenario 1: Image Processing**: Resizing 10,000 product images. Workers ensure CPU isn't starved and the server remains responsive.
- **Scenario 2: Web Scraping**: Crawling thousands of URLs. Workers act as rate-limiters to avoid being blocked by the target site.
- **Scenario 3: Log Analysis**: Parsing GBs of logs. Breaking files into chunks and processing them in parallel using a fixed set of workers.
- **Scenario 4: Database Migration**: Moving millions of records. Workers ensure DB connection limits are respected.
- **Scenario 5: Background Jobs**: Handling asynchronous tasks (like PDF generation or email blasts) in a web server without blocking requests.
