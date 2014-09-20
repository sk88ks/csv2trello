trello-cards-from-csv
=====================

#get app key
`https://trello.com/1/appKey/generate`

#get token
`https://trello.com/1/authorize?key=AppKey&name=<AppName>&response_type=token&scope=read,write`

#get list id
`https://api.trello.com/1/boards/<BoardId>?lists=all&key=<AppKey>&token=<token>`

#usage

```console
go run main.go -key=<key> -token=<token> -l=<listId> -c <csvfile.csv>
```
