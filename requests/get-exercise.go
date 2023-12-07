package requests

import (
	"fmt"
	"io"
	"net/http"
)

func GetExercise(session string, day uint8) string {
	client := &http.Client{}
	uri := fmt.Sprintf("https://adventofcode.com/2023/day/%d/input", day)
	req, _ := http.NewRequest("GET", uri, nil)
	req.AddCookie(&http.Cookie{Name: "session", Value: session})

	resp, _ := client.Do(req)
	body, _ := io.ReadAll(resp.Body)
	resp.Body.Close()

	return string(body)
}
