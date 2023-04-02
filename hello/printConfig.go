package hello

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)

type PrintAppInfo struct {
}

func (*PrintAppInfo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("appinfo")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	name := viper.Get("AppInfo.Name")
	host := viper.Get("AppInfo.Host")
	port := viper.Get("AppInfo.Port")
	desc := viper.Get("AppInfo.Desc")
	fmt.Fprintf(w, "name = %s , host= %s, port=%d, desc=%s", name, host, port, desc)
}

type PrintAuthorInfo struct {
}

func (*PrintAuthorInfo) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	viper.AddConfigPath("configs")
	viper.SetConfigName("authorinfo")
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	name := viper.Get("AuthorInfo.Name")
	gender := viper.Get("AuthorInfo.Gender")
	age := viper.Get("AuthorInfo.Age")
	desc := viper.Get("AuthorInfo.Desc")
	fmt.Fprintf(w, "name = %s , gender = %s, age =%d, desc=%s", name, gender, age, desc)
}
