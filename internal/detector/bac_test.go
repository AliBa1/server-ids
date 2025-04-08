package detector

// func TestBACDetection(t *testing.T) {
// 	tests := []struct {
// 		name        string
// 		request     *http.Request
// 		wasDetected bool
// 		message     string
// 	}{
// 		{
// 			name:        "BAC Detected on unauthorized user changing a role",
// 			request:     httptest.NewRequest(http.MethodPatch, "/users/patrick/role", nil),
// 			wasDetected: true,
// 			message:     "detected in url path: /id=1' OR '1'='1",
// 		},
// 	}

// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			detector := NewDetector()
// 			bacDetection := &BACDetection{}

// 			found, err := bacDetection.Run(httptest.NewRecorder(), test.request, detector)

// 			assert.NoError(t, err)
// 			assert.Equal(t, test.wasDetected, found)
// 			if test.wasDetected && len(detector.Alerts) > 0 {
// 				assert.Equal(t, test.message, detector.Alerts[0].Message)
// 			}
// 		})
// 	}
// }
