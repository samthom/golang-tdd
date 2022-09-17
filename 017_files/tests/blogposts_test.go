package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/samthom/golang_tdd/017_files"
)

type StubFailingFS struct{}

func (s StubFailingFS) Open(name string) (fs.File, error) {
	return nil, errors.New("oh no, i always fail")
}

func assertPost(t *testing.T, got blogposts.Post, want blogposts.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func TestNewBlogPosts(t *testing.T) {
	t.Run("fail test", func(t *testing.T) {
		_, err := blogposts.NewPostsFromFS(StubFailingFS{})

		if err == nil {
			t.Error("got nil error, wanted error")
		}
	})

    const (
        firstBody = `Title: Post 1
Description: Description 1
Tags: go, tdd
---
Hello
World`
        secondBody = `Title: Post 2
Description: Description 2
Tags: go, tdd
---
B
L
M`
    )

	t.Run("Succes test", func(t *testing.T) {
		fs := fstest.MapFS{
			"hello-world.md":  {Data: []byte(firstBody)},
			"hello-world2.md": {Data: []byte(secondBody)},
		}

		posts, err := blogposts.NewPostsFromFS(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("got %d posts, wanted %d posts", len(posts), len(fs))
		}

		got := posts[0]
        want := blogposts.Post{Title: "Post 1", Description: "Description 1", Tags: []string{"go", "tdd"}, Body: `Hello
World`}

        assertPost(t, got, want)
	})
}
