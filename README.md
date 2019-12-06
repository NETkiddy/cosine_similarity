# Project
a test similarity implenmentation by cosine similarity

# About text similarity
https://cloud.tencent.com/developer/article/1389446

# How to test
$ go test -v -test.run TestCosineSimilar

=== RUN   TestCosineSimilar

srcWords: [为什么 我 的 眼里 常含 泪水 ， 因为 我 对 这片 土地 爱得 深沉]

dstWords: [我 深沉 的 爱 着 这片 土地 ， 所以 我 的 眼里 常含 泪水]

allWordsMap:map[为什么:1 因为:1 土地:2 对:1 常含:2 我:4 所以:1 泪水:2 深沉:2 爱:1 爱得:1 的:3 眼里:2 着:1 这片:2 ，:2]

allWordsSlice:[因为 这片 爱得 所以 的 ， 眼里 泪水 对 深沉 爱 我 常含 着 为什么 土地]

srcVector:[1 1 1 0 1 1 1 1 1 1 0 2 1 0 1 1]

dstVector:[0 1 0 1 2 1 1 1 0 1 1 2 1 1 0 1]

--- PASS: TestCosineSimilar (0.62s)

    similarity_test.go:20: CosineSimilar score: 0.7660323462854266
    
PASS

Process finished with exit code 0
