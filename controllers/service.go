package controllers

import (
  beego "github.com/beego/beego/v2/server/web"
  "bytes"
  "fmt"
  "os"
  "os/exec"
  "strings"
)

type ServiceController struct {
  beego.Controller
}

type ServiceResp struct {
  Result string `json:"result"`
  ErrMsg string `json:"errorMsg"`
}

func CmdAndChangeDir(dir string, commandName string, params []string) (string, error) {
  cmd := exec.Command(commandName, params...)
  if len(dir) > 0 {
    cmd.Dir = dir
    fmt.Println("dir", dir, "Cmd: ", cmd.Args)
  } else {
    fmt.Println("Cmd: ", cmd.Args)
  }

  var out bytes.Buffer
  cmd.Stdout = &out
  cmd.Stderr = os.Stderr

  err := cmd.Start()
  if err != nil {
    fmt.Println("Cmd error: ", err.Error())
    return "", err
  }
  err = cmd.Wait()
  outStr := out.String()
  fmt.Println("Cmd out: ", len(outStr), outStr)
  return outStr, err
}

func (c *ServiceController) Post() {
  worker := c.GetString("worker")
  param := c.GetString("param")
  fmt.Println(worker, " param: ", param)
  var params = strings.Split(param, " ")

  // var params = []string { "-s", "46504c590301010000000004020002bb", "16" }
  Result, err := CmdAndChangeDir("playfair", worker, params)
  if err != nil {
    c.Data["json"] = err.Error()
  } else {
    c.Data["json"] = Result
  }

  c.ServeJSON()
}
