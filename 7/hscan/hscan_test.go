// Optional Todo

package hscan

import (
	"testing"
)

func TestGuessSingle(t *testing.T) {
	got := GuessSingle("77f62e3524cd583d698d51fa24fdff4f", "C:/Users/samgo/go/src/github.com/course-materials/materials/lab/7/main/wordlist.txt")
	want := "Nickelback4life"
	if got != want {
		t.Errorf("got %v, wanted %v", got, want)
	}

}

func TestMap(t *testing.T) {
	gotmd5, gotsha := GenHashMaps("C:/Users/samgo/go/src/github.com/course-materials/materials/lab/7/main/wordlist.txt")
	want1 := "map[5f4dcc3b5aa765d61d8327deb882cf99:password 77f62e3524cd583d698d51fa24fdff4f:Nickelback4life c14c9495e50253e317f21db68f273254:Summer2017 dc647eb65e6711e155375218212b3964:Password]"
	want2 := "map[476a319a3fb5e570e614d9b0f6cd20b0d88cc832739634557eb79dd6421665ce:Summer2017 5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8:password 95a5e1547df73abdd4781b6c9e55f3377c15d08884b11738c2727dbd887d4ced:Nickelback4life e7cf3ef4f17c3999a94f2c6f612e8a888e5b1026878e4e19398b23bd38ec221a:Password]"
	if gotmd5 != want1 {
		t.Errorf("got %v, wanted %v", gotmd5, want1)
	} else if gotsha != want2 {
		t.Errorf("got %v, wanted %v", gotsha, want2)
	}

}

func TestMapNoGO(t *testing.T) {
	gotmd5, gotsha := GenHashMap_NoGo("C:/Users/samgo/go/src/github.com/course-materials/materials/lab/7/main/wordlist.txt")
	want1 := "map[5f4dcc3b5aa765d61d8327deb882cf99:password 77f62e3524cd583d698d51fa24fdff4f:Nickelback4life c14c9495e50253e317f21db68f273254:Summer2017 dc647eb65e6711e155375218212b3964:Password]"
	want2 := "map[476a319a3fb5e570e614d9b0f6cd20b0d88cc832739634557eb79dd6421665ce:Summer2017 5e884898da28047151d0e56f8dc6292773603d0d6aabbdd62a11ef721d1542d8:password 95a5e1547df73abdd4781b6c9e55f3377c15d08884b11738c2727dbd887d4ced:Nickelback4life e7cf3ef4f17c3999a94f2c6f612e8a888e5b1026878e4e19398b23bd38ec221a:Password]"
	if gotmd5 != want1 {
		t.Errorf("got %v, wanted %v", gotmd5, want1)
	} else if gotsha != want2 {
		t.Errorf("got %v, wanted %v", gotsha, want2)
	}

}
