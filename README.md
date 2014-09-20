trello-cards-from-csv
=====================
csv2trello is a tool to create cards in a trello board with csv.


#get app key
`https://trello.com/1/appKey/generate`

#get token
`https://trello.com/1/authorize?key=AppKey&name=<AppName>&response_type=token&scope=read,write`

#get list id
`https://api.trello.com/1/boards/<BoardId>?lists=all&key=<AppKey>&token=<token>`

#usage

```console
go get github.com/sk88ks/csv2trello
csv2trello -key=<key> -token=<token> -l=<listId> -c <csvfile.csv>
```

#TODO
Modify to be able to "desc" and "checklist"
