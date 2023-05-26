# custom-skill-builder
Utility to build [Custom Web API Skill](https://learn.microsoft.com/en-us/azure/search/cognitive-search-custom-skill-web-api) for Azure Cognitive Search.

## Usage
[See full example](example/main.go)

### Create a skill

This is a simple example skill that converts the input string to lower case. Type of the args and the return value are arbitrary, but they must be serializable to JSON. And typical types are already defined in `model` package.

```go
lowerSkill := skill.NewSkill(func(d *model.StringData) (*model.StringData, error) {
    if d.Input == "" {
        return nil, model.ErrInputNotFound
    }
    return &model.StringData{Output: strings.ToLower(d.Input)}, nil
})
```

### Register the skill
`skill.Book` is a collection of skills. You can register a skill to the book by calling `Register` method. The first argument is the name of the skill. The second argument is the skill itself. You can register multiple skills to the book. 

`Flatten()` returns `func([]byte) ([]byte, error)`. This is the function that enables type genralization, and both of `[]byte` are serialized JSON. Then you can register the skill to the book.

```go
book := skill.NewBook()
book.Register("lower", lowerSkill.Flatten())
```

### Create and run the server
Once you start the server, it will listen to the port 8080, and `hostname:8080/skills/lower` will be the endpoint of the skill.

```go
svc := service.NewCustomSkillService(book)
svc.Run()
```

### Testing the skill
You can call the skills with the endpoint `skills/{skillName}`, and values of data are json-serialized struct of skill definition.


```http request
POST http://localhost:8080/skills/lower
Content-Type: application/json

{
    "values": [
        {
            "recordId": "a1",
            "data": {
                "input": "UPPER CASE"
            }
        },
        {
            "recordId": "a2",
            "data": {
                "invalidField": "UPPER CASE"
            }
        }
    ]
}
```

then we get the response

```http request
HTTP/1.1 200 OK
Content-Type: application/json
Date: Fri, 26 May 2023 12:26:52 GMT
Content-Length: 120
Connection: close

{
  "values": [
    {
      "recordId": "a1",
      "data": {
        "output": "upper case"
      }
    },
    {
      "recordId": "a2",
      "errors": [
        {
          "message": "input not found"
        }
      ]
    }
  ]
}
```

