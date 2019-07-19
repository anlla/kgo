package kgo

import (
	"strings"
	"testing"
)

func TestNl2br(t *testing.T) {
	str := `hello
world!
你好！`
	res := KStr.Nl2br(str)
	if !strings.Contains(res, "<br />") {
		t.Error("Nl2br fail")
		return
	}
	_ = KStr.Nl2br("")
}

func BenchmarkNl2br(b *testing.B) {
	b.ResetTimer()
	str := `hello
world!
你好！`
	for i := 0; i < b.N; i++ {
		_ = KStr.Nl2br(str)
	}
}

func TestStripTags(t *testing.T) {
	str := `
<h1>Hello world!</h1>
<script>alert('你好！')</scripty>
`
	res := KStr.StripTags(str)
	if strings.Contains(res, "<script>") {
		t.Error("StripTags fail")
		return
	}
	_ = KStr.StripTags("")
}

func BenchmarkStripTags(b *testing.B) {
	b.ResetTimer()
	str := `
<h1>Hello world!</h1>
<script>alert('你好！')</scripty>
`
	for i := 0; i < b.N; i++ {
		_ = KStr.StripTags(str)
	}
}

func TestStringMd5(t *testing.T) {
	str := ""
	res1 := KStr.Md5(str, 32)
	res2 := KStr.Md5(str, 16)
	if res1 != "d41d8cd98f00b204e9800998ecf8427e" {
		t.Error("string Md5 fail")
		return
	}
	if !strings.Contains(res1, res2) {
		t.Error("string Md5 fail")
		return
	}
}

func BenchmarkStringMd5(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		_ = KStr.Md5(str, 32)
	}
}

func TestRandomAlpha(t *testing.T) {
	res := KStr.Random(8, RAND_STRING_ALPHA)
	if !KStr.IsLetter(res) {
		t.Error("RandomAlpha fail")
		return
	}
	KStr.Random(0, RAND_STRING_ALPHA)
	KStr.Random(1, 99)
}

func BenchmarkRandomAlpha(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Random(8, RAND_STRING_ALPHA)
	}
}

func TestRandomNumeric(t *testing.T) {
	str := KStr.Random(8, RAND_STRING_NUMERIC)
	if !KStr.IsNumeric(str) {
		t.Error("RandomNumeric fail")
		return
	}
}

func BenchmarkRandomNumeric(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Random(8, RAND_STRING_NUMERIC)
	}
}

func TestRandomAlphanum(t *testing.T) {
	res := KStr.Random(8, RAND_STRING_ALPHANUM)
	if len(res) != 8 {
		t.Error("RandomAlphanum fail")
		return
	}
}

func BenchmarkRandomAlphanum(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Random(8, RAND_STRING_ALPHANUM)
	}
}

func TestRandomSpecial(t *testing.T) {
	res := KStr.Random(8, RAND_STRING_SPECIAL)
	if len(res) != 8 {
		t.Error("RandomSpecial fail")
		return
	}
}

func BenchmarkRandomSpecial(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Random(8, RAND_STRING_SPECIAL)
	}
}

func TestRandomChinese(t *testing.T) {
	res := KStr.Random(8, RAND_STRING_CHINESE)
	if !KStr.IsChinese(res) {
		t.Error("RandomChinese fail")
		return
	}
}

func BenchmarkRandomChinese(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Random(8, RAND_STRING_CHINESE)
	}
}

func TestStrpos(t *testing.T) {
	str := "hello world!"
	res1 := KStr.Strpos(str, "world", 0)
	res2 := KStr.Strpos(str, "World", 0)
	if res1 < 0 || res2 > 0 {
		t.Error("Strpos fail")
		return
	}
	KStr.Strpos("", "world", 0)
	KStr.Strpos(str, "world", -1)
}

func BenchmarkStrpos(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.Strpos(str, "world", 0)
	}
}

func TestStripos(t *testing.T) {
	str := "hello world!"
	res1 := KStr.Stripos(str, "world", 0)
	res2 := KStr.Stripos(str, "World", 0)
	if res1 < 0 || res2 < 0 {
		t.Error("Stripos fail")
		return
	}
	KStr.Stripos("", "world", 0)
	KStr.Stripos(str, "world", -1)
	KStr.Stripos(str, "haha", 0)
}

func BenchmarkStripos(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.Stripos(str, "World", 0)
	}
}

func TestUcfirst(t *testing.T) {
	str := "hello world!"
	res := KStr.Ucfirst(str)
	if res[0] != 'H' {
		t.Error("Ucfirst fail")
		return
	}
	KStr.Ucfirst("")
}

func BenchmarkUcfirst(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.Ucfirst(str)
	}
}

func TestLcfirst(t *testing.T) {
	str := "HELLOW WORLD!"
	res := KStr.Lcfirst(str)
	if res[0] != 'h' {
		t.Error("Lcfirst fail")
		return
	}
	KStr.Lcfirst("")
}

func BenchmarkLcfirst(b *testing.B) {
	b.ResetTimer()
	str := "HELLOW WORLD!"
	for i := 0; i < b.N; i++ {
		KStr.Lcfirst(str)
	}
}

func TestSubstr(t *testing.T) {
	str := "hello world,welcome to golang!"
	res1 := KStr.Substr(str, 5, 10)
	res2 := KStr.Substr(str, 0, -5)
	res3 := KStr.Substr(str, 5, -1)
	res4 := KStr.Substr(str, 5, 0)

	if len(res1) != 10 || res2 != str || !strings.Contains(str, res3) || res4 != "" {
		t.Error("Substr fail")
		return
	}
	KStr.Substr(str, 10, 50)
}

func BenchmarkSubstr(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.Substr(str, 5, 10)
	}
}

func TestStrrev(t *testing.T) {
	str := "hello world!"
	res1 := KStr.Strrev(str)
	res2 := KStr.Strrev(res1)
	if res2 != str {
		t.Error("Strrev fail")
		return
	}
}

func BenchmarkStrrev(b *testing.B) {
	b.ResetTimer()
	str := "hello world!"
	for i := 0; i < b.N; i++ {
		KStr.Strrev(str)
	}
}

func TestChunkSplit(t *testing.T) {
	str := "Yar?m kilo ?ay, yar?m kilo ?eker"
	res := KStr.ChunkSplit(str, 4, "\r\n")
	if len(res) == 0 {
		t.Error("ChunkSplit fail")
		return
	}
	_ = KStr.ChunkSplit(str, 4, "")
	_ = KStr.ChunkSplit("a", 4, "")
	_ = KStr.ChunkSplit("ab", 64, "")
}

func BenchmarkChunkSplit(b *testing.B) {
	b.ResetTimer()
	str := "Yar?m kilo ?ay, yar?m kilo ?eker"
	for i := 0; i < b.N; i++ {
		KStr.ChunkSplit(str, 4, "")
	}
}

func TestStrlen(t *testing.T) {
	str := "hello world!你好 世界！"
	res := KStr.Strlen(str)
	if res != 28 {
		t.Error("Strlen fail")
		return
	}
}

func BenchmarkStrlen(b *testing.B) {
	b.ResetTimer()
	str := "hello world!你好 世界！"
	for i := 0; i < b.N; i++ {
		KStr.Strlen(str)
	}
}

func TestMbStrlen(t *testing.T) {
	str := "hello world!你好 世界！"
	res := KStr.MbStrlen(str)
	if res != 18 {
		t.Error("MbStrlen fail")
		return
	}
}

func BenchmarkMbStrlen(b *testing.B) {
	b.ResetTimer()
	str := "hello world!你好 世界！"
	for i := 0; i < b.N; i++ {
		KStr.MbStrlen(str)
	}
}

func TestMbStrShuffle(t *testing.T) {
	str := "hello world!你好 世界！"
	res := KStr.StrShuffle(str)
	if res == str {
		t.Error("StrShuffle fail")
		return
	}
}

func BenchmarkStrShuffle(b *testing.B) {
	b.ResetTimer()
	str := "hello world!你好 世界！"
	for i := 0; i < b.N; i++ {
		KStr.StrShuffle(str)
	}
}

func TestTrim(t *testing.T) {
	str := " hello world!你好 世界！　"
	res := KStr.Trim(str)
	if res[0] != 'h' {
		t.Error("Trim fail")
		return
	}
	KStr.Trim(str, "\n")
}

func BenchmarkTrim(b *testing.B) {
	b.ResetTimer()
	str := " hello world!你好 世界！　"
	for i := 0; i < b.N; i++ {
		KStr.Trim(str)
	}
}

func TestLtrim(t *testing.T) {
	str := " hello world!你好 世界！　"
	res := KStr.Ltrim(str)
	if res[0] != 'h' {
		t.Error("Ltrim fail")
		return
	}
	KStr.Ltrim(str, "\n")
}

func BenchmarkLtrim(b *testing.B) {
	b.ResetTimer()
	str := " hello world!你好 世界！　"
	for i := 0; i < b.N; i++ {
		KStr.Ltrim(str)
	}
}

func TestRtrim(t *testing.T) {
	str := " hello world!你好 世界！　"
	res := KStr.Rtrim(str, "　")
	if strings.HasSuffix(res, "　") {
		t.Error("Rtrim fail")
		return
	}
	KStr.Rtrim(str)
}

func BenchmarkRtrim(b *testing.B) {
	b.ResetTimer()
	str := " hello world!你好 世界！　"
	for i := 0; i < b.N; i++ {
		KStr.Rtrim(str)
	}
}

func TestChr(t *testing.T) {
	res := KStr.Chr(65)
	if res != "A" {
		t.Error("Chr fail")
		return
	}
}

func BenchmarkChr(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Chr(int(i))
	}
}

func TestOrd(t *testing.T) {
	res := KStr.Ord("b")
	if res != 98 {
		t.Error("Ord fail")
		return
	}
}

func BenchmarkOrd(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		KStr.Ord("c")
	}
}

func TestJsonEncodeDecode(t *testing.T) {
	obj := make(map[string]interface{})
	obj["k1"] = "abc"
	obj["k2"] = 123
	obj["k3"] = false
	jstr, err := KStr.JsonEncode(obj)
	if err != nil {
		t.Error("JsonEncode fail")
		return
	}

	mp := make(map[string]interface{})
	err2 := KStr.JsonDecode(jstr, &mp)
	if err2 != nil {
		t.Error("JsonDecode fail")
		return
	}
}

func BenchmarkJsonEncode(b *testing.B) {
	b.ResetTimer()
	obj := make(map[string]interface{})
	obj["k1"] = "abc"
	obj["k2"] = 123
	obj["k3"] = false
	for i := 0; i < b.N; i++ {
		_, _ = KStr.JsonEncode(obj)
	}
}

func BenchmarkJsonDecode(b *testing.B) {
	b.ResetTimer()
	str := []byte(`{"k1":"abc","k2":123,"k3":false}`)
	mp := make(map[string]interface{})
	for i := 0; i < b.N; i++ {
		_ = KStr.JsonDecode(str, &mp)
	}
}

func TestAddslashesStripslashes(t *testing.T) {
	str := "Is your name O'reilly?"
	res1 := KStr.Addslashes(str)
	if !strings.Contains(res1, "\\") {
		t.Error("Addslashes fail")
		return
	}

	res2 := KStr.Stripslashes(res1)
	if strings.Contains(res2, "\\") {
		t.Error("Stripslashes fail")
		return
	}
	KStr.Stripslashes(`Is \ your \\name O\'reilly?`)
}

func BenchmarkAddslashes(b *testing.B) {
	b.ResetTimer()
	str := "Is your name O'reilly?"
	for i := 0; i < b.N; i++ {
		KStr.Addslashes(str)
	}
}

func BenchmarkStripslashes(b *testing.B) {
	b.ResetTimer()
	str := `Is your name O\'reilly?`
	for i := 0; i < b.N; i++ {
		KStr.Stripslashes(str)
	}
}

func TestQuotemeta(t *testing.T) {
	str := "Hello world. (can you hear me?)"
	res := KStr.Quotemeta(str)
	if !strings.Contains(res, "\\") {
		t.Error("Quotemeta fail")
		return
	}
}

func BenchmarkQuotemeta(b *testing.B) {
	b.ResetTimer()
	str := "Hello world. (can you hear me?)"
	for i := 0; i < b.N; i++ {
		KStr.Quotemeta(str)
	}
}

func TestHtmlentitiesEncodeDecode(t *testing.T) {
	str := "A 'quote' is <b>bold</b>"
	res1 := KStr.Htmlentities(str)
	if !strings.Contains(res1, "&") {
		t.Error("Htmlentities fail")
		return
	}

	res2 := KStr.HtmlentityDecode(res1)
	if res2 != str {
		t.Error("HtmlentityDecode fail")
		return
	}
}

func BenchmarkHtmlentities(b *testing.B) {
	b.ResetTimer()
	str := "A 'quote' is <b>bold</b>"
	for i := 0; i < b.N; i++ {
		KStr.Htmlentities(str)
	}
}

func BenchmarkHtmlentityDecode(b *testing.B) {
	b.ResetTimer()
	str := `A &#39;quote&#39; is &lt;b&gt;bold&lt;/b&gt;`
	for i := 0; i < b.N; i++ {
		KStr.HtmlentityDecode(str)
	}
}

func TestStringSha1(t *testing.T) {
	str := "apple"
	res := KStr.Sha1(str)
	if res != "d0be2dc421be4fcd0172e5afceea3970e2f3d940" {
		t.Error("String Sha1 fail")
		return
	}
}

func BenchmarkStringSha1(b *testing.B) {
	b.ResetTimer()
	str := "Hello world. (can you hear me?)"
	for i := 0; i < b.N; i++ {
		KStr.Sha1(str)
	}
}

func TestStringSha256(t *testing.T) {
	str := "apple"
	res := KStr.Sha256(str)
	if res != "3a7bd3e2360a3d29eea436fcfb7e44c735d117c42d1c1835420b6b9942dd4f1b" {
		t.Error("String Sha256 fail")
		return
	}
}

func BenchmarkStringSha256(b *testing.B) {
	b.ResetTimer()
	str := "Hello world. (can you hear me?)"
	for i := 0; i < b.N; i++ {
		KStr.Sha256(str)
	}
}

func TestCrc32(t *testing.T) {
	str := "The quick brown fox jumped over the lazy dog"
	res := KStr.Crc32(str)
	if res <= 0 {
		t.Error("Crc32 fail")
		return
	}
}

func BenchmarkCrc32(b *testing.B) {
	b.ResetTimer()
	str := "The quick brown fox jumped over the lazy dog"
	for i := 0; i < b.N; i++ {
		KStr.Crc32(str)
	}
}

func TestSimilarText(t *testing.T) {
	str1 := "The quick brown fox jumped over the lazy dog"
	str2 := "The quick brown fox jumped over the lazy dog"
	var percent float64

	res := KStr.SimilarText(str1, str2, &percent)
	if res <= 0 || percent <= 0 {
		t.Error("Crc32 fail")
		return
	}
	KStr.SimilarText("PHP IS GREAT", "WITH MYSQL", &percent)
	KStr.SimilarText("", "", &percent)
}

func BenchmarkSimilarText(b *testing.B) {
	b.ResetTimer()
	str1 := "The quick brown fox jumped over the lazy dog"
	str2 := "The quick brown fox jumped over the lazy dog"
	var percent float64
	for i := 0; i < b.N; i++ {
		KStr.SimilarText(str1, str2, &percent)
	}
}
