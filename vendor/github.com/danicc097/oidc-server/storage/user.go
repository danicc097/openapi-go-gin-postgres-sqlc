package storage

import (
	"crypto/rsa"
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/fsnotify/fsnotify"
)

type User interface {
	ID() string
	Username() string
	Password() string
	IsAdmin() bool
}

type Service struct {
	keys map[string]*rsa.PublicKey
}

type UserStore[T User] interface {
	GetUserByID(string) *T
	GetUserByUsername(string) *T
	ExampleClientID() string
}

type userStore[T User] struct {
	users   map[string]*T
	dataDir string
	mu      sync.RWMutex
}

var StorageErrors struct {
	Errors []string
	mu     sync.RWMutex
}

func NewUserStore[T User](issuer string, dataDir string) (UserStore[T], error) {
	store := userStore[T]{
		users:   make(map[string]*T),
		dataDir: dataDir,
	}

	err := store.LoadUsersFromJSON()
	if err != nil {
		return nil, fmt.Errorf("could not load users from JSON: %w", err)
	}

	go watchUsersFolder(dataDir, &store)

	return &store, nil
}

func (u *userStore[T]) ExampleClientID() string {
	return "service"
}

func (u *userStore[T]) LoadUsersFromJSON() error {
	u.mu.Lock()
	defer u.mu.Unlock()

	u.users = make(map[string]*T)

	files, err := os.ReadDir(u.dataDir)
	if err != nil {
		return err
	}

	errs := []string{}

	for _, file := range files {
		if !file.IsDir() && filepath.Ext(file.Name()) == ".json" {
			filePath := filepath.Join(u.dataDir, file.Name())
			data, err := os.ReadFile(filePath)
			if err != nil {
				return err
			}

			var uu map[string]*T
			err = json.Unmarshal(data, &uu)
			if err != nil {
				return fmt.Errorf("invalid users in %s: %w", filePath, err)
			}

			for _, user := range uu {
				if _, exists := u.users[(*user).ID()]; exists {
					errMsg := fmt.Sprintf("%s: %s: user with ID %s already exists", filePath, (*user).Username(), (*user).ID())
					errs = append(errs, errMsg)
					log.Println(errMsg)
				}
				u.users[(*user).ID()] = user
			}

			if len(errs) > 0 {
				return errors.New(strings.Join(errs, "\n"))
			}

			log.Printf("loaded users from %s", filePath)
		}
	}

	return nil
}

func (u *userStore[T]) GetUserByID(id string) *T {
	u.mu.RLock()
	defer u.mu.RUnlock()

	return u.users[id]
}

func (u *userStore[T]) GetUserByUsername(username string) *T {
	u.mu.RLock()
	defer u.mu.RUnlock()

	for _, user := range u.users {
		if (*user).Username() == username {
			return user
		}
	}

	return nil
}

func watchUsersFolder[T User](dataDir string, userStore *userStore[T]) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	debounceTimer := time.NewTimer(0)  // Create a timer with no initial delay
	debouncedEvent := fsnotify.Event{} // Stores the latest event to process

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				if event.Has(fsnotify.Write) {
					log.Printf("file modified: %s", event.Name)

					debouncedEvent = event // update the debounced event always to latest

					if !debounceTimer.Stop() {
						// timer has already expired, drain the channel
						select {
						case <-debounceTimer.C:
						default:
						}
					}

					debounceTimer.Reset(50 * time.Millisecond)
				}
			case <-debounceTimer.C:
				if debouncedEvent.Name != "" {
					err := userStore.LoadUsersFromJSON()
					StorageErrors.mu.Lock()
					StorageErrors.Errors = []string{}
					if err != nil {
						errMsg := fmt.Sprintf("error reloading users: %s", err)
						StorageErrors.Errors = append(StorageErrors.Errors, errMsg)
						log.Println(errMsg)
					}
					StorageErrors.mu.Unlock()

					debouncedEvent = fsnotify.Event{}
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Printf("watcher error: %s", err)
			}
		}
	}()

	err = filepath.WalkDir(dataDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			log.Printf("walkDir error: %s", err)
			return err
		}
		if !d.IsDir() {
			err = watcher.Add(path)
			if err != nil {
				log.Printf("watcher error: %s", err)
				return err
			}
		}
		return nil
	})
	if err != nil {
		log.Printf("walk error: %s", err)
	}

	<-done
}
