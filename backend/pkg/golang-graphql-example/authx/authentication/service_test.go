//go:build unit

package authentication

import (
	"net/http"
	"net/http/httptest"
	"regexp"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/stretchr/testify/assert"
)

func Test_getJWTToken(t *testing.T) {
	validAuthorizationHeader := http.Header{}
	validAuthorizationHeader.Add("Authorization", "Bearer TOKEN")
	invalidAuthorizationHeader1 := http.Header{}
	invalidAuthorizationHeader1.Add("Authorization", "TOKEN")
	invalidAuthorizationHeader2 := http.Header{}
	invalidAuthorizationHeader2.Add("Authorization", " TOKEN")
	invalidAuthorizationHeader3 := http.Header{}
	invalidAuthorizationHeader3.Add("Authorization", "Basic TOKEN")
	noHeader := http.Header{}
	validCookie := http.Header{}
	validCookie.Add("Cookie", "oidc=TOKEN")
	type args struct {
		r          *http.Request
		cookieName string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Get token from Authorization header",
			args: args{
				r: &http.Request{
					Header: validAuthorizationHeader,
				},
				cookieName: "oidc",
			},
			want:    "TOKEN",
			wantErr: false,
		},
		{
			name: "Get token from Authorization header (invalid 1)",
			args: args{
				r: &http.Request{
					Header: invalidAuthorizationHeader1,
				},
				cookieName: "oidc",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Get token from Authorization header (invalid 2)",
			args: args{
				r: &http.Request{
					Header: invalidAuthorizationHeader2,
				},
				cookieName: "oidc",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Get token from Authorization header (invalid 3)",
			args: args{
				r: &http.Request{
					Header: invalidAuthorizationHeader3,
				},
				cookieName: "oidc",
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "Get token from cookie without any cookie",
			args: args{
				r: &http.Request{
					Header: noHeader,
				},
				cookieName: "oidc",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Get token from cookie without any cookie",
			args: args{
				r: &http.Request{
					Header: noHeader,
				},
				cookieName: "oidc",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Get token from cookie with valid cookie",
			args: args{
				r: &http.Request{
					Header: validCookie,
				},
				cookieName: "oidc",
			},
			want:    "TOKEN",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getJWTToken(log.NewLogger(), tt.args.r, tt.args.cookieName)
			if (err != nil) != tt.wantErr {
				t.Errorf("getJWTToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getJWTToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidRedirect(t *testing.T) {
	type args struct {
		redirectURLStr string
		reqURLStr      string
	}
	tests := []struct {
		name    string
		args    args
		want    bool
		wantErr bool
	}{
		{
			name: "empty redirect",
			args: args{redirectURLStr: ""},
			want: false,
		},
		{
			name: "singleSlash",
			args: args{redirectURLStr: "/redirect"},
			want: false,
		},
		{
			name: "doubleSlash",
			args: args{redirectURLStr: "//redirect"},
			want: false,
		},
		{
			name: "validHTTP",
			args: args{redirectURLStr: "http://foo.bar/redirect", reqURLStr: "http://foo.bar/"},
			want: true,
		},
		{
			name: "validHTTPS",
			args: args{redirectURLStr: "https://foo.bar/redirect", reqURLStr: "http://foo.bar/"},
			want: true,
		},
		{
			name: "not same domain http",
			args: args{redirectURLStr: "http://foo.bar/redirect", reqURLStr: "http://fake.com/"},
			want: false,
		},
		{
			name: "not same domain https",
			args: args{redirectURLStr: "https://foo.bar/redirect", reqURLStr: "http://fake.com/"},
			want: false,
		},
		{
			name: "openRedirect1",
			args: args{redirectURLStr: "/\\evil.com"},
			want: false,
		},
		{
			name: "openRedirectSpace1",
			args: args{redirectURLStr: "/ /evil.com"},
			want: false,
		},
		{
			name: "openRedirectSpace2",
			args: args{redirectURLStr: "/ \\evil.com"},
			want: false,
		},
		{
			name: "openRedirectTab1",
			args: args{redirectURLStr: "/\t/evil.com"},
			want: false,
		},
		{
			name: "openRedirectTab2",
			args: args{redirectURLStr: "/\t\\evil.com"},
			want: false,
		},
		{
			name: "openRedirectVerticalTab1",
			args: args{redirectURLStr: "/\v/evil.com"},
			want: false,
		},
		{
			name: "openRedirectVerticalTab2",
			args: args{redirectURLStr: "/\v\\evil.com"},
			want: false,
		},
		{
			name: "openRedirectNewLine1",
			args: args{redirectURLStr: "/\n/evil.com"},
			want: false,
		},
		{
			name: "openRedirectNewLine2",
			args: args{redirectURLStr: "/\n\\evil.com"},
			want: false,
		},
		{
			name: "openRedirectCarriageReturn1",
			args: args{redirectURLStr: "/\r/evil.com"},
			want: false,
		},
		{
			name: "openRedirectCarriageReturn2",
			args: args{redirectURLStr: "/\r\\evil.com"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := isValidRedirect(tt.args.redirectURLStr, tt.args.reqURLStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("isValidRedirect() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("isValidRedirect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_service_redirectOrUnauthorized(t *testing.T) {
	fakeMatchingReg := regexp.MustCompile(".*fake")
	type args struct {
		unauthorizedPathRegexList []*regexp.Regexp
	}
	tests := []struct {
		name                   string
		args                   args
		inputRequestPath       string
		expectedStatusCode     int
		expectedLocationHeader string
		expectedBody           string
		checkBody              bool
	}{
		{
			name: "no unauthorized regex list",
			args: args{
				unauthorizedPathRegexList: nil,
			},
			inputRequestPath:       "/faake",
			expectedStatusCode:     307,
			expectedLocationHeader: "/auth/oidc?rd=http%3A%2F%2Fexample.com%2Ffaake",
		},
		{
			name: "not matching unauthorized regex",
			args: args{
				unauthorizedPathRegexList: []*regexp.Regexp{fakeMatchingReg},
			},
			inputRequestPath:       "/faake",
			expectedStatusCode:     307,
			expectedLocationHeader: "/auth/oidc?rd=http%3A%2F%2Fexample.com%2Ffaake",
		},
		{
			name: "unauthorized regex matching",
			args: args{
				unauthorizedPathRegexList: []*regexp.Regexp{fakeMatchingReg},
			},
			inputRequestPath:       "/fake",
			expectedStatusCode:     401,
			expectedLocationHeader: "",
			checkBody:              true,
			expectedBody:           "{\"error\":\"unauthorized\",\"extensions\":{\"code\":\"UNAUTHORIZED\"}}",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create test request
			w := httptest.NewRecorder()
			c, _ := gin.CreateTestContext(w)
			// Create fake request
			req := httptest.NewRequest("GET", tt.inputRequestPath, nil)
			c.Request = req

			// Call function
			redirectOrUnauthorized(c, tt.args.unauthorizedPathRegexList)

			// Check location header
			assert.Equal(t, tt.expectedLocationHeader, w.HeaderMap.Get("Location"))
			// Check status code
			assert.Equal(t, tt.expectedStatusCode, w.Code)
			// Check body
			if tt.checkBody {
				assert.JSONEq(t, tt.expectedBody, w.Body.String())
			}
		})
	}
}
