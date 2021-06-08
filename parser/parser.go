package parser

import (
	"context"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/uoa-compsci-399/error-correction-tool-for-latex/config"
	"github.com/uoa-compsci-399/error-correction-tool-for-latex/nlp"
	"google.golang.org/grpc"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func ProcessTeX(file io.Reader, filename, languageCode string, outHTML bool) (content Report, err error) {
	var (
		nlpservice *nlp.NLP
		report Report = Report{
			Lines: []LineTMPL{},
			ErrorLineIndexes: []uint64{},
			Title: filename,
		}
		cfg *config.ServerConfig
		isDocker bool
		parserURL string
		lines []string
		contentsStr string
	)

	nlpservice = nlp.NewRemoteService(languageCode)

	contents, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}

	contentsStr = string(contents)

	lines = strings.Split(contentsStr, "\n")

	cfg = config.GetServeCmdConfig()

	isDocker = len(os.Getenv("DOCKER_ENVIRONMENT")) == 0
	if isDocker {
		parserURL = cfg.ParserUrl
	} else {
		parserURL = os.Getenv("PARSER_SERVER_ADDR")
	}

	logrus.Info(cfg.ParserUrl)
	conn, err := grpc.Dial(parserURL, grpc.WithInsecure())
	if err != nil {
		logrus.Error(err)
		return
	}
	defer conn.Close()

	client := NewParserClient(conn)
	reply, err := client.Parse(context.Background(), &ParseRequest{Content: contentsStr})
	if err != nil {
		logrus.Error(err)
		return
	}

	var parseResult ParserResult
	err = json.Unmarshal([]byte(reply.Result), &parseResult)
	if err != nil {
		return
	}

	for _, val := range parseResult.Data {
		var start int

		start = strings.Index(lines[val.LineNumber], val.Text)

		line := Line{
			Index: val.LineNumber,
			Raw:               lines[val.LineNumber],
			ProcessedSentence: val.Text,
			Start: 				start,
			Len:               len(val.Text),
		}

		err = line.processText(nlpservice)
		if err != nil {
			return
		}

		if outHTML {
			line.PrintReport()
		}

		reportstr := line.OutputHTMLStr(start)
		if len(reportstr.Matches) > 0 {
			report.Lines = append(report.Lines, reportstr)
			report.ErrorLineIndexes = append(report.ErrorLineIndexes, line.Index)
		}
	}

	return report, nil
}

type LineReply struct {
	Text string `json:"text"`
	LineNumber uint64 `json:"lineNumber"`
	StartPosition uint64 `json:"startPosition"`
}

type ParserResult struct {
	Data []LineReply `json:"data"`
}