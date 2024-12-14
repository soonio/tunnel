# 通过SSH通道连接服务

## 使用

  ```bash
  go get -u github.com/soonio/tunnel
  ```

## 支持内容

  - [x] MYSQL
  - [x] Redis
  - [x] Mongo

## Usage

- Tunnel

  ```go
  tx := tunnel.New(tunnel.Conf{
    User:   "root",
    Secret: "12345678",
    Host:   "192.168.10.11",
    Port:   22,
  })
  if err := tunnelx.Connect(); err != nil {
    panic(err)
  }
  defer tx.Close()
  ```

- Mysql

  ```go
  dbx := mysql.New(mysql.Conf{
      Net:      "mysql+ssh",
      User:     "root",
      Password: "345612345",
      Host:     "xasasqwdqdascaca.aliyuncs.com",
      Port:     3306,
      Name:     "demo",
      Params:   "charset=utf8mb4&parseTime=true&loc=Asia%2FShanghai&multiStatements=true",
  })
  if err := dbx.Connect(UseTunnel().Dialer()); err != nil {
      panic(err)
  }
  defer tx.Close()
  ```
- Redis

  ```go
  redx := red.New(red.Conf{
    Addr:     "drftgyhujikoonbg.redis.rds.aliyuncs.com:6379",
    Username: "demo",
    Password: "2025good",
    Db:       0,
  })
  if err := redx.Connect(UseTunnel()); err != nil {
    panic(err)
  }
  defer redx.Close()
  ```
  
- Mongo
  
  ```go
  monx = mon.New("mongodb://root:asdf1234@dds-ucucucuucucuc.mongodb.rds.aliyuncs.com:3717,dds-ucucucuucucuc2.mongodb.rds.aliyuncs.com:3717/admin?replicaSet=mgset-760809887")
  if err := monx.Connect(UseTunnel().Client()); err != nil {
    panic(err)
  }
  defer monx.Close()
  ```