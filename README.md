1. Few Pre requisites : 

1.1 TO run the code . System should have "make" and "glide" installed.
.2 I have used docker for SQL Database. Steps to install docker and make IceCreams Database up and running is attached. 

2. Steps to start running. (I have tested the code on UBUNTU 16.04)
2.1 First execute steps in "sql_setup_using_docker.txt" . This will install docker and setup SQL server.
2.2 Then go to bin folder and Execute "setup.sh"           
      a. This will execute makefile           
      b. Install all dependencies           
      c. run unit test cases.
      
3. Now Simply start the application using "go run main.go"

4. Postman Link to import APIS : https://www.getpostman.com/collections/84b39d9995521a0ea9f9


5. To Authenticate APIs I have used Json Web Tokens. 
      
      Import Postman Link to seethe APIs









