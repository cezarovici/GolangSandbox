# Package Creator 

## This documentation will help you to use the package easely

### By moment we will use a file text as our package controler 

Here is an example

```
! .
folder1
 subfolder1
 # package main
 file1.go file2.go
    file3.go
 file4.go
 subfolder2
 # t
 file5.go
  subsubfolder1
! /home/apetroaei/ram
 folder2
```

Let me explain for you, what this package can do <br>
First of all, let me present you the keywords or...key characters 

###**Point of reference** ###<br>

`!` change the path to the one that you introduce after it 
in our case : ```! .``` means that we remain in the same path

###**File package**###

`#` can be followed by 2 things like in our example :

####1. something that can be written in all the file from the directory that we are in####

Example : 

a.

```go
package main
```
The command will be `# package main` after the directory that we want to set the package

b.

```c++
#include <iostream>
using namespace std;
```

The command will be `# #include <iostream> \n using namespace std;`

Example :

```
subfolder1
 # package main
 file1.go file2.go
    file3.go
 file4.go
```

After the line `# package main` in the directory that contain all 4 files first of those lines will be `# package main`
 

####2. Can transform all the file bellow, from the current directory to test files (this is working by now only for golang)####

Example :

```
 # t
 file5.go
```

`file5.go` will become `file5_test.go`

###**Line indendation** ###<br>

Every line indentation is important except for key lines which I presented above
Think about it like stairs, every line is idented with a number of spaces that represent the level on that directory

Exemple :

```
folder1
 subfolder1
 subfolder2
  subsubfolder1
 subfolder3
```

In the specified path will be created `folder1` with  `subfolder1`   &    `subfolder2` as his children <br> `subfolder2` will the father of `subsubfolder1` <br><br>
We can even go backwards    `subfolder3`    will be created also in     `folder1`

