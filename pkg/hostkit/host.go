package hostkit

import (
	"bufio"
	"fmt"
	"net/http"
	"strings"
)

func GetGithubKeysByID(id string) string {
	keysURL := fmt.Sprintf("https://github.com/%s.keys", id)
	resp, err := http.Get(keysURL)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	keys := []string{}
	scanner := bufio.NewScanner(resp.Body)
	for i := 0; scanner.Scan() && i < 3; i++ {
		keys = append(keys, fmt.Sprintf("%s https://github.com/%s", scanner.Text(), id))
	}

	return strings.Join(keys, "\n")
}

func GetGithubKeys(userID []string) {
	for _, id := range userID {
		GetGithubKeysByID(id)
	}
}
