golang爬虫功能经常用到的一些函数
爬取结果写入文件中：
 f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
 if err != nil {
    log.Printf("err")
 }
 defer f.Close()
 if _, err = f.WriteString(content); err != nil {
    panic(err)
 }
关于文件操作函数
func OpenFile(name string, flag int, perm FileMode) (file *File, err error)
//以各种方式打开各种存在不存在的文件，具体怎么样看flag和perm。
 O_APPEND int = syscall.O_APPEND // 在文件末尾追加，打开后cursor在文件结尾位置
 O_CREATE int = syscall.O_CREAT  // 如果不存在则创建
 O_WRONLY int = syscall.O_WRONLY // open the file write-only.
perm是文件的unix权限位，可以直接用数字写，如0644。

goquery包来解析爬取的网页内容
doc, err := goquery.NewDocument("http://www.us-proxy.org")
     if err != nil {
        log.Fatal(err)
        panic("get us proxy failed")
     }
    m := make(map[int]string)
    proxyAddr := ""
    doc.Find("tbody").Find("tr").Each(func(i int, s *goquery.Selection) {
            s.Find("td").Each(func(j int, se *goquery.Selection) {
                    if 0 == j {
                        proxyAddr = se.Text()
                    
                    if 1 == j {
                       proxyAddr += ":" + se.Text()
                    }
                })
                m[i] = proxyAddr
        })
        return m
使用外网代理访问网页
func GetByProxy(url_addr, proxy_addr string) (*http.Response, error) {
        request, _ := http.NewRequest("GET", url_addr, nil)
        proxy, err := url.Parse(proxy_addr)
        if err != nil {
                return nil, err
        }
        client := &http.Client{
                Timeout: time.Duration(time.Second * 10),
                Transport: &http.Transport{
                        Proxy: http.ProxyURL(proxy),
                },
        }

        return client.Do(request)
}

go并发控制(使用channel)
var workpool=20  //设置信道缓冲区大小
c := make(chan string,workpool)
maxProcs := runtime.NumCPU()
runtime.GOMAXPROCS(maxProcs) //根据机器核数开启对应的线程。

 for i, ip := range IPList {
     fmt.Println(ip)
     go spider(c, proxyList, ip, i, IPLen)
     c<-"ip:"+ip (向信道中写入数据)
 }

然后在spider函数中从信道中读取数据。
