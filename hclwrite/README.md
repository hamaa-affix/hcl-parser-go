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

