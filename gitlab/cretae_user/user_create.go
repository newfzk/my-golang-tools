package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

// Config 配置结构体
type Config struct {
	GitlabURL    string `yaml:"gitlab_url"`
	PrivateToken string `yaml:"private_token"`
	Users        []User `yaml:"users"`
}

// User 用户结构体
type User struct {
	// 初始密码 12345678
	Email            string `yaml:"email"`
	Username         string `yaml:"username"`
	Name             string `yaml:"name"`
	SkipConfirmation bool   `yaml:"skip_confirmation"`
}

func main() {
	// 读取配置文件
	configFile := "config.yaml"
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}

	configData, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("配置文件读取失败: %v", err)
	}

	var config = Config{}
	if err := yaml.Unmarshal(configData, &config); err != nil {
		log.Fatalf("配置文件解析失败: %v", err)
	}

	// 创建http客户端
	client := &http.Client{}

	// 创建用户
	for _, user := range config.Users {
		if err := createGitlabUser(client, config.GitlabURL, config.PrivateToken, user); err != nil {
			log.Printf("创建用户 %s 失败: %v", user, err)
		} else {
			log.Printf("用户 %s 创建成功", user.Username)
		}
	}

}

func createGitlabUser(client *http.Client, baseURL, token string, user User) error {
	apiURL := fmt.Sprintf("%s/api/v4/users", baseURL)

	// 表单数据构建
	formData := url.Values{}
	formData.Set("username", user.Username)
	formData.Set("email", user.Email)
	formData.Set("name", user.Name)
	formData.Set("skip_confirmation", strconv.FormatBool(user.SkipConfirmation))
	formData.Set("password", "12345678") // 初始密码 12345678

	req, err := http.NewRequest("POST", apiURL, nil)
	if err != nil {
		return err
	}
	// 设置请求头
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("PRIVATE-TOKEN", token)
	req.URL.RawQuery = formData.Encode() // 将表单数据编码为查询字符串

	// 发送请求
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	// defer 推迟 类似于 final
	defer resp.Body.Close()

	// 检查响应状态
	if resp.StatusCode != http.StatusCreated {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("用户创建失败 API响应状态: %s, 响应体: %s", resp.Status, string(body))
	}

	return nil
}
