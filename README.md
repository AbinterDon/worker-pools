# Worker Pools Practice (Go) | Worker Pools 練習 (Go)

[English](#english) | [中文](#中文)

---

## English

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

## 中文

### 什麼是 Worker Pool (工作池)?
**Worker Pool** 是一種在開發併發 (Concurrent) 程式時常用的設計模式。它透過維持固定數量的「工作者」(in Go, these are goroutines) 來處理大量的任務。這種方法可以防止系統因為開啟過多進程而導致記憶體或 CPU 資源耗盡。

### 運作原理
此模式通常包含三個部分：
1.  **任務隊列 (Job Queue)**: 一個存放待處理任務的通道 (Channel)。
2.  **工作者 (Workers)**: 一組固定的 Goroutines，它們會從隊列中取出任務進行處理，並將結果送出。
3.  **結果隊列 (Result Queue)**: 一個存放處理結果的通道。

### 優點
-   **資源控制**: 限制活躍的 Goroutine 數量，避免過度消耗資源。
-   **效能優化**: 減少頻繁建立與銷毀 Goroutines 所產生的額外開銷。
-   **系統穩定性**: 在高負載情況下仍能維持穩定的處理節奏。

### 關於此專案
這是一個簡單的 Go 語言練習專案，展示了 Worker Pool 的基本實作，使用了：
-   `sync.WaitGroup` 來協調併發完成。
-   Buffered channels (有緩衝的通道) 來管理任務與結果。

### 併發模式比較
| 模式 | 資源控制 | 執行效能 | 實作難度 | 適合情境 |
| :--- | :--- | :--- | :--- | :--- |
| **同步處理** | 極佳 | 低 (緩慢) | 低 | 簡單腳本、線性邏輯 |
| **無限制併發** | 差 (有崩潰風險) | 極高 | 中 | 少量且獨立的短任務 |
| **Worker Pool** | **極佳** | **高且穩定** | **高** | **高負載、有外部資源限制** |

> [!IMPORTANT]
> **資源瓶頸觀點**：無限制的併發（如直接使用 `go func()`）可能會壓垮外部資源（如資料庫、外部 API）。Worker Pool 扮演「調節閥」的角色，在提升效能的同時，保護後端基礎設施不被瞬間湧入的請求擊垮。

### Real-world Scenarios (實際應用場景)

#### 1. Synchronous (同步)
- **Scenario**: A tool to back up 5 local configuration files.
- **Why**: The files are few, the operation is fast, and the setup cost of concurrency outweighs the benefits.
- **情境**: 備份 5 個本地設定檔。
- **原因**: 數量極少且執行極快，併發的額外開銷（Setup cost）反而會讓程式變慢。

#### 2. Unrestricted Concurrency (無限制併發 - `go func()`)
- **Scenario**: Sending an email notification after a user signs up.
- **Why**: The task is independent, low frequency, and doesn't share a limited pool of resources.
- **情境**: 使用者註冊後發送一封電子郵件通知。
- **原因**: 任務完全獨立、頻率低且不涉及需要精確限流的共用資源。

#### 3. Worker Pool (受控併發)
- **Scenario**: Resizing 10,000 product images for an e-commerce site.
- **Why**: CPU/Memory is limited. Opening 10,000 goroutines to process heavy images at once would crash the server. A pool of 8 or 16 workers ensures steady progress.
- **情境**: 為電商網站縮放 10,000 張商品圖片。
- **原因**: CPU 與記憶體資源有限。若同時啟動 10,000 個任務處理高解析度圖片會導致伺服器斷電或 OOM (記憶體溢出)。使用固定數量的 Worker（如 8 或 16 個）可確保系統穩定運行。
