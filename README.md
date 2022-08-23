# Config_wrapper
wrap config
#To import

go get github.com/Quddus1916/Config_wrapper

#Functionalities 
1.Incorporate run time changes without restart
2.Get a value (string/int type) against a particular key
3.Set default value if key not present

#Functions

1.NewConfig (file_with_extention string, filepath string) (map[string]interface{}, error) {}



Description: It will create a map from the config file And it will start watching. 
If any changes saved in config file then updates will be
reflected throughout the program without restart

2.GetKeyString(key string, deep_key *string, default_val string) string {}



Description:It will take a key and it will return a value as string. 



NB: deep_key is only used if you store a json against a key
 {
  "app":{
       "port":"8080"
      }
  }
  
  So here key is app and deep_key is port.
  
  

3.GetKeyInt(key string, deep_key *string, default_value string) int {}



Description:Same as GetKeyString and it will return a value as int.


#Usage



        _, err := NewConfig("aconfig.json", ".")

	if err != nil {
		fmt.Println(err.Error())
	}
        val := GetKeyInt("port", nil, "1010")
       val2 := GetKeyString("port", nil, "1010")
  
  
  
  #Limitation
  
  
  1.File name must follow the Format-example app.json/app.env
  
  
  
  2.For Json it only supports upto 2nd level nesting
  
  
  example,
  
        lvl 1 :
              {
              "port":"8080"
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
