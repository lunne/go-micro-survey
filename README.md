# go-micro-survey
API and React Components to create small and easy to use Surveys on any react page

## Development
To generate new models from GraphQL and Proto just run `go generate` or manually run the generation commands in the `main.go` in your terminal.

### MAC
```
brew install protobuf
go github.com/99designs/gqlgen
```
## API
### GraphQL
To crate a new Survey the frontend should use the GraphQL API to get a Template and then create a Form from that.

### GRPC
Backend is driven with Protocol Buffer and GRPC. The CLI can be used to create new Templates and list Surveys

## Datamodel
### Template
A template is created with a set of questiosn and a name and description. This is later used by the Frontend to generate a Survey that can be completed by a user.


 - Template (Name: Edit Device)
    - Questions []:
        - Do you think Edit Device has improved since last time you used it?  [1-7] where one is worse and 7 is alot better! `RAITING`, `REQUIRED`
        - What could improve further? [TEXT], `NOT_REQUIRED`
    Tags: [DEVICE, EDIT]

### Survey
From a Template andits questions a Survey form can be created in the frontend. This is submitted with the `gql  createSurvey`.
A Answer is provided for each Question, something like this:
- Survey
    - Template ID
    - Answers []
        - Question ID
        - Answer (Text, Int, Rating, Option etc)


### Graph
```
     Template ────────► [] Survey ───────► [] Answers
         │                                      │
         │                                      │
         │                                      │
         │                                      │
         ▼                                      │
    [] Question ◄───────────────────────────────┘
```
