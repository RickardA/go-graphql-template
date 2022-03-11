# Go-Graphql-Template #

## Summary ##

Go-Graphql-Template is a template with minimum code to get CRUD and subscriptions working with grapqhl in Golang.
The template uses a mongodb database that is run through docker.

## Requirements ##
- Docker
- VSCode (To use debugger/launch task)
- Golang

## Get Started ##

1. Start mongoDB database specified in docker-compose.yml
```
docker-compose up
```
2. Start the server by going to "Run and Debug" and use the "Launch Server" option.
3. Navigate to "http://localhost:8081/" to experiment with the grapqhl UI. See "Queries And Mutations" on how to use it.

## Graphql Help ##

### Queries And Mutations ###
```grapqhl
mutation createExample($createInput: GQCreateExample!) {
  createExample(input: $createInput)
}

mutation updateExample($updateInput: GQUpdateExample!) {
  updateExample(input: $updateInput) {
    id
    variable1
    variable2
  }
}

mutation deleteExample($id: String!) {
  deleteExample(id: $id)
}

query getByID($id: String!) {
  getByID(id: $id) {
    id
    variable1
    variable2
  }
}

query getAll {
  getAll {
    id
    variable1
    variable2
  }
}
```

### Query Variables ###
```graphql
{
  "createInput": {
    "variable1": "hejsan",
    "variable2": 10
  },
  "updateInput": {
    "id": "622a4da65a80406a0b6d8f9e",
    "variable1": "new value",
    "variable2": 22
  },
  "id": "622a4da65a80406a0b6d8f9e"
}
```

### Subscription ###
```graphql
subscription exampleSubscription {
  exampleSubscription {
    id
    variable1
    variable2
  }
}
```

### Generate from Graphql ###
To generate new files from the specified "schema.graphqls" in the "graph" folder, use the task "Generate Graphql".

1. Press F1
2. Select Run task
3. Select and run "Generate Graphql"

## Author/Maintainer ##
Rickard Andersson
https://github.com/RickardA/go-graphql-template
