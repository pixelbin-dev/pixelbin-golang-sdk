package tests

import (
	"fmt"
	"strings"
	"testing"

	"github.com/pixelbin-io/pixelbin-go/v3/sdk/utils/url"
)

var urls_to_obj = []struct {
	scenario string
	url      string
	opts     []url.UrlToObjOption
	obj      map[string]interface{}
}{
	{
		scenario: "PixelBin CDN with no UrlToObjOption",
		url:      "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/original/W2.jpeg",
		opts:     []url.UrlToObjOption{},
		obj: map[string]interface{}{
			"version":         "v2",
			"cloudName":       "broken-butterfly-3b12f1",
			"pattern":         "original",
			"filePath":        "W2.jpeg",
			"options":         map[string]interface{}{},
			"zone":            nil,
			"baseUrl":         "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{},
			"worker":          false,
			"workerPath":      "",
		},
	},
	{
		scenario: "PixelBin CDN default zone original asset",
		url:      "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/original/W2.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(false),
		},
		obj: map[string]interface{}{
			"version":         "v2",
			"cloudName":       "broken-butterfly-3b12f1",
			"pattern":         "original",
			"filePath":        "W2.jpeg",
			"options":         map[string]interface{}{},
			"zone":            nil,
			"baseUrl":         "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{},
			"worker":          false,
			"workerPath":      "",
		},
	},
	{
		scenario: "PixelBin CDN default zone single transformation",
		url:      "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/t.resize(h:600,w:800)/W2.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(false),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"pattern":   "t.resize(h:600,w:800)",
			"filePath":  "W2.jpeg",
			"options":   map[string]interface{}{},
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "600"},
						{"key": "w", "value": "800"},
					},
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "PixelBin CDN zone with single transformation",
		url:      "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/z-slug/t.resize(h:600,w:800)/W2.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(false),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"pattern":   "t.resize(h:600,w:800)",
			"filePath":  "W2.jpeg",
			"options":   map[string]interface{}{},
			"zone":      "z-slug",
			"baseUrl":   "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "600"},
						{"key": "w", "value": "800"},
					},
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "PixelBin CDN default zone multiple transformations",
		url:      "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/t.resize(h:600,w:800)~t.rotate(a:-249)/W2.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(false),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"pattern":   "t.resize(h:600,w:800)~t.rotate(a:-249)",
			"filePath":  "W2.jpeg",
			"options":   map[string]interface{}{},
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "600"},
						{"key": "w", "value": "800"},
					},
				},
				{
					"plugin": "t",
					"name":   "rotate",
					"values": []map[string]interface{}{
						{"key": "a", "value": "-249"},
					},
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "PixelBin CDN default zone multiple transformations with default values",
		url:      "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/t.resize(h:600,w:800)~t.rotate(a:-249)~t.flip()~t.trim(t:217)/W2.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(false),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"pattern":   "t.resize(h:600,w:800)~t.rotate(a:-249)~t.flip()~t.trim(t:217)",
			"filePath":  "W2.jpeg",
			"options":   map[string]interface{}{},
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "600"},
						{"key": "w", "value": "800"},
					},
				},
				{
					"plugin": "t",
					"name":   "rotate",
					"values": []map[string]interface{}{
						{"key": "a", "value": "-249"},
					},
				},
				{
					"plugin": "t",
					"name":   "flip",
				},
				{
					"plugin": "t",
					"name":   "trim",
					"values": []map[string]interface{}{
						{"key": "t", "value": "217"},
					},
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "PixelBin CDN default zone multiple transformations & presets with variables",
		url:      "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/t.resize(h:200,w:100)~p:preset1(a:100,b:2.1,c:test)/W2.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(false),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"pattern":   "t.resize(h:200,w:100)~p:preset1(a:100,b:2.1,c:test)",
			"filePath":  "W2.jpeg",
			"options":   map[string]interface{}{},
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
					"values": []map[string]interface{}{
						{"key": "a", "value": "100"},
						{"key": "b", "value": "2.1"},
						{"key": "c", "value": "test"},
					},
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "PixelBin CDN default zone multiple transformations & presets with no paranthesis",
		url:      "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/t.resize(h:200,w:100)~p:preset1/W2.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(false),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"pattern":   "t.resize(h:200,w:100)~p:preset1",
			"filePath":  "W2.jpeg",
			"options":   map[string]interface{}{},
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "PixelBin CDN default zone multiple transformations & presets default values",
		url:      "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/t.resize(h:200,w:100)~p:preset1()/W2.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(false),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"pattern":   "t.resize(h:200,w:100)~p:preset1()",
			"filePath":  "W2.jpeg",
			"options":   map[string]interface{}{},
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "PixelBin CDN default zone with options",
		url:      "https://cdn.pixelbin.io/v2/feel/erase.bg(shadow:true)~t.flip()/MZZKB3e1hT48o0NYJ2Kxh.jpeg?dpr=2.5&f_auto=true",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(false),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"baseUrl":   "https://cdn.pixelbin.io",
			"filePath":  "MZZKB3e1hT48o0NYJ2Kxh.jpeg",
			"pattern":   "erase.bg(shadow:true)~t.flip()",
			"cloudName": "feel",
			"options": map[string]interface{}{
				"dpr":    2.5,
				"f_auto": true,
			},
			"zone": nil,
			"transformations": []map[string]interface{}{
				{
					"plugin": "erase",
					"name":   "bg",
					"values": []map[string]interface{}{
						{"key": "shadow", "value": "true"},
					},
				},
				{
					"plugin": "t",
					"name":   "flip",
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "PixelBin CDN default zone with dpr auto",
		url:      "https://cdn.pixelbin.io/v2/feel/erase.bg()/MZZKB3e1hT48o0NYJ2Kxh.jpeg?dpr=auto",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(false),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"baseUrl":   "https://cdn.pixelbin.io",
			"filePath":  "MZZKB3e1hT48o0NYJ2Kxh.jpeg",
			"pattern":   "erase.bg()",
			"cloudName": "feel",
			"options": map[string]interface{}{
				"dpr": "auto",
			},
			"zone": nil,
			"transformations": []map[string]interface{}{
				{
					"plugin": "erase",
					"name":   "bg",
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "PixelBin CDN default zone with f_auto false",
		url:      "https://cdn.pixelbin.io/v2/feel/erase.bg()/asbc.jpeg?f_auto=false",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(false),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"baseUrl":   "https://cdn.pixelbin.io",
			"filePath":  "asbc.jpeg",
			"pattern":   "erase.bg()",
			"cloudName": "feel",
			"options": map[string]interface{}{
				"f_auto": false,
			},
			"zone": nil,
			"transformations": []map[string]interface{}{
				{
					"plugin": "erase",
					"name":   "bg",
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "PixelBin CDN default zone worker",
		url:      "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/wrkr/image.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(false),
		},
		obj: map[string]interface{}{
			"version":         "v2",
			"cloudName":       "broken-butterfly-3b12f1",
			"pattern":         "",
			"filePath":        "",
			"options":         map[string]interface{}{},
			"zone":            nil,
			"baseUrl":         "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{},
			"worker":          true,
			"workerPath":      "image.jpeg",
		},
	},
	{
		scenario: "PixelBin CDN default zone worker depth > 1",
		url:      "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/wrkr/resize:w200,h200/image.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(false),
		},
		obj: map[string]interface{}{
			"version":         "v2",
			"cloudName":       "broken-butterfly-3b12f1",
			"pattern":         "",
			"filePath":        "",
			"options":         map[string]interface{}{},
			"zone":            nil,
			"baseUrl":         "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{},
			"worker":          true,
			"workerPath":      "resize:w200,h200/image.jpeg",
		},
	},
	{
		scenario: "PixelBin CDN zone with worker",
		url:      "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/abcdef/wrkr/image.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(false),
		},
		obj: map[string]interface{}{
			"version":         "v2",
			"cloudName":       "broken-butterfly-3b12f1",
			"pattern":         "",
			"filePath":        "",
			"options":         map[string]interface{}{},
			"zone":            "abcdef",
			"baseUrl":         "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{},
			"worker":          true,
			"workerPath":      "image.jpeg",
		},
	},
	{
		scenario: "PixelBin CDN zone with worker depth > 1",
		url:      "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/abcdef/wrkr/resize:w200,h200/image.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(false),
		},
		obj: map[string]interface{}{
			"version":         "v2",
			"cloudName":       "broken-butterfly-3b12f1",
			"pattern":         "",
			"filePath":        "",
			"options":         map[string]interface{}{},
			"zone":            "abcdef",
			"baseUrl":         "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{},
			"worker":          true,
			"workerPath":      "resize:w200,h200/image.jpeg",
		},
	},

	// custom domain

	{
		scenario: "Custom Domain default zone original asset",
		url:      "https://cdn.twist.vision/v2/original/W2.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(true),
		},
		obj: map[string]interface{}{
			"version":         "v2",
			"cloudName":       nil,
			"pattern":         "original",
			"filePath":        "W2.jpeg",
			"options":         map[string]interface{}{},
			"zone":            nil,
			"baseUrl":         "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{},
			"worker":          false,
			"workerPath":      "",
		},
	},
	{
		scenario: "Custom Domain default zone single transformation",
		url:      "https://cdn.twist.vision/v2/t.resize(h:600,w:800)/W2.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(true),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"cloudName": nil,
			"pattern":   "t.resize(h:600,w:800)",
			"filePath":  "W2.jpeg",
			"options":   map[string]interface{}{},
			"zone":      nil,
			"baseUrl":   "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{
							"key":   "h",
							"value": "600",
						},
						{
							"key":   "w",
							"value": "800",
						},
					},
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "Custom Domain zone with single transformation",
		url:      "https://cdn.twist.vision/v2/z-slug/t.resize(h:600,w:800)/W2.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(true),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"cloudName": nil,
			"pattern":   "t.resize(h:600,w:800)",
			"filePath":  "W2.jpeg",
			"options":   map[string]interface{}{},
			"zone":      "z-slug",
			"baseUrl":   "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{
							"key":   "h",
							"value": "600",
						},
						{
							"key":   "w",
							"value": "800",
						},
					},
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "Custom Domain default zone multiple transformations",
		url:      "https://cdn.twist.vision/v2/t.resize(h:600,w:800)~t.rotate(a:-249)/W2.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(true),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"cloudName": nil,
			"pattern":   "t.resize(h:600,w:800)~t.rotate(a:-249)",
			"filePath":  "W2.jpeg",
			"options":   map[string]interface{}{},
			"zone":      nil,
			"baseUrl":   "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{
							"key":   "h",
							"value": "600",
						},
						{
							"key":   "w",
							"value": "800",
						},
					},
				},
				{
					"plugin": "t",
					"name":   "rotate",
					"values": []map[string]interface{}{
						{
							"key":   "a",
							"value": "-249",
						},
					},
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "Custom Domain default zone multiple transformations with default values",
		url:      "https://cdn.twist.vision/v2/t.resize(h:600,w:800)~t.rotate(a:-249)~t.flip()~t.trim(t:217)/W2.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(true),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"cloudName": nil,
			"pattern":   "t.resize(h:600,w:800)~t.rotate(a:-249)~t.flip()~t.trim(t:217)",
			"filePath":  "W2.jpeg",
			"options":   map[string]interface{}{},
			"zone":      nil,
			"baseUrl":   "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{
							"key":   "h",
							"value": "600",
						},
						{
							"key":   "w",
							"value": "800",
						},
					},
				},
				{
					"plugin": "t",
					"name":   "rotate",
					"values": []map[string]interface{}{
						{
							"key":   "a",
							"value": "-249",
						},
					},
				},
				{
					"plugin": "t",
					"name":   "flip",
				},
				{
					"plugin": "t",
					"name":   "trim",
					"values": []map[string]interface{}{
						{
							"key":   "t",
							"value": "217",
						},
					},
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "Custom Domain default zone multiple transformations & presets with variables",
		url:      "https://cdn.twist.vision/v2/t.resize(h:200,w:100)~p:preset1(a:100,b:2.1,c:test)/W2.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(true),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"cloudName": nil,
			"pattern":   "t.resize(h:200,w:100)~p:preset1(a:100,b:2.1,c:test)",
			"filePath":  "W2.jpeg",
			"options":   map[string]interface{}{},
			"zone":      nil,
			"baseUrl":   "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{
							"key":   "h",
							"value": "200",
						},
						{
							"key":   "w",
							"value": "100",
						},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
					"values": []map[string]interface{}{
						{
							"key":   "a",
							"value": "100",
						},
						{
							"key":   "b",
							"value": "2.1",
						},
						{
							"key":   "c",
							"value": "test",
						},
					},
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "Custom Domain default zone multiple transformations & presets with no paranthesis",
		url:      "https://cdn.twist.vision/v2/t.resize(h:200,w:100)~p:preset1/W2.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(true),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"cloudName": nil,
			"pattern":   "t.resize(h:200,w:100)~p:preset1",
			"filePath":  "W2.jpeg",
			"options":   map[string]interface{}{},
			"zone":      nil,
			"baseUrl":   "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{
							"key":   "h",
							"value": "200",
						},
						{
							"key":   "w",
							"value": "100",
						},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "Custom Domain default zone multiple transformations & presets default values",
		url:      "https://cdn.twist.vision/v2/t.resize(h:200,w:100)~p:preset1()/W2.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(true),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"cloudName": nil,
			"pattern":   "t.resize(h:200,w:100)~p:preset1()",
			"filePath":  "W2.jpeg",
			"options":   map[string]interface{}{},
			"zone":      nil,
			"baseUrl":   "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{
							"key":   "h",
							"value": "200",
						},
						{
							"key":   "w",
							"value": "100",
						},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "Custom Domain default zone with 6 character length path segment",
		url:      "https://cdn.twist.vision/v2/original/z0/orgs/33/skills/icons/Fynd_Platform_Commerce_Extension.png",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(true),
		},
		obj: map[string]interface{}{
			"version":         "v2",
			"baseUrl":         "https://cdn.twist.vision",
			"filePath":        "z0/orgs/33/skills/icons/Fynd_Platform_Commerce_Extension.png",
			"pattern":         "original",
			"cloudName":       nil,
			"options":         map[string]interface{}{},
			"zone":            nil,
			"transformations": []map[string]interface{}{},
			"worker":          false,
			"workerPath":      "",
		},
	},
	{
		scenario: "Custom Domain default zone with options",
		url:      "https://cdn.twist.vision/v2/erase.bg(shadow:true)~t.flip()/MZZKB3e1hT48o0NYJ2Kxh.jpeg?dpr=2.5&f_auto=true",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(true),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"baseUrl":   "https://cdn.twist.vision",
			"filePath":  "MZZKB3e1hT48o0NYJ2Kxh.jpeg",
			"pattern":   "erase.bg(shadow:true)~t.flip()",
			"cloudName": nil,
			"options": map[string]interface{}{
				"dpr":    2.5,
				"f_auto": true,
			},
			"zone": nil,
			"transformations": []map[string]interface{}{
				{
					"plugin": "erase",
					"name":   "bg",
					"values": []map[string]interface{}{
						{
							"key":   "shadow",
							"value": "true",
						},
					},
				},
				{
					"plugin": "t",
					"name":   "flip",
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "Custom Domain default zone with dpr auto",
		url:      "https://cdn.twist.vision/v2/erase.bg()/MZZKB3e1hT48o0NYJ2Kxh.jpeg?dpr=auto",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(true),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"baseUrl":   "https://cdn.twist.vision",
			"filePath":  "MZZKB3e1hT48o0NYJ2Kxh.jpeg",
			"pattern":   "erase.bg()",
			"cloudName": nil,
			"options": map[string]interface{}{
				"dpr": "auto",
			},
			"zone": nil,
			"transformations": []map[string]interface{}{
				{
					"plugin": "erase",
					"name":   "bg",
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "Custom Domain default zone with f_auto false",
		url:      "https://cdn.twist.vision/v2/erase.bg()/asbc.jpeg?f_auto=false",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(true),
		},
		obj: map[string]interface{}{
			"version":   "v2",
			"baseUrl":   "https://cdn.twist.vision",
			"filePath":  "asbc.jpeg",
			"pattern":   "erase.bg()",
			"cloudName": nil,
			"options": map[string]interface{}{
				"f_auto": false,
			},
			"zone": nil,
			"transformations": []map[string]interface{}{
				{
					"plugin": "erase",
					"name":   "bg",
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		scenario: "Custom Domain default zone worker",
		url:      "https://cdn.twist.vision/v2/wrkr/image.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(true),
		},
		obj: map[string]interface{}{
			"version":         "v2",
			"cloudName":       nil,
			"pattern":         "",
			"filePath":        "",
			"options":         map[string]interface{}{},
			"zone":            nil,
			"baseUrl":         "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{},
			"worker":          true,
			"workerPath":      "image.jpeg",
		},
	},
	{
		scenario: "Custom Domain default zone worker depth > 1",
		url:      "https://cdn.twist.vision/v2/wrkr/resize:w200,h200/image.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(true),
		},
		obj: map[string]interface{}{
			"version":         "v2",
			"cloudName":       nil,
			"pattern":         "",
			"filePath":        "",
			"options":         map[string]interface{}{},
			"zone":            nil,
			"baseUrl":         "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{},
			"worker":          true,
			"workerPath":      "resize:w200,h200/image.jpeg",
		},
	},
	{
		scenario: "Custom Domain zone with worker",
		url:      "https://cdn.twist.vision/v2/abcdef/wrkr/image.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(true),
		},
		obj: map[string]interface{}{
			"version":         "v2",
			"cloudName":       nil,
			"pattern":         "",
			"filePath":        "",
			"options":         map[string]interface{}{},
			"zone":            "abcdef",
			"baseUrl":         "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{},
			"worker":          true,
			"workerPath":      "image.jpeg",
		},
	},
	{
		scenario: "Custom Domain zone with worker depth > 1",
		url:      "https://cdn.twist.vision/v2/abcdef/wrkr/resize:w200,h200/image.jpeg",
		opts: []url.UrlToObjOption{
			url.WithCustomDomain(true),
		},
		obj: map[string]interface{}{
			"version":         "v2",
			"cloudName":       nil,
			"pattern":         "",
			"filePath":        "",
			"options":         map[string]interface{}{},
			"zone":            "abcdef",
			"baseUrl":         "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{},
			"worker":          true,
			"workerPath":      "resize:w200,h200/image.jpeg",
		},
	},
}

var obj_to_urls = []map[string]interface{}{
	{
		"url": "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/t.resize(h:600,w:800)/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"options":   map[string]interface{}{},
			"pattern":   "t.resize(h:600,w:800)",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "600"},
						{"key": "w", "value": "800"},
					},
				},
			},
		},
	},
	{
		"url": "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/z-slug/t.resize(h:600,w:800)/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"options":   map[string]interface{}{},
			"pattern":   "t.resize(h:600,w:800)",
			"filePath":  "W2.jpeg",
			"zone":      "z-slug",
			"baseUrl":   "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "600"},
						{"key": "w", "value": "800"},
					},
				},
			},
		},
	},
	{
		"url": "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/t.resize(h:600,w:800)~t.rotate(a:-249)/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"options":   map[string]interface{}{},
			"pattern":   "t.resize(h:600,w:800)~t.rotate(a:-249)",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "600"},
						{"key": "w", "value": "800"},
					},
				},
				{
					"plugin": "t",
					"name":   "rotate",
					"values": []map[string]interface{}{
						{"key": "a", "value": "-249"},
					},
				},
			},
		},
	},
	{
		"url": "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/t.resize(h:600,w:800)~t.rotate(a:-249)~t.flip()~t.trim(t:217)/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"options":   map[string]interface{}{},
			"pattern":   "t.resize(h:600,w:800)~t.rotate(a:-249)~t.flip()~t.trim(t:217)",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "600"},
						{"key": "w", "value": "800"},
					},
				},
				{
					"plugin": "t",
					"name":   "rotate",
					"values": []map[string]interface{}{
						{"key": "a", "value": "-249"},
					},
				},
				{
					"plugin": "t",
					"name":   "flip",
				},
				{
					"plugin": "t",
					"name":   "trim",
					"values": []map[string]interface{}{
						{"key": "t", "value": "217"},
					},
				},
			},
		},
	},
	{
		"url": "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/t.resize(h:200,w:100)~p:preset1(a:100,b:2.1,c:test)/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"options":   map[string]interface{}{},
			"pattern":   "t.resize(h:200,w:100)~p:preset1(a:100,b:2.1,c:test)",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
					"values": []map[string]interface{}{
						{"key": "a", "value": "100"},
						{"key": "b", "value": "2.1"},
						{"key": "c", "value": "test"},
					},
				},
			},
		},
	},
	{
		"url": "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/t.resize(h:200,w:100)~p:preset1/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"options":   map[string]interface{}{},
			"pattern":   "t.resize(h:200,w:100)~p:preset1",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
				},
			},
		},
	},
	{
		"url": "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/t.resize(h:200,w:100)~p:preset1(a:12)/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"options":   map[string]interface{}{},
			"pattern":   "t.resize(h:200,w:100)~p:preset1(a:12)",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
					"values": []map[string]interface{}{
						{"key": "a", "value": "12"},
					},
				},
			},
		},
	},
	{
		"url": "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/t.resize(h:200,w:100)~p:preset1/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"options":   map[string]interface{}{},
			"pattern":   "t.resize(h:200,w:100)~p:preset1(a:12)",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
					"values": []map[string]interface{}{},
				},
			},
		},
	},
	{
		"url": "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/t.resize(h:200,w:100)~p:preset1/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"options":   map[string]interface{}{},
			"pattern":   "t.resize(h:,w:100)~p:preset1(a:12)",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
					"values": []map[string]interface{}{},
				},
			},
		},
		"error": "value not specified",
	},
	{
		"url": "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/t.resize(h:200,w:100)~p:preset1/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"options":   map[string]interface{}{},
			"pattern":   "t.resize(h:200,w:100)~p:preset1(a:12)",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
					"values": []map[string]interface{}{},
				},
			},
		},
		"error": "key not specified",
	},
	{
		"url": "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/t.resize(h:200,w:100)~p:preset1/W2.jpeg",
		"obj": map[string]interface{}{
			"version":   "v2",
			"cloudName": "broken-butterfly-3b12f1",
			"options":   map[string]interface{}{},
			"pattern":   "t.resize(h:200,w:100)~p:preset1(a:12)",
			"filePath":  "W2.jpeg",
			"zone":      nil,
			"baseUrl":   "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
					"values": []map[string]interface{}{},
				},
			},
		},
		"error": "key not specified",
	},
	{
		"url": "https://cdn.pixelbin.io/v2/feel/erase.bg(shadow:true)~t.merge(m:underlay,i:eU44YkFJOHlVMmZrWVRDOUNTRm1D,b:screen,r:true)/MZZKB3e1hT48o0NYJ2Kxh.jpeg?dpr=2&f_auto=true",
		"obj": map[string]interface{}{
			"version":   "v2",
			"baseUrl":   "https://cdn.pixelbin.io",
			"filePath":  "MZZKB3e1hT48o0NYJ2Kxh.jpeg",
			"pattern":   "erase.bg(shadow:true)~t.merge(m:underlay,i:eU44YkFJOHlVMmZrWVRDOUNTRm1D,b:screen,r:true)",
			"cloudName": "feel",
			"options": map[string]interface{}{
				"dpr":    2,
				"f_auto": true,
			},
			"zone": nil,
			"transformations": []map[string]interface{}{
				{
					"values": []map[string]interface{}{
						{
							"key":   "shadow",
							"value": "true",
						},
					},
					"plugin": "erase",
					"name":   "bg",
				},
				{
					"values": []map[string]interface{}{
						{
							"key":   "m",
							"value": "underlay",
						},
						{
							"key":   "i",
							"value": "eU44YkFJOHlVMmZrWVRDOUNTRm1D",
						},
						{
							"key":   "b",
							"value": "screen",
						},
						{
							"key":   "r",
							"value": "true",
						},
					},
					"plugin": "t",
					"name":   "merge",
				},
			},
		},
	},
	{
		"url": "https://cdn.pixelbin.io/v2/feel/erase.bg(shadow:true)~t.merge(m:underlay,i:eU44YkFJOHlVMmZrWVRDOUNTRm1D,b:screen,r:true)/MZZKB3e1hT48o0NYJ2Kxh.jpeg?dpr=2.0&f_auto=true",
		"obj": map[string]interface{}{
			"version":   "v2",
			"baseUrl":   "https://cdn.pixelbin.io",
			"filePath":  "MZZKB3e1hT48o0NYJ2Kxh.jpeg",
			"pattern":   "erase.bg(shadow:true)~t.merge(m:underlay,i:eU44YkFJOHlVMmZrWVRDOUNTRm1D,b:screen,r:true)",
			"cloudName": "feel",
			"options": map[string]interface{}{
				"dpr":    2.0,
				"f_auto": true,
			},
			"zone": nil,
			"transformations": []map[string]interface{}{
				{
					"values": []map[string]interface{}{
						{
							"key":   "shadow",
							"value": "true",
						},
					},
					"plugin": "erase",
					"name":   "bg",
				},
				{
					"values": []map[string]interface{}{
						{
							"key":   "m",
							"value": "underlay",
						},
						{
							"key":   "i",
							"value": "eU44YkFJOHlVMmZrWVRDOUNTRm1D",
						},
						{
							"key":   "b",
							"value": "screen",
						},
						{
							"key":   "r",
							"value": "true",
						},
					},
					"plugin": "t",
					"name":   "merge",
				},
			},
		},
	},
	{
		"url": "https://cdn.pixelbin.io/v2/feel/erase.bg()/MZZKB3e1hT48o0NYJ2Kxh.jpeg?dpr=auto",
		"obj": map[string]interface{}{
			"version":   "v2",
			"baseUrl":   "https://cdn.pixelbin.io",
			"filePath":  "MZZKB3e1hT48o0NYJ2Kxh.jpeg",
			"pattern":   "erase.bg()",
			"cloudName": "feel",
			"options": map[string]interface{}{
				"dpr": "auto",
			},
			"zone": nil,
			"transformations": []map[string]interface{}{
				{
					"plugin": "erase",
					"name":   "bg",
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		"url": "https://cdn.pixelbin.io/v2/feel/erase.bg()/asbc.jpeg?f_auto=false",
		"obj": map[string]interface{}{
			"version":   "v2",
			"baseUrl":   "https://cdn.pixelbin.io",
			"filePath":  "asbc.jpeg",
			"pattern":   "erase.bg()",
			"cloudName": "feel",
			"options": map[string]interface{}{
				"f_auto": false,
			},
			"zone": nil,
			"transformations": []map[string]interface{}{
				{
					"plugin": "erase",
					"name":   "bg",
				},
			},
			"worker":     false,
			"workerPath": "",
		},
	},
	{
		"url": "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/wrkr/image.jpeg",
		"obj": map[string]interface{}{
			"version":         "v2",
			"cloudName":       "broken-butterfly-3b12f1",
			"isCustomDomain":  false,
			"pattern":         "",
			"filePath":        "",
			"options":         map[string]interface{}{},
			"zone":            nil,
			"baseUrl":         "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{},
			"worker":          true,
			"workerPath":      "image.jpeg",
		},
	},
	{
		"url": "https://cdn.pixelbin.io/v2/broken-butterfly-3b12f1/wrkr/resize:w200,h200/image.jpeg",
		"obj": map[string]interface{}{
			"version":         "v2",
			"cloudName":       "broken-butterfly-3b12f1",
			"isCustomDomain":  false,
			"pattern":         "resize:w200,h200",
			"filePath":        "image.jpeg",
			"options":         map[string]interface{}{},
			"zone":            nil,
			"baseUrl":         "https://cdn.pixelbin.io",
			"transformations": []map[string]interface{}{},
			"worker":          true,
			"workerPath":      "resize:w200,h200/image.jpeg",
		},
	},
	{
		"url": "https://cdn.twist.vision/v2/t.resize(h:600,w:800)/W2.jpeg",
		"obj": map[string]interface{}{
			"isCustomDomain": true,
			"version":        "v2",
			"pattern":        "t.resize(h:600,w:800)",
			"filePath":       "W2.jpeg",
			"options":        map[string]interface{}{},
			"zone":           nil,
			"baseUrl":        "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "600"},
						{"key": "w", "value": "800"},
					},
				},
			},
		},
	},
	{
		"url": "https://cdn.twist.vision/v2/z-slug/t.resize(h:600,w:800)/W2.jpeg",
		"obj": map[string]interface{}{
			"isCustomDomain": true,
			"version":        "v2",
			"pattern":        "t.resize(h:600,w:800)",
			"filePath":       "W2.jpeg",
			"options":        map[string]interface{}{},
			"zone":           "z-slug",
			"baseUrl":        "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "600"},
						{"key": "w", "value": "800"},
					},
				},
			},
		},
	},
	{
		"url": "https://cdn.twist.vision/v2/t.resize(h:600,w:800)~t.rotate(a:-249)/W2.jpeg",
		"obj": map[string]interface{}{
			"isCustomDomain": true,
			"version":        "v2",
			"pattern":        "t.resize(h:600,w:800)~t.rotate(a:-249)",
			"filePath":       "W2.jpeg",
			"options":        map[string]interface{}{},
			"zone":           nil,
			"baseUrl":        "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "600"},
						{"key": "w", "value": "800"},
					},
				},
				{
					"plugin": "t",
					"name":   "rotate",
					"values": []map[string]interface{}{
						{"key": "a", "value": "-249"},
					},
				},
			},
		},
	},
	{
		"url": "https://cdn.twist.vision/v2/t.resize(h:600,w:800)~t.rotate(a:-249)~t.flip()~t.trim(t:217)/W2.jpeg",
		"obj": map[string]interface{}{
			"isCustomDomain": true,
			"version":        "v2",
			"pattern":        "t.resize(h:600,w:800)~t.rotate(a:-249)~t.flip()~t.trim(t:217)",
			"filePath":       "W2.jpeg",
			"options":        map[string]interface{}{},
			"zone":           nil,
			"baseUrl":        "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "600"},
						{"key": "w", "value": "800"},
					},
				},
				{
					"plugin": "t",
					"name":   "rotate",
					"values": []map[string]interface{}{
						{"key": "a", "value": "-249"},
					},
				},
				{
					"plugin": "t",
					"name":   "flip",
				},
				{
					"plugin": "t",
					"name":   "trim",
					"values": []map[string]interface{}{
						{"key": "t", "value": "217"},
					},
				},
			},
		},
	},
	{
		"url": "https://cdn.twist.vision/v2/t.resize(h:200,w:100)~p:preset1(a:100,b:2.1,c:test)/W2.jpeg",
		"obj": map[string]interface{}{
			"isCustomDomain": true,
			"version":        "v2",
			"pattern":        "t.resize(h:200,w:100)~p:preset1(a:100,b:2.1,c:test)",
			"filePath":       "W2.jpeg",
			"zone":           nil,
			"baseUrl":        "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
					"values": []map[string]interface{}{
						{"key": "a", "value": "100"},
						{"key": "b", "value": "2.1"},
						{"key": "c", "value": "test"},
					},
				},
			},
		},
	},
	{
		"url": "https://cdn.twist.vision/v2/t.resize(h:200,w:100)~p:preset1/W2.jpeg",
		"obj": map[string]interface{}{
			"isCustomDomain": true,
			"version":        "v2",
			"pattern":        "t.resize(h:200,w:100)~p:preset1",
			"filePath":       "W2.jpeg",
			"zone":           nil,
			"baseUrl":        "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
				},
			},
		},
	},
	{
		"url": "https://cdn.twist.vision/v2/t.resize(h:200,w:100)~p:preset1(a:12)/W2.jpeg",
		"obj": map[string]interface{}{
			"isCustomDomain": true,
			"version":        "v2",
			"pattern":        "t.resize(h:200,w:100)~p:preset1(a:12)",
			"filePath":       "W2.jpeg",
			"zone":           nil,
			"baseUrl":        "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
					"values": []map[string]interface{}{
						{"key": "a", "value": "12"},
					},
				},
			},
		},
	},
	{
		"url": "https://cdn.twist.vision/v2/t.resize(h:200,w:100)~p:preset1/W2.jpeg",
		"obj": map[string]interface{}{
			"isCustomDomain": true,
			"version":        "v2",
			"pattern":        "t.resize(h:200,w:100)~p:preset1(a:12)",
			"filePath":       "W2.jpeg",
			"zone":           nil,
			"baseUrl":        "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": "200"},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
					"values": []map[string]interface{}{},
				},
			},
		},
	},
	{
		"url": "https://cdn.twist.vision/v2/t.resize(h:,w:100)~p:preset1/W2.jpeg",
		"obj": map[string]interface{}{
			"isCustomDomain": true,
			"version":        "v2",
			"pattern":        "t.resize(h:,w:100)~p:preset1(a:12)",
			"filePath":       "W2.jpeg",
			"zone":           nil,
			"baseUrl":        "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"key": "h", "value": ""},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
					"values": []map[string]interface{}{},
				},
			},
		},
		"error": "value not specified for 'h' in 'resize'",
	},
	{
		"url": "https://cdn.twist.vision/v2/t.resize(h:200,w:100)~p:preset1/W2.jpeg",
		"obj": map[string]interface{}{
			"isCustomDomain": true,
			"version":        "v2",
			"pattern":        "t.resize(h:200,w:100)~p:preset1(a:12)",
			"filePath":       "W2.jpeg",
			"zone":           nil,
			"baseUrl":        "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{"value": ""},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
					"values": []map[string]interface{}{},
				},
			},
		},
		"error": "key not specified",
	},
	{
		"url": "https://cdn.twist.vision/v2/t.resize(h:200,w:100)~p:preset1/W2.jpeg",
		"obj": map[string]interface{}{
			"isCustomDomain": true,
			"version":        "v2",
			"pattern":        "t.resize(h:200,w:100)~p:preset1(a:12)",
			"filePath":       "W2.jpeg",
			"zone":           nil,
			"baseUrl":        "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{
				{
					"plugin": "t",
					"name":   "resize",
					"values": []map[string]interface{}{
						{},
						{"key": "w", "value": "100"},
					},
				},
				{
					"plugin": "p",
					"name":   "preset1",
					"values": []map[string]interface{}{},
				},
			},
		},
		"error": "key not specified",
	},
	{
		"url": "https://cdn.twist.vision/v2/erase.bg(shadow:true)~t.merge(m:underlay,i:eU44YkFJOHlVMmZrWVRDOUNTRm1D,b:screen,r:true)/MZZKB3e1hT48o0NYJ2Kxh.jpeg?dpr=2.0&f_auto=true",
		"obj": map[string]interface{}{
			"isCustomDomain": true,
			"version":        "v2",
			"baseUrl":        "https://cdn.twist.vision",
			"filePath":       "MZZKB3e1hT48o0NYJ2Kxh.jpeg",
			"pattern":        "erase.bg(shadow:true)~t.merge(m:underlay,i:eU44YkFJOHlVMmZrWVRDOUNTRm1D,b:screen,r:true)",
			"options": map[string]interface{}{
				"dpr":    2.0,
				"f_auto": true,
			},
			"zone": nil,
			"transformations": []map[string]interface{}{
				{
					"values": []map[string]interface{}{
						{"key": "shadow", "value": "true"},
					},
					"plugin": "erase",
					"name":   "bg",
				},
				{
					"values": []map[string]interface{}{
						{"key": "m", "value": "underlay"},
						{"key": "i", "value": "eU44YkFJOHlVMmZrWVRDOUNTRm1D"},
						{"key": "b", "value": "screen"},
						{"key": "r", "value": "true"},
					},
					"plugin": "t",
					"name":   "merge",
				},
			},
		},
	},
	{
		"url": "https://cdn.twist.vision/v2/erase.bg()/MZZKB3e1hT48o0NYJ2Kxh.jpeg?dpr=auto",
		"obj": map[string]interface{}{
			"isCustomDomain": true,
			"version":        "v2",
			"baseUrl":        "https://cdn.twist.vision",
			"filePath":       "MZZKB3e1hT48o0NYJ2Kxh.jpeg",
			"pattern":        "erase.bg()",
			"options": map[string]interface{}{
				"dpr": "auto",
			},
			"zone":       nil,
			"worker":     false,
			"workerPath": "",
			"transformations": []map[string]interface{}{
				{
					"plugin": "erase",
					"name":   "bg",
				},
			},
		},
	},
	{
		"url": "https://cdn.twist.vision/v2/erase.bg()/asbc.jpeg?f_auto=false",
		"obj": map[string]interface{}{
			"isCustomDomain": true,
			"version":        "v2",
			"baseUrl":        "https://cdn.twist.vision",
			"filePath":       "asbc.jpeg",
			"pattern":        "erase.bg()",
			"options": map[string]interface{}{
				"f_auto": false,
			},
			"zone":       nil,
			"worker":     false,
			"workerPath": "",
			"transformations": []map[string]interface{}{
				{
					"plugin": "erase",
					"name":   "bg",
				},
			},
		},
	},
	{
		"url": "https://cdn.twist.vision/v2/wrkr/image.jpeg",
		"obj": map[string]interface{}{
			"isCustomDomain":  true,
			"version":         "v2",
			"pattern":         "",
			"filePath":        "",
			"options":         map[string]interface{}{},
			"zone":            nil,
			"baseUrl":         "https://cdn.twist.vision",
			"worker":          true,
			"workerPath":      "image.jpeg",
			"transformations": []map[string]interface{}{},
		},
	},
	{
		"url": "https://cdn.twist.vision/v2/wrkr/resize:w200,h200/image.jpeg",
		"obj": map[string]interface{}{
			"isCustomDomain":  true,
			"version":         "v2",
			"pattern":         "",
			"filePath":        "",
			"options":         map[string]interface{}{},
			"zone":            nil,
			"baseUrl":         "https://cdn.twist.vision",
			"worker":          true,
			"workerPath":      "resize:w200,h200/image.jpeg",
			"transformations": []map[string]interface{}{},
		},
	},
	{
		"url": "https://cdn.twist.vision/v2/abcdef/wrkr/image.jpeg",
		"obj": map[string]interface{}{
			"isCustomDomain":  true,
			"version":         "v2",
			"pattern":         "",
			"filePath":        "",
			"options":         map[string]interface{}{},
			"zone":            "abcdef",
			"baseUrl":         "https://cdn.twist.vision",
			"worker":          true,
			"workerPath":      "image.jpeg",
			"transformations": []map[string]interface{}{},
		},
	},
	{
		"url": "https://cdn.twist.vision/v2/abcdef/wrkr/resize:w200,h200/image.jpeg",
		"obj": map[string]interface{}{
			"isCustomDomain":  true,
			"version":         "v2",
			"pattern":         "",
			"filePath":        "",
			"options":         map[string]interface{}{},
			"zone":            "abcdef",
			"baseUrl":         "https://cdn.twist.vision",
			"worker":          true,
			"workerPath":      "resize:w200,h200/image.jpeg",
			"transformations": []map[string]interface{}{},
		},
	},
	{
		"url": "https://cdn.twist.vision/v2/abcdef/wrkr/resize:w200,h200/image.jpeg?dpr=1",
		"obj": map[string]interface{}{
			"isCustomDomain":  true,
			"version":         "v2",
			"pattern":         "",
			"filePath":        "",
			"zone":            "abcdef",
			"baseUrl":         "https://cdn.twist.vision",
			"transformations": []map[string]interface{}{},
			"worker":          true,
			"workerPath":      "resize:w200,h200/image.jpeg",
			"options": map[string]interface{}{
				"dpr": 1,
			},
		},
	},
}

func TestUrlToObjCases(t *testing.T) {
	for i, testcase := range urls_to_obj {
		t.Run(fmt.Sprintf("Case: %s", testcase.scenario), func(t *testing.T) {
			expected := testcase.obj
			obj, err := url.UrlToObj(testcase.url, testcase.opts...)
			if err != nil {
				t.Errorf("Failed case %d ! got err %v", i+1, err)
			} else {
				if fmt.Sprint(obj) == fmt.Sprint(expected) {
					t.Logf("Success case %d", i+1)
				} else {
					t.Errorf("Failed ! case %d expected %v, got %v", i+1, expected, obj)
				}
			}
		})
	}
}

func TestObjToUrlCases(t *testing.T) {
	for i, cases := range obj_to_urls {
		t.Run(fmt.Sprintf("Case %d", i+1), func(t *testing.T) {
			expurlstr := cases["url"].(string)
			Obj := cases["obj"].(map[string]interface{})
			url, err := url.ObjToUrl(Obj)
			if err != nil {
				if err.Error() != cases["error"].(string) {
					t.Errorf("Failed case %d ! got err %v", i+1, err)
				} else {
					t.Logf("Success case %d", i+1)
				}
			} else {
				if url == expurlstr {
					t.Logf("Success case %d", i+1)
				} else {
					t.Errorf("Failed ! case %d expected %v, got %v", i+1, expurlstr, url)
				}
			}
		})
	}
}

func TestFailureOptionDprOfQueryParam(t *testing.T) {
	Obj := map[string]interface{}{
		"baseUrl":   "https://cdn.pixelbin.io",
		"filePath":  "__playground/playground-default.jpeg",
		"version":   "v2",
		"zone":      "z-slug",
		"cloudName": "red-scene-95b6ea",
		"pattern":   nil,
		"options": map[string]interface{}{
			"dpr":    5.5,
			"f_auto": true,
		},
		"transformations": []map[string]interface{}{},
	}
	_, err := url.ObjToUrl(Obj)
	if err != nil {
		errIdentifier := strings.Split(err.Error(), " ")[0]
		if errIdentifier == "DPR" {
			t.Logf("Success")
		} else {
			t.Errorf("Failed ! got err %v", err)
		}
	} else {
		t.Errorf("Failed case")
	}
}

func TestFailureOptionFautoOfQueryParam(t *testing.T) {
	Obj := map[string]interface{}{
		"baseUrl":   "https://cdn.pixelbin.io",
		"filePath":  "__playground/playground-default.jpeg",
		"version":   "v2",
		"zone":      "z-slug",
		"cloudName": "red-scene-95b6ea",
		"pattern":   nil,
		"options": map[string]interface{}{
			"dpr":    5.5,
			"f_auto": true,
		},
		"transformations": []map[string]interface{}{},
	}
	_, err := url.ObjToUrl(Obj)
	if err != nil {
		errIdentifier := strings.Split(err.Error(), " ")[0]
		if errIdentifier == "DPR" {
			t.Logf("Success")
		} else {
			t.Errorf("Failed ! got err %v", err)
		}
	} else {
		t.Errorf("Failed case")
	}
}
