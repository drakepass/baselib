package basefile

import (
	"bufio"
	"io"
	"io/ioutil"
	"os"
)

func ReadAll(filepath string) (str string, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	fileInfo, err := file.Stat()
	if err != nil {
		return
	}
	content := make([]byte, fileInfo.Size())
	_, err = file.Read(content)
	if err != nil {
		return
	}
	str = string(content)
	return
}

func ReadAllByBlock(filepath string, num int) (str string, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	content := make([]byte, 0)
	for {
		tmpcons := make([]byte, num)
		_, err = file.Read(tmpcons)
		if err != nil {
			if err == io.EOF {
				err = nil
				break
			}
			return
		}
		content = append(content, tmpcons...)
	}
	str = string(content)
	return
}

//this function will trim the "\r\n" from reading content
//ReadLine是一个低水平的行数据读取原语。应使用ReadBytes('\n')见下面ReadLineByDelim()或ReadString('\n')代替，或者使用Scanner。
func ReadAllByLine(filepath string) (str string, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	content := make([]byte, 0)
	bufferReader := bufio.NewReader(file)
	for {
		line, _, errs := bufferReader.ReadLine()
		if errs != nil {
			if errs == io.EOF {
				break
			}
			err = errs
			return
		}
		content = append(content, line...)
	}
	str = string(content)
	return
}

//比较可取的方法
//ReadBytes方法在读取到delim之前遇到了错误，它会返回在错误之前读取的数据以及该错误（一般是io.EOF）
func ReadAllByLineDelim(filepath string, delim byte) (str string, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	content := make([]byte, 0)
	bufReader := bufio.NewReader(file)
	for {
		line, errs := bufReader.ReadBytes(delim) //ReadString('\n') 用法一样
		if errs != nil {
			if errs == io.EOF { //出现错误了，仍会返回读取到的数据和错误
				content = append(content, line...)
				break //return
			}
			err = errs
			return
		}
		content = append(content, line...)
	}
	str = string(content)
	return
}

//利用ioutil包方法，要先通过os.open生成句柄
func ReadAllByIoutil(filepath string) (str string, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	str = string(content)
	return
}

//可取的方法
//利用ioutil包直接读取文件内容
func ReadAllByFile(filepath string) (str string, err error) {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		return
	}
	str = string(content)
	return
}

func WriteFile(filepath, content string) (err error) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	//file.Write([]byte(content)) 与下面等同
	file.WriteString(content)
	return
}

func WriteFileByBuf(filepath, content string) (err error) {
	file, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0777)
	if err != nil {
		return
	}
	defer file.Close()
	bufWriter := bufio.NewWriter(file)
	_, err = bufWriter.WriteString(content)
	if err != nil {
		return
	}
	bufWriter.Flush()
	return
}

//没有文件会创建，但会覆盖之前的全部内容
func WriteFileByIoutil(filepath, content string) (err error) {
	err = ioutil.WriteFile(filepath, []byte(content), 0777)
	return
}

//文件拷贝
func CopeFile(dstfile,srcfile string) (writen int64,err error) {
	dst,err := os.OpenFile(dstfile,os.O_RDWR|os.O_CREATE,0777)
	if err != nil {
		return
	}
	src,err := os.Open(srcfile)
	if err != nil {
		return
	}
	writen,err = io.Copy(dst,src)
	return
}

//