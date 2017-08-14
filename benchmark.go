package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const NUM_TRIALS = 10

func main() {
	req, _ := http.NewRequest("GET", "http://localhost:8000/user?access_token=lk4P4yxB_yTWm37W", nil)
	req.Header.Set("Authorization", "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJlbWFpbCI6ImFsZXgrc2FtcGxlQGJsdWVjYW52YXMuaW8iLCJuYW1lIjoiYWxleCtzYW1wbGVAYmx1ZWNhbnZhcy5pbyIsInBpY3R1cmUiOiJodHRwczovL3MuZ3JhdmF0YXIuY29tL2F2YXRhci9iMmZjNGViYzAyNzQyNjAxZmIyZDAyMTAyZGIxZmJhYT9zPTQ4MCZyPXBnJmQ9aHR0cHMlM0ElMkYlMkZjZG4uYXV0aDAuY29tJTJGYXZhdGFycyUyRmFsLnBuZyIsIm5pY2tuYW1lIjoiYWxleCtzYW1wbGUiLCJhcHBfbWV0YWRhdGEiOnsiYXV0aG9yaXphdGlvbiI6eyJncm91cHMiOltdfX0sImF1dGhvcml6YXRpb24iOnsiZ3JvdXBzIjpbXX0sImdyb3VwcyI6W10sImVtYWlsX3ZlcmlmaWVkIjp0cnVlLCJjbGllbnRJRCI6ImtieXVGRGlkTExtMjgwTEl3VkZpYXpPcWpPM3R5OEtIIiwidXBkYXRlZF9hdCI6IjIwMTctMDgtMTRUMTU6MTA6MzcuNTg4WiIsInVzZXJfaWQiOiJhdXRoMHw1OTZmMjdjMmMzNzA5NjYxZTljZWEzN2QiLCJpZGVudGl0aWVzIjpbeyJ1c2VyX2lkIjoiNTk2ZjI3YzJjMzcwOTY2MWU5Y2VhMzdkIiwicHJvdmlkZXIiOiJhdXRoMCIsImNvbm5lY3Rpb24iOiJVc2VybmFtZS1QYXNzd29yZC1BdXRoZW50aWNhdGlvbiIsImlzU29jaWFsIjpmYWxzZX1dLCJjcmVhdGVkX2F0IjoiMjAxNy0wNy0xOVQwOTozNDo1OC4yMjlaIiwiaXNzIjoiaHR0cHM6Ly9zYW1wbGVzLmF1dGgwLmNvbS8iLCJzdWIiOiJhdXRoMHw1OTZmMjdjMmMzNzA5NjYxZTljZWEzN2QiLCJhdWQiOiJrYnl1RkRpZExMbTI4MExJd1ZGaWF6T3FqTzN0eThLSCIsImlhdCI6MTUwMjcyMzQ0OCwiZXhwIjoxNTAyNzU5NDQ4fQ.oSpcx4A50jdR4GCKDdyaxb5zcDKp2xDK08pi5jgmadk")

	start := time.Now().UnixNano() / 1e6
	for i := 0; i < NUM_TRIALS; i++ {
		if resp, err := http.DefaultClient.Do(req); err != nil {
			panic(err)
		} else {
			_, _ = ioutil.ReadAll(resp.Body)
			resp.Body.Close()
		}
	}
	end := time.Now().UnixNano() / 1e6

	fmt.Printf("%v req/sec\n", NUM_TRIALS*1000/(end-start))
}
