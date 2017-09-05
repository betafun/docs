访问一个API并把返回的结果存到DB中。
首先需要获取到API返回数据的结构，可以用在线的转换工具：http://json2struct.mervine.net/
什么情况下使用http.client，什么情况下使用transport
For control over HTTP client headers, redirect policy, and other settings, create a Client:
For control over proxies, TLS configuration, keep-alives, compression, and other settings, create a Transport:
import(
"encoding/json"
"net/http"
"io/ioutil"
"time"
"log"
"database/sql"
_ "github.com/go-sql-driver/mysql"
)

type MyJsonName struct {
        Data []struct {
                CurrA       string      `json:"curr_a"`
                CurrB       string      `json:"curr_b"`
                CurrSuffix  string      `json:"curr_suffix"`
                Marketcap   string      `json:"marketcap"`
                Name        string      `json:"name"`
                NameCn      string      `json:"name_cn"`
                NameEn      string      `json:"name_en"`
                No          int         `json:"no"`
                Pair        string      `json:"pair"`
                Plot        interface{} `json:"plot"`
                Rate        string      `json:"rate"`
                RatePercent string      `json:"rate_percent"`
                Supply      int         `json:"supply"`
                Symbol      string      `json:"symbol"`
                Trend       string      `json:"trend"`
                VolA        float32     `json:"vol_a"`
                VolB        string      `json:"vol_b"`
        } `json:"data"`
        Result string `json:"result"`
}

然后模拟http客户端发起请求。

func Httpget(url_addr string) MyJsonName {
    //url_addr:="http://data.bter.com/api/1/marketlist"
    //设置超时时间
    client := &http.Client{
    Timeout: time.Duration(time.Second * 60 ),
    }

    req, err := http.NewRequest("GET", url_addr, nil)
    if err != nil {
       log.Println(err)
    }
    //设置一些请求头信息
   //req.Header.Set("Content-Type","application/json;charset=UTF-8")
   //req.Header.Set("User-Agent", "MIG-Patrol")
   //req.Header.Set("Referer", "Deanlzhang" )
   //异常捕获
   defer func() {
    if err:=recover();err!=nil{
       //log.Println(  err )
    }
  }()
  var r MyJsonName
  //timestamp:=time.Now().Format("2006-01-02 15:04:05")
  //失败重试三次
  for i := 0; i < 3; i++ {
    resp, err := client.Do(req)
    if err != nil {
      log.Println( err )
    }
    defer resp.Body.Close()
    //读取响应数据并解析
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
      log.Println( err )
    }
    //解析json数据
    err = json.Unmarshal([]byte(string(body)), &r)
    if "true" == r.Result {
        break
    }
  }
  return r //返回结果
}

将请求到的结果存到DB中
func InsertIntoDB(r MyJsonName) {
    timestamp:=time.Now().Format("2006-01-02 15:04:05")
   //连接DB
    db, err := sql.Open("mysql", "用户名:密码@tcp(IP:PORT)/db")
    if err != nil {
        log.Println(err)
    }
    //插入数据
    for _,val := range r.Data {
        _,err = db.Exec("INSERT INTO t_coin_info values(?,?,?,?,?,?,?,?,?,?,?)",val.Symbol,val.NameCn,val.NameEn,val.Rate,val.Marketcap,val.VolA,val.VolB,val.Supply,val.RatePercent,timestamp,"bter")
        if err != nil {
          log.Println(err)
        }
    }
    //关闭DB连接
    defer db.Close()
}

从DB中读取相应的数据，并保存到slice中
func SelectfromDB() []string {
    var result []string
    db, err := sql.Open("mysql", "用户:密码@tcp(IP:Port)/DB")
    if err != nil {
        log.Println(err)
    }   

    rows,err:=db.Query("select distinct(name) from t_coin_info where source='bter'")
    if err != nil {
        log.Println(err)
    }   
    defer rows.Close()
    for rows.Next() {
        var name string
        if err := rows.Scan(&name); err != nil {
                log.Println(err)
        }   
        result=append(result,strings.ToLower(name))
    }   
    //返回rows在迭代过程中返回的错误
    if err := rows.Err(); err != nil {
       log.Println(err)
    }   
    return result
}



                                                            
