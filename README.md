# gowatch

### Basic structure

[![Build Status](https://travis-ci.org/dapangmao/gowatch.svg?branch=master)](https://travis-ci.org/dapangmao/gowatch)
[![Coverage Status](https://coveralls.io/repos/github/dapangmao/gowatch/badge.svg)](https://coveralls.io/github/dapangmao/gowatch)

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
