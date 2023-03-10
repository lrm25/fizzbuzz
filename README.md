# FizzBuzz Application

## Backend

To run backend in Windows (assuming you have go):

```
go mod tidy
$env:FIZZ = "<'Fizz', or other value>"
$env:BUZZ = "<'Buzz', or other value>"
$env:FIZZBUZZ_SERVERPORT = "<local port to run server on>"
$env:FIZZBUZZ_CLIENTURL = "<URL of client, including scheme, default would be http://localhost:3000>"
go run main.go
```

## Frontend

To run frontend:

1.  In `frontend/fizzbuzz/.env` file, set `REACT_APP_BACKEND_URL` parameter to address of Golang server, including port.
2.  Go to `frontend/fizzbuzz` folder and run `npm run build`.
3.  Run program with `npm start`.

## Future Improvements (If This Was A Serious Project)

* Prevent GUI from shifting when "Fizz" or "Buzz" appears
* Use docker-compose, kubernetes, etc. to deploy/manage simultaneously
