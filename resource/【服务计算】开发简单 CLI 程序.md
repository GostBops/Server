
这次作业对我来说很难，踩了很多坑，读了很多东西，总算是勉强完成了基本要求，有了一些了解

>老师的博客地址：[CLI 命令行实用程序开发基础][1]
>我的代码地址：[selpg.go][2]

那么这次博客的主题是解释selpg.c源程序，了解实现go程序的包和工具，测试程序运行

@[toc]

## 0、命令行程序

刚开始做作业的时候可能一头雾水，需要了解的内容非常多，需要实现的功能很复杂，其实通篇阅读下来，主要内容还是在第一个链接

[开发 Linux 命令行实用程序][3]

那么CLI是什么呢？

>CLI（Command Line Interface）实用程序是Linux下应用开发的基础。正确的编写命令行程序让应用与操作系统融为一体，通过shell或script使得应用获得最大的灵活性与开发效率。

我们平时在Linux终端输入的ls,cd,mkdir等命令，就是CLI，命令行程序，我们需要做的便是按照上面那个链接的要求，用go实现名称为selpg的命令行工具

在阅读源码之前，先举一些简单的例子
比如我们应该如何正确处理ls -a呢
每种后端语言都有识别输入参数的接口，“ls”和“-a”就对应两个输入参数，获取到它们后，我们解析这些参数，然后实现对应的功能，这就是命令行程序的工作

更规范的命令行格式请阅读上面的链接，如你所见，这只是一个小白教程

## 1、阅读selpg.c源码

[selpg.c][4]

selpg源码分为三个部分，注释都很详尽，作者也给了充足的解释，我主要从如何翻译成go的角度来解释一下这份源码

### process_args() 函数

#### 以注释“==== check the command-line arguments ===”开始的行

这个部分和接下来两个部分**“handle 1st arg”**和**“handle 2nd arg”**，在go中，我们需要使用os.Args来解析前两个参数
os.Args是一个储存了所有参数的string数组，我们可以使用下标来访问参数

    var Args []string

比如selpg -s 10 -e 10

    os.Args[0] == selpg
    os.Args[1] == -s
    os.Args[2] == 10


#### 以注释“now handle optional args”开始的行

这一部分和**“there is one more arg”**由c处理起来很复杂，但是go为我们提供了flag工具
flag可以自动帮我们解析绑定在它身上的给定参数，以及为我们生成合适的help信息

##### flag

参考资料：
[Golang之使用Flag和Pflag][5]
[Package pflag][6]

相关代码如下：
```go
flag.IntVarP(&sa.start_page,"start",  "s", -1, "start page(>1)")
flag.IntVarP(&sa.end_page,"end", "e",  -1, "end page(>=start_page)")
flag.IntVarP(&sa.page_len,"len", "l", 72, "page len")
flag.StringVarP(&sa.page_type,"type", "f", "l", "'l' for lines-delimited, 'f' for form-feed-delimited. default is 'I'")
flag.Lookup("type").NoOptDefVal = "f"
flag.StringVarP(&sa.print_dest,"dest", "d", "", "print dest")

flag.Usage = func() {
	fmt.Fprintf(os.Stderr,
		"USAGE: \n%s -s start_page -e end_page [ -f | -llines_per_page ]" + 
		" [ -ddest ] [ in_filename ]\n", pr测试
	flag.PrintDefaults()
}
flag.Parse()
```

这里使用的是pflag包，和flag用法基本相同

这里xxxVarP()函数把某个变量和参数名绑定

    func IntVarP(p *int, name, shorthand string, value int, usage string)
第一个参数为变量
第二个参数为命令行参数名，如"--start"
第三个参数为该参数的短名，如“-s”
第四个参数当该参数没有在命令行出现时的默认值，比如value=1，若命令行中没有“--start”或“-s”，则默认“--start=1”
第五个参数是help的帮助信息

    flag.Parse()
必须在所有flags被定义后，以及被执行前调用

```go
if len(flag.Args()) == 1 {
	_, err := os.Stat(flag.Args()[0])
	/* check if file exists */
	if err != nil && os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "\n%s: input file \"%s\" does not exist\n",
				progname, flag.Args()[0]);
		os.Exit(6);
	}
	sa.in_filename = flag.Args()[0]
}
```
flag.Args()返回非flag型的命令行参数，即前缀不是“--”和“-”的命令行参数，比如“input_file”，也是一个string数组

    func Args() []string

os.Stat（）检查了文件是否存在

### process_input（）函数

#### 以注释“set the input source”开始的行

我们需要阅读go的os包，了解go的文件读写机制，用os.Open（）打开一个文件
我们来可以使用bufio.NewReader（）带缓冲地读取，并且调用ReadString（）可以按规则读取一段数据

##### os

参考资料：
[Go 文件读写][7]

```go
var fin *os.File 
if len(sa.in_filename) == 0 {
	fin = os.Stdin
} else {
	var err error
	fin, err = os.Open(sa.in_filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "\n%s: could not open input file \"%s\"\n",
			progname, sa.in_filename)
		os.Exit(7)
	}
	defer fin.Close()
}
bufFin := bufio.NewReader(fin)
```

十分简单，小白友好型代码

#### 以注释“set the output destination”开始的行

我们需要os/exec来执行lp命令，创建进程，并开启管道实现selpg进程和lp进程间的通信

##### Command

参考资料：
[golang中os/exec包的用法][8]

```go
var fout io.WriteCloser
cmd := &exec.Cmd{}

if len(sa.print_dest) == 0 {
	fout = os.Stdout
} else {
	cmd = exec.Command("cat")
	//没法测试lp，用cat代替测试
	//cmd = exec.COmmand("lp", "-d", sa.print_dest)
	var err error
	cmd.Stdout, err = os.OpenFile(sa.print_dest, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	fout, err = cmd.StdinPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "\n%s: can't open pipe to \"lp -d%s\"\n",
			progname, sa.print_dest)
		os.Exit(8)
	}
	cmd.Start()
}
```

通过调用exec.Command()执行命令，返回一个cmd的结构体指针
cmd.StdinPipe()返回一个连接到command标准输入的管道
cmd.StdoutPipe()返回一个连接到command标准输出的管道pipe
下面我们会使用fout通过command管道向print_dest文件写入数据
cmd.Start()使某个命令开始执行，但是并不等到他执行结束，与cmd.Wait()配套使用

另外还要注意的一点是，os.Open()打开的文件只有读权限，所以要写数据应该使用os.OpenFile()，并且填入读写权限参数

参考资料：
[golang bad file descriptor][9]

#### 以注释“begin one of two main loops”开始的行

这里我们使用ReadString（）逐页地读取input_file，输入到管道中，并输出到指定地点

```go
if sa.page_type == "l" {
	line_ctr := 0
	page_ctr = 1
	for {
		line,  crc := bufFin.ReadString('\n')
		if crc != nil {
			break
		}
		line_ctr++
		if line_ctr > sa.page_len {
			page_ctr++
			line_ctr = 1
		}

		if (page_ctr >= sa.start_page) && (page_ctr <= sa.end_page) {
			_, err := fout.Write([]byte(line))
			if err != nil {
				fmt.Println(err)
				os.Exit(9)
			}
	 	}
	}  
}
```

这是page_type为“l”时的读取情况，小白友好型代码

## 2、selpg使用测试

	$ selpg
![0](https://img-blog.csdn.net/20181004222252310?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)


	$ selpg -s 10 -e 20 input_file
该命令将把“input_file”的第 1 页写至标准输出（也就是屏幕），因为这里没有重定向或管道。
![1](https://img-blog.csdn.net/20181004215354239?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

	$ selpg -s 10 -e 20 < input_file
该命令与示例 1 所做的工作相同，但在本例中，selpg 读取标准输入，而标准输入已被 shell／内核重定向为来自“input_file”而不是显式命名的文件名参数。输入的第 1 页被写至屏幕。

![2](https://img-blog.csdn.net/20181004215749153?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

	$ other_command | selpg -s 10 -e 20
“other_command”的标准输出被 shell／内核重定向至 selpg 的标准输入。将第 10 页到第 20 页写至 selpg 的标准输出（屏幕）。

![3](https://img-blog.csdn.net/20181004215940491?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

	$ selpg -s 10 -e 20 input_file >output_file
selpg 将第 10 页到第 20 页写至标准输出；标准输出被 shell／内核重定向至“output_file”。
![4](https://img-blog.csdn.net/20181004220313809?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

![5](https://img-blog.csdn.net/2018100422032596?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

	$ selpg -s 10 -e 20 input_file 2>error_file
selpg 将第 10 页到第 20 页写至标准输出（屏幕）；所有的错误消息被 shell／内核重定向至“error_file”。
![6](https://img-blog.csdn.net/20181004220619573?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)
![7](https://img-blog.csdn.net/20181004220629376?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

	$ selpg -s 10 -e 20 input_file >output_file 2>error_file
selpg 将第 10 页到第 20 页写至标准输出，标准输出被重定向至“output_file”；selpg 写至标准错误的所有内容都被重定向至“error_file”。当“input_file”很大时可使用这种调用；您不会想坐在那里等着 selpg 完成工作，并且您希望对输出和错误都进行保存。
![8](https://img-blog.csdn.net/20181004220936952?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)
![9](https://img-blog.csdn.net/20181004220950658?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)
![10](https://img-blog.csdn.net/20181004221008870?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

	$ selpg -s 10 -e 20 input_file >output_file 2>/dev/null
selpg 将第 10 页到第 20 页写至标准输出，标准输出被重定向至“output_file”；selpg 写至标准错误的所有内容都被重定向至 /dev/null（空设备），这意味着错误消息被丢弃了。设备文件 /dev/null 废弃所有写至它的输出，当从该设备文件读取时，会立即返回 EOF。

![11](https://img-blog.csdn.net/20181004221237829?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

![12](https://img-blog.csdn.net/20181004221248804?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

	$ selpg -s 10 -e 20 input_file >/dev/null
selpg 将第 10 页到第 20 页写至标准输出，标准输出被丢弃；错误消息在屏幕出现。这可作为测试 selpg 的用途，此时您也许只想（对一些测试情况）检查错误消息，而不想看到正常输出。
![13](https://img-blog.csdn.net/20181004221350920?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

	$ selpg -s 10 -e 20 input_file | other_command
selpg 的标准输出透明地被 shell／内核重定向，成为“other_command”的标准输入，第 10 页到第 20 页被写至该标准输入。“other_command”的示例可以是 lp，它使输出在系统缺省打印机上打印。“other_command”的示例也可以 wc，它会显示选定范围的页中包含的行数、字数和字符数。“other_command”可以是任何其它能从其标准输入读取的命令。错误消息仍在屏幕显示。

![14](https://img-blog.csdn.net/20181004221734497?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

	$ selpg -s 10 -e 20 input_file 2>error_file | other_command

![15](https://img-blog.csdn.net/20181004222012605?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)
![16](https://img-blog.csdn.net/20181004222053908?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

	$ selpg -s 10 -e 20 -l 66 input_file
该命令将页长设置为 66 行，这样 selpg 就可以把输入当作被定界为该长度的页那样处理。第 10 页到第 20 页被写至 selpg 的标准输出（屏幕）。

![17](https://img-blog.csdn.net/20181004222603426?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

	$ selpg -s10 -e20 -f input_file
假定页由换页符定界。第 10 页到第 20 页被写至 selpg 的标准输出（屏幕）。
(由于TXT文件没有换页符，程序改用换行符测试)

![18](https://img-blog.csdn.net/20181004222952693?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)

	$ selpg -s 10 -e 20 -d lp1 input_file

第 10 页到第 20 页由管道输送至命令“lp -dlp1”，该命令将使输出在打印机 lp1 上打印。
（由于没有连接打字机设备，改换“cat”命令，该命令将使输出在文件lp1上打印）

![19](https://img-blog.csdn.net/20181004224007904?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)
![20](https://img-blog.csdn.net/20181004224103578?watermark/2/text/aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2hjbV8wMDc5/font/5a6L5L2T/fontsize/400/fill/I0JBQkFCMA==/dissolve/70)


  [1]: https://pmlpml.github.io/ServiceComputingOnCloud/ex-cli-basic
  [2]: https://github.com/GostBop/Go-Learning/blob/master/selpg/selpg.go
  [3]: https://www.ibm.com/developerworks/cn/linux/shell/clutil/index.html
  [4]: https://www.ibm.com/developerworks/cn/linux/shell/clutil/selpg.c
  [5]: https://o-my-chenjian.com/2017/09/20/Using-Flag-And-Pflag-With-Golang/
  [6]: https://godoc.org/github.com/spf13/pflag#Parse
  [7]: http://wiki.jikexueyuan.com/project/the-way-to-go/12.2.html
  [8]: https://blog.csdn.net/chenbaoke/article/details/42556949
  [9]: https://blog.csdn.net/benben_2015/article/details/80607425
