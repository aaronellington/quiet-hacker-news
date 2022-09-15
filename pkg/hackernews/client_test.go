package hackernews

import (
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestClient_TopStories(t *testing.T) {
	responseBytes := []byte("")
	responseCode := http.StatusOK
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(responseBytes)
		w.WriteHeader(responseCode)
	}))
	c := Client{
		BaseURL: ts.URL,
	}

	tests := []struct {
		name        string
		want        []int
		wantErr     bool
		preTestHook func()
	}{
		{
			want:    []int{1, 2, 3},
			wantErr: false,
			preTestHook: func() {
				responseCode = http.StatusOK
				responseBytes = []byte("[1,2,3]")
			},
		},
		{
			want:    []int{},
			wantErr: true,
			preTestHook: func() {
				responseCode = http.StatusOK
				responseBytes = []byte("")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.preTestHook()
			got, err := c.TopStories()
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.TopStories() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.TopStories() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_Item(t *testing.T) {
	responseBytes := []byte("")
	responseCode := http.StatusOK
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(responseBytes)
		w.WriteHeader(responseCode)
	}))
	c := Client{
		BaseURL: ts.URL,
	}

	type args struct {
		id int
	}

	tests := []struct {
		name        string
		args        args
		want        Item
		wantErr     bool
		preTestHook func()
	}{
		{
			args: args{id: 1},
			want: Item{
				Title: "test title",
				URL:   "https://example.com/",
				Host:  "example.com",
			},
			wantErr: false,
			preTestHook: func() {
				responseCode = http.StatusOK
				responseBytes = []byte("{\"title\": \"test title\",\"url\": \"https://example.com/\"}")
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.preTestHook()

			got, err := c.Item(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Client.Item() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.Item() = %v, want %v", got, tt.want)
			}
		})
	}
}
