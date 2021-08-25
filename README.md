## logrusen

Wrapper for logrus package  

### Usage

```go
import "github.com/teamseodo/logrusen"

func main() {
	logger := logrusen.New()
	err := logger.Setup()

	if err != nil {
		logger.Error("Logger initialization error", err, nil)
	}
	logger.Info("Info testing", nil)
	logger.Debug("Debug testing", logrusen.Fields{"test": "testvalue"})
	logger.Warn("Warn logging", err, nil)
	logger.Fatal("Fatal logging", err, nil)
}

```


### Tests

```bash
$ go test -v *.go
```

```bash
=== RUN   TestNew
=== RUN   TestNew/create_standard_logger
--- PASS: TestNew (0.00s)
    --- PASS: TestNew/create_standard_logger (0.00s)
=== RUN   Test_standardLogger_Setup
=== RUN   Test_standardLogger_Setup/invalid_env_name_(env:development)
=== RUN   Test_standardLogger_Setup/valid_env_name_(env:dev)
=== RUN   Test_standardLogger_Setup/valid_env_name_but_dsn_is_invalid_(env:prod,_dsn:123456789)
=== RUN   Test_standardLogger_Setup/invalid_env_name_(env:test)
--- PASS: Test_standardLogger_Setup (0.01s)
    --- PASS: Test_standardLogger_Setup/invalid_env_name_(env:development) (0.00s)
    --- PASS: Test_standardLogger_Setup/valid_env_name_(env:dev) (0.00s)
    --- PASS: Test_standardLogger_Setup/valid_env_name_but_dsn_is_invalid_(env:prod,_dsn:123456789) (0.01s)
    --- PASS: Test_standardLogger_Setup/invalid_env_name_(env:test) (0.00s)
=== RUN   Test_stdFields
=== RUN   Test_stdFields/foo_and_topic1
=== RUN   Test_stdFields/fib_and_topic2
--- PASS: Test_stdFields (0.00s)
    --- PASS: Test_stdFields/foo_and_topic1 (0.00s)
    --- PASS: Test_stdFields/fib_and_topic2 (0.00s)
PASS
ok  	command-line-arguments	0.544s
```



