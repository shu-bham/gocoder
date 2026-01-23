# Senior Golang Developer Interview Questions

## Microservices Architecture & Design

### Question 1: Circuit Breaker Implementation
Your microservice is calling a downstream payment service that occasionally becomes unavailable, causing cascading failures across your system. How would you implement a circuit breaker pattern in Go? What states would it have, and how would you handle partial failures?

<details>
<summary>Answer</summary>

A circuit breaker should have three states: Closed (normal operation), Open (failing fast), and Half-Open (testing recovery).

**Implementation approach:**
- Use a counter for consecutive failures with a threshold (e.g., 5 failures)
- Implement timeout for state transitions (e.g., 30 seconds in Open state before trying Half-Open)
- Use mutex for thread-safe state transitions
- Include metrics/logging for observability

```go
type CircuitBreaker struct {
    maxFailures  int
    timeout      time.Duration
    failures     int
    lastFailTime time.Time
    state        State
    mu           sync.RWMutex
}
```

**Key considerations:**
- Exponential backoff for retry attempts
- Different thresholds for different error types (timeouts vs 500s)
- Health check endpoint to force state transitions
- Bulkhead pattern to isolate different downstream services

**Suggested Reading:**
- "Release It!" by Michael Nygard (Circuit Breaker pattern)
- github.com/sony/gobreaker package documentation
- Martin Fowler's Circuit Breaker article
</details>

### Question 2: Service Discovery in Kubernetes
You're migrating from a monolith to microservices on Kubernetes. Services need to discover and communicate with each other. How would you implement service discovery? What are the tradeoffs between client-side and server-side discovery?

<details>
<summary>Answer</summary>

**Kubernetes-native approach:**
- Use Kubernetes Services with DNS (server-side discovery)
- ClusterIP for internal services
- Headless services for StatefulSets
- Use environment variables or DNS for service endpoints

**Client-side discovery:**
- More control over load balancing algorithms
- Lower latency (no proxy hop)
- More complex client code
- Example: Using etcd/Consul with custom client library

**Server-side discovery (recommended for K8s):**
- Simpler client code
- Centralized configuration
- K8s handles health checks and routing
- Use Ingress/Service mesh (Istio, Linkerd) for advanced routing

**Implementation:**
```go
// Using K8s DNS
serviceURL := "http://payment-service.default.svc.cluster.local:8080"

// Using environment variables
serviceHost := os.Getenv("PAYMENT_SERVICE_HOST")
servicePort := os.Getenv("PAYMENT_SERVICE_PORT")
```

**Suggested Reading:**
- "Kubernetes Patterns" by Bilgin Ibryam
- Kubernetes Service documentation
- "Building Microservices" by Sam Newman (Service Discovery chapter)
</details>

### Question 3: Distributed Tracing Implementation
Your microservices system has high latency, but you can't pinpoint where the bottleneck is. How would you implement distributed tracing in Go? What context propagation strategy would you use?

<details>
<summary>Answer</summary>

**Implementation with OpenTelemetry:**
- Use OpenTelemetry SDK for Go
- Propagate trace context via HTTP headers (W3C Trace Context)
- Use context.Context for propagation within services
- Export to Jaeger/Zipkin/Tempo

```go
import (
    "go.opentelemetry.io/otel"
    "go.opentelemetry.io/otel/trace"
)

func processRequest(ctx context.Context) {
    tracer := otel.Tracer("service-name")
    ctx, span := tracer.Start(ctx, "operation-name")
    defer span.End()
    
    // Pass ctx to downstream calls
    makeHTTPCall(ctx, url)
}
```

**Context propagation:**
- Extract trace context from incoming requests
- Inject into outgoing HTTP/gRPC calls
- Use middleware for automatic instrumentation
- Include trace ID in logs for correlation

**Best practices:**
- Sample traces (e.g., 1% in production)
- Add custom attributes for business context
- Set span status on errors
- Use span events for important milestones

**Suggested Reading:**
- OpenTelemetry Go documentation
- "Distributed Tracing in Practice" by Austin Parker
- Jaeger documentation and best practices
</details>

## Performance & Debugging

### Question 4: Memory Leak Investigation
Your Go service's memory usage keeps growing over time until the pod gets OOMKilled. How would you investigate and fix this? What tools would you use?

<details>
<summary>Answer</summary>

**Investigation steps:**

1. **Enable pprof:**
```go
import _ "net/http/pprof"
go func() {
    http.ListenAndServe("localhost:6060", nil)
}()
```

2. **Capture heap profiles:**
```bash
curl http://localhost:6060/debug/pprof/heap > heap.prof
go tool pprof heap.prof
```

3. **Compare snapshots:**
```bash
go tool pprof -base heap1.prof heap2.prof
```

**Common causes:**
- Goroutine leaks (blocked on channel, infinite loops)
- Global variables holding references
- Unclosed resources (files, HTTP connections)
- Large slices that aren't released
- Timers not stopped
- Context not being cancelled

**Detection:**
```bash
# Check goroutine count
curl http://localhost:6060/debug/pprof/goroutine?debug=2

# Check heap allocation
go tool pprof -alloc_space http://localhost:6060/debug/pprof/heap
```

**Prevention:**
- Use `defer` for cleanup
- Set `SetMaxIdleConns` and `SetMaxIdleConnsPerHost` on HTTP clients
- Use context with timeout
- Monitor goroutine count in production

**Suggested Reading:**
- "Go Performance Workshop" by Dave Cheney
- Official Go pprof documentation
- "Profiling Go Programs" blog post by Go team
</details>

### Question 5: High CPU Usage Debugging
Your service is consuming 300% CPU in production. How would you identify the hot path without impacting production traffic?

<details>
<summary>Answer</summary>

**Profiling approach:**

1. **CPU profiling with minimal overhead:**
```bash
# 30-second CPU profile
curl http://localhost:6060/debug/pprof/profile?seconds=30 > cpu.prof
go tool pprof -http=:8080 cpu.prof
```

2. **Analyze flame graphs:**
- Identify functions consuming most CPU time
- Look for unexpected regex compilation, JSON marshaling, or reflection

**Common culprits:**
- Inefficient JSON marshaling in hot path
- Regular expression compilation in loops
- Excessive allocations triggering GC
- Unbounded goroutines
- Lock contention

**Production-safe techniques:**
- Use continuous profiling (Pyroscope, Parca)
- Sample-based profiling (low overhead)
- Export metrics to Prometheus
- Use runtime metrics: `runtime.NumGoroutine()`, `runtime.ReadMemStats()`

**Example optimization:**
```go
// Bad: Compiles regex in loop
for _, item := range items {
    re := regexp.MustCompile(pattern)
    if re.MatchString(item) { ... }
}

// Good: Compile once
re := regexp.MustCompile(pattern)
for _, item := range items {
    if re.MatchString(item) { ... }
}
```

**Suggested Reading:**
- "Profiling Go Programs" by Go Blog
- pprof documentation
- "The Go Programming Language" Chapter 11 (Testing)
- Continuous profiling best practices
</details>

### Question 6: Goroutine Leak Detection
You notice your service has 50,000+ goroutines after running for a few hours. How do you identify which goroutines are leaking and why?

<details>
<summary>Answer</summary>

**Detection:**

```bash
# Get goroutine dump
curl http://localhost:6060/debug/pprof/goroutine?debug=2 > goroutines.txt

# Analyze with pprof
go tool pprof http://localhost:6060/debug/pprof/goroutine
```

**Common leak patterns:**

1. **Channel blocking:**
```go
// Leak: goroutine waits forever
go func() {
    ch <- data  // blocks if no receiver
}()

// Fix: use select with timeout
go func() {
    select {
    case ch <- data:
    case <-time.After(5 * time.Second):
        return
    case <-ctx.Done():
        return
    }
}()
```

2. **HTTP request without timeout:**
```go
// Leak
resp, _ := http.Get(url)

// Fix
client := &http.Client{Timeout: 10 * time.Second}
resp, err := client.Get(url)
```

3. **Worker pool without shutdown:**
```go
// Use errgroup for proper cleanup
g, ctx := errgroup.WithContext(ctx)
for i := 0; i < workers; i++ {
    g.Go(func() error {
        return processWork(ctx)
    })
}
g.Wait()
```

**Prevention:**
- Always use context with cancellation
- Set timeouts on blocking operations
- Use `goleak` package in tests
- Monitor `runtime.NumGoroutine()` metric

**Suggested Reading:**
- "Concurrency in Go" by Katherine Cox-Buday
- uber-go/goleak documentation
- "Common Goroutine Leaks" blog posts
</details>

## Concurrency & Synchronization

### Question 7: Rate Limiter Implementation
You need to implement a distributed rate limiter that allows 1000 requests per minute per user across multiple service instances. How would you design this in Go?

<details>
<summary>Answer</summary>

**Design options:**

**1. Token Bucket Algorithm (Redis):**
```go
func allowRequest(userID string) (bool, error) {
    key := fmt.Sprintf("rate_limit:%s", userID)
    now := time.Now().Unix()
    
    // Use Redis pipeline
    pipe := redisClient.Pipeline()
    pipe.ZRemRangeByScore(ctx, key, "-inf", now-60)
    pipe.ZCard(ctx, key)
    pipe.ZAdd(ctx, key, redis.Z{Score: float64(now), Member: now})
    pipe.Expire(ctx, key, 61*time.Second)
    
    cmds, err := pipe.Exec(ctx)
    count := cmds[1].(*redis.IntCmd).Val()
    
    return count < 1000, err
}
```

**2. Sliding Window Counter:**
- More accurate than fixed window
- Use Redis sorted sets with timestamp scores
- Remove old entries outside window

**3. Local rate limiting with sync/rate:**
```go
import "golang.org/x/time/rate"

limiter := rate.NewLimiter(rate.Limit(1000.0/60.0), 1000)
if !limiter.Allow() {
    return errors.New("rate limit exceeded")
}
```

**Distributed considerations:**
- Redis for shared state
- Handle Redis failures gracefully (fail open vs closed)
- Use Lua scripts for atomic operations
- Consider eventual consistency

**Suggested Reading:**
- "Token Bucket Algorithm" explanation
- Redis Lua scripting guide
- golang.org/x/time/rate package
- "Stripe's rate limiting in practice" blog post
</details>

### Question 8: Worker Pool with Graceful Shutdown
Design a worker pool that processes jobs from a queue and implements graceful shutdown when receiving a SIGTERM. How do you ensure no jobs are lost?

<details>
<summary>Answer</summary>

**Implementation:**

```go
type WorkerPool struct {
    workers   int
    jobQueue  chan Job
    ctx       context.Context
    cancel    context.CancelFunc
    wg        sync.WaitGroup
}

func (p *WorkerPool) Start() {
    p.ctx, p.cancel = context.WithCancel(context.Background())
    
    for i := 0; i < p.workers; i++ {
        p.wg.Add(1)
        go p.worker()
    }
}

func (p *WorkerPool) worker() {
    defer p.wg.Done()
    
    for {
        select {
        case job := <-p.jobQueue:
            p.processJob(job)
        case <-p.ctx.Done():
            // Process remaining jobs in queue
            for {
                select {
                case job := <-p.jobQueue:
                    p.processJob(job)
                default:
                    return
                }
            }
        }
    }
}

func (p *WorkerPool) Shutdown(timeout time.Duration) error {
    close(p.jobQueue) // Stop accepting new jobs
    p.cancel()        // Signal workers to stop
    
    done := make(chan struct{})
    go func() {
        p.wg.Wait()
        close(done)
    }()
    
    select {
    case <-done:
        return nil
    case <-time.After(timeout):
        return errors.New("shutdown timeout")
    }
}
```

**Signal handling:**
```go
sigCh := make(chan os.Signal, 1)
signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

<-sigCh
pool.Shutdown(30 * time.Second)
```

**Suggested Reading:**
- "Concurrency in Go" by Katherine Cox-Buday
- errgroup package documentation
- "Graceful Shutdown in Go" blog posts
</details>

### Question 9: Preventing Race Conditions
You have a cache that's accessed by multiple goroutines. Some read, some write, and some do read-modify-write. How do you prevent race conditions while maintaining good performance?

<details>
<summary>Answer</summary>

**Solutions by use case:**

**1. Read-heavy workload (sync.RWMutex):**
```go
type Cache struct {
    mu    sync.RWMutex
    items map[string]interface{}
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mu.RLock()
    defer c.mu.RUnlock()
    val, ok := c.items[key]
    return val, ok
}

func (c *Cache) Set(key string, val interface{}) {
    c.mu.Lock()
    defer c.mu.Unlock()
    c.items[key] = val
}
```

**2. High contention (sync.Map):**
```go
var cache sync.Map

cache.Store(key, value)
val, ok := cache.Load(key)
cache.LoadOrStore(key, value)
```

**3. Atomic operations for counters:**
```go
var counter int64
atomic.AddInt64(&counter, 1)
count := atomic.LoadInt64(&counter)
```

**4. Read-modify-write pattern:**
```go
func (c *Cache) Increment(key string) {
    c.mu.Lock()
    defer c.mu.Unlock()
    
    val, ok := c.items[key]
    if !ok {
        c.items[key] = 1
        return
    }
    c.items[key] = val.(int) + 1
}
```

**Testing for races:**
```bash
go test -race ./...
```

**Suggested Reading:**
- Go Memory Model documentation
- sync package documentation
- "The Go Memory Model" by Russ Cox
- "Mutex or Channel" Go blog post
</details>

## Database & Data Layer

### Question 10: Connection Pool Tuning
Your service experiences intermittent database timeouts under load. How would you tune the database connection pool? What metrics would you monitor?

<details>
<summary>Answer</summary>

**Connection pool configuration:**

```go
db.SetMaxOpenConns(25)        // Max connections
db.SetMaxIdleConns(25)        // Keep all connections idle
db.SetConnMaxLifetime(5 * time.Minute)  // Recycle connections
db.SetConnMaxIdleTime(5 * time.Minute)  // Close idle connections
```

**Tuning principles:**

1. **MaxOpenConns:**
   - Formula: `available_db_connections / number_of_service_instances`
   - Start with 25-50, adjust based on load
   - Too high: exhaust DB connections
   - Too low: connection wait time increases

2. **MaxIdleConns:**
   - Should equal MaxOpenConns for stable workloads
   - Prevents connection churn
   - Lower for variable traffic patterns

3. **ConnMaxLifetime:**
   - Prevents stale connections
   - Handle DB-side connection recycling
   - Set to less than DB's connection timeout

**Monitoring metrics:**
```go
stats := db.Stats()
log.Printf("InUse: %d, Idle: %d, WaitCount: %d, WaitDuration: %v",
    stats.InUse,
    stats.Idle,
    stats.WaitCount,
    stats.WaitDuration)
```

**Key metrics:**
- `WaitCount` and `WaitDuration` (should be low)
- `InUse` (should be < MaxOpenConns)
- Query duration percentiles (p50, p95, p99)

**Common issues:**
- Long-running transactions holding connections
- Missing context timeouts on queries
- Connection leaks (not closing rows/statements)

**Suggested Reading:**
- "Configuring sql.DB for Better Performance"
- PostgreSQL connection pooling best practices
- PgBouncer documentation
</details>

### Question 11: Handling Database Deadlocks
Your service occasionally encounters database deadlocks causing transaction failures. How do you detect, prevent, and handle them in Go?

<details>
<summary>Answer</summary>

**Detection:**

```go
func executeTransaction(db *sql.DB) error {
    tx, err := db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer tx.Rollback()
    
    // ... perform operations
    
    if err := tx.Commit(); err != nil {
        if isDeadlock(err) {
            // Log and retry
            return retryTransaction(db)
        }
        return err
    }
    return nil
}

func isDeadlock(err error) bool {
    // PostgreSQL
    if pqErr, ok := err.(*pq.Error); ok {
        return pqErr.Code == "40P01"
    }
    // MySQL
    if mysqlErr, ok := err.(*mysql.MySQLError); ok {
        return mysqlErr.Number == 1213
    }
    return false
}
```

**Prevention strategies:**

1. **Consistent lock ordering:**
```go
// Always acquire locks in same order (e.g., by ID)
func transfer(from, to int, amount decimal.Decimal) error {
    if from > to {
        from, to = to, from
    }
    // Lock accounts in consistent order
}
```

2. **Keep transactions short:**
- Minimize work inside transactions
- Move non-critical operations outside
- Use `SELECT FOR UPDATE` judiciously

3. **Use appropriate isolation levels:**
```go
tx, err := db.BeginTx(ctx, &sql.TxOptions{
    Isolation: sql.LevelReadCommitted,
})
```

**Retry logic with exponential backoff:**
```go
func withRetry(fn func() error, maxRetries int) error {
    for i := 0; i < maxRetries; i++ {
        err := fn()
        if err == nil {
            return nil
        }
        if !isDeadlock(err) {
            return err
        }
        backoff := time.Duration(math.Pow(2, float64(i))) * 100 * time.Millisecond
        time.Sleep(backoff)
    }
    return errors.New("max retries exceeded")
}
```

**Suggested Reading:**
- Database-specific deadlock documentation
- "High Performance MySQL" (Deadlocks chapter)
- "Designing Data-Intensive Applications" by Martin Kleppmann
</details>

### Question 12: Optimistic Locking Implementation
You need to implement optimistic locking to handle concurrent updates to the same record. How would you design this in Go?

<details>
<summary>Answer</summary>

**Database schema:**
```sql
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255),
    price DECIMAL(10,2),
    version INTEGER NOT NULL DEFAULT 1
);
```

**Go implementation:**

```go
type Product struct {
    ID      int
    Name    string
    Price   decimal.Decimal
    Version int
}

func UpdateProduct(db *sql.DB, product *Product) error {
    result, err := db.Exec(`
        UPDATE products 
        SET name = $1, price = $2, version = version + 1
        WHERE id = $3 AND version = $4
    `, product.Name, product.Price, product.ID, product.Version)
    
    if err != nil {
        return err
    }
    
    rowsAffected, err := result.RowsAffected()
    if err != nil {
        return err
    }
    
    if rowsAffected == 0 {
        return ErrConcurrentUpdate
    }
    
    product.Version++
    return nil
}
```

**Retry logic:**
```go
func UpdateProductWithRetry(db *sql.DB, id int, updateFn func(*Product)) error {
    const maxRetries = 3
    
    for i := 0; i < maxRetries; i++ {
        product, err := GetProduct(db, id)
        if err != nil {
            return err
        }
        
        updateFn(product)
        
        err = UpdateProduct(db, product)
        if err == nil {
            return nil
        }
        if err != ErrConcurrentUpdate {
            return err
        }
        
        // Exponential backoff
        time.Sleep(time.Duration(1<<uint(i)) * 100 * time.Millisecond)
    }
    
    return errors.New("max retries exceeded")
}
```

**When to use:**
- Low contention scenarios
- Read-heavy workloads
- Avoid pessimistic locks (SELECT FOR UPDATE)

**Suggested Reading:**
- "Optimistic vs Pessimistic Locking" articles
- Martin Fowler's "Patterns of Enterprise Application Architecture"
- Database transaction isolation levels
</details>

## HTTP & gRPC

### Question 13: HTTP Server Graceful Shutdown
Implement a production-ready HTTP server in Go with graceful shutdown, proper timeouts, and context propagation. What edge cases should you handle?

<details>
<summary>Answer</summary>

**Complete implementation:**

```go
type Server struct {
    httpServer *http.Server
}

func NewServer(addr string, handler http.Handler) *Server {
    return &Server{
        httpServer: &http.Server{
            Addr:         addr,
            Handler:      handler,
            ReadTimeout:  15 * time.Second,
            WriteTimeout: 15 * time.Second,
            IdleTimeout:  60 * time.Second,
            // Prevent Slowloris attacks
            ReadHeaderTimeout: 5 * time.Second,
            MaxHeaderBytes:    1 << 20, // 1 MB
        },
    }
}

func (s *Server) Start() error {
    return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
    // Stop accepting new connections
    // Wait for active requests to complete
    return s.httpServer.Shutdown(ctx)
}

func main() {
    handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        // Use request context
        ctx := r.Context()
        
        // Simulate work
        select {
        case <-time.After(5 * time.Second):
            w.WriteHeader(http.StatusOK)
        case <-ctx.Done():
            // Client disconnected or timeout
            return
        }
    })
    
    srv := NewServer(":8080", handler)
    
    go func() {
        if err := srv.Start(); err != http.ErrServerClosed {
            log.Fatal(err)
        }
    }()
    
    // Wait for interrupt signal
    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
    <-sigCh
    
    // Graceful shutdown with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := srv.Shutdown(ctx); err != nil {
        log.Printf("Shutdown error: %v", err)
    }
}
```

**Middleware for context:**
```go
func timeoutMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        ctx, cancel := context.WithTimeout(r.Context(), 10*time.Second)
        defer cancel()
        
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}
```

**Edge cases:**
- WebSocket connections (need special handling)
- Long-polling requests
- File uploads/downloads
- Streaming responses

**Suggested Reading:**
- "The complete guide to Go net/http timeouts"
- Cloudflare blog on timeouts
- Go net/http documentation
</details>

### Question 14: gRPC Interceptors for Observability
Implement gRPC server and client interceptors for logging, metrics, and tracing. How do you handle streaming RPCs differently?

<details>
<summary>Answer</summary>

**Unary server interceptor:**

```go
func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
    return func(
        ctx context.Context,
        req interface{},
        info *grpc.UnaryServerInfo,
        handler grpc.UnaryHandler,
    ) (interface{}, error) {
        start := time.Now()
        
        // Extract metadata
        md, _ := metadata.FromIncomingContext(ctx)
        traceID := md.Get("trace-id")
        
        log.Printf("gRPC call: %s, trace: %v", info.FullMethod, traceID)
        
        // Call handler
        resp, err := handler(ctx, req)
        
        // Record metrics
        duration := time.Since(start)
        grpcRequestDuration.WithLabelValues(
            info.FullMethod,
            grpcStatus(err),
        ).Observe(duration.Seconds())
        
        return resp, err
    }
}
```

**Stream server interceptor:**

```go
func StreamServerInterceptor() grpc.StreamServerInterceptor {
    return func(
        srv interface{},
        ss grpc.ServerStream,
        info *grpc.StreamServerInfo,
        handler grpc.StreamHandler,
    ) error {
        start := time.Now()
        
        // Wrap stream to count messages
        wrapped := &wrappedStream{
            ServerStream: ss,
            method:       info.FullMethod,
        }
        
        err := handler(srv, wrapped)
        
        duration := time.Since(start)
        log.Printf("Stream %s completed in %v, messages: %d",
            info.FullMethod, duration, wrapped.messageCount)
        
        return err
    }
}

type wrappedStream struct {
    grpc.ServerStream
    method       string
    messageCount int
}

func (w *wrappedStream) RecvMsg(m interface{}) error {
    err := w.ServerStream.RecvMsg(m)
    if err == nil {
        w.messageCount++
    }
    return err
}
```

**Client interceptor with retry:**

```go
func UnaryClientInterceptor() grpc.UnaryClientInterceptor {
    return func(
        ctx context.Context,
        method string,
        req, reply interface{},
        cc *grpc.ClientConn,
        invoker grpc.UnaryInvoker,
        opts ...grpc.CallOption,
    ) error {
        // Add metadata
        ctx = metadata.AppendToOutgoingContext(ctx,
            "trace-id", getTraceID(ctx))
        
        // Retry logic
        var err error
        for i := 0; i < 3; i++ {
            err = invoker(ctx, method, req, reply, cc, opts...)
            if err == nil || !isRetryable(err) {
                return err
            }
            time.Sleep(backoff(i))
        }
        return err
    }
}
```

**Suggested Reading:**
- gRPC-Go interceptor examples
- OpenTelemetry gRPC instrumentation
- "gRPC Up and Running" by Kasun Indrasiri
</details>

## Testing & Quality (Continued)

### Question 15: Testing Database Code
How do you write integration tests for database code in Go? What strategies ensure tests are fast, isolated, and reliable?

<details>
<summary>Answer</summary>

**Strategies:**

**1. Using testcontainers:**
```go
import "github.com/testcontainers/testcontainers-go/modules/postgres"

func TestUserRepository(t *testing.T) {
    ctx := context.Background()
    
    pgContainer, err := postgres.RunContainer(ctx,
        testcontainers.WithImage("postgres:15"),
        postgres.WithDatabase("testdb"),
        postgres.WithUsername("test"),
        postgres.WithPassword("test"),
    )
    require.NoError(t, err)
    defer pgContainer.Terminate(ctx)
    
    connStr, err := pgContainer.ConnectionString(ctx)
    require.NoError(t, err)
    
    db, err := sql.Open("postgres", connStr)
    require.NoError(t, err)
    defer db.Close()
    
    // Run migrations
    migrate(db)
    
    // Test repository
    repo := NewUserRepository(db)
    user := &User{Name: "Alice", Email: "alice@example.com"}
    err = repo.Create(ctx, user)
    assert.NoError(t, err)
    assert.NotZero(t, user.ID)
}
```

**2. Transaction rollback pattern:**
```go
func TestWithTx(t *testing.T) {
    tx, err := db.Begin()
    require.NoError(t, err)
    defer tx.Rollback() // Always rollback
    
    // Use transaction for test
    repo := NewUserRepository(tx)
    // ... test code
}
```

**3. Fixtures and cleanup:**
```go
func setupTest(t *testing.T) (*sql.DB, func()) {
    db := connectDB(t)
    
    // Setup fixtures
    seedTestData(t, db)
    
    cleanup := func() {
        cleanupTestData(t, db)
        db.Close()
    }
    
    return db, cleanup
}

func TestUser(t *testing.T) {
    db, cleanup := setupTest(t)
    defer cleanup()
    
    // Test code
}
```

**4. Using sqlmock for unit tests:**
```go
import "github.com/DATA-DOG/go-sqlmock"

func TestGetUser(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()
    
    rows := sqlmock.NewRows([]string{"id", "name", "email"}).
        AddRow(1, "Alice", "alice@example.com")
    
    mock.ExpectQuery("SELECT (.+) FROM users WHERE id = ?").
        WithArgs(1).
        WillReturnRows(rows)
    
    repo := NewUserRepository(db)
    user, err := repo.GetByID(context.Background(), 1)
    
    assert.NoError(t, err)
    assert.Equal(t, "Alice", user.Name)
    assert.NoError(t, mock.ExpectationsWereMet())
}
```

**Best practices:**
- Use separate test database
- Parallel test execution with `t.Parallel()` (careful with shared DB)
- Use test helpers to reduce boilerplate
- Clean up after each test
- Use database migrations in tests

**Suggested Reading:**
- testcontainers-go documentation
- "Testing database interactions in Go" blog posts
- sqlmock GitHub repository
- "Learn Go with Tests" by Chris James
</details>

### Question 16: Table-Driven Tests and Subtests
You need to test a complex validation function with many edge cases. How do you structure tests to be maintainable and provide clear failure messages?

<details>
<summary>Answer</summary>

**Table-driven test implementation:**

```go
func TestValidateEmail(t *testing.T) {
    tests := []struct {
        name    string
        email   string
        want    bool
        wantErr error
    }{
        {
            name:    "valid email",
            email:   "user@example.com",
            want:    true,
            wantErr: nil,
        },
        {
            name:    "missing @",
            email:   "userexample.com",
            want:    false,
            wantErr: ErrInvalidEmail,
        },
        {
            name:    "empty string",
            email:   "",
            want:    false,
            wantErr: ErrEmptyEmail,
        },
        {
            name:    "multiple @ symbols",
            email:   "user@@example.com",
            want:    false,
            wantErr: ErrInvalidEmail,
        },
        {
            name:    "no domain",
            email:   "user@",
            want:    false,
            wantErr: ErrInvalidEmail,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := ValidateEmail(tt.email)
            
            if tt.wantErr != nil {
                assert.ErrorIs(t, err, tt.wantErr)
            } else {
                assert.NoError(t, err)
            }
            
            assert.Equal(t, tt.want, got)
        })
    }
}
```

**Parallel execution:**
```go
func TestValidateEmailParallel(t *testing.T) {
    tests := []struct {
        name  string
        email string
        want  bool
    }{
        // ... test cases
    }
    
    for _, tt := range tests {
        tt := tt // Capture range variable
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel() // Run tests in parallel
            
            got, err := ValidateEmail(tt.email)
            assert.NoError(t, err)
            assert.Equal(t, tt.want, got)
        })
    }
}
```

**Complex scenario with setup/teardown:**
```go
func TestUserService(t *testing.T) {
    tests := []struct {
        name    string
        setup   func(t *testing.T, repo *UserRepository)
        input   *User
        want    *User
        wantErr bool
    }{
        {
            name: "create new user",
            setup: func(t *testing.T, repo *UserRepository) {
                // Setup preconditions
            },
            input: &User{Name: "Alice"},
            want:  &User{ID: 1, Name: "Alice"},
            wantErr: false,
        },
        {
            name: "duplicate email",
            setup: func(t *testing.T, repo *UserRepository) {
                repo.Create(context.Background(), &User{
                    Email: "test@example.com",
                })
            },
            input: &User{Email: "test@example.com"},
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            repo := setupTestRepo(t)
            if tt.setup != nil {
                tt.setup(t, repo)
            }
            
            got, err := repo.Create(context.Background(), tt.input)
            
            if tt.wantErr {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tt.want.Name, got.Name)
            }
        })
    }
}
```

**Benefits:**
- Clear test case names
- Easy to add new cases
- Isolated test execution
- Better failure messages
- Code reuse

**Suggested Reading:**
- "TableDrivenTests" Go Wiki
- "Advanced Testing in Go" by Mitchell Hashimoto
- Official Go testing documentation
</details>

### Question 17: Benchmarking and Performance Testing
You need to optimize a hot path in your code. How do you write benchmarks to measure performance improvements? What metrics matter?

<details>
<summary>Answer</summary>

**Basic benchmark:**

```go
func BenchmarkJSONMarshal(b *testing.B) {
    data := &User{
        ID:    1,
        Name:  "Alice",
        Email: "alice@example.com",
    }
    
    b.ResetTimer() // Reset timer after setup
    
    for i := 0; i < b.N; i++ {
        _, err := json.Marshal(data)
        if err != nil {
            b.Fatal(err)
        }
    }
}
```

**Memory allocation tracking:**

```go
func BenchmarkStringConcat(b *testing.B) {
    b.Run("using +", func(b *testing.B) {
        b.ReportAllocs() // Report memory allocations
        
        for i := 0; i < b.N; i++ {
            s := ""
            for j := 0; j < 100; j++ {
                s += "a"
            }
        }
    })
    
    b.Run("using strings.Builder", func(b *testing.B) {
        b.ReportAllocs()
        
        for i := 0; i < b.N; i++ {
            var sb strings.Builder
            for j := 0; j < 100; j++ {
                sb.WriteString("a")
            }
            _ = sb.String()
        }
    })
}
```

**Running benchmarks:**
```bash
# Run all benchmarks
go test -bench=. -benchmem

# Run specific benchmark
go test -bench=BenchmarkJSONMarshal -benchmem

# With CPU profiling
go test -bench=. -cpuprofile=cpu.prof

# With memory profiling
go test -bench=. -memprofile=mem.prof

# Compare benchmarks
go test -bench=. -benchmem > old.txt
# make changes
go test -bench=. -benchmem > new.txt
benchstat old.txt new.txt
```

**Sub-benchmarks for comparison:**

```go
func BenchmarkCache(b *testing.B) {
    sizes := []int{100, 1000, 10000}
    
    for _, size := range sizes {
        b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
            cache := NewCache(size)
            
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                key := fmt.Sprintf("key%d", i%size)
                cache.Get(key)
            }
        })
    }
}
```

**Parallel benchmarks:**

```go
func BenchmarkCacheParallel(b *testing.B) {
    cache := NewCache(1000)
    
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            key := fmt.Sprintf("key%d", rand.Intn(1000))
            cache.Get(key)
        }
    })
}
```

**Key metrics:**
- **ns/op**: Nanoseconds per operation (lower is better)
- **B/op**: Bytes allocated per operation
- **allocs/op**: Number of allocations per operation
- **MB/s**: Throughput for data processing benchmarks

**Optimization tips:**
- Reduce allocations (use sync.Pool, preallocate slices)
- Avoid interface conversions in hot paths
- Use benchcmp/benchstat for comparison
- Profile before optimizing (pprof)

**Suggested Reading:**
- "Profiling Go Programs" by Go Blog
- Dave Cheney's "High Performance Go Workshop"
- benchstat tool documentation
- "Writing and Optimizing Go Code" presentations
</details>

## Production Issues & Monitoring

### Question 18: Panic Recovery and Error Handling
Your production service occasionally panics, causing pod restarts. How do you implement panic recovery without hiding bugs? What's your strategy for error handling in microservices?

<details>
<summary>Answer</summary>

**Panic recovery middleware:**

```go
func RecoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                // Get stack trace
                stack := debug.Stack()
                
                // Log panic with context
                log.Printf("PANIC: %v\nStack:\n%s", err, stack)
                
                // Send to error tracking (Sentry, etc.)
                captureException(fmt.Errorf("panic: %v", err), stack)
                
                // Increment panic metric
                panicCounter.Inc()
                
                // Return 500 to client
                http.Error(w, "Internal Server Error", 
                    http.StatusInternalServerError)
            }
        }()
        
        next.ServeHTTP(w, r)
    })
}
```

**Goroutine panic recovery:**

```go
func SafeGo(fn func()) {
    go func() {
        defer func() {
            if err := recover(); err != nil {
                stack := debug.Stack()
                log.Printf("Goroutine panic: %v\n%s", err, stack)
                
                // Don't hide the panic - alert on-call
                alertOnCall("goroutine_panic", fmt.Sprintf("%v", err))
            }
        }()
        
        fn()
    }()
}

// Usage
SafeGo(func() {
    processJob(job)
})
```

**Error handling strategy:**

```go
// Custom error types with context
type AppError struct {
    Code       string
    Message    string
    Err        error
    StatusCode int
    Fields     map[string]interface{}
}

func (e *AppError) Error() string {
    return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
    return e.Err
}

// Error categories
var (
    ErrNotFound = &AppError{
        Code:       "NOT_FOUND",
        Message:    "Resource not found",
        StatusCode: http.StatusNotFound,
    }
    
    ErrValidation = &AppError{
        Code:       "VALIDATION_ERROR",
        Message:    "Validation failed",
        StatusCode: http.StatusBadRequest,
    }
)

// Error wrapping with context
func GetUser(ctx context.Context, id int) (*User, error) {
    user, err := repo.FindByID(ctx, id)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, fmt.Errorf("%w: user_id=%d", ErrNotFound, id)
        }
        return nil, fmt.Errorf("failed to get user: %w", err)
    }
    return user, nil
}
```

**Centralized error handling:**

```go
func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
    // Unwrap to find AppError
    var appErr *AppError
    if errors.As(err, &appErr) {
        respondJSON(w, appErr.StatusCode, map[string]interface{}{
            "error": appErr.Code,
            "message": appErr.Message,
            "fields": appErr.Fields,
        })
        return
    }
    
    // Unknown error - log and return 500
    log.Printf("Unexpected error: %v", err)
    http.Error(w, "Internal Server Error", 
        http.StatusInternalServerError)
}
```

**When NOT to recover:**
- In main goroutine (let it crash and restart)
- During initialization (fail fast)
- In tests (panics indicate bugs)

**Best practices:**
- Log panics with full context
- Alert on-call for production panics
- Use error wrapping with `%w`
- Return errors, don't panic (except for programmer errors)
- Use `errors.Is` and `errors.As` for error checking

**Suggested Reading:**
- "Error handling and Go" by Go Blog
- "Working with Errors in Go 1.13" by Go Blog
- Dave Cheney's "Don't just check errors, handle them gracefully"
- pkg/errors documentation
</details>

### Question 19: Structured Logging and Observability
Design a logging strategy for a microservices system. How do you correlate logs across services? What log levels and structured fields should you include?

<details>
<summary>Answer</summary>

**Structured logging with zerolog:**

```go
import "github.com/rs/zerolog/log"

func init() {
    // Production: JSON format
    // Development: Console format
    if os.Getenv("ENV") == "production" {
        zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
    } else {
        log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
    }
}

func ProcessOrder(ctx context.Context, orderID string) error {
    logger := log.With().
        Str("order_id", orderID).
        Str("trace_id", getTraceID(ctx)).
        Str("user_id", getUserID(ctx)).
        Logger()
    
    logger.Info().Msg("Processing order")
    
    if err := validateOrder(order); err != nil {
        logger.Error().
            Err(err).
            Str("validation_step", "payment").
            Msg("Order validation failed")
        return err
    }
    
    logger.Info().
        Dur("duration_ms", time.Since(start)).
        Msg("Order processed successfully")
    
    return nil
}
```

**Context-based logging:**

```go
type contextKey string

const (
    loggerKey   contextKey = "logger"
    traceIDKey  contextKey = "trace_id"
    requestIDKey contextKey = "request_id"
)

// Middleware to add logger to context
func LoggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        requestID := generateRequestID()
        traceID := r.Header.Get("X-Trace-ID")
        if traceID == "" {
            traceID = generateTraceID()
        }
        
        logger := log.With().
            Str("request_id", requestID).
            Str("trace_id", traceID).
            Str("method", r.Method).
            Str("path", r.URL.Path).
            Str("remote_addr", r.RemoteAddr).
            Logger()
        
        ctx := context.WithValue(r.Context(), loggerKey, &logger)
        ctx = context.WithValue(ctx, traceIDKey, traceID)
        
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Get logger from context
func GetLogger(ctx context.Context) *zerolog.Logger {
    if logger, ok := ctx.Value(loggerKey).(*zerolog.Logger); ok {
        return logger
    }
    return &log.Logger
}

// Usage in handlers
func HandleOrder(w http.ResponseWriter, r *http.Request) {
    logger := GetLogger(r.Context())
    logger.Info().Msg("Handling order request")
}
```

**Log levels strategy:**

```go
// ERROR: Something failed, requires attention
logger.Error().
    Err(err).
    Str("user_id", userID).
    Msg("Failed to process payment")

// WARN: Something unexpected but handled
logger.Warn().
    Str("cache_key", key).
    Msg("Cache miss, fetching from database")

// INFO: Important business events
logger.Info().
    Str("order_id", orderID).
    Float64("amount", amount).
    Msg("Order created")

// DEBUG: Detailed diagnostic info (disabled in prod)
logger.Debug().
    Interface("request", req).
    Msg("Received request")
```

**Correlation across services:**

```go
// Propagate trace ID in HTTP headers
client := &http.Client{}
req, _ := http.NewRequestWithContext(ctx, "POST", url, body)

traceID := ctx.Value(traceIDKey).(string)
req.Header.Set("X-Trace-ID", traceID)
req.Header.Set("X-Request-ID", requestID)

resp, err := client.Do(req)
```

**Essential log fields:**
- `timestamp`: ISO8601 or Unix timestamp
- `level`: ERROR, WARN, INFO, DEBUG
- `service_name`: Service identifier
- `trace_id`: Distributed tracing ID
- `request_id`: Unique per request
- `user_id`: User context
- `error`: Error message and stack trace
- `duration_ms`: Operation duration
- `environment`: prod, staging, dev

**Sampling for high-volume logs:**

```go
sampler := &zerolog.LevelSampler{
    DebugSampler: &zerolog.BurstSampler{
        Burst:       5,
        Period:      time.Second,
        NextSampler: &zerolog.BasicSampler{N: 100},
    },
}

logger := zerolog.New(os.Stdout).
    Sample(sampler).
    With().
    Timestamp().
    Logger()
```

**Suggested Reading:**
- "The Twelve-Factor App" (Logs section)
- zerolog/zap documentation
- "Observability Engineering" by Charity Majors
- OpenTelemetry logging specification
</details>

### Question 20: Metrics and Alerting
What metrics would you expose for a typical HTTP/gRPC service? How do you instrument your code and what alerts would you set up?

<details>
<summary>Answer</summary>

**Prometheus metrics implementation:**

```go
import "github.com/prometheus/client_golang/prometheus"

var (
    // HTTP metrics
    httpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "path", "status"},
    )
    
    httpRequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "HTTP request duration in seconds",
            Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
        },
        []string{"method", "path", "status"},
    )
    
    httpRequestSize = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_size_bytes",
            Help:    "HTTP request size in bytes",
            Buckets: prometheus.ExponentialBuckets(100, 10, 8),
        },
        []string{"method", "path"},
    )
    
    // Business metrics
    ordersCreated = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "orders_created_total",
            Help: "Total number of orders created",
        },
        []string{"status"},
    )
    
    orderValue = prometheus.NewHistogram(
        prometheus.HistogramOpts{
            Name:    "order_value_usd",
            Help:    "Order value in USD",
            Buckets: []float64{10, 50, 100, 500, 1000, 5000, 10000},
        },
    )
    
    // Database metrics
    dbConnectionsInUse = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Name: "db_connections_in_use",
            Help: "Number of database connections in use",
        },
    )
    
    dbQueryDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "db_query_duration_seconds",
            Help:    "Database query duration",
            Buckets: prometheus.DefBuckets,
        },
        []string{"query_type"},
    )
)

func init() {
    prometheus.MustRegister(
        httpRequestsTotal,
        httpRequestDuration,
        httpRequestSize,
        ordersCreated,
        orderValue,
        dbConnectionsInUse,
        dbQueryDuration,
    )
}
```

**Instrumentation middleware:**

```go
func MetricsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Wrap response writer to capture status code
        wrapped := &responseWriter{
            ResponseWriter: w,
            statusCode:     http.StatusOK,
        }
        
        next.ServeHTTP(wrapped, r)
        
        duration := time.Since(start).Seconds()
        status := fmt.Sprintf("%d", wrapped.statusCode)
        
        httpRequestsTotal.WithLabelValues(
            r.Method,
            r.URL.Path,
            status,
        ).Inc()
        
        httpRequestDuration.WithLabelValues(
            r.Method,
            r.URL.Path,
            status,
        ).Observe(duration)
        
        httpRequestSize.WithLabelValues(
            r.Method,
            r.URL.Path,
        ).Observe(float64(r.ContentLength))
    })
}

type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}
```

**Database metrics collection:**

```go
func collectDBMetrics(db *sql.DB) {
    go func() {
        ticker := time.NewTicker(15 * time.Second)
        defer ticker.Stop()
        
        for range ticker.C {
            stats := db.Stats()
            dbConnectionsInUse.Set(float64(stats.InUse))
            dbConnectionsIdle.Set(float64(stats.Idle))
            dbConnectionsWaitCount.Set(float64(stats.WaitCount))
            dbConnectionsWaitDuration.Set(stats.WaitDuration.Seconds())
        }
    }()
}
```

**Essential metrics categories:**

**1. RED Method (for services):**
- Rate: Requests per second
- Errors: Error rate
- Duration: Latency percentiles (p50, p95, p99)

**2. USE Method (for resources):**
- Utilization: % time resource is busy
- Saturation: Queue depth
- Errors: Error count

**3. Golden Signals (Google SRE):**
- Latency
- Traffic
- Errors
- Saturation

**Alert examples (Prometheus AlertManager):**

```yaml
groups:
  - name: service_alerts
    interval: 30s
    rules:
      # High error rate
      - alert: HighErrorRate
        expr: |
          (sum(rate(http_requests_total{status=~"5.."}[5m])) 
           / sum(rate(http_requests_total[5m]))) > 0.05
        for: 5m
        labels:
          severity: critical
        annotations:
          summary: "High error rate ({{ $value }}%)"
          
      # High latency
      - alert: HighLatency
        expr: |
          histogram_quantile(0.95, 
            rate(http_request_duration_seconds_bucket[5m])) > 1
        for: 10m
        labels:
          severity: warning
        annotations:
          summary: "P95 latency above 1s"
          
      # Database connection exhaustion
      - alert: DBConnectionPoolExhausted
        expr: db_connections_in_use / db_connections_max > 0.9
        for: 5m
        labels:
          severity: warning
          
      # Service down
      - alert: ServiceDown
        expr: up{job="my-service"} == 0
        for: 1m
        labels:
          severity: critical
```

**Suggested Reading:**
- "Site Reliability Engineering" by Google (Monitoring chapter)
- Prometheus best practices documentation
- "The Four Golden Signals" by Google SRE
- RED method documentation
</details>

## Testing & Quality (Continued)

### Question 15: Testing Database Code
How do you write integration tests for database code in Go? What strategies ensure tests are fast, isolated, and reliable?

<details>
<summary>Answer</summary>

**Strategies:**

**1. Using testcontainers:**
```go
import "github.com/testcontainers/testcontainers-go/modules/postgres"

func TestUserRepository(t *testing.T) {
    ctx := context.Background()
    
    pgContainer, err := postgres.RunContainer(ctx,
        testcontainers.WithImage("postgres:15"),
        postgres.WithDatabase("testdb"),
        postgres.WithUsername("test"),
        postgres.WithPassword("test"),
    )
    require.NoError(t, err)
    defer pgContainer.Terminate(ctx)
    
    connStr, err := pgContainer.ConnectionString(ctx)
    require.NoError(t, err)
    
    db, err := sql.Open("postgres", connStr)
    require.NoError(t, err)
    defer db.Close()
    
    // Run migrations
    migrate(db)
    
    // Test repository
    repo := NewUserRepository(db)
    user := &User{Name: "Alice", Email: "alice@example.com"}
    err = repo.Create(ctx, user)
    assert.NoError(t, err)
    assert.NotZero(t, user.ID)
}
```

**2. Transaction rollback pattern:**
```go
func TestWithTx(t *testing.T) {
    tx, err := db.Begin()
    require.NoError(t, err)
    defer tx.Rollback() // Always rollback
    
    // Use transaction for test
    repo := NewUserRepository(tx)
    // ... test code
}
```

**3. Fixtures and cleanup:**
```go
func setupTest(t *testing.T) (*sql.DB, func()) {
    db := connectDB(t)
    
    // Setup fixtures
    seedTestData(t, db)
    
    cleanup := func() {
        cleanupTestData(t, db)
        db.Close()
    }
    
    return db, cleanup
}

func TestUser(t *testing.T) {
    db, cleanup := setupTest(t)
    defer cleanup()
    
    // Test code
}
```

**4. Using sqlmock for unit tests:**
```go
import "github.com/DATA-DOG/go-sqlmock"

func TestGetUser(t *testing.T) {
    db, mock, err := sqlmock.New()
    require.NoError(t, err)
    defer db.Close()
    
    rows := sqlmock.NewRows([]string{"id", "name", "email"}).
        AddRow(1, "Alice", "alice@example.com")
    
    mock.ExpectQuery("SELECT (.+) FROM users WHERE id = ?").
        WithArgs(1).
        WillReturnRows(rows)
    
    repo := NewUserRepository(db)
    user, err := repo.GetByID(context.Background(), 1)
    
    assert.NoError(t, err)
    assert.Equal(t, "Alice", user.Name)
    assert.NoError(t, mock.ExpectationsWereMet())
}
```

**Best practices:**
- Use separate test database
- Parallel test execution with `t.Parallel()` (careful with shared DB)
- Use test helpers to reduce boilerplate
- Clean up after each test
- Use database migrations in tests

**Suggested Reading:**
- testcontainers-go documentation
- "Testing database interactions in Go" blog posts
- sqlmock GitHub repository
- "Learn Go with Tests" by Chris James
</details>

### Question 16: Table-Driven Tests and Subtests
You need to test a complex validation function with many edge cases. How do you structure tests to be maintainable and provide clear failure messages?

<details>
<summary>Answer</summary>

**Table-driven test implementation:**

```go
func TestValidateEmail(t *testing.T) {
    tests := []struct {
        name    string
        email   string
        want    bool
        wantErr error
    }{
        {
            name:    "valid email",
            email:   "user@example.com",
            want:    true,
            wantErr: nil,
        },
        {
            name:    "missing @",
            email:   "userexample.com",
            want:    false,
            wantErr: ErrInvalidEmail,
        },
        {
            name:    "empty string",
            email:   "",
            want:    false,
            wantErr: ErrEmptyEmail,
        },
        {
            name:    "multiple @ symbols",
            email:   "user@@example.com",
            want:    false,
            wantErr: ErrInvalidEmail,
        },
        {
            name:    "no domain",
            email:   "user@",
            want:    false,
            wantErr: ErrInvalidEmail,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got, err := ValidateEmail(tt.email)
            
            if tt.wantErr != nil {
                assert.ErrorIs(t, err, tt.wantErr)
            } else {
                assert.NoError(t, err)
            }
            
            assert.Equal(t, tt.want, got)
        })
    }
}
```

**Parallel execution:**
```go
func TestValidateEmailParallel(t *testing.T) {
    tests := []struct {
        name  string
        email string
        want  bool
    }{
        // ... test cases
    }
    
    for _, tt := range tests {
        tt := tt // Capture range variable
        t.Run(tt.name, func(t *testing.T) {
            t.Parallel() // Run tests in parallel
            
            got, err := ValidateEmail(tt.email)
            assert.NoError(t, err)
            assert.Equal(t, tt.want, got)
        })
    }
}
```

**Complex scenario with setup/teardown:**
```go
func TestUserService(t *testing.T) {
    tests := []struct {
        name    string
        setup   func(t *testing.T, repo *UserRepository)
        input   *User
        want    *User
        wantErr bool
    }{
        {
            name: "create new user",
            setup: func(t *testing.T, repo *UserRepository) {
                // Setup preconditions
            },
            input: &User{Name: "Alice"},
            want:  &User{ID: 1, Name: "Alice"},
            wantErr: false,
        },
        {
            name: "duplicate email",
            setup: func(t *testing.T, repo *UserRepository) {
                repo.Create(context.Background(), &User{
                    Email: "test@example.com",
                })
            },
            input: &User{Email: "test@example.com"},
            wantErr: true,
        },
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            repo := setupTestRepo(t)
            if tt.setup != nil {
                tt.setup(t, repo)
            }
            
            got, err := repo.Create(context.Background(), tt.input)
            
            if tt.wantErr {
                assert.Error(t, err)
            } else {
                assert.NoError(t, err)
                assert.Equal(t, tt.want.Name, got.Name)
            }
        })
    }
}
```

**Benefits:**
- Clear test case names
- Easy to add new cases
- Isolated test execution
- Better failure messages
- Code reuse

**Suggested Reading:**
- "TableDrivenTests" Go Wiki
- "Advanced Testing in Go" by Mitchell Hashimoto
- Official Go testing documentation
</details>

### Question 17: Benchmarking and Performance Testing
You need to optimize a hot path in your code. How do you write benchmarks to measure performance improvements? What metrics matter?

<details>
<summary>Answer</summary>

**Basic benchmark:**

```go
func BenchmarkJSONMarshal(b *testing.B) {
    data := &User{
        ID:    1,
        Name:  "Alice",
        Email: "alice@example.com",
    }
    
    b.ResetTimer() // Reset timer after setup
    
    for i := 0; i < b.N; i++ {
        _, err := json.Marshal(data)
        if err != nil {
            b.Fatal(err)
        }
    }
}
```

**Memory allocation tracking:**

```go
func BenchmarkStringConcat(b *testing.B) {
    b.Run("using +", func(b *testing.B) {
        b.ReportAllocs() // Report memory allocations
        
        for i := 0; i < b.N; i++ {
            s := ""
            for j := 0; j < 100; j++ {
                s += "a"
            }
        }
    })
    
    b.Run("using strings.Builder", func(b *testing.B) {
        b.ReportAllocs()
        
        for i := 0; i < b.N; i++ {
            var sb strings.Builder
            for j := 0; j < 100; j++ {
                sb.WriteString("a")
            }
            _ = sb.String()
        }
    })
}
```

**Running benchmarks:**
```bash
# Run all benchmarks
go test -bench=. -benchmem

# Run specific benchmark
go test -bench=BenchmarkJSONMarshal -benchmem

# With CPU profiling
go test -bench=. -cpuprofile=cpu.prof

# With memory profiling
go test -bench=. -memprofile=mem.prof

# Compare benchmarks
go test -bench=. -benchmem > old.txt
# make changes
go test -bench=. -benchmem > new.txt
benchstat old.txt new.txt
```

**Sub-benchmarks for comparison:**

```go
func BenchmarkCache(b *testing.B) {
    sizes := []int{100, 1000, 10000}
    
    for _, size := range sizes {
        b.Run(fmt.Sprintf("size=%d", size), func(b *testing.B) {
            cache := NewCache(size)
            
            b.ResetTimer()
            for i := 0; i < b.N; i++ {
                key := fmt.Sprintf("key%d", i%size)
                cache.Get(key)
            }
        })
    }
}
```

**Parallel benchmarks:**

```go
func BenchmarkCacheParallel(b *testing.B) {
    cache := NewCache(1000)
    
    b.RunParallel(func(pb *testing.PB) {
        for pb.Next() {
            key := fmt.Sprintf("key%d", rand.Intn(1000))
            cache.Get(key)
        }
    })
}
```

**Key metrics:**
- **ns/op**: Nanoseconds per operation (lower is better)
- **B/op**: Bytes allocated per operation
- **allocs/op**: Number of allocations per operation
- **MB/s**: Throughput for data processing benchmarks

**Optimization tips:**
- Reduce allocations (use sync.Pool, preallocate slices)
- Avoid interface conversions in hot paths
- Use benchcmp/benchstat for comparison
- Profile before optimizing (pprof)

**Suggested Reading:**
- "Profiling Go Programs" by Go Blog
- Dave Cheney's "High Performance Go Workshop"
- benchstat tool documentation
- "Writing and Optimizing Go Code" presentations
</details>

## Production Issues & Monitoring

### Question 18: Panic Recovery and Error Handling
Your production service occasionally panics, causing pod restarts. How do you implement panic recovery without hiding bugs? What's your strategy for error handling in microservices?

<details>
<summary>Answer</summary>

**Panic recovery middleware:**

```go
func RecoveryMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        defer func() {
            if err := recover(); err != nil {
                // Get stack trace
                stack := debug.Stack()
                
                // Log panic with context
                log.Printf("PANIC: %v\nStack:\n%s", err, stack)
                
                // Send to error tracking (Sentry, etc.)
                captureException(fmt.Errorf("panic: %v", err), stack)
                
                // Increment panic metric
                panicCounter.Inc()
                
                // Return 500 to client
                http.Error(w, "Internal Server Error", 
                    http.StatusInternalServerError)
            }
        }()
        
        next.ServeHTTP(w, r)
    })
}
```

**Goroutine panic recovery:**

```go
func SafeGo(fn func()) {
    go func() {
        defer func() {
            if err := recover(); err != nil {
                stack := debug.Stack()
                log.Printf("Goroutine panic: %v\n%s", err, stack)
                
                // Don't hide the panic - alert on-call
                alertOnCall("goroutine_panic", fmt.Sprintf("%v", err))
            }
        }()
        
        fn()
    }()
}

// Usage
SafeGo(func() {
    processJob(job)
})
```

**Error handling strategy:**

```go
// Custom error types with context
type AppError struct {
    Code       string
    Message    string
    Err        error
    StatusCode int
    Fields     map[string]interface{}
}

func (e *AppError) Error() string {
    return fmt.Sprintf("%s: %s", e.Code, e.Message)
}

func (e *AppError) Unwrap() error {
    return e.Err
}

// Error categories
var (
    ErrNotFound = &AppError{
        Code:       "NOT_FOUND",
        Message:    "Resource not found",
        StatusCode: http.StatusNotFound,
    }
    
    ErrValidation = &AppError{
        Code:       "VALIDATION_ERROR",
        Message:    "Validation failed",
        StatusCode: http.StatusBadRequest,
    }
)

// Error wrapping with context
func GetUser(ctx context.Context, id int) (*User, error) {
    user, err := repo.FindByID(ctx, id)
    if err != nil {
        if errors.Is(err, sql.ErrNoRows) {
            return nil, fmt.Errorf("%w: user_id=%d", ErrNotFound, id)
        }
        return nil, fmt.Errorf("failed to get user: %w", err)
    }
    return user, nil
}
```

**Centralized error handling:**

```go
func ErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
    // Unwrap to find AppError
    var appErr *AppError
    if errors.As(err, &appErr) {
        respondJSON(w, appErr.StatusCode, map[string]interface{}{
            "error": appErr.Code,
            "message": appErr.Message,
            "fields": appErr.Fields,
        })
        return
    }
    
    // Unknown error - log and return 500
    log.Printf("Unexpected error: %v", err)
    http.Error(w, "Internal Server Error", 
        http.StatusInternalServerError)
}
```

**When NOT to recover:**
- In main goroutine (let it crash and restart)
- During initialization (fail fast)
- In tests (panics indicate bugs)

**Best practices:**
- Log panics with full context
- Alert on-call for production panics
- Use error wrapping with `%w`
- Return errors, don't panic (except for programmer errors)
- Use `errors.Is` and `errors.As` for error checking

**Suggested Reading:**
- "Error handling and Go" by Go Blog
- "Working with Errors in Go 1.13" by Go Blog
- Dave Cheney's "Don't just check errors, handle them gracefully"
- pkg/errors documentation
</details>

### Question 19: Structured Logging and Observability
Design a logging strategy for a microservices system. How do you correlate logs across services? What log levels and structured fields should you include?

<details>
<summary>Answer</summary>

**Structured logging with zerolog:**

```go
import "github.com/rs/zerolog/log"

func init() {
    // Production: JSON format
    // Development: Console format
    if os.Getenv("ENV") == "production" {
        zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
    } else {
        log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
    }
}

func ProcessOrder(ctx context.Context, orderID string) error {
    logger := log.With().
        Str("order_id", orderID).
        Str("trace_id", getTraceID(ctx)).
        Str("user_id", getUserID(ctx)).
        Logger()
    
    logger.Info().Msg("Processing order")
    
    if err := validateOrder(order); err != nil {
        logger.Error().
            Err(err).
            Str("validation_step", "payment").
            Msg("Order validation failed")
        return err
    }
    
    logger.Info().
        Dur("duration_ms", time.Since(start)).
        Msg("Order processed successfully")
    
    return nil
}
```

**Context-based logging:**

```go
type contextKey string

const (
    loggerKey   contextKey = "logger"
    traceIDKey  contextKey = "trace_id"
    requestIDKey contextKey = "request_id"
)

// Middleware to add logger to context
func LoggerMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        requestID := generateRequestID()
        traceID := r.Header.Get("X-Trace-ID")
        if traceID == "" {
            traceID = generateTraceID()
        }
        
        logger := log.With().
            Str("request_id", requestID).
            Str("trace_id", traceID).
            Str("method", r.Method).
            Str("path", r.URL.Path).
            Str("remote_addr", r.RemoteAddr).
            Logger()
        
        ctx := context.WithValue(r.Context(), loggerKey, &logger)
        ctx = context.WithValue(ctx, traceIDKey, traceID)
        
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Get logger from context
func GetLogger(ctx context.Context) *zerolog.Logger {
    if logger, ok := ctx.Value(loggerKey).(*zerolog.Logger); ok {
        return logger
    }
    return &log.Logger
}

// Usage in handlers
func HandleOrder(w http.ResponseWriter, r *http.Request) {
    logger := GetLogger(r.Context())
    logger.Info().Msg("Handling order request")
}
```

**Log levels strategy:**

```go
// ERROR: Something failed, requires attention
logger.Error().
    Err(err).
    Str("user_id", userID).
    Msg("Failed to process payment")

// WARN: Something unexpected but handled
logger.Warn().
    Str("cache_key", key).
    Msg("Cache miss, fetching from database")

// INFO: Important business events
logger.Info().
    Str("order_id", orderID).
    Float64("amount", amount).
    Msg("Order created")

// DEBUG: Detailed diagnostic info (disabled in prod)
logger.Debug().
    Interface("request", req).
    Msg("Received request")
```

**Correlation across services:**

```go
// Propagate trace ID in HTTP headers
client := &http.Client{}
req, _ := http.NewRequestWithContext(ctx, "POST", url, body)

traceID := ctx.Value(traceIDKey).(string)
req.Header.Set("X-Trace-ID", traceID)
req.Header.Set("X-Request-ID", requestID)

resp, err := client.Do(req)
```

**Essential log fields:**
- `timestamp`: ISO8601 or Unix timestamp
- `level`: ERROR, WARN, INFO, DEBUG
- `service_name`: Service identifier
- `trace_id`: Distributed tracing ID
- `request_id`: Unique per request
- `user_id`: User context
- `error`: Error message and stack trace
- `duration_ms`: Operation duration
- `environment`: prod, staging, dev

**Sampling for high-volume logs:**

```go
sampler := &zerolog.LevelSampler{
    DebugSampler: &zerolog.BurstSampler{
        Burst:       5,
        Period:      time.Second,
        NextSampler: &zerolog.BasicSampler{N: 100},
    },
}

logger := zerolog.New(os.Stdout).
    Sample(sampler).
    With().
    Timestamp().
    Logger()
```

**Suggested Reading:**
- "The Twelve-Factor App" (Logs section)
- zerolog/zap documentation
- "Observability Engineering" by Charity Majors
- OpenTelemetry logging specification
</details>

### Question 20: Metrics and Alerting
What metrics would you expose for a typical HTTP/gRPC service? How do you instrument your code and what alerts would you set up?

<details>
<summary>Answer</summary>

**Prometheus metrics implementation:**

```go
import "github.com/prometheus/client_golang/prometheus"

var (
    // HTTP metrics
    httpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests",
        },
        []string{"method", "path", "status"},
    )
    
    httpRequestDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_duration_seconds",
            Help:    "HTTP request duration in seconds",
            Buckets: []float64{.005, .01, .025, .05, .1, .25, .5, 1, 2.5, 5, 10},
        },
        []string{"method", "path", "status"},
    )
    
    httpRequestSize = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "http_request_size_bytes",
            Help:    "HTTP request size in bytes",
            Buckets: prometheus.ExponentialBuckets(100, 10, 8),
        },
        []string{"method", "path"},
    )
    
    // Business metrics
    ordersCreated = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "orders_created_total",
            Help: "Total number of orders created",
        },
        []string{"status"},
    )
    
    orderValue = prometheus.NewHistogram(
        prometheus.HistogramOpts{
            Name:    "order_value_usd",
            Help:    "Order value in USD",
            Buckets: []float64{10, 50, 100, 500, 1000, 5000, 10000},
        },
    )
    
    // Database metrics
    dbConnectionsInUse = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Name: "db_connections_in_use",
            Help: "Number of database connections in use",
        },
    )
    
    dbQueryDuration = prometheus.NewHistogramVec(
        prometheus.HistogramOpts{
            Name:    "db_query_duration_seconds",
            Help:    "Database query duration",
            Buckets: prometheus.DefBuckets,
        },
        []string{"query_type"},
    )
)

func init() {
    prometheus.MustRegister(
        httpRequestsTotal,
        httpRequestDuration,
        httpRequestSize,
        ordersCreated,
        orderValue,
        dbConnectionsInUse,
        dbQueryDuration,
    )
}
```

**Instrumentation middleware:**

```go
func MetricsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        
        // Wrap response writer to capture status code
        wrapped := &responseWriter{
            ResponseWriter: w,
            statusCode:     http.StatusOK,
        }
        
        next.ServeHTTP(wrapped, r)
        
        duration := time.Since(start).Seconds()
        status := fmt.Sprintf("%d", wrapped.statusCode)
        
        httpRequestsTotal.WithLabelValues(
            r.Method,
            r.URL.Path,
            status,
        ).Inc()
        
        httpRequestDuration.WithLabelValues(
            r.Method,
            r.URL.Path,
            status,
        ).Observe(duration)
        
        httpRequestSize.WithLabelValues(
            r.Method,
            r.URL.Path,
        ).Observe(float64(r.ContentLength))
    })
}

type responseWriter struct {
    http.ResponseWriter
    statusCode int
}

func (rw *responseWriter) WriteHeader(code int) {
    rw.statusCode = code
    rw.ResponseWriter.WriteHeader(code)
}
```

**Database metrics collection:**

```go
func collectDBMetrics(db *sql.DB) {
    go func() {
        ticker := time.NewTicker(15 * time.Second)
        defer ticker.Stop()
        
        for range ticker.C {
            stats := db.Stats()
            dbConnectionsInUse.Set(float64(stats.InUse))
            dbConnectionsIdle.Set(float64(stats.Idle))
            dbConnectionsWaitCount.Set(float64(stats.WaitCount))
            dbConnectionsWaitDuration.Set(stats.WaitDuration.Seconds())
        }
    }()
}
```

**Essential metrics categories:**

**1. RED Method (for services):**
- Rate: Requests per second
- Errors: Error rate
- Duration: Latency percentiles (p50, p95, p99)

**2. USE Method (for resources):**
- Utilization: % time resource is busy
- Saturation: Queue depth
- Errors: Error count

**3. Golden Signals (Google SRE):**
- Latency
- Traffic
- Errors
- Saturation

**Alert examples (Prometheus AlertManager):**

```yaml
groups:
  - name: service_alerts
    interval: 30s
    rules:
      # High error rate
      - alert: HighErrorRate
        expr: |
          (sum(rate(http_requests_total{status=~"5.."}[5m])) 
           / sum(rate(http_requests_total[5m]))) > 0.05
        for: 5m
        labels:
          severity: critical
        annotations:
          summary: "High error rate ({{ $value }}%)"
          
      # High latency
      - alert: HighLatency
        expr: |
          histogram_quantile(0.95, 
            rate(http_request_duration_seconds_bucket[5m])) > 1
        for: 10m
        labels:
          severity: warning
        annotations:
          summary: "P95 latency above 1s"
          
      # Database connection exhaustion
      - alert: DBConnectionPoolExhausted
        expr: db_connections_in_use / db_connections_max > 0.9
        for: 5m
        labels:
          severity: warning
          
      # Service down
      - alert: ServiceDown
        expr: up{job="my-service"} == 0
        for: 1m
        labels:
          severity: critical
```

**Suggested Reading:**
- "Site Reliability Engineering" by Google (Monitoring chapter)
- Prometheus best practices documentation
- "The Four Golden Signals" by Google SRE
- RED method documentation
</details>

### Question 21: Zero-Downtime Deployments
How do you implement zero-downtime deployments for a Go service in Kubernetes? What health checks and readiness probes do you configure?

<details>
<summary>Answer</summary>

**Health check endpoints:**

```go
type HealthChecker struct {
    db          *sql.DB
    redis       *redis.Client
    startTime   time.Time
    isReady     atomic.Bool
}

func NewHealthChecker(db *sql.DB, redis *redis.Client) *HealthChecker {
    hc := &HealthChecker{
        db:        db,
        redis:     redis,
        startTime: time.Now(),
    }
    hc.isReady.Store(false)
    return hc
}

// Liveness probe - is the service alive?
func (h *HealthChecker) LivenessHandler(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "status": "alive",
        "uptime": time.Since(h.startTime).Seconds(),
    })
}

// Readiness probe - can the service handle traffic?
func (h *HealthChecker) ReadinessHandler(w http.ResponseWriter, r *http.Request) {
    if !h.isReady.Load() {
        http.Error(w, "Not ready", http.StatusServiceUnavailable)
        return
    }
    
    ctx, cancel := context.WithTimeout(r.Context(), 2*time.Second)
    defer cancel()
    
    health := map[string]string{
        "database": "healthy",
        "redis":    "healthy",
    }
    
    // Check database
    if err := h.db.PingContext(ctx); err != nil {
        health["database"] = "unhealthy: " + err.Error()
        w.WriteHeader(http.StatusServiceUnavailable)
    }
    
    // Check Redis
    if err := h.redis.Ping(ctx).Err(); err != nil {
        health["redis"] = "unhealthy: " + err.Error()
        w.WriteHeader(http.StatusServiceUnavailable)
    }
    
    if w.(*responseWriter).statusCode == 0 {
        w.WriteHeader(http.StatusOK)
    }
    
    json.NewEncoder(w).Encode(map[string]interface{}{
        "status":     "ready",
        "components": health,
    })
}

// Mark service as ready after warmup
func (h *HealthChecker) MarkReady() {
    h.isReady.Store(true)
}
```

**Graceful shutdown with connection draining:**

```go
func main() {
    srv := &http.Server{
        Addr:    ":8080",
        Handler: router,
    }
    
    healthChecker := NewHealthChecker(db, redis)
    
    // Warmup phase
    go func() {
        // Preload caches, warm up connections
        warmupCaches()
        warmupConnectionPools()
        
        // Mark as ready after warmup
        healthChecker.MarkReady()
        log.Info().Msg("Service is ready to accept traffic")
    }()
    
    // Start server
    go func() {
        log.Info().Msg("Starting server on :8080")
        if err := srv.ListenAndServe(); err != http.ErrServerClosed {
            log.Fatal().Err(err).Msg("Server failed")
        }
    }()
    
    // Wait for interrupt signal
    sigCh := make(chan os.Signal, 1)
    signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)
    <-sigCh
    
    log.Info().Msg("Shutdown signal received")
    
    // Mark as not ready (stop receiving new traffic)
    healthChecker.isReady.Store(false)
    
    // Give load balancer time to remove from rotation
    log.Info().Msg("Waiting for load balancer to deregister...")
    time.Sleep(5 * time.Second)
    
    // Graceful shutdown with timeout
    ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
    defer cancel()
    
    if err := srv.Shutdown(ctx); err != nil {
        log.Error().Err(err).Msg("Forced shutdown")
    }
    
    log.Info().Msg("Server stopped gracefully")
}
```

**Kubernetes deployment configuration:**

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-service
spec:
  replicas: 3
  strategy:
    type: RollingUpdate
    rollingUpdate:
      maxSurge: 1        # Max pods above desired
      maxUnavailable: 0  # Keep all pods available
  template:
    spec:
      containers:
      - name: app
        image: my-service:v1.0.0
        ports:
        - containerPort: 8080
        
        # Liveness probe - restart if unhealthy
        livenessProbe:
          httpGet:
            path: /healthz/live
            port: 8080
          initialDelaySeconds: 10
          periodSeconds: 10
          timeoutSeconds: 2
          failureThreshold: 3
        
        # Readiness probe - remove from service if not ready
        readinessProbe:
          httpGet:
            path: /healthz/ready
            port: 8080
          initialDelaySeconds: 5
          periodSeconds: 5
          timeoutSeconds: 2
          failureThreshold: 2
        
        # Lifecycle hooks
        lifecycle:
          preStop:
            exec:
              command: ["/bin/sh", "-c", "sleep 15"]
        
        # Resource limits
        resources:
          requests:
            memory: "128Mi"
            cpu: "100m"
          limits:
            memory: "512Mi"
            cpu: "500m"
      
      # Graceful termination
      terminationGracePeriodSeconds: 45
```

**PreStop hook for graceful shutdown:**

```go
// Alternative: HTTP endpoint for preStop hook
func PreStopHandler(w http.ResponseWriter, r *http.Request) {
    log.Info().Msg("PreStop hook called")
    
    // Mark as not ready
    healthChecker.isReady.Store(false)
    
    // Wait for connections to drain
    time.Sleep(10 * time.Second)
    
    w.WriteHeader(http.StatusOK)
}
```

**Deployment best practices:**
1. **Use rolling updates** with maxUnavailable: 0
2. **Configure proper health checks** (liveness vs readiness)
3. **Set adequate termination grace period** (30-60s)
4. **Use preStop hook** to drain connections
5. **Implement connection pooling** with proper timeouts
6. **Monitor deployment progress** with kubectl rollout status

**Common pitfalls:**
- Not waiting for load balancer deregistration
- Liveness probe too aggressive (causing restart loops)
- No readiness probe (serving traffic before ready)
- Short termination grace period
- Not handling SIGTERM

**Suggested Reading:**
- "Kubernetes Patterns" by Bilgin Ibryam (Health Probe pattern)
- Kubernetes best practices for zero-downtime deployments
- "Production Kubernetes" by Josh Rosso
- Google's "The Prod Engineer's Survival Guide"
</details>

## Security & Authentication

### Question 22: JWT Authentication and Authorization
Implement JWT-based authentication for a microservices system. How do you handle token refresh, revocation, and propagation between services?

<details>
<summary>Answer</summary>

**JWT middleware:**

```go
import "github.com/golang-jwt/jwt/v5"

type Claims struct {
    UserID   string   `json:"user_id"`
    Email    string   `json:"email"`
    Roles    []string `json:"roles"`
    jwt.RegisteredClaims
}

type JWTMiddleware struct {
    secretKey []byte
    redis     *redis.Client
}

func (m *JWTMiddleware) Authenticate(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        authHeader := r.Header.Get("Authorization")
        if authHeader == "" {
            http.Error(w, "Missing authorization header", 
                http.StatusUnauthorized)
            return
        }
        
        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        
        // Parse and validate token
        token, err := jwt.ParseWithClaims(
            tokenString,
            &Claims{},
            func(token *jwt.Token) (interface{}, error) {
                // Verify signing algorithm
                if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
                    return nil, fmt.Errorf("unexpected signing method")
                }
                return m.secretKey, nil
            },
        )
        
        if err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }
        
        claims, ok := token.Claims.(*Claims)
        if !ok || !token.Valid {
            http.Error(w, "Invalid token claims", http.StatusUnauthorized)
            return
        }
        
        // Check token revocation (blacklist in Redis)
        ctx := r.Context()
        revoked, err := m.redis.Exists(ctx, 
            fmt.Sprintf("revoked:%s", claims.ID)).Result()
        if err == nil && revoked > 0 {
            http.Error(w, "Token revoked", http.StatusUnauthorized)
            return
        }
        
        // Add claims to context
        ctx = context.WithValue(ctx, "claims", claims)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

// Authorization middleware
func RequireRole(role string) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            claims := r.Context().Value("claims").(*Claims)
            
            hasRole := false
            for _, r := range claims.Roles {
                if r == role {
                    hasRole = true
                    break
                }
            }
            
            if !hasRole {
                http.Error(w, "Forbidden", http.StatusForbidden)
                return
            }
            
            next.ServeHTTP(w, r)
        })
    }
}
```

**Token generation with refresh tokens:**

```go
type TokenPair struct {
    AccessToken  string `json:"access_token"`
    RefreshToken string `json:"refresh_token"`
    ExpiresIn    int64  `json:"expires_in"`
}

func GenerateTokenPair(userID, email string, roles []string) (*TokenPair, error) {
    // Access token (short-lived: 15 minutes)
    accessClaims := &Claims{
        UserID: userID,
        Email:  email,
        Roles:  roles,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            ID:        uuid.New().String(),
        },
    }
    
    accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims)
    accessString, err := accessToken.SignedString(secretKey)
    if err != nil {
        return nil, err
    }
    
    // Refresh token (long-lived: 7 days)
    refreshClaims := &Claims{
        UserID: userID,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(7 * 24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            ID:        uuid.New().String(),
        },
    }
    
    refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
    refreshString, err := refreshToken.SignedString(secretKey)
    if err != nil {
        return nil, err
    }
    
    // Store refresh token in Redis with TTL
    err = redisClient.Set(
        context.Background(),
        fmt.Sprintf("refresh:%s", refreshClaims.ID),
        userID,
        7*24*time.Hour,
    ).Err()
    
    return &TokenPair{
        AccessToken:  accessString,
        RefreshToken: refreshString,
        ExpiresIn:    900, // 15 minutes
    }, nil
}
```

**Token refresh endpoint:**

```go
func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {
    var req struct {
        RefreshToken string `json:"refresh_token"`
    }
    
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
        http.Error(w, "Invalid request", http.StatusBadRequest)
        return
    }
    
    // Validate refresh token
    token, err := jwt.ParseWithClaims(
        req.RefreshToken,
        &Claims{},
        func(token *jwt.Token) (interface{}, error) {
            return secretKey, nil
        },
    )
    
    if err != nil || !token.Valid {
        http.Error(w, "Invalid refresh token", http.StatusUnauthorized)
        return
    }
    
    claims := token.Claims.(*Claims)
    
    // Check if refresh token exists in Redis
    ctx := r.Context()
    exists, err := redisClient.Exists(ctx,
        fmt.Sprintf("refresh:%s", claims.ID)).Result()
    if err != nil || exists == 0 {
        http.Error(w, "Refresh token not found", http.StatusUnauthorized)
        return
    }
    
    // Get user data to generate new tokens
    user, err := getUserByID(ctx, claims.UserID)
    if err != nil {
        http.Error(w, "User not found", http.StatusUnauthorized)
        return
    }
    
    // Generate new token pair
    tokens, err := GenerateTokenPair(user.ID, user.Email, user.Roles)
    if err != nil {
        http.Error(w, "Token generation failed", 
            http.StatusInternalServerError)
        return
    }
    
    // Revoke old refresh token
    redisClient.Del(ctx, fmt.Sprintf("refresh:%s", claims.ID))
    
    json.NewEncoder(w).Encode(tokens)
}
```

**Token revocation (logout):**

```go
func RevokeTokenHandler(w http.ResponseWriter, r *http.Request) {
    claims := r.Context().Value("claims").(*Claims)
    
    // Add token to blacklist
    ttl := time.Until(claims.ExpiresAt.Time)
    err := redisClient.Set(
        r.Context(),
        fmt.Sprintf("revoked:%s", claims.ID),
        true,
        ttl,
    ).Err()
    
    if err != nil {
        http.Error(w, "Revocation failed", http.StatusInternalServerError)
        return
    }
    
    w.WriteHeader(http.StatusNoContent)
}
```

**Service-to-service authentication:**

```go
// Propagate token to downstream services
func CallDownstreamService(ctx context.Context, url string) error {
    claims := ctx.Value("claims").(*Claims)
    
    req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
    if err != nil {
        return err
    }
    
    // Generate service token or propagate user token
    req.Header.Set("Authorization", "Bearer "+generateServiceToken(claims))
    
    resp, err := http.DefaultClient.Do(req)
    if err != nil {
        return err
    }
    defer resp.Body.Close()
    
    return nil
}
```

**Best practices:**
- Use short-lived access tokens (15 min)
- Rotate refresh tokens on use
- Store refresh tokens securely (httpOnly cookies or secure storage)
- Implement token blacklist for revocation
- Use HTTPS only
- Sign tokens with strong secret (or use asymmetric RS256)
- Validate all claims (exp, iat, iss, aud)

**Suggested Reading:**
- JWT specification (RFC 7519)
- OWASP JWT Cheat Sheet
- "Stop using JWT for sessions" blog post
- golang-jwt/jwt documentation
</details>

### Question 23: Secrets Management
How do you handle secrets (API keys, database passwords) in a Go microservice? What's your strategy for local development vs production?

<details>
<summary>Answer</summary>

**Secret management architecture:**

```go
type SecretProvider interface {
    GetSecret(ctx context.Context, key string) (string, error)
    GetSecretWithDefault(ctx context.Context, key, defaultVal string) string
}

// Environment-based secret provider (local dev)
type EnvSecretProvider struct{}

func (p *EnvSecretProvider) GetSecret(ctx context.Context, key string) (string, error) {
    val := os.Getenv(key)
    if val == "" {
        return "", fmt.Errorf("secret %s not found", key)
    }
    return val, nil
}

// AWS Secrets Manager provider (production)
type AWSSecretProvider struct {
    client *secretsmanager.Client
    cache  *cache.Cache
}

func NewAWSSecretProvider() *AWSSecretProvider {
    cfg, err := config.LoadDefaultConfig(context.Background())
    if err != nil {
        log.Fatal().Err(err).Msg("Failed to load AWS config")
    }
    
    return &AWSSecretProvider{
        client: secretsmanager.NewFromConfig(cfg),
        cache:  cache.New(5*time.Minute, 10*time.Minute),
    }
}

func (p *AWSSecretProvider) GetSecret(ctx context.Context, key string) (string, error) {
    // Check cache first
    if cached, found := p.cache.Get(key); found {
        return cached.(string), nil
    }
    
    input := &secretsmanager.GetSecretValueInput{
        SecretId: aws.String(key),
    }
    
    result, err := p.client.GetSecretValue(ctx, input)
    if err != nil {
        return "", fmt.Errorf("failed to get secret %s: %w", key, err)
    }
    
    secret := *result.SecretString
    p.cache.Set(key, secret, cache.DefaultExpiration)
    
    return secret, nil
}

// Kubernetes Secret provider
type K8sSecretProvider struct {
    basePath string
}

func NewK8sSecretProvider(basePath string) *K8sSecretProvider {
    return &K8sSecretProvider{basePath: basePath}
}

func (p *K8sSecretProvider) GetSecret(ctx context.Context, key string) (string, error) {
    path := filepath.Join(p.basePath, key)
    data, err := os.ReadFile(path)
    if err != nil {
        return "", fmt.Errorf("failed to read secret %s: %w", key, err)
    }
    return strings.TrimSpace(string(data)), nil
}
```

**Factory pattern for environment-specific providers:**

```go
func NewSecretProvider() SecretProvider {
    env := os.Getenv("ENVIRONMENT")
    
    switch env {
    case "production":
        return NewAWSSecretProvider()
    case "staging":
        return NewAWSSecretProvider()
    case "kubernetes":
        return NewK8sSecretProvider("/etc/secrets")
    default:
        return &EnvSecretProvider{}
    }
}
```

**Configuration struct with secrets:**

```go
type Config struct {
    DatabaseURL      string
    RedisURL         string
    APIKey           string
    JWTSecret        []byte
    EncryptionKey    []byte
}

func LoadConfig(ctx context.Context, provider SecretProvider) (*Config, error) {
    dbURL, err := provider.GetSecret(ctx, "DATABASE_URL")
    if err != nil {
        return nil, err
    }
    
    redisURL, err := provider.GetSecret(ctx, "REDIS_URL")
    if err != nil {
        return nil, err
    }
    
    apiKey, err := provider.GetSecret(ctx, "API_KEY")
    if err != nil {
        return nil, err
    }
    
    jwtSecret, err := provider.GetSecret(ctx, "JWT_SECRET")
    if err != nil {
        return nil, err
    }
    
    encryptionKey, err := provider.GetSecret(ctx, "ENCRYPTION_KEY")
    if err != nil {
        return nil, err
    }
    
    return &Config{
        DatabaseURL:   dbURL,
        RedisURL:      redisURL,
        APIKey:        apiKey,
        JWTSecret:     []byte(jwtSecret),
        EncryptionKey: []byte(encryptionKey),
    }, nil
}
```

**Local development with .env file:**

```go
import "github.com/joho/godotenv"

func init() {
    if os.Getenv("ENVIRONMENT") == "" {
        // Load .env file in development
        if err := godotenv.Load(); err != nil {
            log.Warn().Msg("No .env file found")
        }
    }
}
```

**Kubernetes secret mounting:**

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: app-secrets
type: Opaque
stringData:
  DATABASE_URL: "postgresql://user:pass@host:5432/db"
  JWT_SECRET: "your-secret-key"
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: my-service
spec:
  template:
    spec:
      containers:
      - name: app
        image: my-service:latest
        volumeMounts:
        - name: secrets
          mountPath: "/etc/secrets"
          readOnly: true
        env:
        - name: DATABASE_URL
          valueFrom:
            secretKeyRef:
              name: app-secrets
              key: DATABASE_URL
      volumes:
      - name: secrets
        secret:
          secretName: app-secrets
```

**Sensitive data encryption in code:**

```go
import "golang.org/x/crypto/nacl/secretbox"

func EncryptData(data []byte, key *[32]byte) ([]byte, error) {
    var nonce [24]byte
    if _, err := rand.Read(nonce[:]); err != nil {
        return nil, err
    }
    
    encrypted := secretbox.Seal(nonce[:], data, &nonce, key)
    return encrypted, nil
}

func DecryptData(encrypted []byte, key *[32]byte) ([]byte, error) {
    if len(encrypted) < 24 {
        return nil, errors.New("ciphertext too short")
    }
    
    var nonce [24]byte
    copy(nonce[:], encrypted[:24])
    
    decrypted, ok := secretbox.Open(nil, encrypted[24:], &nonce, key)
    if !ok {
        return nil, errors.New("decryption failed")
    }
    
    return decrypted, nil
}
```

**Best practices:**
- Never commit secrets to git
- Use different secrets per environment
- Rotate secrets regularly
- Use IAM roles instead of hardcoded credentials (AWS)
- Encrypt secrets at rest
- Limit secret access with RBAC
- Audit secret access
- Use secret scanning tools (git-secrets, truffleHog)

**Suggested Reading:**
- "Secrets Management Cheat Sheet" by OWASP
- AWS Secrets Manager documentation
- HashiCorp Vault documentation
- "12-Factor App" (Config section)
</details>

## Advanced Patterns

### Question 24: Implementing Saga Pattern
You need to implement a distributed transaction across multiple microservices (order, payment, inventory). How do you implement the saga pattern in Go?

<details>
<summary>Answer</summary>

**Saga orchestrator implementation:**

```go
type SagaStep struct {
    Name        string
    Execute     func(context.Context, *SagaContext) error
    Compensate  func(context.Context, *SagaContext) error
}

type SagaContext struct {
    OrderID   string
    UserID    string
    Amount    decimal.Decimal
    ProductID string
    Data      map[string]interface{}
}

type SagaOrchestrator struct {
    steps []SagaStep
}

func NewOrderSaga() *SagaOrchestrator {
    return &SagaOrchestrator{
        steps: []SagaStep{
            {
                Name: "create_order",
                Execute: func(ctx context.Context, sc *SagaContext) error {
                    order, err := orderService.CreateOrder(ctx, &Order{
                        UserID:    sc.UserID,
                        ProductID: sc.ProductID,
                        Amount:    sc.Amount,
                    })
                    if err != nil {
                        return fmt.Errorf("create order failed: %w", err)
                    }
                    sc.OrderID = order.ID
                    sc.Data["order"] = order
                    return nil
                },
                Compensate: func(ctx context.Context, sc *SagaContext) error {
                    return orderService.CancelOrder(ctx, sc.OrderID)
                },
            },
            {
                Name: "reserve_inventory",
                Execute: func(ctx context.Context, sc *SagaContext) error {
                    reservation, err := inventoryService.Reserve(ctx, 
                        sc.ProductID, 1)
                    if err != nil {
                        return fmt.Errorf("inventory reservation failed: %w", err)
                    }
                    sc.Data["reservation_id"] = reservation.ID
                    return nil
                },
                Compensate: func(ctx context.Context, sc *SagaContext) error {
                    reservationID := sc.Data["reservation_id"].(string)
                    return inventoryService.ReleaseReservation(ctx, reservationID)
                },
            },
            {
                Name: "process_payment",
                Execute: func(ctx context.Context, sc *SagaContext) error {
                    payment, err := paymentService.Charge(ctx, &Payment{
                        OrderID: sc.OrderID,
                        UserID:  sc.UserID,
                        Amount:  sc.Amount,
                    })
                    if err != nil {
                        return fmt.Errorf("payment failed: %w", err)
                    }
                    sc.Data["payment_id"] = payment.ID
                    return nil
                },
                Compensate: func(ctx context.Context, sc *SagaContext) error {
                    paymentID := sc.Data["payment_id"].(string)
                    return paymentService.Refund(ctx, paymentID)
                },
            },
            {
                Name: "confirm_order",
                Execute: func(ctx context.Context, sc *SagaContext) error {
                    return orderService.ConfirmOrder(ctx, sc.OrderID)
                },
                Compensate: func(ctx context.Context, sc *SagaContext) error {
                    // Already compensated by cancel order
                    return nil
                },
            },
        },
    }
}

func (so *SagaOrchestrator) Execute(ctx context.Context, sc *SagaContext) error {
    completedSteps := []int{}
    
    // Execute steps forward
    for i, step := range so.steps {
        log.Info().
            Str("saga_step", step.Name).
            Str("order_id", sc.OrderID).
            Msg("Executing saga step")
        
        if err := step.Execute(ctx, sc); err != nil {
            log.Error().
                Err(err).
                Str("saga_step", step.Name).
                Msg("Saga step failed, starting compensation")
            
            // Compensate completed steps in reverse order
            return so.compensate(ctx, sc, completedSteps)
        }
        
        completedSteps = append(completedSteps, i)
    }
    
    log.Info().
        Str("order_id", sc.OrderID).
        Msg("Saga completed successfully")
    
    return nil
}

func (so *SagaOrchestrator) compensate(
    ctx context.Context,
    sc *SagaContext,
    completedSteps []int,
) error {
    var compensationErrors []error
    
    // Compensate in reverse order
    for i := len(completedSteps) - 1; i >= 0; i-- {
        step := so.steps[completedSteps[i]]
        
        log.Info().
            Str("saga_step", step.Name).
            Msg("Compensating saga step")
        
        if err := step.Compensate(ctx, sc); err != nil {
            log.Error().
                Err(err).
                Str("saga_step", step.Name).
                Msg("Compensation failed")
            compensationErrors = append(compensationErrors, err)
        }
    }
    
    if len(compensationErrors) > 0 {
        return fmt.Errorf("compensation errors: %v", compensationErrors)
    }
    
    return errors.New("saga failed and compensated")
}
```

**Event-driven saga (choreography):**

```go
type EventBus interface {
    Publish(ctx context.Context, event Event) error
    Subscribe(eventType string, handler func(Event) error) error
}

type Event struct {
    ID        string
    Type      string
    Payload   interface{}
    Timestamp time.Time
}

// Order service
func (s *OrderService) HandleOrderCreated(event Event) error {
    order := event.Payload.(*Order)
    
    // Publish event for next step
    return s.eventBus.Publish(context.Background(), Event{
        Type: "inventory.reserve.requested",
        Payload: map[string]interface{}{
            "order_id":   order.ID,
            "product_id": order.ProductID,
            "quantity":   1,
        },
    })
}

func (s *OrderService) HandlePaymentFailed(event Event) error {
    orderID := event.Payload.(map[string]interface{})["order_id"].(string)
    
    // Compensate: cancel order
    return s.CancelOrder(context.Background(), orderID)
}

// Inventory service
func (s *InventoryService) HandleReserveRequested(event Event) error {
    data := event.Payload.(map[string]interface{})
    productID := data["product_id"].(string)
    
    err := s.ReserveInventory(context.Background(), productID, 1)
    if err != nil {
        // Publish failure event
        return s.eventBus.Publish(context.Background(), Event{
            Type: "inventory.reserve.failed",
            Payload: data,
        })
    }
    
    // Publish success event
    return s.eventBus.Publish(context.Background(), Event{
        Type: "inventory.reserve.succeeded",
        Payload: data,
    })
}

// Payment service
func (s *PaymentService) HandleInventoryReserved(event Event) error {
    data := event.Payload.(map[string]interface{})
    orderID := data["order_id"].(string)
    
    err := s.ProcessPayment(context.Background(), orderID)
    if err != nil {
        return s.eventBus.Publish(context.Background(), Event{
            Type: "payment.failed",
            Payload: data,
        })
    }
    
    return s.eventBus.Publish(context.Background(), Event{
        Type: "payment.succeeded",
        Payload: data,
    })
}
```

**Saga state persistence:**

```go
type SagaState struct {
    ID              string
    Status          string // pending, completed, compensating, failed
    CurrentStep     int
    Context         *SagaContext
    CompletedSteps  []int
    CreatedAt       time.Time
    UpdatedAt       time.Time
}

func (so *SagaOrchestrator) ExecuteWithPersistence(
    ctx context.Context,
    sc *SagaContext,
) error {
    state := &SagaState{
        ID:      uuid.New().String(),
        Status:  "pending",
        Context: sc,
    }
    
    // Save initial state
    if err := so.saveState(ctx, state); err != nil {
        return err
    }
    
    for i, step := range so.steps {
        state.CurrentStep = i
        
        if err := step.Execute(ctx, sc); err != nil {
            state.Status = "compensating"
            so.saveState(ctx, state)
            
            return so.compensate(ctx, sc, state.CompletedSteps)
        }
        
        state.CompletedSteps = append(state.CompletedSteps, i)
        so.saveState(ctx, state)
    }
    
    state.Status = "completed"
    so.saveState(ctx, state)
    
    return nil
}
```

**Best practices:**
- Idempotent operations (use idempotency keys)
- Timeout handling
- Retry with exponential backoff
- Dead letter queue for failed compensations
- Monitoring and alerting
- State persistence for recovery
- Use event sourcing for audit trail

**Suggested Reading:**
- "Microservices Patterns" by Chris Richardson (Saga chapter)
- "Enterprise Integration Patterns" by Gregor Hohpe
- "Building Microservices" by Sam Newman
- Temporal.io documentation (workflow orchestration)
</details>

### Question 25: Event Sourcing Implementation
Implement an event-sourced aggregate in Go. How do you handle event versioning, snapshots, and projections?

<details>
<summary>Answer</summary>

**Event and aggregate base:**

```go
type Event interface {
    EventType() string
    AggregateID() string
    Version() int
    Timestamp() time.Time
}

type BaseEvent struct {
    Type        string    `json:"type"`
    AggrID      string    `json:"aggregate_id"`
    Ver         int       `json:"version"`
    Time        time.Time `json:"timestamp"`
    Payload     json.RawMessage `json:"payload"`
}

func (e *BaseEvent) EventType() string    { return e.Type }
func (e *BaseEvent) AggregateID() string  { return e.AggrID }
func (e *BaseEvent) Version() int         { return e.Ver }
func (e *BaseEvent) Timestamp() time.Time { return e.Time }

// Account aggregate
type Account struct {
    ID             string
    Balance        decimal.Decimal
    Owner          string
    Status         string
    version        int
    uncommittedEvents []Event
}

// Domain events
type AccountOpened struct {
    BaseEvent
    AccountID string
    Owner     string
    InitialBalance decimal.Decimal
}

type MoneyDeposited struct {
    BaseEvent
    Amount decimal.Decimal
}

type MoneyWithdrawn struct {
    BaseEvent
    Amount decimal.Decimal
}

type AccountClosed struct {
    BaseEvent
    Reason string
}
```

**Aggregate methods:**

```go
func (a *Account) Open(id, owner string, initialBalance decimal.Decimal) error {
    if initialBalance.LessThan(decimal.Zero) {
        return errors.New("initial balance cannot be negative")
    }
    
    event := &AccountOpened{
        BaseEvent: BaseEvent{
            Type:   "account.opened",
            AggrID: id,
            Ver:    1,
            Time:   time.Now