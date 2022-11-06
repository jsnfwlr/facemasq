package apiclient

import (
	"bytes"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"text/template"

	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type APIClient struct {
	Engine       http.Client
	LooseCookies []*http.Cookie
}

// Request is a struct used to store the collective data of the API/Web Request used by Do() and Build()
type Request struct {
	URL             string
	Method          string
	EndPoint        string
	EndPointParams  URLParams
	DynamicEndPoint bool
	Headers         Headers
	Path            []string
	Query           url.Values
	Body            BodyParams
	BodyJSON        io.Reader
	Auth            Auth
	BodyText        string
	Cookies         []*http.Cookie
}

type Response struct {
	Original   *http.Response
	Body       string
	Cookies    string
	Cookie     *http.Cookie
	Jar        *http.CookieJar
	StatusCode int
	Header     *http.Header
}

// URLParams is a map of strings, with a string key, used to store the parameters of the request url
type URLParams map[string]string

// Headers is a map of strings, with a string key, used to store the parameters of the request headers
type Headers map[string]string

// BodyParams is a map of strings, with a string key, used to store the parameters of the request body
type BodyParams map[string]string

// Auth is a struct used to store the parameters of the request authentication values
type Auth struct {
	Type     string
	Username string
	Password string
}

// Prepare sets-up the client used in Do()
func Prepare() (client *APIClient) {
	jar, _ := cookiejar.New(nil)
	client = &APIClient{
		Engine: http.Client{
			Jar: jar,
		},
	}
	return
}

// Build creates a http.Request from an Request and returns a pointer to it, ready for use in non-http calls to our own handlers
func (request *Request) build() (action *http.Request, err error) {
	if request.DynamicEndPoint {
		var tmpl *template.Template
		tmpl, err = template.New("Dynamic").Parse(request.EndPoint)
		if err != nil {
			err = errors.Wrap(err, "Could not parse endpoint template")
			return
		}
		buffer := new(bytes.Buffer)
		err = tmpl.Execute(buffer, request.EndPointParams)
		if err != nil {
			err = errors.Wrap(err, "Could not execute endpoint template")
			return
		}
		request.EndPoint = buffer.String()
	}
	request.EndPoint = request.URL + request.EndPoint
	if request.Query != nil {
		request.EndPoint = request.EndPoint + "?" + request.Query.Encode()
	}

	action, err = http.NewRequest(request.Method, request.EndPoint, request.BodyJSON)
	if err != nil {
		err = errors.Wrap(err, "Encountered error while preparing API action.")
		return
	}

	if request.Body != nil {
		var body []byte
		body, err = json.Marshal(request.Body)
		if err != nil {
			err = errors.Wrap(err, "Encountered error while preparing Body Params.")
			return
		}
		action, err = http.NewRequest(request.Method, request.EndPoint, bytes.NewReader(body))
		if err != nil {
			err = errors.Wrap(err, "Encountered error while preparing API action.")
			return
		}
	}

	if request.BodyText != "" {
		body := []byte(request.BodyText)
		action, err = http.NewRequest(request.Method, request.EndPoint, bytes.NewReader(body))
		if err != nil {
			err = errors.Wrap(err, "Encountered error while preparing API action.")
			return
		}
	}

	for key, value := range request.Headers {
		action.Header.Add(key, value)
	}

	if request.Auth != (Auth{}) {
		switch request.Auth.Type {
		case "Basic":
			action.SetBasicAuth(request.Auth.Username, request.Auth.Password)
		case "Bearer":
			action.Header.Add("Authorization", "Bearer "+request.Auth.Username)
		}
	}

	for i := range request.Cookies {
		action.AddCookie(request.Cookies[i])
	}

	return
}

// Do creates a http.Request from an Request and executes the request,
func (client *APIClient) Do(request *Request) (response *http.Response, err error) {
	var action *http.Request
	request.Cookies = append(request.Cookies, client.LooseCookies...)
	action, err = request.build()
	if err != nil {
		err = errors.Wrap(err, "Encountered error while building request.")
		return
	}

	response, err = client.Engine.Do(action)
	if err != nil {
		err = errors.Wrap(err, "Encountered error while executing request.")
		return
	}
	if response.StatusCode > 299 {
		err = errors.New(http.StatusText(response.StatusCode))
		return
	}
	return
}

func ParseBody(response *http.Response) (contents string, err error) {
	var body []byte
	defer response.Body.Close()
	body, err = io.ReadAll(response.Body)
	if err != nil {
		return
	}
	contents = string(body)
	return
}

func (client *APIClient) ClearLooseCookies() {
	client.LooseCookies = []*http.Cookie{}
}
