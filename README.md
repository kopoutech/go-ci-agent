# go-ci-agent
 Go ci agent is a program that parses go ci format to run tasks

Go CI format 

```json
{"source" : {
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
}}
```