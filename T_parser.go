package vo

// Model represents a 3D model with vertices, faces, materials, UV coordinates, and normals.
type Model struct {
	ID        string     `json:"id"`        // Model identifier
	Vertices  []Vertex   `json:"vertices"`  // List of vertices
	Faces     []Face     `json:"faces"`     // List of faces
	Materials []Material `json:"materials"` // List of materials
	UVCoords  []UV       `json:"uv_coords"` // UV coordinates
	Normals   []Normal   `json:"normals"`   // Model normals
}

// Vertex represents a single point in 3D space.
type Vertex struct {
	X float32 `json:"x"` // X coordinate
	Y float32 `json:"y"` // Y coordinate
	Z float32 `json:"z"` // Z coordinate
}

// Face represents a polygon formed by indices of vertices.
type Face struct {
	VertexIndices []int32 `json:"vertex_indices"` // Indices of the vertices forming the face
}

// Material represents the material properties of a model.
type Material struct {
	Name        string  `json:"name"`        // Material name
	Color       string  `json:"color"`       // Material color (hex or rgba)
	Specularity float32 `json:"specularity"` // Specularity
}

// UV represents a UV coordinate.
type UV struct {
	U float32 `json:"u"` // U coordinate
	V float32 `json:"v"` // V coordinate
}

// Normal represents a normal vector in 3D space.
type Normal struct {
	NX float32 `json:"nx"` // Normal X coordinate
	NY float32 `json:"ny"` // Normal Y coordinate
	NZ float32 `json:"nz"` // Normal Z coordinate
}
