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
