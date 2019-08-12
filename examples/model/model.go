package model

// Face hello kratos.
type Face struct {
	FaceID    string
	FaceName  string
	FaceImage string
}

// Faces get faces
type Faces struct {
	Top   int
	Faces []Face
}
