trello-cards-from-csv
=====================

#get token
https://trello.com/1/authorize?key=AppKey&name=<AppName>&response_type=token&scope=read,write

#get list id
api.trello.com/1/boards/<BoardId>?lists=all&key=<AppKey>&token=<token>
