package orchestrator_test

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"hexa/pkg/hawk_support"
	"hexa/pkg/orchestrator"
	"hexa/pkg/web_support"
	"io"
	"log"
	"net/http"
	"testing"
)

func setup(key string) func(t *testing.T) {
	store := hawk_support.NewCredentialStore(key)
	handler := orchestrator.NewApplicationsHandler()
	server := web_support.Create("localhost:8883", func(router *mux.Router) {
		router.HandleFunc("/applications", hawk_support.HawkMiddleware(handler.Applications, store, "localhost:8883")).Methods("GET")
	}, web_support.Options{})

	go web_support.Start(server)
	web_support.WaitForHealthy(server)

	return func(t *testing.T) {
		defer web_support.Stop(server)
	}
}

func TestApplications(t *testing.T) {
	hash := sha256.Sum256([]byte("aKey"))
	key := hex.EncodeToString(hash[:])
	teardownTestCase := setup(key)
	defer teardownTestCase(t)

	resp, err := hawk_support.HawkGet(&http.Client{},"anId", key, "http://localhost:8883/applications")
	if err != nil {
		log.Fatalln(err)
	}
	body, _ := io.ReadAll(resp.Body)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	assert.Equal(t, "{\"applications\":[{\"name\":\"anApp\"}]}", string(body))
}
