## The Go command

go build -> Complies a bunch of Go source code.
go run -> Compiles and executes one or two files
go fmt -> Formats all the code in each file in the current directory
go intall -> Compiles and installs a package
go get -> Downloads the raw source code of somone else's package
go test -> Runs any test associated with the current project.

## Types of packages

1. Executable - Generates a file that we can run
2. Reusable - Code used as 'helpers', reusable logic. 'packages'

The name of the package at the top of the file determines whether or not you are making an executable or reusable type package.

In main.go the package name is 'main'. When we built main.go it created the main executable. This is a reserved name and how we make executable. NB it must aslo have a func main() in the file.

Any other name will create a reusable package.

## Organisation

1. package statement
2. import statements
3. code
