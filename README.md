# goschedule

### Basic structure

- Job 

```go
type Job struct {
    command chan interface{}
    url string
    selector map[string]string
    isStopped bool
    id int
}

func (j *job) run()

func (j *job) crawl [][]string 

```

