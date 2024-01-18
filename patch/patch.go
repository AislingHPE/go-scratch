package patch

import (
	"fmt"

	jsonpatch "github.com/evanphx/json-patch"
)

func Main() {
	firstJson := `{
	"A": 1,	
	"B": 2,	
	"C": {
		"D":3
	},
	"E": [
		{"F":4},
		{"F":5}
	],
	"H": [{"I":1}]
}`
	patchJson := `[
{"op":"replace","path":"/A","value":6},
{"op":"replace","path":"/B","value":7},
{"op":"replace","path":"/C/D","value":8},
{"op":"replace","path":"/E/0/F","value":9},
{"op":"replace","path":"/E/1/F","value":10},
{"op":"replace","path":"/A","value":11},

{"op":"add","path":"/E/0/G","value":null},

{"op":"remove","path":"/H/0/I"}
]`

	patch, err := jsonpatch.DecodePatch([]byte(patchJson))
	if err != nil {
		fmt.Println(err)
	}
	newJson, err := patch.Apply([]byte(firstJson))
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(newJson))
}

func Merge() {
	firstJSON := `{"A":10,"B":20} `
	newJson := `{"R":20}`
	newJson3 := `{"A":null,"R":20}`
	merge1, err := jsonpatch.CreateMergePatch([]byte(firstJSON), []byte(newJson))
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(merge))

	merge2, err := jsonpatch.CreateMergePatch([]byte(firstJSON), []byte(newJson3))
	if err != nil {
		fmt.Println(err)
	}
	// fmt.Println(string(merge))
	// setting to a null value is treated the same as deleting, which is why this is bad
	a, _ := jsonpatch.MergePatch([]byte(firstJSON), []byte(merge1))
	fmt.Println(string(a))
	b, _ := jsonpatch.MergePatch([]byte(firstJSON), []byte(merge2))
	fmt.Println(string(b))
}

func NullTest() {
	firstJson := `{"A":10,"B":20} `
	patch1 := `[{"op":"remove","path":"/A"},{"op":"remove","path":"/B"},{"op":"add","path":"/R","value":20}]`
	patch2 := `[{"op":"replace","path":"/A","value":null},{"op":"remove","path":"/B"},{"op":"add","path":"/R","value":20}]`

	jpatch1, err := jsonpatch.DecodePatch([]byte(patch1))
	if err != nil {
		fmt.Println(err)
	}
	jpatch2, _ := jsonpatch.DecodePatch([]byte(patch2))

	newJson1, _ := jpatch1.Apply([]byte(firstJson))
	newJson2, _ := jpatch2.Apply([]byte(firstJson))

	fmt.Println(string(newJson1))
	fmt.Println(string(newJson2))
	// jsonpatch is cool
}
