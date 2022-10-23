go-questionnaire
-

Introduction
--

The questionnaire will ask you random questions from a yaml file you provide.

Usage
--

Basic usage:

```shell
./questionnaire -file questions.yaml
```

Questionnaire will ask 10 random questions by default. If you want to choose the number of questions please do:

```shell
./questionnaire -file questions.yaml -n 5
```

Questions file
--

Questionnaire accepts yaml question files tha adhere to a schema:

```go
type T struct {
    Questions []struct {
        Q string
        A []string
    }
}
```

Example:

```yaml
questions:
    - q: How  many days are there in a week?
      a: 7
```
