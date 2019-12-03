package config

import "os"

//StaticPath static base path of project
var StaticPath = os.Getenv("StaticPath")

//ServerHostWithPort server host
var ServerHostWithPort = os.Getenv("ServerHost") + ":" + os.Getenv("ServerPort")

//StorageBaseURL storage service base url
var StorageBaseURL = "http://" + ServerHostWithPort + "/storage/"

//DBHost db host url
var DBHost = os.Getenv("DBHost")

//DBPort db port
var DBPort = os.Getenv("DBPort")

//DBName db user
var DBName = os.Getenv("DBName")

//DBUser db user
var DBUser = os.Getenv("DBUser")

//DBPassword db password
var DBPassword = os.Getenv("DBPassword")

//AllowedOrigins allowed origins
var AllowedOrigins = []string{"*"}

//DocToPngServiceURL doc to png service url
var DocToPngServiceURL = os.Getenv("DocToPngServiceURL")

//DocTempDirPath dir to store
var DocTempDirPath = os.Getenv("DocTempDirPath")
