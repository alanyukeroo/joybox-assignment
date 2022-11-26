# joybox-assignment

How to run :
1. ``git clone https://github.com/alanyukeroo/joybox-assignment``
2. inside repository ``go mod tidy``
3. ``go run main.go``


there are 2 endpoints:
1. ``/list`` to get Title, Author, and Edition Number of a book.
sample cURL
```
curl -X GET \
  'http://localhost:9000/list?subject=education' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)'
```

2. ``/submit`` to set an appointment for lend the book
```
curl -X POST \
  'http://localhost:9000/submit?pickup_date=2021-11-13&editionCount=772&subject=education&lender_name=Momo' \
  --header 'Accept: */*' \
  --header 'User-Agent: Thunder Client (https://www.thunderclient.com)'
```
