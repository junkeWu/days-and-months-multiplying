package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/pkg/errors"
)

type (
	EndPoint struct {
		Query         string    `json:"query"`
		Variables     Variables `json:"variables"`
		OperationName string    `json:"operationName"`
	}
	Variables struct {
		FloraExternalId string `json:"flora_external_id"`
		Offset          int    `json:"offset"`
		Limit           int    `json:"limit"`
		InvalidAnswer   bool   `json:"invalid_answer"`
	}
)

func main() {
	sources, err := downloadParticipation("f6bd2439-a022-4baf-a66b-590725b0fdf9")
	if err != nil {
		panic(err)
	}

	for _, source := range sources {
		if err := download("./downloads/", source); err != nil {
			panic(err)
		}
	}
}

type Result struct {
	Data struct {
		CubicSurvey struct {
			StatisticAnalysis struct {
				Participation struct {
					ErrorMsg interface{} `json:"error_msg"`
					Result   struct {
						Answers []struct {
							AnswerData []struct {
								Data     string `json:"data"`
								Question struct {
									FloraExternalId string `json:"flora_external_id"`
									Name            string `json:"name"`
									Number          int    `json:"number"`
								} `json:"question"`
								UploadedFiles []struct {
									DownloadURL string `json:"download_url"`
								} `json:"uploaded_files"`
							} `json:"answer_data"`
						} `json:"answers"`
					} `json:"result"`
				} `json:"participation"`
			} `json:"statistic_analysis"`
		} `json:"cubic_survey"`
	} `json:"data"`
}

func downloadParticipation(externalID string) (downSource []DownSource, err error) {
	by, err := json.Marshal(EndPoint{
		Query: "query queryparticipation($flora_external_id:String!,$offset:Int!,$limit:Int!,$invalid_answer:Boolean){cubic_survey{statistic_analysis{participation(flora_external_id:$flora_external_id,offset:$offset,limit:$limit,invalid_answer:$invalid_answer){status error_msg result{page_info{total_count offset}questions{name number kind flora_external_id}answers{flora_external_id submit_time ip_info answer_data{data uploaded_files{file_name download_url}question{flora_external_id number name}}}}}}}}",
		Variables: Variables{
			FloraExternalId: externalID,
			Offset:          0,
			Limit:           10000,
			InvalidAnswer:   false,
		},
		OperationName: "queryparticipation",
	})
	if err != nil {
		return nil, errors.WithStack(err)
	}

	req, err := http.NewRequest(http.MethodPost, "https://survey.sigs.tsinghua.edu.cn/gql/endpoint", bytes.NewBuffer(by))
	if err != nil {
		return nil, errors.WithStack(err)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	defer resp.Body.Close()

	by, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	var result Result
	if err := json.Unmarshal(by, &result); err != nil {
		return nil, errors.WithStack(err)
	}

	answers := result.Data.CubicSurvey.StatisticAnalysis.Participation.Result.Answers
	for _, answer := range answers {
		// var filename []byte
		var args []string
		var downloadUrls []string

		for _, datum := range answer.AnswerData {
			if len(datum.UploadedFiles) == 0 && !strings.HasPrefix(datum.Data, "http") {
				args = append(args, datum.Data)
			}
			for _, file := range datum.UploadedFiles {
				downloadUrls = append(downloadUrls, file.DownloadURL)
			}
		}
		filename := replacer.Replace(strings.Join(args, "-"))
		for i, url := range downloadUrls {
			downSource = append(downSource, DownSource{
				Filename: fmt.Sprintf("%s-%d.png", filename, i),
				Url:      url,
			})
		}
	}
	fmt.Println("downSource", downSource)
	return downSource, nil
}

type DownSource struct {
	Filename string `json:"filename"`
	Url      string `json:"url"`
}

var replacer = strings.NewReplacer(" ", "", "\t", "", "\n", "")

func download(dir string, source DownSource) error {
	if dir == "" {
		dir = "./"
	}

	resp, err := http.Get(source.Url)
	if err != nil {
		return errors.WithStack(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic("resp err")
		return nil
	}
	file, err := os.Create(dir + source.Filename)
	if err != nil {
		return errors.WithStack(err)
	}

	defer file.Close()
	if _, err := io.Copy(file, resp.Body); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
