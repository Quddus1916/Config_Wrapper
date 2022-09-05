# Config_wrapper
wrap config


#To import

go get github.com/Quddus1916/Config_wrapper

#Functionalities 


1.Incorporate run time changes without restart


2.Get a value (string/int/float type) against a particular key


3.Populate a struct from the config file


4.Set default value if key not present

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


4. GetParamAsStruct("app", nil, "10110", &conf)



Description: it will map the values from config file to a particular struct

      
      
      
                        type C struct {
	                      Name  string `mapstructure:"name"`
	                      Port  int64  `mapstructure:"port"`
	                      Page  int64  `mapstructure:"page"`
	                      Limit int64  `mapstructure:"limit"`
                                      }
				      
				      func main() {
				      var conf = new(C)
				      err := InitConfig("./config.dev.json", pair)
				      GetConfig().GetParamAsStruct("app", nil, "10110", &conf)
				      fmt.Println(conf.Limit)
				      }



#Usage for .env nad .json file:=



        func Prints1() {
	fmt.Println("db updated")
                      }
        var pair []KeyFunc
	pair = []KeyFunc{
		{Key: "db", CallBackFunc: Prints1},
	}
	err := InitConfig("./config.dev.json", pair)
	p := "limit"
	val := Getconfig().GetConfigParamAsString("app", &p, "1010")
	val2 := Getconfig().GetConfigParamAsInt64("app", &p, "1010")
	val3 := Getconfig().GetConfigParamAsFloat64("port", nil, "1010")
	GetConfig().GetParamAsStruct("app", nil, "10110", &conf)
  
  
 
  
  #Limitation
  
  
  
 1.For Json it only supports upto 2nd level nesting
 
 
 
 2.Callback functions only work for parent keys like for example 1 all and for example 2 port and smtp
 
 
 
 3.Only supports .env and .json
 
 
 4.GetParamAsStruct doesn't work for .env file only works for nested json like example lvl 2
  
  
  
  
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
