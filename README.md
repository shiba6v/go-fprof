# FProf
## About
FProf is a simple function level profiling tool.

## Usage
Add
e.GET("/fprof_result", GetAnalizeFProfResult) // FPROF_CODE
to endpoints.
Add
InitFProf() // FPROF_CODE
to main().

After benchmark,
curl localhost:9000/fprof_result > /tmp/fprof_result.txt
