package parser



const (
	Normal = "normal"
	Error = "error"
)

type LineTMPL struct {
	Matches []MatchTMPL `json:"matches"`
	LineIndex uint64 `json:"lineIndex"`
	Segments []Segment `json:"segments"`
	Sentence string `json:"sentence"`
}

type Segment struct {
	ID string `json:"id"`
	Text string `json:"text"`
	Type string `json:"type"`
}

type MatchTMPL struct {
	ID string `json:"id"`
	Reason string `json:"reason"`
	MatchPreString string `json:"matchPreString"`
	MatchPostString string `json:"matchPostString"`
	ErrorString string `json:"errorString"`
	GlobalErrorStartPosition int `json:"globalErrorStartPosition"`
	GlobalErrorEndPosition int `json:"globalErrorEndPosition"`
	Replacements []ReplacementTMPL `json:"replacements"`
	IssueType string `json:"issueType"`
	IssueDescription string `json:"issueDescription"`
}

type ReplacementTMPL struct {
	Index int `json:"index"`
	ReplacementString string `json:"replacementString"`
}

type Report struct {
	Lines []LineTMPL `json:"lines"`
	ErrorLineIndexes []uint64 `json:"errorLineIndexes"`
	Title string `json:"title"`
}


const REPORTTEMPLATE = `
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <title>.Title</title>
    <style>
        .report-container {
            display: flex;
            flex-direction: column;
            align-items: flex-start;
        }

        .line-header {
            display:flex;
            flex-direction: column;
            align-items:flex-start;
            height: 25px;
            color: #1E88E5;
            font-weight: bold;
        }

        .prompt {
            margin-left: 10px;
            margin-right: 5px;
            height: 25px;
            display:inline-flex;
            flex-direction: column;
            align-items:flex-start;
            color: #1E88E5;
            font-weight: bold;
        }

        .error {
            color: #ef5350;
            font-weight: bold;
        }

        .prompt-gap {
            margin-top: 6px;
            margin-bottom: 6px;
        }

        .replacement-gap {
            margin-top: 4px;
            margin-bottom: 4px;
        }

        .index {
            color: #FF9800;
            margin-left: 20px;
            margin-right: 5px;

        }

        .suggestion-container {
            display:flex;
            flex-direction:column;
            align-items: flex-start;
        }

        .suggestion {
            color: #FF9800;
            font-weight: bold;
        }
    </style>
</head>
<body>
{{range .Lines}}
<div class="report-container">
    <div class="line-header">
        <span>Line number: {{.LineIndex}}</span>
    </div>
    {{range .Matches}}
    <div class = "prompt-gap">
        <span class="prompt">Sentence: </span>{{.MatchPreString}} <span class="error">{{.ErrorString}}</span> {{.MatchPostString}}
    </div>
    <div class = "prompt-gap">
        <span class="prompt">Reason: </span>{{.Reason}}
    </div>
    <div>
        <div class = "prompt-gap"><span class="prompt">Suggestions: </span></div>
        <div class="suggestion-container">
        {{range .Replacements }}
            <div>
                <span class="index">{{.Index}})</span>
                <span class="suggestion">{{.ReplacementString}}</span>
            </div>
            <div class="replacement-gap"></div>
        {{end}}
        </div>
    </div>
    <br>
    {{end}}
</div>
<br>
<br>
{{end}}
</body>
</html>
`