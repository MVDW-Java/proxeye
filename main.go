package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type GelbooruResponse struct {
	Post []struct {
		FileURL string `json:"file_url"`
	} `json:"post"`
}

var client = &http.Client{
	Timeout: 20 * time.Second,
}

func main() {
	godotenv.Load()

	http.HandleFunc("/", handlePost)
	port := os.Getenv("PROXEYE_PORT")
	if port == "" {
		port = "8080"
	}

	log.Println("Listening on :", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handlePost(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/")
	if path == "" {
		http.Error(w, "Missing post ID", http.StatusBadRequest)
		return
	}
	parts := strings.Split(path, ".")
	postID := parts[0]

	apiKey := os.Getenv("GELBOORU_API_KEY")
	userID := os.Getenv("GELBOORU_USER_ID")

	apiURL := fmt.Sprintf(
		"https://gelbooru.com/index.php?page=dapi&s=post&q=index&id=%s&json=1&api_key=%s&user_id=%s",
		postID,
		apiKey,
		userID,
	)

	resp, err := client.Get(apiURL)
	if err != nil {
		http.Error(w, "Failed to fetch post metadata", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Error(w, "Post not found", http.StatusNotFound)
		return
	}

	var data GelbooruResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		http.Error(w, "Failed to decode API response", http.StatusInternalServerError)
		return
	}

	if len(data.Post) == 0 || data.Post[0].FileURL == "" {
		http.Error(w, "No file found for post", http.StatusNotFound)
		return
	}

	mediaURL := data.Post[0].FileURL

	req, err := http.NewRequest("GET", mediaURL, nil)
	if err != nil {
		http.Error(w, "Failed to create media request", http.StatusInternalServerError)
		return
	}

	req.Header.Set("Referer",
		fmt.Sprintf("https://gelbooru.com/index.php?page=post&s=view&id=%s", postID),
	)

	mediaResp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to fetch media", http.StatusBadGateway)
		return
	}
	defer mediaResp.Body.Close()

	if mediaResp.StatusCode != http.StatusOK {
		http.Error(w, "Media not accessible", http.StatusBadGateway)
		return
	}

	for k, v := range mediaResp.Header {
		switch strings.ToLower(k) {
		case "content-type", "content-length", "last-modified", "etag":
			for _, val := range v {
				w.Header().Add(k, val)
			}
		}
	}

	w.WriteHeader(mediaResp.StatusCode)

	_, err = io.Copy(w, mediaResp.Body)
	if err != nil {
		log.Println("Streaming error:", err)
	}
}
