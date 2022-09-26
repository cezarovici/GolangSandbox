# Package Creator 

### **Change Directory** <br>

`! path` = cd path

### **File package**

`#` can be followed by two things :

#### 1. something that can be written in all the files from the directory that we are in

Example : 

```go
package main
```
The command will be `# package main` after the directory that we want to set the package

Example :

```
subfolder1
 # package main
 file1.go file2.go
    file3.go
 file4.go
```

After the line `# package main` in the directory that contains all four files first of those lines will be `# package main`
 

#### 2. Can transform all the files below, from the current directory to test files ####

Example :

```
 # t
 file5.go
```

`file5.go` will become `file5_test.go`

### **Line indendation** <br>

Example :

```
folder1
 subfolder1
 subfolder2
  subsubfolder1
 subfolder3
```

In the specified path will be created `folder1` with  `subfolder1`   &    `subfolder2` as his children <br> `subfolder2` will the father of `subsubfolder1` <br><br>
We can even go backward `subfolder3`    will be created also in     `folder1`

