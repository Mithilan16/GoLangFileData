# GoLangFileData

This is a Go library that is comprised of various functions that provide information about a given file.

## Contributing

Anyone is welcome to contribute to this library. If you wish to contribute to this project:
* Fork this repo.
* Clone this repo locally.
* Add and commit your changes/contributions to your local copy.
* Make a pull request with the original repo as the base.


## What to contribute?
GoLangFileData welcomes any contributions that provide new functionality or improves on the already existing utilities in this library. All new functions should keep with the theme and provide file information or additional utility when working with files.  
If you would like ideas on how to contribute to this project, please have a look at the issues section of this repo.

## Using this project in your projects  
To use this library, simply copy thie fileUtils folder into the root of your GO project. 
* You must import the package like so:  
```c
import "/fileUtils"
```
## Example Code

The **GetSha1(filePath string)** function returns the sha1 digest of the given file.

```c
fmt.Printf("SHA1: %x\n", fileUtils.GetSha1("./testfile.txt"))
```

The **GetMd5(filePath string)** function returns the md5 digest of the given file.

```c
fmt.Printf("MD5: %x\n", fileUtils.GetMd5("./testfile.txt"))
```

The **ParseFileName(filePath string)** function returns the name of the file, without the path.

```c
fmt.Println("file name: " + fileUtils.ParseFileName("./testfile.txt"))
```