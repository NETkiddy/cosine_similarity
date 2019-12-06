package cosine_similarity

import (
	"fmt"
	"testing"
)

func TestCosineSimilar(t *testing.T) {
	g := NewGoJieba()
	srcStr := "为什么我的眼里常含泪水，因为我对这片土地爱得深沉"
	srcStr = RemoveHtml(srcStr)
	srcWords := g.C.Cut(srcStr, true)
	dstStr := "我深沉的爱着这片土地，所以我的眼里常含泪水"
	dstStr = RemoveHtml(dstStr)
	dstWords := g.C.Cut(dstStr, true)

	fmt.Printf("srcWords: %v\ndstWords: %v\n", srcWords, dstWords)
	score := CosineSimilar(srcWords, dstWords)

	t.Logf("CosineSimilar score: %v", score)
}
