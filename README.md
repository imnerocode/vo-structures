# T_parser.go

This file contains the Go package `vo` which defines structures for representing 3D models. The package includes definitions for vertices, faces, materials, UV coordinates, and normals.

## Structures

### Model

The `Model` structure represents a 3D model with the following fields:

- `ID`: A string identifier for the model.
- `Vertices`: A slice of `Vertex` structures representing the vertices of the model.
- `Faces`: A slice of `Face` structures representing the faces of the model.
- `Materials`: A slice of `Material` structures representing the materials of the model.
- `UVCoords`: A slice of `UV` structures representing the UV coordinates of the model.
- `Normals`: A slice of `Normal` structures representing the normals of the model.

### Vertex

The `Vertex` structure represents a single point in 3D space with the following fields:

- `X`: The X coordinate (float32).
- `Y`: The Y coordinate (float32).
- `Z`: The Z coordinate (float32).

### Face

The `Face` structure represents a polygon formed by indices of vertices with the following field:

- `VertexIndices`: A slice of int32 representing the indices of the vertices forming the face.

### Material

The `Material` structure represents the material properties of a model with the following fields:

- `Name`: The name of the material (string).
- `Color`: The color of the material (string, hex or rgba).
- `Specularity`: The specularity of the material (float32).

### UV

The `UV` structure represents a UV coordinate with the following fields:

- `U`: The U coordinate (float32).
- `V`: The V coordinate (float32).

### Normal

The `Normal` structure represents a normal vector in 3D space with the following fields:

- `NX`: The normal X coordinate (float32).
- `NY`: The normal Y coordinate (float32).
- `NZ`: The normal Z coordinate (float32).

## Usage

To use these structures, import the `vo` package in your Go project:

```go
import "github.com/imnerocode/vo-structures/vo"
```

You can then create and manipulate 3D models using the provided structures.

## License

This project is licensed under the MIT License.

This `README.md` file provides an overview of the structures defined in the `T_parser.go` file, their fields, and usage instructions.