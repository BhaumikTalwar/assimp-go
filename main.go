package main

import (
	"bytes"
	"fmt"
	"image/png"

	"github.com/bloeys/assimp-go/asig"
)

func main() {

	scene, release, err := asig.ImportFile("../../depedencies/backpack.obj", asig.PostProcessTriangulate|
		asig.PostProcessGenNormals|
		asig.PostProcessOptimizeMeshes|
		asig.PostProcessJoinIdenticalVertices|
		asig.PostProcessGenUVCoords|
		asig.PostProcessFlipUVs|
		asig.PostProcessCalcTangentSpace,
	)

	if err != nil {
		panic(err)
	}
	defer release()

	fmt.Printf("RootNode: %+v\n\n", scene.RootNode)

	var vertCount uint32 = 0
	for i := 0; i < len(scene.Meshes); i++ {
		println("Mesh:", i, "; Verts:", len(scene.Meshes[i].Vertices), "; Normals:", len(scene.Meshes[i].Normals), "; MatIndex:", scene.Meshes[i].MaterialIndex)

		vertCount += uint32(len(scene.Meshes[i].Vertices))
		for j := 0; j < len(scene.Meshes[i].Vertices); j++ {
			fmt.Printf("V(%v): (%v, %v, %v)\n", j, scene.Meshes[i].Vertices[j].X(), scene.Meshes[i].Vertices[j].Y(), scene.Meshes[i].Vertices[j].Z())

		}
	}

	for i := 0; i < len(scene.Materials); i++ {

		m := scene.Materials[i]
		println("Material:", i, "; Props:", len(scene.Materials[i].Properties))
		texCount := asig.GetMaterialTextureCount(m, asig.TextureTypeDiffuse)
		fmt.Println("Material Texture count:", texCount)

		if texCount > 0 {

			texInfo, err := asig.GetMaterialTexture(m, asig.TextureTypeDiffuse, 0)
			if err != nil {
				panic(err)
			}

			fmt.Printf("%v", texInfo)
		}
	}

	ts := scene.Textures
	for i := 0; i < len(ts); i++ {
		t := ts[i]

		fmt.Printf("T(%v): Name=%v, Hint=%v, Width=%v, Height=%v, NumTexels=%v\n", i, t.Filename, t.FormatHint, t.Width, t.Height, len(t.Data))

		if t.FormatHint == "png" {
			decodePNG(t.Data)
		}
	}

	fmt.Println("\n---------------------------------------------")
	fmt.Println("Total Mesh Count:", len(scene.Meshes))
	fmt.Println("Total Vert Count:", vertCount)
	fmt.Println("Total Material Count:", len(scene.Materials))
	fmt.Println("Total Texture Count:", len(scene.Textures))
}

func decodePNG(texels []byte) {

	img, err := png.Decode(bytes.NewReader(texels))
	if err != nil {
		panic("wow2: " + err.Error())
	}

	println("C:", img.At(100, 100))
}
