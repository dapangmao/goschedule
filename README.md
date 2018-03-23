# gowatch

### Basic structure

[![Build Status](https://travis-ci.org/dapangmao/gowatch.svg?branch=master)](https://travis-ci.org/dapangmao/gowatch)

```
+---------------+             +------------------+
|  Webui        |             |  Scheduler       |
|  + parse      +-----------> |                  |
|               |             |                  |
+---------------+             +----------+-------+
                                         |
                                         |
                                         |
                                    +----v-----+
       +--------------+             |  Job     |
       | Main         |             |          |
       |              |             |          |
       | run goroutine|             +-+--------+
       +--------------+               |
                                  +---v-----+
                                  | Fetcher |
                                  |         |
                                  |         |
                                  |         |
                                  |         |
                                  +---------+
                                        |
                                    +---v-----+
                                    | BlotDB  |
                                    |         |
                                    |         |
                                    |         |
                                    |         |
                                    +---------+
```
