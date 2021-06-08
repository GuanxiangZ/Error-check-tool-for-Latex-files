package nlp

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type Language struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

var (
	// SupportedLanguages is a list of the supported languages of the tool
	SupportedLanguages = []Language{
		{
			Name: "Arabic",
			Code: "ar",
		},
		{
			Name: "Asturian",
			Code: "ast-ES",
		},
		{
			Name: "Belarusian",
			Code: "be-BY",
		},
		{
			Name: "Breton",
			Code: "br-FR",
		},
		{
			Name: "Catalan",
			Code: "ca-ES",
		},
		{
			Name: "Catalan (Valencian)",
			Code: "ca-ES-valencia",
		},
		{
			Name: "Chinese",
			Code: "zh-CN",
		},
		{
			Name: "Danish",
			Code: "da-DK",
		},
		{
			Name: "Dutch",
			Code: "nl",
		},
		{
			Name: "English (Australia)",
			Code: "en-AU",
		},
		{
			Name: "English (GB)",
			Code: "en-GB",
		},
		{
			Name: "English (Canada)",
			Code: "en-CA",
		},
		{
			Name: "English (New Zealand)",
			Code: "en-NZ",
		},
		{
			Name: "English (South Africa)",
			Code: "en-ZA",
		},
		{
			Name: "English (US)",
			Code: "en-US",
		},
		{
			Name: "Esperanto",
			Code: "eo",
		},
		{
			Name: "French",
			Code: "fr",
		},
		{
			Name: "Galician",
			Code: "gl-ES",
		},
		{
			Name: "German (Austria)",
			Code: "de-AT",
		},
		{
			Name: "German (Germany)",
			Code: "de-DE",
		},
		{
			Name: "German (Switzerland)",
			Code: "de-CH",
		},
		{
			Name: "Greek",
			Code: "el-GR",
		},
		{
			Name: "Irish",
			Code: "ga-IE",
		},
		{
			Name: "Italian",
			Code: "it",
		},
		{
			Name: "Japanese",
			Code: "js-JP",
		},
		{
			Name: "Khmer",
			Code: "km-KH",
		},
		{
			Name: "Norwegian",
			Code: "no",
		},
		{
			Name: "Persian",
			Code: "fa",
		},
		{
			Name: "Polish",
			Code: "pl-PL",
		},
		{
			Name: "Portuguese (Angola)",
			Code: "pt-AO",
		},
		{
			Name: "Portuguese (Brazil)",
			Code: "pt-BR",
		},
		{
			Name: "Portuguese (Mozambique)",
			Code: "pt-MZ",
		},
		{
			Name: "Portuguese (Portugal)",
			Code: "pt-PT",
		},
		{
			Name: "Romanian",
			Code: "ro-RO",
		},
		{
			Name: "Russian",
			Code: "ru-RU",
		},
		{
			Name: "Slovak",
			Code: "sk-SK",
		},
		{
			Name: "Slovenian",
			Code: "sl-SI",
		},
		{
			Name: "Spanish",
			Code: "es",
		},
		{
			Name: "Swedish",
			Code: "sv",
		},
		{
			Name: "Tagalog",
			Code: "tl-PH",
		},
		{
			Name: "Tamil",
			Code: "ta-IN",
		},
		{
			Name: "Ukrainian",
			Code: "uk-UA",
		},
	}
)

type Replacement struct{
	Value string `json:"value"`
}

// Match contains the grammatical errors information
type Match struct {
	Message      string `json:"message"`
	Replacements []Replacement `json:"replacements"`
	LineOffset int `json:"offset"`
	Length int `json:"length"`
	Context Context `json:"context"`
	Sentence string `json:"sentence"`
	Rule Rule `json:"rule"`
}

type Context struct {
	Offset int `json:"offset"`
	Length int `json:"length"`
	Text string `json:"text"`
}

type Rule struct {
	IssueType string `json:"issueType"`
	Description string `json:"description"`
}

// Suggestion contains the data of grammatical recommendations and reason
type Suggestion struct {
	Reason          string   `json:"reason"`
	Recommendations []string `json:"recommendations"`
	Segments []string `json:"segments"`
}

// NLP is the language tool service configuration
type NLP struct {
	Language string
	URL      string
}

// NewDefaultService will create a new NLP instance with default arguments
func NewDefaultService() *NLP {
	return &NLP{
		Language: "en-US",
		URL:      "https://api.languagetool.org/v2/",
	}
}

// NewRemoteService will create a new NLP instance with default arguments
func NewRemoteService(languageCode string) *NLP {
	return &NLP{
		Language: languageCode,
		URL:      "http://localhost:8081/v2/",
	}
}


// NewService will create a new NLP instance with given arguments
func NewService(url, language string) *NLP {
	return &NLP{
		Language: language,
		URL:      url + "/v2/",
	}
}

// SetLanguage will set the language for NLP instance
func (s *NLP) SetLanguage(code string) {
	s.Language = code
}

// SetURL will set the LanguageTool server address for NLP instance
func (s *NLP) SetURL(url string) {
	s.URL = url
}

// Check will check text and give the grammatical correction suggestions for the given text.
func (s NLP) Check(text string) (matches []Match, err error) {
	resp, err := s.check(text)
	if err != nil {
		return nil, err
	}

	return resp.Matches, nil
}

func (s NLP) check(text string) (checkResp *CheckResponse, err error) {

	resp, err := http.PostForm(s.URL+"check", url.Values{
		"language": {s.Language},
		"text":     {text},
	})
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&checkResp)
	if err != nil {
		return nil, err
	}

	return
}

// CheckResponse is the response of LanguageTool API
type CheckResponse struct {
	Matches []Match `json:"matches"`
}