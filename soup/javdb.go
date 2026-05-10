package soup

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/anaskhan96/soup"
)

var (
	html  = "https://javdb.com/search?q"
	proxy = "http://127.0.0.1:8889"
)

// createProxyClient 创建带代理的 HTTP 客户端
func createProxyClient(proxyURL string) (*http.Client, error) {
	proxyParsed, err := url.Parse(proxyURL)
	if err != nil {
		return nil, fmt.Errorf("解析代理地址失败: %w", err)
	}

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyParsed),
	}

	client := &http.Client{
		Transport: transport,
		Timeout:   30 * time.Second,
	}

	return client, nil
}

func Javdb(keyword string) (string, error) {
	query := strings.Join([]string{html, keyword}, "=")
	soup.SetDebug(	true)
log.Printf("请求的网址为: %v\n",query)
	// 创建带代理的客户端
	client, err := createProxyClient(proxy)
	if err != nil {
		return "", fmt.Errorf("创建代理客户端失败: %w", err)
	}

	// 使用自定义客户端发起请求
	resp, err := soup.GetWithClient(query, client)
	if err != nil {
		return "", fmt.Errorf("请求失败: %w", err)
	}
	//在这里把这次请求到的网址保存为tmp.html文件
	err = os.WriteFile("tmp.html", []byte(resp), 0644)
	if err != nil {
		log.Printf("保存HTML文件失败: %v\n", err)
	} else {
		log.Println("HTML已保存到 tmp.html")
	}
	root := soup.HTMLParse(resp)
	log.Printf("root is : %v\n", root)
	return resp, nil
}
