# goschedule

### Basic structure

- Job 

```go
type Job struct {
    fn func()
    kind jobKind
    timer *time
    period time.Duration
    command chan string
}

type JobInterface interface {
    updateTime() 
    startTime()
}

```

- Scheduler

```go 

type Scheduler struct {
    funcMap map[string]*Job
    funcArray []*Job 

}
