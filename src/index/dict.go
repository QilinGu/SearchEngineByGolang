package index

import (
  "fmt"
  "os"
  "bufio"
)

const maxbufsize = 4096

type dictionary struct {
  dict map[string]int
  curSize int
  reader *bufio.Reader
  writer *bufio.Writer
  path string
}

func (d *dictionary) genTmpIndexFile(key string, docid int) {
  wrdid := d.getWordValue(key)

}


func (d *dictionary)getWordValue(key string) int {
  // 查找键值是否存在
  if v, ok := d.dict[key]; !ok {
    d.dick[key] = ++d.curSize
  	//fmt.Println(v)
    return d.curSize
  }
  return d.dict[key]
}

func (d *dictionary)loadDictFile(filename string) map[string]uint32 {
  fi, err := os.Open(filename)
  if err != nil {
    panic(err)
  }
  defer fi.Close()
  dict := make(map[string]uint32)
  d.reader = bufio.NewReader(fi)
  //read buffer to map
  for {
		line, err := d.reader.ReadString('\n')
		line = strings.TrimSpace(line)
		//handler(line)
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
	}
  return dict
}

func (d *dictionary)writeDictFile(filename string) {
  if checkFileIsExist(filename) {  //如果文件存在
    fo, _ = os.OpenFile(filename, os.O_APPEND, 0666)  //打开文件
    fmt.Println("file exist!");
   }else {
    fo, _ = os.Create(filename)  //创建文件
    fmt.Println("file is not exist!");
   }
   /*strkey := []byte(key)
   if len(key) > maxkeylen {
     strkey = strings(key[0:maxkeylen-1])
   }*/
   d.writer = bufio.NewWriter(fo)
   buf := make([]byte, maxkeylen+6)
   for k, v := range m1 {
     strkey = strings(k[0:maxkeylen])
     strval := strconv.Itoa(v)
     str := strkey + "," + strval + "\n"//strings.Join({strkey, strval}, ",")
     err := d.writer.WriteFile(str)  //写入文件(字节数组)
     if !err {
       fmt.Println("write file error!");
     }
   }
}

/**
 * 判断文件是否存在  存在返回 true 不存在返回false
 */
func checkFileIsExist(filename string) (bool) {
 var exist = true;
 if _, err := os.Stat(filename); os.IsNotExist(err) {
  exist = false;
 }
 return exist;
}