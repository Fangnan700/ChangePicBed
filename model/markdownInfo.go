package model

type MarkdownInfo struct {
	FileName     string
	FileSize     int64
	ContentLines []string
	ImagesInfo   []Images
}

type Images struct {
	LineIndex int
	ImageMd5  string
	ImageSrc  string
	ImageDes  string
}
