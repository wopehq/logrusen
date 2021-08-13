## logrusen

sentry integrated logrus package for our internal projects

### Usage

```go
import "github.com/teamseodo/logrusen"

func main() {
    logger := logrusen.New()
    logger, err := logger.Setup("dev", "dsn")
    if err != nil {
	    logger.Fatal("main", "", "logger setting up error", fmt.Sprintf("%s", err))
	}
    logger.Info("main", "", "logger initialized..")
}
```


