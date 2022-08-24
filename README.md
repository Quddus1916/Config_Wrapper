# Config_wrapper
wrap config
#To import

go get github.com/Quddus1916/Config_wrapper

#Functionalities 
1.Incorporate run time changes without restart


2.Get a value (string/int/float type) against a particular key


3.Set default value if key not present

#Functions

1.InitConfig (filepath string) (config, error) {}



Description: It will create a map from the config file And it will start watching. 
If any changes saved in config file then updates will be
reflected throughout the program without restart

2.GetConfigParamAsString(key string, deep_key *string, default_val string) string {}



Description:It will take a key and it will return a value as string. 



NB: deep_key is only used if you store a json against a key
 {
  "app":{
       "port":"8080"
      }
  }
  
  So here key is app and deep_key is port.
  
  

3.GetConfigParamAsInt64(key string, deep_key *string, default_value string) int64 {}



Description:Same as GetKeyString and it will return a value as int64.



3.GetConfigParamAsFloat64(key string, deep_key *string, default_value string) float64 {}



Description:Same as GetKeyString and it will return a value as float64.



#Usage



       config, err := InitConfig("./config.dev.json")

	if err != nil {
		fmt.Println(err.Error())
	}
	p := "limit"
	val := config.GetConfigParamAsString("app", &p, "1010")
	val2 := config.GetConfigParamAsInt64("app", &p, "1010")
	val3 := config.GetConfigParamAsFloat64("port", nil, "1010")
  
  
  
  #Limitation
  
  
  
 1.For Json it only supports upto 2nd level nesting
  
  
  
  
  example,
  
        lvl 1 :
              {
              "port":"8080",
	          "smtp_port":"555",
              "smtp_user":"abc",
              "smtp_pass":"eanded98a7c"
              }
  
        lvl 2 :
              {
              "port":"8080"
              "smtp":{
                       "smtp_port":"555",
                       "smtp_user":"abc",
                       "smtp_pass":"eanded98a7c"
                     }
               }
