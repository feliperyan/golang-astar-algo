# Tips

### Using
- Spacebar will generate a new "dungeon"
- P will show the a* path 

## Some stuff to keep in mind about GO

1. If I want functions to be viewed across files, like calling a function in utilities.go from within main.go, I have to "run" both of them.

2. Only types starting with capital letters are accessible in a package, there doesn't seem to be a public/private declaration.

3. Init is a "magical function" that will always get called when running a go script, this really confused me when reading Ebiten examples!