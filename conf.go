package amock

// Conf configures the generation process.
type Conf struct {
	Package string // Package of the generated mock implementation.
	Name    string // Name of the generated file and mock implementation type.
	Path    string // Path of the generated file.
}
