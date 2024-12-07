package dao

import (
	"context"
	"net/http"
	"translate/internal/model"
)

type LLMDao struct {
	client *llmClient
}

type llmClient struct {
	Client *http.Client
	Url    string
	Key    string
}

func InitLLM() *LLMDao {
	return &LLMDao{
		client: &llmClient{
			// new http client for llm
			Client: &http.Client{},
		},
	}
}

func (dao *LLMDao) Translate(ctx context.Context, content string) (result *model.LLMResult, err error) {
	// TODO curl llm, demo
	//resp, err := call LLM

	result = &model.LLMResult{
		LLMId:       "xxxx",
		Status:      0,
		ReqContent:  "",
		RespContent: "",
	}
	return
}

func (dao *LLMDao) GetResult(ctx context.Context, llmId string) (result *model.LLMResult, err error) {
	// TODO curl llm, demo
	//resp, err := call LLM
	result = &model.LLMResult{
		LLMId:       "xxxx",
		Status:      1,
		ReqContent:  "123",
		RespContent: "one two three",
	}
	return
}
