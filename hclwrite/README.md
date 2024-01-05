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
