package e2e_test

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"testing"

	qt "github.com/frankban/quicktest"
	"github.com/komuw/kama"
)

var longText = `AT last the sleepy atmosphere was stirred—and vigorously: the murder trial came on in the court. It became the absorbing topic of village talk immediately. Tom could not get away from it. Every reference to the murder sent a shudder to his heart, for his troubled conscience and fears almost persuaded him that these remarks were put forth in his hearing as “feelers”; he did not see how he could be suspected of knowing anything about the murder, but still he could not be comfortable in the midst of this gossip. It kept him in a cold shiver all the time. He took Huck to a lonely place to have a talk with him. It would be some relief to unseal his tongue for a little while; to divide his burden of distress with another sufferer. Moreover, he wanted to assure himself that Huck had remained discreet.
“Huck, have you ever told anybody about—that?”
“’Bout what?”
“You know what.”
“Oh—’course I haven’t.”
“Never a word?”
“Never a solitary word, so help me. What makes you ask?”
“Well, I was afeard.”
“Why, Tom Sawyer, we wouldn’t be alive two days if that got found out. You know that.”
Tom felt more comfortable. After a pause:
“Huck, they couldn’t anybody get you to tell, could they?”
“Get me to tell? Why, if I wanted that halfbreed devil to drownd me they could get me to tell. They ain’t no different way.”
“Well, that’s all right, then. I reckon we’re safe as long as we keep mum. But let’s swear again, anyway. It’s more surer.”
“I’m agreed.”
So they swore again with dread solemnities.
“What is the talk around, Huck? I’ve heard a power of it.”
“Talk? Well, it’s just Muff Potter, Muff Potter, Muff Potter all the time. It keeps me in a sweat, constant, so’s I want to hide som’ers.”
“That’s just the same way they go on round me. I reckon he’s a goner. Don’t you feel sorry for him, sometimes?”
“Most always—most always. He ain’t no account; but then he hain’t ever done anything to hurt anybody. Just fishes a little, to get money to get drunk on—and loafs around considerable; but lord, we all do that—leastways most of us—preachers and such like. But he’s kind of good—he give me half a fish, once, when there warn’t enough for two; and lots of times he’s kind of stood by me when I was out of luck.”
“Well, he’s mended kites for me, Huck, and knitted hooks on to my line. I wish we could get him out of there.”
“My! we couldn’t get him out, Tom. And besides, ’twouldn’t do any good; they’d ketch him again.”
“Yes—so they would. But I hate to hear ’em abuse him so like the dickens when he never done—that.”
“I do too, Tom. Lord, I hear ’em say he’s the bloodiest looking villain in this country, and they wonder he wasn’t ever hung before.”
“Yes, they talk like that, all the time. I’ve heard ’em say that if he was to get free they’d lynch him.”
“And they’d do it, too.”
The boys had a long talk, but it brought them little comfort. As the twilight drew on, they found themselves hanging about the neighborhood of the little isolated jail, perhaps with an undefined hope that something would happen that might clear away their difficulties. But nothing happened; there seemed to be no angels or fairies interested in this luckless captive.
The boys did as they had often done before—went to the cell grating and gave Potter some tobacco and matches. He was on the ground floor and there were no guards.`

type Distance uint64

func bigMap() map[int]string {
	y := map[int]string{}
	for i := 0; i < 10_000; i++ {
		y[i] = fmt.Sprintf("%d", i)
	}
	return y
}

func bigChan() chan int {
	z := make(chan int, 10_000)
	for i := 0; i < 122; i++ {
		// TODO: will be fixed by https://github.com/sanity-io/litter/pull/42
		z <- i
	}
	return z
}

func bigSlice() []int {
	bigSlice := []int{}
	for i := 0; i < 10_000; i++ {
		bigSlice = append(bigSlice, i)
	}
	return bigSlice
}

type FuncWithReturn func(http.ResponseWriter) (uint16, error)

type SomeStruct struct {
	SomeInt            int16
	SomeUintptr        uintptr
	SliceOfHttpRequest []http.Request
	OneHttpRequest     http.Request
	EmptyString        string
	SmallString        string
	LargeString        string
	DistinctType       Distance
	SomeError          error
	LargeSlice         []int
	LargeMap           map[int]string
	UndirectedChan     chan int
	DirectedChan       chan<- bool
	SomeBool           bool

	FuncClosure    func() (io.ReadCloser, error)
	FuncFromStdLib http.HandlerFunc
	FuncWithReturn FuncWithReturn

	ZeroPointerStruct           *url.URL // we won't initiliaze this
	NonZeroPointerStruct        *url.URL
	EvenMoreUrl                 *url.URL
	SliceOfNonZeroPointerStruct []*url.URL
}

func TestDir(t *testing.T) {
	t.Parallel()

	t.Run("slice on its own is not compacted", func(t *testing.T) {
		t.Parallel()
		c := qt.New(t)
		expected := `
[
NAME: []int
KIND: slice
SIGNATURE: [[]int]
FIELDS: []
METHODS: []
SNIPPET: []int{
   0,
   1,
   2,
   3,
   4,
   5,
 ...<9994 more redacted>..}
]
`

		mySlice := bigSlice()

		res := kama.Dir(mySlice)
		c.Assert(res, qt.Equals, expected)
	})

	t.Run("slice in a struct is compacted", func(t *testing.T) {
		t.Parallel()
		c := qt.New(t)
		expected := `
[
NAME: github.com/komuw/kama/e2e_test.some
KIND: struct
SIGNATURE: [e2e_test.some *e2e_test.some]
FIELDS: [
	XX []int 
	]
METHODS: []
SNIPPET: some{
  XX: []int{0,1,2,3,4,5, ...<9994 more redacted>..},
}
]
`
		type some struct {
			XX []int
		}
		s := some{XX: bigSlice()}

		res := kama.Dir(s)
		c.Assert(res, qt.Equals, expected)
	})
}
