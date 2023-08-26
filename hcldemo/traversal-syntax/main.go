package main

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclwrite"
)

func main() {
	file := hclwrite.NewEmptyFile()
	body := file.Body()

	body.AppendNewBlock("resource", []string{"aws_s3_bucket_acl", "main"})
	for _, b := range body.Blocks() {

		//aws_s3_bucket.main.idを生成
		b.Body().SetAttributeTraversal("bucket", hcl.Traversal{
			hcl.TraverseRoot{Name: "aws_s3_bucket"},
            hcl.TraverseAttr{Name: "main"},
            hcl.TraverseAttr{Name: "id"},
		})
	}

	by := file.Bytes()
	fmt.Println(string(by))
}
