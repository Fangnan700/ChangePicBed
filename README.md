# ChangePicBed
A tool to change the image bed for Markdown images in bulk.



## Introduction

This is a tool for modifying markdown image links in batches, which can quickly dump the images in markdown to the specified image bed.



## Function

1. Batch export Yuque documents
2. Modify markdown image links in batches
3. Currently supported image beds:
   1. Tencent Cloud COS



## Quick Start

**Get the latest**

Download the latest version from the Release page.



**Get the yuque's cookies**

Log in to the Yuque web page and get the following cookies:

![image-20230615112156883](https://yvling-typora-image-1257337367.cos.ap-nanjing.myqcloud.com/typora/image-20230615112156883.png)



**Fill the config file**

The configuration file is located in the config directory.

Fill in the example as follows:

```yaml
# Input file directory
input_dir: "input"

# Output file directory
output_dir: "output"

# Temporary file directory
temp_dir: "temp"

# Bed service provider, optional (COS)
pic_bed: "cos"

# Tencent Cloud COS Configuration
cos_config:
  bucket_name: "yvling-typora-image-125xxxxxxx"
  bucket_area: "ap-nanjing"
  pic_path: "typora"
  secret_id: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  secret_key: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

# Yuque Configuration
yuque_config:
  _yuque_session: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  yuque_ctoken: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
  export_path: "yuque_export"
```







## project structure

```
.
│  .gitignore
│  go.mod
│  go.sum
│  LICENSE
│  main.go
│  README.md
│
├─config
│      config.yaml
│
├─model
│      config.go
│      markdownInfo.go
│      yuqueBookStacks.go
│      readFile.go
│      uploadCOS.go
│      writeFile.go
│
├─utils
│      changePicBed.go
│      check.go
│      clear.go
│      downloadImages.go
│      printConfig.go
│      readFile.go
│      uploadCOS.go
│      writeFile.go
│
└─yuque
       exportBook.go
       getBookStacks.go
```











