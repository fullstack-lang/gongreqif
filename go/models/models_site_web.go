package models

type StaticWebSite struct {
	Name string

	//gong:text
	//gong:width 900 gong:height 600
	MarkdownContent string

	Chapters []*StaticWebSiteChapter

	InputImagesDir     string
	OutputStaticWebDir string
}

type StaticWebSiteChapter struct {
	Name string

	//gong:text
	//gong:width 900 gong:height 600
	MarkdownContent string

	Paragraphs []*Paragraph
}

type Paragraph struct {
	Name string

	//gong:text
	//gong:width 300 gong:height 300
	LegendMarkdownContent string

	Image *Image
}

// Iamge needs to be copied to "/static/images"
type Image struct {
	Name                string
	SourceDirectoryPath string // for instance "../../../images/logo.svg"
	Height              int
}

type GeneratedImageMetamodel struct {
	Name string

	ImageName   string
	IsMetamodel bool

	//gong:text
	//gong:width 300 gong:height 300
	LegendMarkdownContent string
}
