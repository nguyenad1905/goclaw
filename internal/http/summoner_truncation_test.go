package http

import (
	"context"
	"strings"
	"testing"

	"github.com/nextlevelbuilder/goclaw/internal/providers"
)

// summonStubProvider returns a canned response and records the requested max_tokens.
type summonStubProvider struct {
	resp      providers.ChatResponse
	maxTokens any
}

func (p *summonStubProvider) Chat(_ context.Context, req providers.ChatRequest) (*providers.ChatResponse, error) {
	p.maxTokens = req.Options["max_tokens"]
	r := p.resp
	return &r, nil
}

func (p *summonStubProvider) ChatStream(ctx context.Context, req providers.ChatRequest, _ func(providers.StreamChunk)) (*providers.ChatResponse, error) {
	return p.Chat(ctx, req)
}

func (p *summonStubProvider) DefaultModel() string { return "stub-model" }
func (p *summonStubProvider) Name() string         { return "stub" }

// TestGenerateFiles_RejectsTruncatedResponse ensures a finish_reason=length response
// is not stored as partial files, while a complete response still parses.
func TestGenerateFiles_RejectsTruncatedResponse(t *testing.T) {
	const body = "<file name=\"SOUL.md\">\nsoul\n</file>\n<file name=\"CAPABILITIES.md\">\ncut off"

	stub := &summonStubProvider{resp: providers.ChatResponse{Content: body, FinishReason: "length"}}
	reg := providers.NewRegistry(nil)
	reg.Register(stub)
	s := NewAgentSummoner(nil, reg, nil, nil)

	files, err := s.generateFiles(context.Background(), "stub", "stub-model", "prompt")
	if err == nil || !strings.Contains(err.Error(), "truncated") {
		t.Fatalf("want truncation error, got files=%v err=%v", files, err)
	}
	if stub.maxTokens != summonMaxTokens {
		t.Errorf("max_tokens = %v, want %d", stub.maxTokens, summonMaxTokens)
	}

	stub.resp.FinishReason = "stop"
	files, err = s.generateFiles(context.Background(), "stub", "stub-model", "prompt")
	if err != nil || files["SOUL.md"] != "soul" {
		t.Fatalf("complete response: files=%v err=%v", files, err)
	}
}
