//go:build integration

package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	cmocks "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config/mocks"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	smocks "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/signalhandler/mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/goleak"
	"go.uber.org/mock/gomock"
)

func TestInternalServer_generateInternalRouter(t *testing.T) {
	tests := []struct {
		name            string
		inputMethod     string
		inputURL        string
		expectedCode    int
		expectedBody    string
		notExpectedBody string
	}{
		{
			name:         "Should be ok to call /health",
			inputMethod:  "GET",
			inputURL:     "http://localhost/health",
			expectedCode: 200,
			expectedBody: "{}\n",
		},
		{
			name:         "Should be ok to call /metrics",
			inputMethod:  "GET",
			inputURL:     "http://localhost/metrics",
			expectedCode: 200,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create go mock controller
			ctrl := gomock.NewController(t)
			cfgManagerMock := cmocks.NewMockManager(ctrl)

			cfgManagerMock.EXPECT().GetConfig().Return(&config.Config{
				InternalServer: &config.ServerConfig{},
			})

			svr := &InternalServer{
				logger:     log.NewLogger(),
				cfgManager: cfgManagerMock,
				metricsSvc: metricsCtx,
			}
			got, err := svr.generateInternalRouter()
			if err != nil {
				t.Error(err)
				return
			}

			w := httptest.NewRecorder()
			req, err := http.NewRequest(
				tt.inputMethod,
				tt.inputURL,
				nil,
			)
			if err != nil {
				t.Error(err)
				return
			}
			got.ServeHTTP(w, req)
			if tt.expectedCode != w.Code {
				t.Errorf("Integration test on generateInternalRouter() status code = %v, expected status code %v", w.Code, tt.expectedCode)
				return
			}

			if tt.expectedBody != "" {
				body := w.Body.String()
				if tt.expectedBody != body {
					t.Errorf("Integration test on generateInternalRouter() body = \"%v\", expected body \"%v\"", body, tt.expectedBody)
					return
				}
			}

			if tt.notExpectedBody != "" {
				body := w.Body.String()
				if tt.notExpectedBody == body {
					t.Errorf("Integration test on generateInternalRouter() body = \"%v\", not expected body \"%v\"", body, tt.notExpectedBody)
					return
				}
			}
		})
	}
}

func TestInternal_Server_Listen(t *testing.T) {
	// Verify there isn't any go routine leak
	defer goleak.VerifyNone(
		t,
		// Ignore database specific
		goleak.IgnoreTopFunction("database/sql.(*DB).connectionOpener"),
		// Ignore gorm prometheus plugin because it creates an infinite go routine
		goleak.IgnoreTopFunction("gorm.io/plugin/prometheus.(*Prometheus).Initialize.func1.1"),
	)

	// Create go mock controller
	ctrl := gomock.NewController(t)
	cfgManagerMock := cmocks.NewMockManager(ctrl)
	signalHandlerMock := smocks.NewMockService(ctrl)

	// Load configuration in manager
	cfgManagerMock.EXPECT().GetConfig().AnyTimes().Return(&config.Config{
		InternalServer: &config.ServerConfig{
			ListenAddr: "",
			Port:       8080,
		},
	})

	svr := NewInternalServer(log.NewLogger(), cfgManagerMock, metricsCtx, signalHandlerMock)
	// Generate server
	err := svr.GenerateServer()
	assert.NoError(t, err)

	var wg sync.WaitGroup
	// Add a wait
	wg.Add(1)
	// Listen and synchronize wait
	go func() {
		wg.Done()
		err := svr.Listen()
		if !errors.Is(err, http.ErrServerClosed) {
			assert.NoError(t, err)
		}
	}()
	// Wait server up and running
	wg.Wait()
	// Sleep 1 second in order to wait again start server
	time.Sleep(time.Second)

	// Do a request
	req, err := http.NewRequest("GET", "http://localhost:8080/health", nil)
	assert.NoError(t, err)
	// Create client
	httpc := &http.Client{
		Timeout: 200 * time.Millisecond,
	}
	resp, err := httpc.Do(req)
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, 200, resp.StatusCode)
	// Defer close server
	err = svr.server.Close()
	assert.NoError(t, err)

	// Wait a bit
	time.Sleep(200 * time.Millisecond)
}
