package detector_test

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"server-ids/internal/auth"
	"server-ids/internal/database"
	"server-ids/internal/detector"
	"server-ids/internal/sessions"
	"server-ids/internal/template"
	"server-ids/internal/user"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestBACDetection(t *testing.T) {
	tests := []struct {
		name        string
		request     *http.Request
		wasDetected bool
		message     string
	}{
		{
			name: "BAC Detected on unauthorized user changing a role",
			request: func() *http.Request {
				formData := url.Values{}
				formData.Set("newRole", "admin")
				r, _ := http.NewRequest(http.MethodPatch, "/users/{username}/role", strings.NewReader(formData.Encode()))
				r = mux.SetURLVars(r, map[string]string{
					"username": "patrick",
				})
				return r
			}(),
			wasDetected: true,
			message:     "unauthenticated person tried to change patrick's role to a admin",
		},
		{
			name: "BAC Detected on guest user changing a role",
			request: func() *http.Request {
				formData := url.Values{}
				formData.Set("newRole", "admin")
				r, _ := http.NewRequest(http.MethodPatch, "/users/{username}/role", strings.NewReader(formData.Encode()))
				r = mux.SetURLVars(r, map[string]string{
					"username": "patrick",
				})
				r.AddCookie(&http.Cookie{
					Name:  "session_key",
					Value: "00000000-0000-0000-0000-000000000000",
				})
				return r
			}(),
			wasDetected: true,
			message:     "secure21 tried to change patrick's role to a admin",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tmpl := template.NewTestTemplate()
			d := detector.NewDetector(tmpl)
			db := database.CreateMockDB()
			defer db.Close()
			ar := auth.NewAuthRepository(db)
			ur := user.NewUserRepository(db)
			user, err := ur.GetUser("secure21")
			assert.NoError(t, err)
			sessions := sessions.NewSessions(db)
			sessionID := []byte("00000000-0000-0000-0000-000000000000")
			sessionUUID, _ := uuid.FromBytes(sessionID)
			ar.AddSession(sessionUUID, *user)
			bacDetection := &detector.BACDetection{
				Sessions: sessions,
			}

			found, err := bacDetection.Run(httptest.NewRecorder(), test.request, d)

			assert.NoError(t, err)
			assert.Equal(t, test.wasDetected, found)
			messages := []string{}
			for _, alert := range d.Alerts {
				messages = append(messages, alert.Message)
			}
			assert.Contains(t, messages, test.message)
		})
	}
}
