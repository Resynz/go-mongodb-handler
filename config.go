/**
 * @Author: Resynz
 * @Date: 2020/3/5 16:10
 */
package go_mongodb_handler

type Option struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type Config struct {
	Host     string   `json:"host"`
	Port     uint     `json:"port"`
	UserName string   `json:"user_name"`
	Password string   `json:"password"`
	Database string   `json:"database"`
	Options  []Option `json:"options"`
}
