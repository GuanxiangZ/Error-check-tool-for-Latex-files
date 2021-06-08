package parser

import (
	"encoding/json"
	"fmt"
	"github.com/uoa-compsci-399/error-correction-tool-for-latex/config"
	"github.com/uoa-compsci-399/error-correction-tool-for-latex/nlp"
	"github.com/uoa-compsci-399/error-correction-tool-for-latex/utils"
)

type LineType int

const (
	Comment LineType = iota
	Blank
	Other
)



const (
	text = "Text"
	sentence = "Sentence"
	suggestions = "Suggestions"
	reason = "Reason"
	linenumber = "Line"
)

type Line struct {
	Index uint64 `json:"index" expr:"the line number"`
	Start    int    `json:"start"`
	End      int    `json:"End"`
	Raw      string `json:"raw"`
	ProcessedSentence      string `json:"processed_sentence"`
	Len      int    `json:"len"`
	Matches []nlp.Match `json:"matches"`
}

func (l Line) String() (res string) {
	bs, _ := json.MarshalIndent(l, "", "	")
	return string(bs)
}

func (l *Line) OutputHTMLStr(start int) (linetmpl LineTMPL)  {
	if len(l.Matches) <= 0 {
		return
	}

	linetmpl.LineIndex = l.Index
	linetmpl.Matches = []MatchTMPL{}
	linetmpl.Segments = []Segment{}

	var (
		cursor, pad int
		cfg = config.GetServeCmdConfig()
	)

	if start > 0 {
		pad = start
		linetmpl.Segments = append(linetmpl.Segments, Segment{
			Text: l.Raw[:pad],
			Type: "normal",
		})
	} else {
		pad = 0
	}

	cursor = 0
	for _, match := range l.Matches {
		var (
			matchtmpl MatchTMPL
			pre, post, errorstr string
			lineOffset int
			id string
		)

		id = fmt.Sprintf("%d", cfg.NextID())

		lineOffset = match.LineOffset

		if cursor != lineOffset {
			linetmpl.Segments = append(linetmpl.Segments, Segment{
				Text: l.ProcessedSentence[cursor : lineOffset],
				Type: "normal",
			})
		}
		linetmpl.Segments = append(linetmpl.Segments, Segment{
			Text: l.ProcessedSentence[lineOffset : lineOffset + match.Length],
			Type: "error",
			ID: id,
		})
		cursor = lineOffset + match.Length
		pre = match.Context.Text[:match.Context.Offset]

		errorstr = match.Context.Text[match.Context.Offset:match.Context.Offset+ match.Context.Length]

		post = match.Context.Text[match.Context.Offset + match.Context.Length:]

		matchtmpl.ID = id
		matchtmpl.MatchPreString = pre
		matchtmpl.ErrorString = errorstr
		matchtmpl.MatchPostString = post

		matchtmpl.Reason = match.Message
		matchtmpl.IssueType = match.Rule.IssueType
		matchtmpl.IssueDescription = match.Rule.Description
		matchtmpl.GlobalErrorStartPosition = l.Start + match.LineOffset
		matchtmpl.GlobalErrorEndPosition = l.Start + match.LineOffset + match.Context.Length
		matchtmpl.Replacements = []ReplacementTMPL{}

		for idx, repl := range match.Replacements {
			if idx >= 20 {
				break
			}

			var replace = ReplacementTMPL{
				Index: idx + 1,
				ReplacementString: repl.Value,
			}

			matchtmpl.Replacements = append(matchtmpl.Replacements, replace)
		}

		linetmpl.Matches = append(linetmpl.Matches, matchtmpl)
	}

	linetmpl.Segments = append(linetmpl.Segments, Segment{
		Text: l.ProcessedSentence[cursor:],
		Type: "normal",
	})

	if start > 0 {
		if pad + len(l.ProcessedSentence) < len(l.Raw) {
			linetmpl.Segments = append(linetmpl.Segments, Segment{
				Text: l.Raw[pad + len(l.ProcessedSentence):],
				Type: "normal",
			})
		}
	}

	return
}

func (l *Line) PrintReport() {
	if len(l.Matches) <= 0 {
		return
	}

	fmt.Printf("%s%d\n", utils.Prompt(linenumber), l.Index)
	fmt.Println(l.Start, l.End)
	for _, match := range l.Matches {

		fmt.Printf("  %s", utils.Prompt(sentence))
		fmt.Printf("%s", match.Context.Text[:match.Context.Offset])
		fmt.Printf("%s", utils.Error(match.Context.Text[match.Context.Offset:match.Context.Offset+ match.Context.Length]))
		fmt.Println(match.Context.Text[match.Context.Offset + match.Context.Length:])
		fmt.Printf("  %s", utils.Prompt(reason))
		fmt.Println(match.Message)
		fmt.Printf("  %s\n", utils.Prompt(suggestions))
		for idx, repl := range match.Replacements {
			if idx >= 20 {
				break
			}

			fmt.Printf("%s", utils.Index(idx + 1))
			fmt.Printf("%s", utils.Suggestion(repl.Value))
			fmt.Println()
		}
		fmt.Println()
	}
}

func (l *Line) processText(svc *nlp.NLP) (err error) {
	l.Matches, err = svc.Check(l.ProcessedSentence)
	return
}