# ass_makeGoServer
This was a project done on the directions of Josh Univ.
COncepts used:
- Go
- Go http package
- made a server
- did not make a client
- used a curl to hit up with request.
- was using curl on windows, faced issue as windows doen't use curl as curl but as Invoke-WebRequest. used the command: Remove-item alias:curl
curl worked therafter. however, another issue popped up; while making a PUT request, the data being sent was not being recognized ig the issue stems from something related to 
problem with "" and '' and json and curl on windows.
