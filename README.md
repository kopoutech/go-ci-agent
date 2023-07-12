# go-ci-agent

Go ci agent is a program that parses go ci format to run tasks

## Goal

To create a program that allows running CI tasks in multiple platforms. It allows one to capture the output and any other artifacts generated. 

Unlike CI programs like Jenkins, CircleCI you need to write manifest file only once for all the platforms. 

## Supported OS

- [x] Linux
- [x] Windows

## Supported Platforms

- [x] Local
- [] Remote Machine (SSH)
- [] Local Docker
- [] Local Podman
- [] Kubernetes


Sample Go CI format 

```json
{
  "source" : {
    "url" : "https://github.com/realpython/python-scripts.git",
    "branch" : "master"
  },
  "program" : {
    "executable" : "python",
    "args" : ["29_json_to_yaml.py" , "29_json_test.json"],
    "dir" : "scripts"
  },
  "capture" : {
    "stdout" : true,
    "stdoutFile" : "stdout.io",
    "stderr" : true,
    "stderrFile" : "stderr.io",
    "console" : true
  },
  "config" : {
    "baseDir" : "/tmp/"
  }
}
```
