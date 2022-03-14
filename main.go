package main

import "github.com/adiatma85/go-tutorial-gorm/internal/api"

type Source struct {
	Label   string `json:"label"`
	Info    string `json:"info"`
	Version int    `json:"version"`
}

type Sink struct {
	Label string
	Info  string
}

type HereticSink struct {
	NahLabel string `json:"label"`
	HahaInfo string `json:"info"`
	Version  string `json:"heretic_version"`
}

func main() {
	api.Run("")
	// source := Source{
	// 	Label:   "source",
	// 	Info:    "the origin",
	// 	Version: 1,
	// }
	// fmt.Println("source:", source)
	// mapped := smapping.MapFields(&source)
	// fmt.Println("mapped:", mapped)

	// sink := Sink{}
	// err := smapping.FillStruct(&sink, mapped)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("sink:", sink)

	// maptags := smapping.MapTags(&source, "json")
	// fmt.Println("maptags:", maptags)
	// hereticsink := HereticSink{}
	// err = smapping.FillStructByTags(&hereticsink, maptags, "json")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("heretic sink:", hereticsink)
}
