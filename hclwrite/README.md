## hclwrite
- hcl2の書き込み専用のパッケージ

### hclfileのパース
```go
filepath := "./hoge.tf"
fileByteData, err := os.ReadFile(filepath)
if err != nil {}

hclFile, diag := hclwrite.ParseConfig(fileByteData, filepath, hcl.InitialPos)
if diag != nil {}
```
### hclfileをパース後のblockの取得
```go
hclFile, diag := hclwrite.ParseConfig(fileByteData, filepath, hcl.InitialPos)
if diag != nil {}

for _, block := range hclFile.Body().Blocks() {}
```
### blockのtypeとlabelを取得
```go
// type => resourceなど
// label => "aws_vpc", "hoge"など
for _, block := range hclFile.Body().Blocks() {
    labels := block.Labels()
    resourceType := block.Type()
    fmt.Printf("type: %s, labels: %v\n", resourceType, labels)
    /*
        <output>
        type: resource, labels: [aws_vpc prod]
        type: resource, labels: [aws_vpc staging]
        type: output, labels: [hoge]
    */

    // labelとtypeを.区切りで文字列結合(block path生成)
    labels := block.Labels()
    resourceType := block.Type()
    var joinString []string
    joinString = append(joinString, resourceType)
    joinString = append(joinString, labels...)
    blockPath := strings.Join(joinString, ".")
    fmt.Println(blockPath)
    /*
        <output>
        resource.aws_vpc.prod
        resource.aws_vpc.staging
        output.hoge
    */
}
```
### 新しいblockとattributeを追加する
```go
for _, block := range hclFile.Body().Blocks() {
    // blockのbodyに新しいblockを追加
    testBlock := block.Body().AppendNewBlock("test", nil)
    testBlock.Body().SetAttributeValue("hoge", cty.StringVal("fuga"))
    /*
        これができる
        test {
            hoge = "huga"
        }
    */
}
```
### 新しいattributeを作成して、aws_vpc.hoge.idみたいなexpressionを追加する
```go
// hcl fileの解析、body, blocksの取得
for _, block := range hclFile.Body().Blocks() {
    block.Body().SetAttributeTraversal("traversal", hcl.Traversal{
        hcl.TraverseRoot{Name: "root"},
        hcl.TraverseAttr{Name: "attr"},
        hcl.TraverseAttr{Name: "hgoe"},
    })
}

// output
resource "aws_vpc" "staging" {
  ~中略~
  traversal = root.attr.hgoe
}
```

### blockの中のattributeを参照する
```go
hclFile, diag := hclwrite.ParseConfig(fileByteData, filepath, hcl.InitialPos)
if diag != nil {}

for _, block := range hclFile.Body().Blocks() {
    for key, attr := range block.Body().Attributes() {}
}
```
### attributeのexpressionを取得
```go
hclFile, diag := hclwrite.ParseConfig(fileByteData, filepath, hcl.InitialPos)
if diag != nil {}

for _, block := range hclFile.Body().Blocks() {
    for key, attr := range block.Body().Attributes() {
        expr := string(attr.Expr().BuildTokens(nil).Bytes()
        fmt.Printf("左辺(attribute): %s, 右辺(expression)%s\n", key, expr)
    }
}
```
### expressionを書き換える
```go
	// hcl fileの解析、body, blocksの取得
	for _, block := range hclFile.Body().Blocks() {
		// blockのbodyのattributeを取得
		for _, attribute := range block.Body().Attributes() {
			// attributeのexpressionを取得
			tokens := attribute.Expr().BuildTokens(nil)
            // 取得してexpressionをbuildTokenに変換しつつ、
            // loopしながら条件にあったものだけを書き換える
			for tokenIndex, token := range tokens {
				hclToken := string(token.Bytes)
				if hclToken == "10.0.0.0/16" {
					tokens[tokenIndex].Bytes = []byte(strings.ReplaceAll(hclToken, hclToken, "127.0.0.0/32"))
				}
			}

			fmt.Println(string(tokens.Bytes()))
		}
	}
```
### formatを整える
```go
hclwrite.Fomat(hclFile.BUildTokens(nil).Byte())
```
