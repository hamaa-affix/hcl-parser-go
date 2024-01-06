package main

import (
	"testing"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/zclconf/go-cty/cty"
)

func TestCreateBlockPath(t *testing.T) {
	block := hclwrite.NewBlock("resource", []string{"datadog_monitor", "test"})

	result := "resource.datadog_monitor.test"
	blockpath := createBlockPath(block)
	if result != blockpath {
		t.Errorf("failed createBlockPath function test\n, result: %s, expected: %s", blockpath, result)
	}
}

func TestCreateNewBlock(t *testing.T) {
	body := hclwrite.NewEmptyFile().Body()
	newBlock := createNewBlock(body, "test", nil)
	if newBlock.Type() != "test" {
		t.Errorf("failed createNewBlock function test\n, result: %s, expected: %s", newBlock.Type(), "test")
	}
}

func TestReWriteToken(t *testing.T) {
	body := hclwrite.NewEmptyFile().Body()
	body.SetAttributeValue("cidr_block", cty.StringVal("10.0.0.0/16"))
	attributes := body.Attributes()

	conditionString := "10.0.0.0/16"
	replaceString := "127.0.0.0.0/16"

	rewriteToken(attributes, conditionString, replaceString)
	tokens := body.GetAttribute("cidr_block").Expr().BuildTokens(nil)
	for _, token := range tokens {
		hclToken := string(token.Bytes)
		if hclToken == conditionString {
			t.Errorf("failed rewriteToken function test\n, result: %s, expected: %s", hclToken, replaceString)
		}
	}
}
