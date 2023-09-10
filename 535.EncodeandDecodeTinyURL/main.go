package main

import "fmt"

// TinyURL is a URL shortening service where you enter a URL such as https://leetcode.com/problems/design-tinyurl and it returns a short URL such as http://tinyurl.com/4e9iAk. Design a class to encode a URL and decode a tiny URL.

// There is no restriction on how your encode/decode algorithm should work. You just need to ensure that a URL can be encoded to a tiny URL and the tiny URL can be decoded to the original URL.

// Implement the Solution class:

//     Solution() Initializes the object of the system.
//     String encode(String longUrl) Returns a tiny URL for the given longUrl.
//     String decode(String shortUrl) Returns the original long URL for the given shortUrl. It is guaranteed that the given shortUrl was encoded by the same object.

// Example 1:

// Input: url = "https://leetcode.com/problems/design-tinyurl"
// Output: "https://leetcode.com/problems/design-tinyurl"

// Explanation:
// Solution obj = new Solution();
// string tiny = obj.encode(url); // returns the encoded tiny url.
// string ans = obj.decode(tiny); // returns the original url after decoding it.

type Codec struct {
	encodeStore map[string]string
	decodeStore map[string]string
	baseUrl     string
}

func Constructor() Codec {
	return Codec{
		encodeStore: make(map[string]string, 0),
		decodeStore: make(map[string]string),
		baseUrl:     "http://tinyurl.com/",
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	if _, ok := this.encodeStore[longUrl]; ok {
		return this.encodeStore[longUrl]
	}

	fmt.Println(len(this.encodeStore) + 1)
	shortUrl := fmt.Sprintf("%s%d", this.baseUrl, len(this.encodeStore)+1)
	this.encodeStore[longUrl] = shortUrl
	this.decodeStore[shortUrl] = longUrl

	return shortUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {

	return this.decodeStore[shortUrl]
}

/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */

func main() {
	longUrl := "https://leetcode.com/problems/design-tinyurl"

	obj := Constructor()
	url := obj.encode(longUrl)
	ans := obj.decode(url)

	fmt.Println(ans)
}
